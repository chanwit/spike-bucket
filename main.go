package main

import (
	"context"
	"log"
	"net"
	"net/http/httptest"
	"os"
	"os/signal"
	"syscall"

	"github.com/johannesboyne/gofakes3"
	"github.com/johannesboyne/gofakes3/backend/s3mem"
)

const version = "2"

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM, os.Kill)

	defer cancel()
	logger := log.New(os.Stdout, "", 0)
	// backend, err := s3afero.MultiBucket(afero.NewMemMapFs())
	// if err != nil {
	//	log.Fatal(err)
	// }
	backend := s3mem.New()
	s3 := gofakes3.New(backend,
		gofakes3.WithAutoBucket(true),
		gofakes3.WithLogger(gofakes3.StdLog(logger,
			gofakes3.LogErr, gofakes3.LogWarn, gofakes3.LogInfo, gofakes3.LogDebug)))

	// create a listener with the desired port.
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	ts := httptest.NewUnstartedServer(s3.Server())
	if err := ts.Listener.Close(); err != nil {
		return
	}
	ts.Listener = listener

	// Start the server.
	ts.Start()
	defer ts.Close()

	logger.Println(ts.URL)

	<-ctx.Done()
}
