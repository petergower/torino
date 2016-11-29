package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cmd := exec.Command("./_test/main")
	//cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
	startEngine()
}

func startEngine() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d containers found\n", len(containers))

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}
