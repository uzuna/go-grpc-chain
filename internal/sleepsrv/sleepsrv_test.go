package sleepsrv_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"github.com/uzuna/go-grpc-chain/internal/sleepsrv"
	"github.com/uzuna/go-grpc-chain/pb/sleep"
)

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
			checkError(t, err)
			assert.Equal(t, req.DummyBytes, res.DummyBytes)
		}(i)
	}
	wg.Wait()
}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Logf("%+v", err)
		t.FailNow()
	}
}
