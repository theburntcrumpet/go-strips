package service

import (
	"archive/zip"
	"io"
	"strings"
)

type ComicPage struct {
	Filename string
	Data     []byte
}

type Comic struct {
	Filename string
	Pages    []ComicPage
}

type ComicParser interface {
	ParseComic(filename string) (Comic, error)
}

type CbzComicParser struct{}

func (p CbzComicParser) ParseComic(filename string) (Comic, error) {
	comic := Comic{Filename: filename, Pages: []ComicPage{}}

	// extract cbt
	r, err := zip.OpenReader(filename)
	if err != nil {
		return comic, err
	}

	defer r.Close()

	for _, f := range r.File {
		if !strings.HasSuffix(f.Name, ".jpg") {
			continue
		}
		d, err := f.Open()
		if err != nil {
			return comic, err
		}
		data, err := io.ReadAll(d)
		if err != nil {
			return comic, err
		}
		comic.Pages = append(comic.Pages, ComicPage{Filename: f.Name, Data: data})
		defer d.Close()
	}

	return comic, nil
}
