package contexto

import (
	"context"
	"fmt"
	"time"
)

type keyName string
type keyInt string

const (
	key_name   keyName = "api-key"
	my_key_int keyInt  = "my-key-int"
)

func Contexto() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, key_name, "super-secret-api-key")
	ctx = context.WithValue(ctx, my_key_int, 5)
	viewContext(ctx)

	ctx2, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	myProcess(ctx2)
	defer cancel()

}
func viewContext(ctx context.Context) {
	fmt.Printf("My value is %s\n", ctx.Value(key_name))
	fmt.Printf("My value is %d\n", ctx.Value(my_key_int))
}
func myProcess(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("mi mensaje de error")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("soy default")
		}
	}
}
