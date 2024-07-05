package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	/*
		go func() {
			time.Sleep(time.Second * 5)
			cancel()
		}()
	*/
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Tempo excedido para bookear o quarto.")
	case <-time.After(time.Second * 5):
		fmt.Println("Quarto reservado com sucesso.")
	}
}
