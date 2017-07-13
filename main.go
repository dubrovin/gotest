package main

import (
	"flag"
	"fmt"
	"github.com/dubrovin/gotest/appconf"
	"github.com/dubrovin/gotest/cache"
	"github.com/dubrovin/gotest/server"
	"github.com/dubrovin/gotest/storage"
	"github.com/dubrovin/gotest/utils"
	"log"
	"time"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
	dir  = flag.String("dir", "storage", "storage file path")
	ttl  = flag.String("ttl", "10s", "time for counting")
)

func main() {
	flag.Parse()
	tmpFile := "zip.zip"

	defaultTTL, err := time.ParseDuration(*ttl)
	stor, err := storage.NewStorage(*dir)

	utils.CreateTestFile(*dir, tmpFile)
	f, err := storage.NewZipFile(fmt.Sprintf("%s/%s", *dir, tmpFile))
	f.Read()
	stor.Add(f)
	c := cache.NewCache(stor, defaultTTL)
	go c.Checker()
	api := server.NewAPI(c)
	api.RegisterHandlers()
	serv, err := server.NewServer(
		&appconf.Config{ListenAddress: *addr, RootDirectory: *dir, DefaultTTL: defaultTTL},
		api,
	)
	if err != nil {
		log.Fatal(err)
	}
	serv.Run()

	for {
		time.Sleep(time.Second * 10)
		log.Println("Heartbeat")
	}
}
