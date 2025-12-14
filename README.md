# SinDB v0.1

SinDB 是一个从零开始构建的数据库项目，目标是演进为一套 **可插拔存储引擎的商业级数据库系统**。

当前版本 **v0.1** 是项目的第一个稳定里程碑，专注于构建一个：

> **接口清晰、行为确定、可长期演进的内存型 Key-Value 数据库内核**

---

## v0.1 版本定位

- 类型：In-Memory KV Database
- 存储模型：Key → Value
- 运行方式：单进程、交互式 CLI
- 设计目标：**正确性优先、接口先行、能力最小化**

v0.1 不是 Demo，也不是一次性代码，而是整个数据库系统的 **基线版本（baseline）**。

---

## 已支持功能

- 内存型 Key-Value 存储（基于 `sync.Map`）
- 并发安全的基础 KV 操作
    - `Get`
    - `Put`
    - `Delete`
    - `Has`
- 数据库生命周期管理
    - `Open`
    - `Close`
    - `Ping`
- 基础运行统计
    - Key 数量统计（近似）
- 交互式命令行（CLI）
- 支持系统信号的安全退出
    - `SIGINT`（Ctrl+C）
    - `SIGTERM`

---

## 明确不支持的功能（v0.1 设计如此）

以下能力在 v0.1 **明确不提供**，以保证语义清晰：

- ❌ 持久化（WAL / SST）
- ❌ 事务（Transaction）
- ❌ 快照（Snapshot）
- ❌ 迭代器（Iterator）
- ❌ 压缩 / Compaction
- ❌ SQL / 表 / Schema

这些能力将在后续版本中 **逐步、可验证地引入**，而不是一次性堆叠。

---

## 项目结构

```text
SinDB/
├── db/              # 数据库协议定义（对外稳定接口）
├── engine/mem/      # v0.1 内存存储引擎实现
├── cmd/sindb/       # CLI 程序入口
└── README.md
```

快速开始
1. 初始化 Go Module
```text
go mod init github.com/mikeximu/SinDB
go mod tidy
```
2. 启动 CLI
```text
go run ./cmd/sindb
```
启动成功后：
```text
SinDB v0.1 (mem)
>

```

| 命令    | 说明           | 示例               |
| ----- | ------------ | ---------------- |
| put   | 写入 key/value | `put name alice` |
| get   | 获取 value     | `get name`       |
| del   | 删除 key       | `del name`       |
| has   | 判断 key 是否存在  | `has name`       |
| stats | 查看数据库统计      | `stats`          |
| exit  | 退出程序         | `exit`           |

### 使用示例
```text
> put a 1
> get a
1
> has a
true
> del a
> has a
false
> stats
{KeyCount:0}
> exit

```

### 数据库设计文档 v0.1
---
####  设计原则

* **第一性原理**：只实现当前成立的能力。
* **接口先行**：协议独立于具体引擎。
* **显式降级**：不支持的能力明确拒绝。
* **工程严谨**：所有生命周期路径可控。
* **可演进性**：不为未来提前背技术债。

#### 版本定位

* **当前版本**：v0.1.0

* **类型**：内存型 KV 数据库

* **定位**：长期演进的数据库内核起点

后续版本将在 不破坏 v0.1 接口语义 的前提下，引入：

* **WAL** 与持久化

* **Batch** 写入

* **Snapshot / MVCC**

* **多存储引擎实现**

* **SQL / 关系模型（上层）**