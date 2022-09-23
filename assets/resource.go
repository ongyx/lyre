package assets

import (
	"image"
	_ "image/png"
	"io/fs"
)

type Resource struct {
	path string
	fs   fs.FS
}

func NewResource(path string) *Resource {
	return &Resource{path: path, fs: filesystem}
}

func (r *Resource) Open() (fs.File, error) {
	return r.fs.Open(r.path)
}

func (r *Resource) OpenImage() (image.Image, error) {
	f, err := r.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}
