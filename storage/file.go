package storage

import (
	"archive/zip"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
)

const (
	ZipExtension = ".zip"
)

type File struct {
	Name string
	Size int64
	Ext  string
}

type ZipFile struct {
	File
	Path       string
	InnerFiles []string
}

func NewZipFile(path string) (*ZipFile, error) {
	f, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	ext := filepath.Ext(path)
	if ext != ZipExtension {
		return nil, errors.New(fmt.Sprintf("expected %s extention, received %s extention", ZipExtension, ext))
	}
	return &ZipFile{
		Path: path,
		File: File{
			Name: f.Name(),
			Size: f.Size(),
			Ext:  filepath.Ext(path),
		},
	}, nil
}

func (zp *ZipFile) fillInnerFiles(r *zip.ReadCloser) {
	for _, f := range r.File {
		// fill slice of inner files of zipped file
		zp.InnerFiles = append(zp.InnerFiles, f.Name)
	}
}

func (zp *ZipFile) Read() (map[string][]byte, error) {
	// Open a zip archive for reading.
	r, err := zip.OpenReader(zp.Path)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	files := make(map[string][]byte, len(r.File))

	if len(zp.InnerFiles) == 0 {
		zp.fillInnerFiles(r)
	}

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		// читаем файл
		readedFile, err := ioutil.ReadAll(rc)
		if err != nil {
			return nil, err
		}
		files[f.Name] = readedFile

		// закрываем файл
		err = rc.Close()
		if err != nil {
			return nil, err
		}
	}
	return files, nil
}

func (zp *ZipFile) Bytes() ([]byte, error) {
	return ioutil.ReadFile(zp.Path)
}