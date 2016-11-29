package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/strslice"

	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	var exist bool
	var contID string
	for _, cnt := range containers {
		fmt.Printf("ID: %s Image:  %s\n", cnt.ID[:10], cnt.Image)

		if cnt.Names[0] == "test-container" {
			exist = true
			contID = cnt.ID
		}

	}

	if !exist {

		fmt.Println("Container does not exist.")

		config := container.Config{}
		config.Image = "docker/whalesay"
		config.Cmd = strslice.StrSlice{"cowsay", "boo boo boo"}

		config.AttachStdout = true

		// In the context is possible to send a timeout so...
		cont, err := cli.ContainerCreate(context.Background(), &config, &container.HostConfig{}, &network.NetworkingConfig{}, "")

		if err != nil {
			panic(err)
		}

		contID = cont.ID
	}

	opt := types.ContainerStartOptions{}

	err = cli.ContainerStart(context.Background(), contID, opt)

	if err != nil {
		panic(err)
	}
	lOpt := types.ContainerLogsOptions{}
	lOpt.ShowStdout = true
	rc, err := cli.ContainerLogs(context.Background(), contID, lOpt)

	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)
	s := buf.String()

	fmt.Println(s)
}
