package mtoi

import (
	"testing"
	"time"

	"github.com/WindomZ/testify/assert"
)

var kv *KV

const demo = "abcdefghijklmnopqrstuvwxyz"

func TestKV_NewKV(t *testing.T) {
	kv = NewKV(20)
}

func TestKV_Put(t *testing.T) {
	for i := 0; i < 10; i++ {
		s := string(demo[i])
		kv.Put(s, s)
	}
}

func TestKV_Get(t *testing.T) {
	time.Sleep(time.Millisecond * 100)

	for i := 0; i < 10; i++ {
		k := string(demo[i])
		v, ok := kv.Get(k)
		if assert.True(t, ok) {
			s, ok := v.(string)
			assert.True(t, ok)
			assert.NotEmpty(t, s)
			assert.Equal(t, 1, len(s))
		}
	}
}

func TestKV_Contain(t *testing.T) {
	assert.True(t, kv.Contain("a"))
	assert.False(t, kv.Contain("z"))
}

func TestKV_MulPut(t *testing.T) {
	f, stop := kv.MulPut()
	for i := 10; i < 20; i++ {
		s := string(demo[i])
		f(s, s)
	}
	stop()
}

func TestKV_Get2(t *testing.T) {
	time.Sleep(time.Millisecond * 100)

	for i := 10; i < 20; i++ {
		k := string(demo[i])
		v, ok := kv.Get(k)
		if assert.True(t, ok) {
			s, ok := v.(string)
			assert.True(t, ok)
			assert.NotEmpty(t, s)
			assert.Equal(t, 1, len(s))
		}
	}
}

func TestKV_Close(t *testing.T) {
	kv.Close()
}
