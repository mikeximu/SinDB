package mem

import "github.com/mikeximu/SinDB/db"

/*
Get 根据 key 获取 value。

语义说明：
- key 不存在返回 ErrNotFound
- 返回的 value 必须是独立拷贝
*/
func (e *Engine) Get(key []byte, _ *db.ReadOptions) ([]byte, error) {
	if err := e.ensureOpen(); err != nil {
		return nil, err
	}

	v, ok := e.data.Load(string(key))
	if !ok {
		return nil, db.ErrNotFound
	}

	val := v.([]byte)
	cp := make([]byte, len(val))
	copy(cp, val)
	return cp, nil
}

/*
Put 写入或覆盖一个 key/value。

v0.1 行为：
- 覆盖写
- 不支持 TTL / WAL / Sync
*/
func (e *Engine) Put(key, value []byte, _ *db.WriteOptions) error {
	if err := e.ensureOpen(); err != nil {
		return err
	}
	_, existed := e.data.Load(string(key))
	cp := make([]byte, len(value))
	copy(cp, value)
	e.data.Store(string(key), cp)
	if !existed {
		e.keyCount.Add(1)
	}
	return nil
}

/*
Delete 删除指定 key。

v0.1 行为：
- key 不存在视为成功
*/
func (e *Engine) Delete(key []byte, _ *db.WriteOptions) error {
	if err := e.ensureOpen(); err != nil {
		return err
	}

	if _, ok := e.data.LoadAndDelete(string(key)); ok {
		e.keyCount.Add(^uint64(0)) // 原子 -1
	}
	return nil
}

/*
Has 判断 key 是否存在。
*/
func (e *Engine) Has(key []byte, _ *db.ReadOptions) (bool, error) {
	if err := e.ensureOpen(); err != nil {
		return false, err
	}

	_, ok := e.data.Load(string(key))
	return ok, nil
}
