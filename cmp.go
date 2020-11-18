package matchers

import (
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

type cmpEqMatcher struct {
	x interface{}
}

func (e cmpEqMatcher) Matches(x interface{}) bool {
	return cmp.Equal(e.x, x)
}

func (e cmpEqMatcher) String() string {
	return fmt.Sprintf("is equal to %v", e.x)
}

func CmpEq(x interface{}) gomock.Matcher { return &cmpEqMatcher{x} }
