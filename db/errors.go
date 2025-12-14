package db

import "errors"

/*
v0.1 错误设计原则：

1. 错误是“状态的事实”，不是“实现的细节”
2. 错误数量保持最小但完备
3. 不为未来能力提前定义错误
*/

// =============================
// Core KV Errors
// =============================

var (
	// ErrNotFound 表示 key 不存在
	ErrNotFound = errors.New("kv: not found")

	// ErrInvalidKey 表示 key 非法（nil / 空等）
	ErrInvalidKey = errors.New("kv: invalid key")

	// ErrInvalidValue 表示 value 非法
	ErrInvalidValue = errors.New("kv: invalid value")
)

// =============================
// Database State Errors
// =============================

var (
	// ErrClosed 表示数据库已关闭
	ErrClosed = errors.New("kv: database closed")

	// ErrNotSupported 表示当前引擎或版本不支持该操作
	ErrNotSupported = errors.New("kv: operation not supported")
)
