package cache

import (
	"github.com/dubrovin/gotest/storage"
	"github.com/dubrovin/gotest/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"fmt"
)

func TestNewCache(t *testing.T) {
	dir := "dir"
	stor, err := storage.NewStorage(dir)
	require.NoError(t, err)
	c := NewCache(stor, 1000)
	t.Log(c)
	utils.DeleteDir(dir)
}

func TestCache_GetZipFile(t *testing.T) {
	dir := "dir"
	tmpDir := "tmp"
	tmpFile := "zip.zip"

	// создаем хранилище
	stor, err := storage.NewStorage(dir)
	require.NoError(t, err)
	require.NotEmpty(t, stor)

	// создаем тестовый зип файл и заполняем его
	utils.CreateTestFile(tmpDir, tmpFile)
	f, err := storage.NewZipFile(fmt.Sprintf("%s/%s", tmpDir, tmpFile))
	require.NoError(t, err)
	require.Equal(t, tmpFile, f.Name)

	// добавляем зип файл в хранилище
	stor.Add(f)
	require.Len(t, stor.Files, 1)

	// создаем кэш со стором
	c := NewCache(stor, 1000)

	b, err := c.GetZipFile(tmpFile)
	require.NoError(t, err)
	require.NotNil(t, b)

	utils.DeleteDir(tmpDir)
	utils.DeleteDir(dir)
}

func TestCache_GetFiles(t *testing.T) {
	dir := "dir"
	tmpDir := "tmp"
	tmpFile := "zip.zip"

	// создаем хранилище
	stor, err := storage.NewStorage(dir)
	require.NoError(t, err)
	require.NotEmpty(t, stor)

	// создаем тестовый зип файл и заполняем его
	utils.CreateTestFile(tmpDir, tmpFile)
	f, err := storage.NewZipFile(fmt.Sprintf("%s/%s", tmpDir, tmpFile))
	require.NoError(t, err)
	require.Equal(t, tmpFile, f.Name)

	// добавляем зип файл в хранилище
	stor.Add(f)
	require.Len(t, stor.Files, 1)

	// создаем кэш со стором
	c := NewCache(stor, 1000)

	// получаем часть файлов
	filesMap, err := c.GetFiles(tmpFile, []string{"readme.txt"})
	require.NoError(t, err)
	b, ok := filesMap["readme.txt"]
	require.True(t, ok)
	require.NotNil(t, b)

	utils.DeleteDir(tmpDir)
	utils.DeleteDir(dir)
}
