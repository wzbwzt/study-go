package singleton_test

import (
	"testing"

	singleton "github.com/wzbwzt/studyGo/design_pattern/CreateType/singleton_pattern"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	assert.Equal(t, singleton.GetInstance(), singleton.GetInstance())
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetInstance() != singleton.GetInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
