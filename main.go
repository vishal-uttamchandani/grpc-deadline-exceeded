package main

import (
	"context"
	"flag"
	"grpc-deadline-exceeded/pb"
	"grpc-deadline-exceeded/server"
	"io"
	"log"
	"net"
	"os"
	"time"

	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var (
	serverMode *bool
)

const (
	timeOutDuration = time.Duration(5 * time.Second)
	workDuration    = time.Duration(1 * time.Second)
)

func init() {
	serverMode = flag.Bool("server", false, "start in server mode ?")
	flag.Parse()
}

func main() {
	if *serverMode {
		runServer()
	}
	runClient()
}

func runServer() {
	service := &server.DebugService{}
	server := grpc.NewServer()

	pb.RegisterDebugServer(server, service)

	l, err := net.Listen("tcp", ":9191")
	handleErr(err)

	log.Println("server is listening at port 9191")
	log.Fatal(server.Serve(l))
}

func runClient() {
	conn, err := grpc.Dial(":9191", grpc.WithInsecure())
	handleErr(err)

	client := pb.NewDebugClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), timeOutDuration)
	defer cancel()

	stream, err := client.FetchMessages(ctx, &google_protobuf.Empty{})

	handleErr(err)

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		// Simulate some work delay
		time.Sleep(workDuration)

		log.Println(msg.Id)
	}
}

func getEnv(env, fallback string) string {
	input := os.Getenv(env)
	if input == "" {
		return fallback
	}
	return input
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
