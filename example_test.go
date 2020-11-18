package matchers_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"

	"github.com/higebu/matchers"
	"github.com/higebu/matchers/internal/example"
	"github.com/higebu/matchers/internal/testpb"
)

func ExampleCmpEq() {
	t := &testing.T{} // provided by test
	ctrl := gomock.NewController(t)
	m := example.NewMockFoo(ctrl)
	m.EXPECT().Bar(matchers.CmpEq([]*example.A{{Name: "blah"}})).Return(nil)
	m.Bar([]example.A{{Name: "blah"}})
}

func ExampleProtoMsgEq() {
	t := &testing.T{} // provided by test
	ctrl := gomock.NewController(t)
	m := example.NewMockFoo(ctrl)
	m.EXPECT().Bar(matchers.ProtoMsgEq(&testpb.Message{Name: "blah"})).Return(nil)
	m.Baz(&testpb.Message{Name: "blah"})
}

func ExampleTimestampEq() {
	t := &testing.T{} // provided by test
	ctrl := gomock.NewController(t)
	m := example.NewMockFoo(ctrl)
	now := ptypes.TimestampNow()
	m.EXPECT().Qux(matchers.TimestampEq(now)).Return(nil)
	m.Qux(now)
}
