package term

import (
	"reflect"
	"testing"
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
			name:  "test_le_re",
			input: &Bound{LeftInclude: false, RightInclude: false},
			bType: LEFT_EXCLUDE_RIGHT_EXCLUDE,
		},
		{
			name:  "test_le_ri",
			input: &Bound{LeftInclude: false, RightInclude: true},
			bType: LEFT_EXCLUDE_RIGHT_INCLUDE,
		},
		{
			name:  "test_li_re",
			input: &Bound{LeftInclude: true, RightInclude: false},
			bType: LEFT_INCLUDE_RIGHT_EXCLUDE,
		},
		{
			name:  "test_li_ri",
			input: &Bound{LeftInclude: true, RightInclude: true},
			bType: LEFT_INCLUDE_RIGHT_INCLUDE,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if tt.input.GetBoundType() != tt.bType {
				t.Errorf("expect got: %v, but got: %v", tt.bType, tt.input.GetBoundType())
			}
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
			if tt.input.String() != tt.wantS {
				t.Errorf("expect %s, but %s", tt.wantS, tt.input.String())
			}
			if s, _ := tt.input.Value(func(s string) (interface{}, error) { return s, nil }); !reflect.DeepEqual(s, tt.value) {
				t.Errorf("expect %v, but %v", tt.value, s)
			}
		})
	}

	var s *RangeValue
	if s.String() != "" {
		t.Errorf("expect empty")
	}
	if x, _ := s.Value(func(s string) (interface{}, error) { return s, nil }); !reflect.DeepEqual(x, nil) {
		t.Errorf("expect empty")
	}
	s = &RangeValue{}
	if s.String() != "" {
		t.Errorf("expect empty")
	}
	if x, _ := s.Value(func(s string) (interface{}, error) { return s, nil }); !reflect.DeepEqual(x, "") {
		t.Errorf("expect empty")
	}
}
