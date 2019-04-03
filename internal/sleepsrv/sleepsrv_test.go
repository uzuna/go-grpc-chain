package sleepsrv_test

import (
	"context"
	"log"
	"net"
	"os"
	"sync"
	"testing"
	"time"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes"
	"github.com/hashicorp/logutils"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/uzuna/go-grpc-chain/internal/sleepsrv"
	"github.com/uzuna/go-grpc-chain/pb/sleep"
)

func TestMain(m *testing.M) {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("DEBUG"),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	code := m.Run()
	os.Exit(code)
}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Logf("%+v", err)
		t.FailNow()
	}
}

func TestSleep(t *testing.T) {
	srv := &sleepsrv.SleepService{}

	wg := new(sync.WaitGroup)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer func() {
				cancel()
				wg.Done()
			}()
			dur := ptypes.DurationProto(time.Millisecond * 50)
			req := &sleep.SleepRequest{
				Duration:   dur,
				DummyBytes: int32(i),
			}
			res, err := srv.Sleep(ctx, req)
			checkError(t, errors.WithStack(err))
			assert.Equal(t, req.DummyBytes, res.DummyBytes)
		}(i)
	}
	wg.Wait()
}

func TestClient(t *testing.T) {
	adrs := "localhost:8000"
	// server start
	l, err := net.Listen("tcp", adrs)
	checkError(t, errors.WithStack(err))
	grpcs := grpc.NewServer()
	srv := &sleepsrv.SleepService{}
	sleep.RegisterSleepServiceServer(grpcs, srv)

	wg := new(sync.WaitGroup)

	// start server
	wg.Add(1)
	go func() {
		wg.Done()
		t.Logf("Start Server")
		err := grpcs.Serve(l)
		checkError(t, errors.WithStack(err))
	}()
	wg.Wait()

	// Client
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(adrs, opts...)

	client := sleep.NewSleepServiceClient(conn)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer func() {
				cancel()
				wg.Done()
			}()
			dur := ptypes.DurationProto(time.Millisecond * 1000)
			req := &sleep.SleepRequest{
				Duration:   dur,
				DummyBytes: int32(i),
			}
			res, err := client.Sleep(ctx, req)
			checkError(t, errors.WithStack(err))
			assert.Equal(t, req.DummyBytes, res.DummyBytes)
		}(i)
	}
	wg.Wait()
}
