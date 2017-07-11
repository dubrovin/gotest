package storage

import (
	"os"
	"path/filepath"
)

type Storage struct {
	Files   map[string]*ZipFile
	RootDir string
}

func NewStorage(dir string) (*Storage, error) {
	rootDir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		err := os.Mkdir(rootDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	return &Storage{
		Files:   make(map[string]*ZipFile, 0),
		RootDir: rootDir,
	}, nil
}

func (s *Storage) Add(file *ZipFile) {
	s.Files[file.Name] = file
}

func (s *Storage) ListFileNames() (names []string) {
	for name := range s.Files {
		names = append(names, name)
	}
	return names
}
