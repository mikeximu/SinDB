package db

/*
DB 是 v0.1 版本对外暴露的数据库协议。

设计原则：
- 只定义“已经成立”的能力
- 不引入未来语义
- 不泄露实现细节
*/

// =============================
// Options（v0.1 为空壳，占位但不承诺语义）
// =============================

// ReadOptions 控制一次读取行为。
// v0.1 中保留类型，但不定义任何语义。
type ReadOptions struct{}

// WriteOptions 控制一次写入行为。
// v0.1 中保留类型，但不定义任何语义。
type WriteOptions struct{}

// =============================
// Stats
// =============================

// DBStats 表示数据库的运行统计信息。
type DBStats struct {
	KeyCount uint64 // 当前 key 数量（近似）
}

// =============================
// DB Interface (v0.1)
// =============================

type DB interface {
	// -------- Core KV --------

	// Get 根据 key 获取 value。
	// 若 key 不存在，返回 ErrNotFound。
	Get(key []byte, opts *ReadOptions) ([]byte, error)

	// Put 写入或覆盖一个 key/value。
	Put(key, value []byte, opts *WriteOptions) error

	// Delete 删除指定 key。
	// 若 key 不存在，视为成功。
	Delete(key []byte, opts *WriteOptions) error

	// Has 判断 key 是否存在。
	Has(key []byte, opts *ReadOptions) (bool, error)

	// -------- Admin --------

	// Stats 返回数据库统计信息。
	Stats() DBStats

	// Size 返回数据库规模的近似值。
	Size() int64

	// -------- Lifecycle --------

	// Close 关闭数据库。
	Close() error

	// IsClosed 判断数据库是否已关闭。
	IsClosed() bool

	// Ping 用于健康检查。
	Ping() error
}
