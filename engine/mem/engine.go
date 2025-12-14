package mem

import (
	"github.com/mikeximu/SinDB/db"
	"sync"
	"sync/atomic"
)

/*
Engine 是 v0.1 版本的内存型 KV 引擎实现。

设计目标：
- 提供最小可用的 KV 能力
- 接口语义正确、行为明确
- 不引入任何未来假设（WAL / MVCC / LSM）

存储模型：
- 使用 sync.Map 实现并发安全的 key-value 映射
- key 使用 string 形式（v0.1 不定义 comparator 语义）
- value 始终做 defensive copy，避免外部修改
*/
type Engine struct {
	// data 存储所有 KV 对
	// key: string
	// value: []byte（拷贝后存储）
	data sync.Map

	// closed 标识数据库是否已关闭
	closed atomic.Bool

	// keyCount 用于粗略统计 key 数量
	keyCount atomic.Uint64
}

// Open 创建一个新的内存数据库实例。
func Open() *Engine {
	return &Engine{}
}

// ensureOpen 用于所有对外方法的前置检查。
func (e *Engine) ensureOpen() error {
	if e.closed.Load() {
		return db.ErrClosed
	}
	return nil
}
