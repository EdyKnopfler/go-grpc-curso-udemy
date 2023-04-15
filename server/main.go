package main

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net"

	pb "com.derso/curso_creuto/grpc/myprotos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	// Incorporação de structs
	pb.UnimplementedSignVerifyServer
}

var (
	privateKey *rsa.PrivateKey
)

func GetPrivateKey() (*rsa.PrivateKey, error) {
	key, err := ioutil.ReadFile("./private_key.pem")

	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(key)
	der, err := x509.ParsePKCS8PrivateKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	return der.(*rsa.PrivateKey), err
}

// Função que implementa a interface do servidor
func (s *server) Sign(ctx context.Context, input *pb.SignRequest) (*pb.SignResponse, error) {
	hashed := sha256.Sum256([]byte(input.Text))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])

	if err != nil {
		panic(err)
	}

	output := hex.EncodeToString(signature)
	return &pb.SignResponse{Signature: output}, nil
}

func main() {
	// privateKey já declarada logo acima, então declaramos err antecipadamente também
	// e não usamos :=
	var err error
	privateKey, err = GetPrivateKey()

	if err != nil {
		panic(err)
	}

	// Escuta em socket bruto
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Printf("Erro ao subir servidor: %v\n", err)
		panic(err)
	}

	// Autenticando:
	// https://grpc.io/docs/guides/auth/
	// https://linuxize.com/post/creating-a-self-signed-ssl-certificate/
	// https://jfrog.com/knowledge-base/general-what-should-i-do-if-i-get-an-x509-certificate-relies-on-legacy-common-name-field-error/
	creds, _ := credentials.NewServerTLSFromFile("./mycert.crt", "./mycert.key")
	s := grpc.NewServer(grpc.Creds(creds))

	// Usando as interfaces geradas pelo "protoc"
	pb.RegisterSignVerifyServer(s, &server{})
	fmt.Printf("Servidor ouvindo em %v\n", lis.Addr())

	// Eu gosto deste idiomismo mas nem sempre concordo com ele...
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Falha ao iniciar gRPC: %v\n", err)
	}
}
