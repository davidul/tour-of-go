package main

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	list, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range list {
		fmt.Printf("%s %s\n", container.ID, container.Image)
	}

	build()
}

func build(){
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	dockerFile := "c:\\Users\\davidul\\IdeaProjects\\_work\\_angie_demo\\i15n-api-ms\\Dockerfile"
	dockerFileReader, err := os.Open(dockerFile)
	if err != nil {
		panic(err)
	}

	readDockerFile, err := ioutil.ReadAll(dockerFileReader)

	header := &tar.Header{
		Name: dockerFile,
		Size: int64(len(readDockerFile)),
	}

	err = tw.WriteHeader(header)
	if err != nil {
		panic(err)
	}

	dockerFileTarReader := bytes.NewReader(buf.Bytes())

	imageBuildResponse, err := cli.ImageBuild(
		ctx,
		dockerFileTarReader,
		types.ImageBuildOptions{
			Context:    dockerFileTarReader,
			Dockerfile: dockerFile,
			Remove:     true})
	if err != nil {
		log.Fatal(err, " :unable to build docker image")
	}
	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		log.Fatal(err, " :unable to read image build response")
	}
}
