package storage

import "testing"

type testCase struct {
	name  string
	key   string
	value string
}

func TestSetGet(t *testing.T) {

	cases := []testCase{
		{"hello world", "hello", "world"},
	}

	s, err := NewStorage()
	if err != nil {
		t.Error("new storage: %v", err)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s.Set(c.key, c.value)

			sValue := s.Get(c.key)

			if *sValue != c.value {
				t.Error("values not equal")
			}
		})
	}
}
