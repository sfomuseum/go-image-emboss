package main

import (
	"context"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/aaronland/go-colours/app/review"
	"github.com/sfomuseum/go-image-emboss"
)

func main() {

	ctx := context.Background()

	var embosser_uri string

	fs := review.DefaultFlagSet()
	fs.StringVar(&embosser_uri, "embosser-uri", "grpc://localhost:8080", "A valid sfomuseum/go-image-emboss.Embosser URI.")

	opts, err := review.RunOptionsFromFlagSet(fs)

	if err != nil {
		log.Fatal(err)
	}

	em, err := emboss.NewEmbosser(ctx, embosser_uri)

	if err != nil {
		log.Fatalf("Failed to create new embosser, %v", err)
	}

	root_dir, err := os.MkdirTemp("", "colours")

	if err != nil {
		log.Fatalf("Failed to create tmp dir, %v", err)
	}

	defer os.RemoveAll(root_dir)
	abs_root := root_dir

	// START OF rewrite list of images to review
	// by passing each item in opts.Images through the image "embosser" and appending
	// image segmentations

	images := make([]string, 0)

	for _, path := range opts.Images {

		if strings.HasPrefix(path, "https://") || strings.HasPrefix(path, "http"){

			if !opts.AllowRemote {
				log.Fatalf("Remote URIs forbidden")
			}
			
			rsp, err := http.Get(path)

			if err != nil {
				log.Fatalf("Failed to fetch %s, %v", path, err)
			}

			defer rsp.Body.Close()

			new_fname := filepath.Base(path)
			new_path := filepath.Join(abs_root, new_fname)

			new_wr, err := os.OpenFile(new_path, os.O_RDWR|os.O_CREATE, 0644)

			if err != nil {
				log.Fatalf("Failed to open %s for writing, %v", new_path, err)
			}

			_, err = io.Copy(new_wr, rsp.Body)

			if err != nil {
				log.Fatalf("Failed to copy %s to %s, %v", path, new_path, err)
			}

			err = new_wr.Close()

			if err != nil {
				log.Fatalf("Failed to close %s after writing, %v", new_path, err)
			}

			path = new_wr.Name()
		}

		fname := filepath.Base(path)
		ext := filepath.Ext(fname)
		fname = strings.Replace(fname, ext, ".png", 1)

		r, err := os.Open(path)

		if err != nil {
			log.Fatalf("Failed to open %s for reading, %v", path, err)
		}

		defer r.Close()

		im, _, err := image.Decode(r)

		if err != nil {
			log.Fatalf("Failed to decode %s, %v", path, err)
		}

		im_path := filepath.Join(abs_root, fname)

		im_wr, err := os.OpenFile(im_path, os.O_RDWR|os.O_CREATE, 0644)

		if err != nil {
			log.Fatalf("Failed to open %s for writing, %v", im_path, err)
		}

		err = png.Encode(im_wr, im)

		if err != nil {
			log.Fatalf("Failed to encode %s, %v", im_path, err)
		}

		err = im_wr.Close()

		if err != nil {
			log.Fatalf("Failed to close %s after writing, %v", im_path, err)
		}

		images = append(images, im_path)

		combined := false

		rsp, err := em.EmbossImage(ctx, path, combined)

		if err != nil {
			log.Fatalf("Failed to extract image from %s, %v", path, err)
		}

		for idx, im := range rsp {

			im_fname := fmt.Sprintf("%03d-%s", idx+1, fname)
			im_path := filepath.Join(abs_root, im_fname)

			im_wr, err := os.OpenFile(im_path, os.O_RDWR|os.O_CREATE, 0644)

			if err != nil {
				log.Fatalf("Failed to open %s for writing, %v", im_path, err)
			}

			err = png.Encode(im_wr, im)

			if err != nil {
				log.Fatalf("Failed to encode %s, %v", im_path, err)
			}

			err = im_wr.Close()

			if err != nil {
				log.Fatalf("Failed to close %s after writing, %v", im_path, err)
			}

			images = append(images, im_path)
		}
	}

	opts.Images = images

	// END OF rewrite list of images to review

	err = review.RunWithOptions(ctx, opts)

	if err != nil {
		log.Fatal(err)
	}
}
