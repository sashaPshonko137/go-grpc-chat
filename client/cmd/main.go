package main

import (
	desc "client/pkg/chat_v1" // Импортируйте ваш сгенерированный protobuf пакет
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        fmt.Printf("Ошибка подключения: %v", err)
        return
    }
    defer conn.Close()

	c := desc.NewChatV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	r, err := c.WriteMessage(ctx, &desc.WriteRequest{From: "сеня", Message: "УМНЫЙ СЕНЯ"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r)
}
