package interceptor

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strconv"
	"time"
)

func ClientAuthInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	startTime := time.Now()

	userInfo := make(map[string]string)
	userid := ctx.Value("userid")
	if userid != nil {
		i, ok := userid.(int64)
		if ok {
			userInfo["userid"] = strconv.FormatInt(i, 10)
		}
	}

	role := ctx.Value("role")
	if role != nil {
		userInfo["role"] = role.(string)
	}

	menu := ctx.Value("menu")

	if menu != nil {
		menuByte, _ := jsoniter.Marshal(menu)
		userInfo["menu"] = string(menuByte)
	}

	data := ctx.Value("data")
	if data != nil {
		dataByte, _ := jsoniter.Marshal(data)
		userInfo["data"] = string(dataByte)
	}

	md := metadata.New(userInfo)
	ctx = metadata.NewOutgoingContext(ctx, md)

	// 拦截前
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}

	fmt.Printf("调用 %s 方法 耗时: %v\n", method, time.Since(startTime))
	return nil
}
