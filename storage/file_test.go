package storage

import (
	"fmt"
	"github.com/dubrovin/gotest/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewZipFile(t *testing.T) {
	tmpDir := "tmp"
	tmpFile := "zip.zip"
	utils.CreateTestFile(tmpDir, tmpFile)

	f, err := NewZipFile(fmt.Sprintf("%s/%s", tmpDir, tmpFile))
	require.NoError(t, err)
	require.Equal(t, tmpFile, f.Name)

	utils.DeleteDir(tmpDir)
}

func TestRead(t *testing.T) {
	tmpDir := "tmp"
	tmpFile := "zip.zip"
	utils.CreateTestFile(tmpDir, tmpFile)

	f, err := NewZipFile(fmt.Sprintf("%s/%s", tmpDir, tmpFile))
	require.NoError(t, err)
	require.Equal(t, tmpFile, f.Name)
	files, err := f.Read()
	require.NoError(t, err)
	require.Len(t, files, 3)
	utils.DeleteDir(tmpDir)
}
