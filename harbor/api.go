package harbor

import "context"

type Interface interface {
	GetMetadata(ctx context.Context, image string)
}

type harborImpl struct {
}

// GetMetadata implements Interface.
func (*harborImpl) GetMetadata(ctx context.Context, image string) {
	panic("unimplemented")
}

func New() Interface {
	return &harborImpl{}
}
