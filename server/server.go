package server

import (
	"grpc-deadline-exceeded/pb"
	"log"
	"time"

	google_protobuf "github.com/golang/protobuf/ptypes/empty"
)

const (
	noOfMessages = 20
)

// DebugService represents grpc service
type DebugService struct{}

// FetchMessages will stream messages back to the client
func (*DebugService) FetchMessages(e *google_protobuf.Empty, stream pb.Debug_FetchMessagesServer) error {
	defer trackTime(time.Now())
	// generate a few messages to stream to the client
	for msg := range genMessages(noOfMessages, 1*time.Second) {
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
	log.Printf("all messages sent to the client. total = %d\n", noOfMessages)
	return nil
}

func genMessages(noOfMessages int, delay time.Duration) <-chan *pb.Message {
	ch := make(chan *pb.Message, noOfMessages)
	go func() {
		defer close(ch)
		for i := 1; i <= noOfMessages; i++ {
			// time.Sleep(delay)
			ch <- &pb.Message{Id: int32(i)}
		}
	}()
	return ch
}

func trackTime(start time.Time) {
	log.Printf("time taken = %s", time.Since(start))
}
