package storage

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewStorage(t *testing.T) {
	dir := "dir"
	stor, err := NewStorage(dir)
	require.NoError(t, err)
	require.NotEmpty(t, stor)

	DeleteDir(dir)
}

func TestStorage_Add(t *testing.T) {
	tmpDir := "tmp"
	tmpFile := "zip.zip"
	dir := "dir"
	stor, err := NewStorage(dir)
	require.NoError(t, err)
	require.NotEmpty(t, stor)
	CreateTestFile(tmpDir, tmpFile)

	f, err := NewZipFile(fmt.Sprintf("%s/%s", tmpDir, tmpFile))
	require.NoError(t, err)
	require.Equal(t, tmpFile, f.Name)

	stor.Add(f)
	require.Len(t, stor.Files, 1)

	DeleteDir(tmpDir)
	DeleteDir(dir)
}
