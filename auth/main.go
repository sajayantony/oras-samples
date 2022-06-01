package main

import (
	"context"
	"fmt"

	registry "oras.land/oras-go/v2/registry/remote"
	"oras.land/oras-go/v2/registry/remote/auth"
)

func main() {
	ctx := context.Background()
	ref := "localhost:5000/busybox:latest"
	repo, err := registry.NewRepository(ref)
	if err != nil {
		panic(err)
	}

	// Prepare the creds function
	creds := func(ctx context.Context, hostname string) (auth.Credential, error) {
		return auth.Credential{
			Username: "myuser",
			Password: "mypass",
		}, nil
	}

	repo.Client = &auth.Client{
		Credential: creds,
	}

	repo.PlainHTTP = true

	desc, err := repo.Manifests().Resolve(ctx, ref)
	if err != nil {
		panic(err)
	}

	fmt.Println(desc.Digest)
}
