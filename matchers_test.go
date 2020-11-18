package matchers_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
	"github.com/higebu/matchers"
	"github.com/higebu/matchers/testpb"
)

func TestMatchers(t *testing.T) {
	n := time.Now().UTC()
	nint64 := int64(n.Unix()) + int64(n.Nanosecond()/1e6)
	now := time.Unix(nint64, 0)
	pnow, _ := ptypes.TimestampProto(now)
	after5nanos, _ := ptypes.TimestampProto(time.Now().Add(5 * time.Nanosecond))
	after5s, _ := ptypes.TimestampProto(time.Now().Add(5 * time.Second))
	type e interface{}
	tests := []struct {
		name    string
		matcher gomock.Matcher
		yes, no []e
	}{
		{"test CmpEq", matchers.CmpEq(4), []e{4}, []e{3, "blah", nil, int64(4)}},
		{"test ProtoMsgEq", matchers.ProtoMsgEq(&testpb.Message{Name: "blah"}), []e{&testpb.Message{Name: "blah"}}, []e{&testpb.Message{Name: "blahblah"}, "blah", nil, int64(4)}},
		{"test TimestampEq", matchers.TimestampEq(pnow), []e{pnow}, []e{after5s, after5nanos, "blah", nil, int64(4)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, x := range tt.yes {
				if !tt.matcher.Matches(x) {
					t.Errorf(`"%v %s": got false, want true.`, x, tt.matcher)
				}
			}
			for _, x := range tt.no {
				if tt.matcher.Matches(x) {
					t.Errorf(`"%v %s": got true, want false.`, x, tt.matcher)
				}
			}
		})
	}
}
