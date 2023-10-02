package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/sfomuseum/go-image-emboss"
)

func main() {

	embosser_uri := flag.String("embosser-uri", "local:///usr/local/sfomuseum/bin/image-emboss", "A valid sfomuseum/go-image-emboss.Embosser URI.")

	flag.Parse()

	ctx := context.Background()

	em, err := emboss.NewEmbosser(ctx, *embosser_uri)

	if err != nil {
		log.Fatalf("Failed to create new embosser, %v", err)
	}

	defer em.Close(ctx)

	for _, path := range flag.Args() {

		rsp, err := em.EmbossImage(ctx, path)

		if err != nil {
			log.Fatalf("Failed to extract image from %s, %v", path, err)
		}

		fmt.Println(rsp)
	}
}
