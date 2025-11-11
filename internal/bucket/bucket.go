package bucket

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

type BucketType int

const (
	AwsProvider BucketType = iota
)

type BucketInterface interface {
	Upload(io.Reader, string) error
	Download(string, string) (os.File, error)
	Delete(string) error
}

type Bucket struct {
	p BucketInterface
}

func New(bt BucketType, cfg any) (b *Bucket, err error) {
	// Implementação futura para diferentes tipos de bucket
	rt := reflect.TypeOf(cfg)

	switch bt {
	case AwsProvider:
		if rt.Name() != "AwsConfig" {
			return nil, fmt.Errorf("invalid config type for AWS S3")
		}
		return &Bucket{}, nil
	default:
		return nil, fmt.Errorf("invalid bucket type")
	}
}

func (b *Bucket) Upload(file io.Reader, key string) error {
	return b.p.Upload(file, key)
}

func (b *Bucket) Download(src, dest string) (file os.File, err error) {
	return b.p.Download(src, dest)
}

func (b *Bucket) Delete(key string) error {
	return b.p.Delete(key)
}
