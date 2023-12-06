package mai

import (
	"context"
	"fmt"

	registry "oras.land/oras-go/v2/registry/remote"
)

func main() {
	ctx := context.Background()
	ref := "ghcr.io/oras-project/oras:v1.1.0"
	bogusref := "ghcr.io/oras-project/oras:invalid"

	repo, err := registry.NewRepository(ref)
	if err != nil {
		panic(err)
	}

	desc, err := repo.Manifests().Resolve(ctx, ref)
	if err != nil {
		panic(err)
	}

	fmt.Println(desc.Digest)

	_, err = repo.Resolve(ctx, bogusref)
	fmt.Printf("%v\n", err)

}
