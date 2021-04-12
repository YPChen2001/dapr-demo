package main

import (
	"context"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	"time"
)

func main() {
	// create a Dapr service server
	fmt.Println("client start...")

	dc, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	// 获取一个状态
	item, err := dc.GetState(ctx, "statestore", "statestore")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(item.Value))
	// start the server
	for {
		content := &dapr.DataContent{
			ContentType: "text/plain",
			Data:        []byte("hellow"),
		}
		resp, err := dc.InvokeMethodWithContent(ctx, "svc1", "echo", "post", content)
		if err != nil {
			panic(err)
		}
		fmt.Printf("service method invoked, response: %s\n", string(resp))
		time.Sleep(time.Duration(2) * time.Second)
	}
}
