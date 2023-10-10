package s3

import "context"

type Interface interface {
	Get(ctx context.Context, obj string)
	Delete(ctx context.Context, obj string)
	Download(ctx context.Context, obj string)
}

type s3Wrapper struct {
	bucket string
}

// Delete implements Interface.
func (*s3Wrapper) Delete(ctx context.Context, obj string) {
	panic("unimplemented")
}

// Download implements Interface.
func (*s3Wrapper) Download(ctx context.Context, obj string) {
	panic("unimplemented")
}

// Get implements Interface.
func (*s3Wrapper) Get(ctx context.Context, obj string) {
	panic("unimplemented")
}

func New(bkt string) Interface {
	return &s3Wrapper{bucket: bkt}
}
