package cache

import (
	"github.com/dubrovin/gotest/storage"
	"github.com/dubrovin/gotest/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewCache(t *testing.T) {
	dir := "dir"
	stor, err := storage.NewStorage(dir)
	require.NoError(t, err)
	c := NewCache(stor, 1000)
	t.Log(c)
	utils.DeleteDir(dir)
}
