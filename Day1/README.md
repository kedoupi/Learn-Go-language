# Go 教学教案 - Day 1：Go 初识与 Hello World

## 🎯 今日目标（目标导向）

完成 Go 编程环境搭建，理解包结构与标准库导入，并编写第一个 Go 程序。

---

## 🧠 核心概念讲解

### 1️⃣ Go 是什么语言？

- Go 是 Google 开发的编译型语言，主打：简单、并发、跨平台、高性能
- 常用于：CLI 工具、Web 后端、云原生服务、WASM 等

---

### 2️⃣ `package` 是什么？

```go
package main
```

- 每个 `.go` 文件都属于一个包（package）
- `main` 是特殊的包，用于可执行程序
- Go 会从 `main()` 函数开始执行

---

### 3️⃣ `import` 是怎么工作的？

```go
import "fmt"
```

- `fmt` 是 Go 的标准库之一，提供格式化输出能力（类似 JS 中 `console.log`）
- 所有导入的包必须使用，否则编译报错
- 可以一次导入多个包：

```go
import (
  "fmt"
  "os"
)
```

---

### 4️⃣ `main()` 函数是程序入口

```go
func main() {
    fmt.Println("Hello, Go!")
}
```

- 可执行程序必须包含一个 `main()` 函数

---

## 💻 示例代码讲解版

```go
package main        // 声明本文件属于 main 包（必须）

import "fmt"        // 导入标准库 fmt（格式化输出）

// 函数：程序入口点
func main() {
    fmt.Println("Hello, Go!") // 打印一行文字并换行
}
```

### 编译与运行

```bash
go run main.go     # 直接运行
go build -o hello  # 编译为可执行文件
./hello            # 运行程序
```

---

## 📝 今日练习：写一个自我介绍函数

### 需求：

- 接收 `name` 和 `age` 参数
- 使用 `fmt.Printf()` 输出格式化信息

### 示例代码：

```go
package main

import "fmt"

func introduce(name string, age int) {
    fmt.Printf("你好，我叫 %s，今年 %d 岁，是一名前端开发者。\n", name, age)
}

func main() {
    introduce("Alice", 28)
}
```

---

## 🚀 延伸挑战（可选）

1. 使用 `fmt.Sprintf` 返回字符串而不是直接打印：

```go
str := fmt.Sprintf("你好，我叫 %s，今年 %d 岁。", name, age)
```

2. 尝试从命令行读取参数（用 `os.Args`）

---

## 🔚 小结与思考

| 概念 | 说明 |
|------|------|
| `package main` | 可执行程序的入口标识 |
| `import "fmt"` | 引入格式化输出库 |
| `func main()` | 程序启动函数 |
| `fmt.Println` / `Printf` | 打印输出到终端 |

---
