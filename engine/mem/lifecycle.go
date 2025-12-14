package mem

/*
Close 关闭数据库实例。

v0.1 行为：
- 仅设置 closed 状态
- 不释放资源（内存由 GC 管理）
*/
func (e *Engine) Close() error {
	e.closed.Store(true)
	return nil
}

func (e *Engine) IsClosed() bool {
	return e.closed.Load()
}

/*
Ping 用于健康检查。
*/
func (e *Engine) Ping() error {
	return e.ensureOpen()
}
