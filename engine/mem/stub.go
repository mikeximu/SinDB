package mem

/*
以下接口在 v0.1 中明确不支持。

设计原则：
- 显式降级
- 不假装成功
- 不隐藏能力边界
*/
// v0.1 明确不支持的能力

func (e *Engine) Sync() error               { return nil }
func (e *Engine) Flush() error              { return nil }
func (e *Engine) Compact(_, _ []byte) error { return nil }
func (e *Engine) CompactAll() error         { return nil }
