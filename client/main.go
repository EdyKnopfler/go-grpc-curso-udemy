package main

import (
	"context"
	"fmt"
	"time"

	pb "com.derso/curso_creuto/grpc/myprotos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddr := "localhost:8080" // Para fins didáticos isto serve!

	// Não curti este INSECURE!
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println(err)
		panic("Não conectou!")
	}

	defer conn.Close()
	c := pb.NewSignVerifyClient(conn)

	// Olha um idiomismo comum com contextos em Go! :)
	// A função disto é definir um tempo limite para a execução e enviar um cancel()
	// para a goroutine que realiza a requisição.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	signRequest := pb.SignRequest{Text: "Pague o aluguel!"}
	signResponse, err := c.Sign(ctx, &signRequest)

	if err != nil {
		panic(err)
	}

	fmt.Println("A assinatura é: ", signResponse.Signature)
}
