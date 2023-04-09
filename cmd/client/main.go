package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	hellopb "github.com/inahym196/xo_connect5/pkg/grpc"
)

var (
	scanner *bufio.Scanner
	client  hellopb.GameServiceClient
)

func main() {
	fmt.Println("start gPRC Client")

	scanner = bufio.NewScanner(os.Stdin)
	address := "localhost:9000"
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()
	client = hellopb.NewGameServiceClient(conn)
	for {
		fmt.Println("1: send Request")
		fmt.Println("2: exit")
		fmt.Print("please enter >")
		scanner.Scan()
		in := scanner.Text()
		switch in {
		case "1":
			Hello()
		case "2":
			fmt.Println("bye.")
			goto M
		}
	}
M:
}

func Hello() {
	fmt.Println("Please enter your message.")
	scanner.Scan()
	message := scanner.Text()
	req := &hellopb.HelloRequest{Body: message}
	res, err := client.Hello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetBody())
	}

}
