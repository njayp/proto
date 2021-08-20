package server

import (
	"io"
	"log"
	"net"

	"github.com/njayp/proto/proto/generated"
	"google.golang.org/grpc"
)

type implServer struct {
	generated.UnimplementedServiceServer
}

func (s implServer) ChunkStream(server generated.Service_ChunkStreamServer) error {
	data := make([]byte, 0)
	msgcnt := 0
	for {
		msgcnt += 1
		msg, err := server.Recv()
		data = append(data, msg.GetChunk()...)
		if err != nil {
			if err == io.EOF {
				println("got", len(data))
				break
			} else {
				log.Fatalf("msgcnt: %v, err: %v", msgcnt, err)
			}
		}
	}

	responce := &generated.AkMessage{}
	responce.Ak = "Acknowlegg"
	err := server.SendMsg(responce)
	if err != nil {
		panic(err)
	}

	println("Ackkers")
	return nil
}

func Start() {
	addr := "localhost:3001"
	conn, err := net.Listen("tcp", addr)

	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	generated.RegisterServiceServer(server, implServer{})
	println("Serving on", addr)
	println(server.Serve(conn))

}
