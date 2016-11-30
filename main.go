package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/fsouza/go-dockerclient"
)

func main() {
	endpoint := "unix:///var/run/docker.sock"

	// test with NewClientWithEnv for custom params
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	// http://stackoverflow.com/questions/36279253/go-compiled-binary-wont-run-in-an-alpine-docker-container-on-ubuntu-host

	outputbuf, err := buildImage(client)
	if err != nil {
		panic(err)
	}
	fmt.Println(outputbuf.String())

	container, err := createContainer(client, "test")
	if err != nil {

		panic(err)
	}

	fmt.Println("Container created")

	err = client.StartContainer(container.ID, nil)

	if err != nil {
		panic(err)
	}

	err = client.StopContainer(container.ID, 1)
	if err != nil {
		panic(err)

	}
	logsOpt := docker.LogsOptions{
		OutputStream: os.Stdout,
		Container:    container.ID,
		Stdout:       true,
	}

	err = client.Logs(logsOpt)
	if err != nil {

		panic(err)
	}
	fmt.Println("Done")
}

func buildImage(client *docker.Client) (*bytes.Buffer, error) {
	// https://docs.docker.com/engine/reference/api/docker_remote_api_v1.23/#start-a-container
	//inputbuf,
	outputbuf := bytes.NewBuffer(nil)
	//, bytes.NewBuffer(nil)
	opts := docker.BuildImageOptions{
		Name: "test",
		//InputStream:  inputbuf,
		OutputStream: outputbuf,
		ContextDir:   "images/torino",
		Dockerfile:   "Dockerfile",
	}

	return outputbuf, client.BuildImage(opts)

}

func createContainer(client *docker.Client, imageName string) (*docker.Container, error) {

	opts := docker.CreateContainerOptions{}

	opts.Config = &docker.Config{
		Image:        imageName,
		Cmd:          []string{"/bin/torino"},
		AttachStdout: true,
	}

	return client.CreateContainer(opts)

}
