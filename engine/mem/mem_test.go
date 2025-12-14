package mem

import (
	"github.com/mikeximu/SinDB/db"
	"strconv"
	"sync"
	"testing"
)

func TestMemEngineBasic(t *testing.T) {
	e := Open()
	defer e.Close()

	key := []byte("key1")
	val := []byte("value1")
	// Put 测试
	if err := e.Put(key, val, nil); err != nil {
		t.Fatalf("Put failed: %v", err)
	}
	// Get 测试
	got, err := e.Get(key, nil)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if string(got) != string(val) {
		t.Errorf("expected %s, got %s", val, got)
	}

	// Has 测试
	exists, err := e.Has(key, nil)
	if err != nil {
		t.Fatalf("Has failed: %v", err)
	}
	if !exists {
		t.Errorf("expected key to exist")
	}

	// Delete 测试
	if err := e.Delete(key, nil); err != nil {
		t.Errorf("Delete failed: %v", err)
	}
	exists, _ = e.Has(key, nil)
	if exists {
		t.Errorf("expected key to be deleted")
	}

	// Stats & Size 测试
	if s := e.Stats(); s.KeyCount != 0 {
		t.Errorf("expected KeyCount=0, got %d", s.KeyCount)
	}
	if sz := e.Size(); sz != 0 {
		t.Errorf("expected Size=0, got %d", sz)
	}

	// Ping 测试
	if err := e.Ping(); err != nil {
		t.Errorf("Ping failed: %v", err)
	}

	// Close & IsClosed 测试
	if err := e.Close(); err != nil {
		t.Errorf("Close failed: %v", err)
	}
	if !e.IsClosed() {
		t.Errorf("expected DB to be closed")
	}

	// ensureOpen 关闭后报错测试
	if err := e.Put([]byte("k2"), []byte("v2"), nil); err != db.ErrClosed {
		t.Errorf("expected ErrClosed, got %v", err)
	}
	if _, err := e.Get([]byte("k2"), nil); err != db.ErrClosed {
		t.Errorf("expected ErrClosed, got %v", err)
	}
	if _, err := e.Has([]byte("k2"), nil); err != db.ErrClosed {
		t.Errorf("expected ErrClosed, got %v", err)
	}
	if err := e.Delete([]byte("k2"), nil); err != db.ErrClosed {
		t.Errorf("expected ErrClosed, got %v", err)
	}
}

func TestMemEngineConcurrency(t *testing.T) {
	e := Open()
	defer e.Close()

	wg := sync.WaitGroup{}
	numKeys := 1000

	// 并发写入
	for i := 0; i < numKeys; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := []byte("key" + strconv.Itoa(i))
			val := []byte("val" + strconv.Itoa(i))
			if err := e.Put(key, val, nil); err != nil {
				t.Errorf("Put failed: %v", err)
			}
		}(i)
	}
	wg.Wait()

	// 检查 KeyCount
	if s := e.Stats(); s.KeyCount != uint64(numKeys) {
		t.Errorf("expected KeyCount=%d, got %d", numKeys, s.KeyCount)
	}
	if sz := e.Size(); sz != int64(numKeys) {
		t.Errorf("expected Size=%d, got %d", numKeys, sz)
	}
}
