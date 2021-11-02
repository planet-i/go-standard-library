package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func main() {
	// 检验携带数据ctx
	ProcessEnter(NewContextWithTraceID())
	// 检验自带超时ctx
	HttpHandler()
	// 检验取消控制
	HandlerCancel()
}

/**
检验携带数据ctx
**/
// 通过trace_id串联所有的日志
const (
	KEY = "trace_id"
)

// 随机生成uuid
func NewRequestID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

// 基于context.Background创建一个携带teace_id的ctx
func NewContextWithTraceID() context.Context {
	ctx := context.WithValue(context.Background(), KEY, NewRequestID())
	return ctx
}

// 传递ctx
func ProcessEnter(ctx context.Context) {
	PrintLog(ctx, "Context包学习")
}

// 定义输出日志格式
func PrintLog(ctx context.Context, message string) {
	fmt.Printf("%s|info|trace_id=%s|%s", time.Now().Format("2006-01-02 15:04:05"), GetContextValue(ctx, KEY), message)
}

// 获取ctx的trace_id
func GetContextValue(ctx context.Context, k string) string {
	// interface{}强制转换成string
	v, ok := ctx.Value(k).(string)
	if !ok {
		return ""
	}
	return v
}

/**
检验自带超时ctx
**/
// 基于context.Background创建一个3s后超时的ctx
func NewContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Second)
}

// 传递ctx
func HttpHandler() {
	ctx, cancel := NewContextWithTimeout()
	defer cancel() //调用ctx取消子ctx的方法cancel。(可以在最上层调用，也可以随着ctx传递下去调用)
	deal(ctx)
}

// 监听超时退出
func deal(ctx context.Context) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Printf("deal time is %d\n", i)
			// cancel() //可以没有达到超时时间自己终止接下来的执行
		}
	}
}

/**
检验取消控制
**/
func HandlerCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go Speak(ctx)
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func Speak(ctx context.Context) {
	// 每隔1s循环一次
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			fmt.Println("我要闭嘴了")
			return
		default:
			fmt.Println("balabalabalabala")
		}
	}
}
