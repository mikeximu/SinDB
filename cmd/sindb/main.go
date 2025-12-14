package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/mikeximu/SinDB/engine/mem"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	db := mem.Open()
	defer db.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	in := bufio.NewScanner(os.Stdin)
	fmt.Println("SinDB v0.1 (mem)")

	for {
		select {
		case <-ctx.Done():
			// 收到系统终止信号，安全退出
			fmt.Println("shutdown signal received, exiting...")
			return
		default:
		}

		fmt.Print("> ")
		if !in.Scan() {
			// stdin 关闭或发生错误，直接退出
			return
		}

		args := strings.Fields(in.Text())
		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "put":
			db.Put([]byte(args[1]), []byte(args[2]), nil)
		case "get":
			v, err := db.Get([]byte(args[1]), nil)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(v))
			}
		case "del":
			db.Delete([]byte(args[1]), nil)
		case "has":
			ok, _ := db.Has([]byte(args[1]), nil)
			fmt.Println(ok)
		case "stats":
			fmt.Printf("%+v \n", db.Stats())
		case "exit":
			return
		}
	}
}
