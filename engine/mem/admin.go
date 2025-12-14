package mem

import "github.com/mikeximu/SinDB/db"

/*
Stats 返回数据库运行统计信息。

v0.1 说明：
- 仅提供 key 数量
- 不保证精确性
*/
func (e *Engine) Stats() db.DBStats {
	return db.DBStats{
		KeyCount: e.keyCount.Load(),
	}
}

/*
Size 返回数据库占用大小。

v0.1：
- 返回 key 数量作为近似值
*/
func (e *Engine) Size() int64 {
	return int64(e.keyCount.Load())
}

func (e *Engine) Properties() map[string]string {
	return map[string]string{
		"engine":  "mem",
		"version": "0.1",
	}
}
