package main

import (
	"context"
	"fmt"
	"time"

	pb "com.derso/curso_creuto/grpc/myprotos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	serverAddr := "localhost:8080" // Para fins didáticos isto serve!

	// Não curti este INSECURE!
	//conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// Autenticando:
	// https://grpc.io/docs/guides/auth/
	// https://linuxize.com/post/creating-a-self-signed-ssl-certificate/
	// https://jfrog.com/knowledge-base/general-what-should-i-do-if-i-get-an-x509-certificate-relies-on-legacy-common-name-field-error/
	creds, _ := credentials.NewClientTLSFromFile("./mycert.crt", "")
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(creds))

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
