//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE -self_package=github.com/higebu/matchers/$GOPACKAGE

package example

import (
	"github.com/higebu/matchers/testpb"
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
