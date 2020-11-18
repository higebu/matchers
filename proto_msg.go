package matchers

import (
	"fmt"

	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/proto"
)

// protoMsgMatcher implements the gomock.Matcher interface
type protoMsgMatcher struct {
	msg proto.Message
}

func (p *protoMsgMatcher) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, p.msg)
}

func (p *protoMsgMatcher) String() string {
	return fmt.Sprintf("is equal to %v", p.msg)
}

func ProtoMsgEq(x proto.Message) gomock.Matcher { return &protoMsgMatcher{x} }
