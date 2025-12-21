package contexto

import (
	"context"
	"fmt"
	"time"
)

type keyCustom string

var key keyCustom = "key"

func Contexto() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, key, "value")
	viewContext(ctx)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	go myProcess(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("-> ", ctx.Err())

	}

}

func viewContext(ctx context.Context) {
	fmt.Printf("my value is %s\n", ctx.Value(key))
}

func myProcess(ctx context.Context) {
	//espera que el contexto expire para terminar
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("Doing some process")
		}
		//controlamos la velocidad del loop
		time.Sleep(500 * time.Millisecond)
	}
}
