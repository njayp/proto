package client

import (
	"context"
	"log"
	"time"

	"github.com/njayp/proto/proto/generated"
	"google.golang.org/grpc"
)

func SendBytes(blob []byte) {
	CHUNKSIZE := 1024 * 64
	addr := "localhost:3001"

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	client := generated.NewServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.ChunkStream(ctx)
	if err != nil {
		log.Fatalf("fail to : %v", err)
	}

	msg := &generated.ChunkMessage{}

	for i := 0; i < len(blob); i += CHUNKSIZE {
		j := i + CHUNKSIZE
		println("sending chunk:", j/CHUNKSIZE)

		if j < len(blob) {
			msg.Chunk = blob[i:j]
		} else {
			msg.Chunk = blob[i:]
		}
		if err = stream.Send(msg); err != nil {
			log.Fatalf("send err %v", err)
		}
	}
	ak, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("fail : %v", err)
	}

	println("recieved from server:", ak.GetAk())
}
