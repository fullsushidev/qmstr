package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/QMSTR/qmstr/pkg/master"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

func CopyResults(ctx context.Context, cli *client.Client, container string, destinationPath string) error {
	data, stat, err := cli.CopyFromContainer(ctx, container, master.ServerOutputDir)
	if err != nil {
		return err
	}
	defer data.Close()

	srcInfo := archive.CopyInfo{
		Path:   master.ServerOutputDir,
		Exists: true,
		IsDir:  stat.Mode.IsDir(),
	}

	return archive.CopyTo(data, srcInfo, destinationPath)
}

func WriteContainerFile(ctx context.Context, cli *client.Client, data []byte, container string, destination string) error {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	hdr := &tar.Header{
		Name: destination,
		Mode: 0664,
		Size: int64(len(data)),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		return err
	}
	if _, err := tw.Write(data); err != nil {
		return err
	}
	if err := tw.Close(); err != nil {
		return err
	}

	log.Printf("Writing %d bytes to %s", hdr.Size, destination)
	reader := bytes.NewReader(buf.Bytes())

	err := cli.CopyToContainer(ctx, container, "/", reader, types.CopyToContainerOptions{})
	return err
}

func GetMasterConfig(ctx context.Context, cli *client.Client, container string, containerConfigPath string) ([]byte, error) {
	data, _, err := cli.CopyFromContainer(ctx, container, containerConfigPath)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	config, err := ioutil.ReadAll(data)
	if err != nil {
		return nil, err
	}

	// poor man's untar this might be replaced by a proper untar function
	tarPattern := regexp.MustCompile("^.*package:")
	config = tarPattern.ReplaceAll(config, []byte("package:"))
	config = bytes.Trim(config, "\x00")

	return config, nil
}