package sleeproute_test

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
	"github.com/uzuna/go-grpc-chain/internal/sleeproute"
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

func TestClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	// back server start
	backAdrs := "localhost:8000"
	err := createServer(ctx, backAdrs, &sleepsrv.SleepService{})
	checkError(t, errors.WithStack(err))

	// Cleate Router
	connEnd, err := createConn(backAdrs)
	checkError(t, errors.WithStack(err))
	sleepRouter := sleeproute.NewSleepRouteService(connEnd)
	rtAdrs := "localhost:8001"
	err = createServer(ctx, rtAdrs, sleepRouter)
	checkError(t, errors.WithStack(err))

	// Create test
	connFront, err := createConn(rtAdrs)
	checkError(t, errors.WithStack(err))
	client := sleep.NewSleepServiceClient(connFront)
	wg := new(sync.WaitGroup)
	for i := 0; i < 100; i++ {
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

func createConn(adrs string) (*grpc.ClientConn, error) {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	return grpc.Dial(adrs, opts...)
}

func createServer(ctx context.Context, adrs string, srv sleep.SleepServiceServer) error {
	l, err := net.Listen("tcp", adrs)
	if err != nil {
		return err
	}
	grpcs := grpc.NewServer()
	sleep.RegisterSleepServiceServer(grpcs, srv)

	wg := new(sync.WaitGroup)

	// start server
	wg.Add(1)
	go func() {
		wg.Done()
		log.Printf("[DEBUG] Start Server %s", adrs)
		err := grpcs.Serve(l)
		if err != nil {
			log.Printf("[ERROR] Stop Server %s", err.Error())
		}
	}()
	// Close by Context
	go func() {
		<-ctx.Done()
		grpcs.GracefulStop()
		l.Close()
	}()
	wg.Wait()
	return nil
}
