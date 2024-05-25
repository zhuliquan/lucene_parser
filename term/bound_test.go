package term

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoundType(t *testing.T) {
	type test struct {
		name  string
		input *Bound
		bType BoundType
	}
	for _, tt := range []test{
		{
			name:  "test_unknown",
			input: nil,
			bType: UNKNOWN_BOUND_TYPE,
		},
		{
			name:  "test_left_exclude_right_exclude",
			input: &Bound{LeftInclude: false, RightInclude: false},
			bType: LEFT_EXCLUDE_RIGHT_EXCLUDE,
		},
		{
			name:  "test_left_exclude_right_include",
			input: &Bound{LeftInclude: false, RightInclude: true},
			bType: LEFT_EXCLUDE_RIGHT_INCLUDE,
		},
		{
			name:  "test_left_include_right_exclude",
			input: &Bound{LeftInclude: true, RightInclude: false},
			bType: LEFT_INCLUDE_RIGHT_EXCLUDE,
		},
		{
			name:  "test_left_include_right_include",
			input: &Bound{LeftInclude: true, RightInclude: true},
			bType: LEFT_INCLUDE_RIGHT_INCLUDE,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.bType, tt.input.GetBoundType())
		})
	}
}

func TestRangeValue(t *testing.T) {
	type test struct {
		name  string
		input *RangeValue
		value interface{}
		wantS string
	}

	for _, tt := range []test{
		{
			name:  "test_inf",
			input: &RangeValue{InfinityVal: "*"},
			value: "*",
			wantS: "*",
		},
		{
			name:  "test_phrase",
			input: &RangeValue{PhraseValue: []string{"1", "+", "1"}},
			value: "1+1",
			wantS: `"1+1"`,
		},
		{
			name:  "test_single",
			input: &RangeValue{SingleValue: []string{"a", "+", "b"}},
			value: "a+b",
			wantS: `a+b`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantS, tt.input.String())
			s, _ := tt.input.Value(func(s string) (interface{}, error) { return s, nil })
			assert.Equal(t, tt.value, s)
		})
	}

	var s *RangeValue
	assert.Empty(t, s.String())
	x, _ := s.Value(func(s string) (interface{}, error) { return s, nil })
	assert.Equal(t, nil, x)
	s = &RangeValue{}
	assert.Empty(t, s.String())
	x, _ = s.Value(func(s string) (interface{}, error) { return s, nil })
	assert.Equal(t, "", x)
}
