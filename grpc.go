package emboss

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"image"
	_ "image/png"
	"io"
	"io/ioutil"
	_ "log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	emboss_grpc "github.com/sfomuseum/go-image-emboss/v2/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GrpcEmbosser struct {
	conn   *grpc.ClientConn
	client emboss_grpc.ImageEmbosserClient
}

func init() {
	ctx := context.Background()
	RegisterEmbosser(ctx, "grpc", NewGrpcEmbosser)
}

func NewGrpcEmbosser(ctx context.Context, uri string) (Embosser, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	q := u.Query()

	q_tls_cert := q.Get("tls-certificate")
	q_tls_key := q.Get("tls-key")
	q_tls_ca := q.Get("tls-ca-certificate")
	q_tls_insecure := q.Get("tls-insecure")

	addr := u.Host

	opts := make([]grpc.DialOption, 0)

	if q_tls_cert != "" && q_tls_key != "" {

		cert, err := tls.LoadX509KeyPair(q_tls_cert, q_tls_key)

		if err != nil {
			return nil, fmt.Errorf("Failed to load TLS pair, %w", err)
		}

		tls_config := &tls.Config{
			Certificates: []tls.Certificate{cert},
		}

		if q_tls_ca != "" {

			ca_cert, err := ioutil.ReadFile(q_tls_ca)

			if err != nil {
				return nil, fmt.Errorf("Failed to create CA certificate, %w", err)
			}

			cert_pool := x509.NewCertPool()

			ok := cert_pool.AppendCertsFromPEM(ca_cert)

			if !ok {
				return nil, fmt.Errorf("Failed to append CA certificate, %w", err)
			}

			tls_config.RootCAs = cert_pool

		} else if q_tls_insecure != "" {

			v, err := strconv.ParseBool(q_tls_insecure)

			if err != nil {
				return nil, fmt.Errorf("Failed to parse ?tls-insecure= parameter, %w", err)
			}

			tls_config.InsecureSkipVerify = v
		}

		tls_credentials := credentials.NewTLS(tls_config)
		opts = append(opts, grpc.WithTransportCredentials(tls_credentials))

	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// This is necessary to disable the dynamicWindow flags in internal/http2_client.go
	// which is what triggers the BDP estimator code which, in turn, is what triggers the
	// too many ping and subsequent ENHANCE_YOUR_CALM errros. See also:
	// https://github.com/grpc/grpc-swift-2/issues/5#issuecomment-2984421768
	//
	// 65535 is the initial window size in internal/transport/defaults.go and anything
	// greater than this will disable the BDP estimator stuff

	window_sz := int32(65535 + 1)
	opts = append(opts, grpc.WithInitialWindowSize(window_sz))

	conn, err := grpc.NewClient(addr, opts...)

	if err != nil {
		return nil, fmt.Errorf("Failed to dial '%s', %w", addr, err)
	}

	// defer conn.Close()

	client := emboss_grpc.NewImageEmbosserClient(conn)

	e := &GrpcEmbosser{
		conn:   conn,
		client: client,
	}

	return e, nil
}

func (e *GrpcEmbosser) EmbossImage(ctx context.Context, path string, combined bool) ([]image.Image, error) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	im_r, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("Failed to open %s for reading, %w", path, err)
	}

	defer im_r.Close()

	return e.EmbossImageWithReader(ctx, im_r, path, combined)
}

func (e *GrpcEmbosser) EmbossImageWithReader(ctx context.Context, im_r io.Reader, path string, combined bool) ([]image.Image, error) {

	fname := filepath.Base(path)

	body, err := io.ReadAll(im_r)

	if err != nil {
		return nil, fmt.Errorf("Failed to read %s, %w", path, err)
	}

	req := &emboss_grpc.EmbossImageRequest{
		Filename: fname,
		Body:     body,
		Combined: combined,
	}

	rsp, err := e.client.EmbossImage(ctx, req)

	if err != nil {
		return nil, fmt.Errorf("Failed to emboss image, %w", err)
	}

	images := make([]image.Image, len(rsp.Body))

	for idx, data := range rsp.Body {

		r := bytes.NewReader(data)
		im, _, err := image.Decode(r)

		if err != nil {
			return nil, fmt.Errorf("Failed to decode image at offset %d, %w", idx, err)
		}

		images[idx] = im
	}

	return images, nil
}

func (e *GrpcEmbosser) Close(ctx context.Context) error {
	return e.conn.Close()
}
