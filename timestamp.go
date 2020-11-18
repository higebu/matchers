package matchers

import (
	"fmt"
	"reflect"

	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type timestampMatcher struct {
	x *timestamppb.Timestamp
}

func (t *timestampMatcher) Matches(x interface{}) bool {
	if reflect.TypeOf(x) != reflect.TypeOf((*timestamppb.Timestamp)(nil)) {
		return false
	}
	a := x.(*timestamppb.Timestamp)
	if t.x.Seconds != a.Seconds {
		return false
	}
	return t.x.Nanos == a.Nanos
}

func (t *timestampMatcher) String() string {
	return fmt.Sprintf("is equal to %v", t.x)
}

func TimestampEq(x *timestamppb.Timestamp) gomock.Matcher { return &timestampMatcher{x} }
