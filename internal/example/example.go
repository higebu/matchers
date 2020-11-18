//go:generate go run github.com/golang/mock/mockgen -package=example -source=$GOFILE -destination=mock_$GOFILE

package example

import (
	"github.com/higebu/matchers/internal/testpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type A struct {
	Name string
}

type Foo interface {
	Bar(x []A) error
	Baz(msg *testpb.Message) error
	Qux(t *timestamppb.Timestamp) error
}
