package sleepsrv

import (
	"context"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/uzuna/go-grpc-chain/pb/sleep"
)

type SleepService struct{}

func (s *SleepService) Sleep(ctx context.Context, req *sleep.SleepRequest) (*sleep.SleepResponce, error) {
	dur, err := ptypes.Duration(req.Duration)
	if err != nil {
		err = errors.WithStack(err)
		return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
	}
	select {
	case <-ctx.Done():
		errMsg := ctx.Err().Error()
		if strings.Contains(errMsg, "deadline") {
			return nil, grpc.Errorf(codes.DeadlineExceeded, errMsg)
		} else if strings.Contains(errMsg, "cancel") {
			return nil, grpc.Errorf(codes.Canceled, errMsg)
		}
		return nil, nil
	case <-time.After(dur):
		res := &sleep.SleepResponce{
			Duration:   req.Duration,
			DummyBytes: req.DummyBytes,
			Data:       []byte{},
		}
		return res, nil
	}
}
