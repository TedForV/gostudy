package stringsvc

import (
	"context"
	"errors"
	"github.com/go-kit/kit/log"
	"strings"
	"time"
)

type StringService interface {
	Uppsercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}

type stringService struct {
}

type stringService2 struct {
	logger log.Logger
	next   StringService
}

func (mw stringService2) Uppsercase(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "uppsercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin))
	}(time.Now())

	output, err = mw.next.Uppsercase(ctx, s)
	return

}

func (mw stringService2) Count(ctx context.Context, s string) (n int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())
	n = mw.next.Count(ctx, s)
	return
}

func (stringService) Uppsercase(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(_ context.Context, s string) int {
	return len(s)
}

//ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")
