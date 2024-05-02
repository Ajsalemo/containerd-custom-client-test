package main

import (
	"log"
	"context"

	ctr "containerd-custom-client/ctr"
	containerd "github.com/containerd/containerd/v2/client"
)

// Todo - replace loggers with something like `zapp`
func main() {
	log.Println("Executing main.go")
	client, ctx, err := ctr.ContainerdClient()
	if err != nil {
		log.Printf("An error occurred when using the containerd client..")
		log.Fatal(err)
	}
	// Close the client later on
	defer client.Close()

	log.Println("Preparing to pull image..")
	err2 := pullImage(client, ctx)
	if err2 != nil {
		log.Println("ERROR: Image pull failed..")
		log.Fatal(err)
	}

	err3 := listImages(client, ctx)
	if err3 != nil {
		log.Println("ERROR: Listing images failed..")
		log.Fatal(err)
	}
}

func pullImage(client *containerd.Client, ctx context.Context) error {
	image, err := client.Pull(ctx, "docker.io/library/redis:latest", containerd.WithPullUnpack)
	if err != nil {
		return err
	}

	log.Printf("Successfully pulled %s image\n", image.Name())

	return nil
}

func listImages(client *containerd.Client, ctx context.Context) error {
	images, err := client.ListImages(ctx)

	if err != nil {
		return err
	}
	log.Println("Listing currently downloaded images..")
	for _, image := range images {
		log.Println(" - " + image.Name())
	}

	return nil
}