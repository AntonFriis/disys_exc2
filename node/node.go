package main

import (
	"bufio"
	"context"
	pb "disys_exc2/p2p"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
)

type noden struct {
	Addr string
	Con  string
	Id   int64
	Ids  string
}

var node *noden

type server struct {
	pb.UnimplementedGreeterServer
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Arguments required: <id> as int 0-4")
		os.Exit(1)
	}
	Id, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		log.Panic(err)
	}

	//Setup the file for log outputs
	LogFile := "./systemlogs/node.log"
	// open log file
	logFile, err := os.OpenFile(LogFile, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			log.Fatalf("File not found: %v", err)
		}
	}(logFile)

	log.SetOutput(logFile)
	node = &noden{
		Id:  Id,
		Ids: strconv.Itoa(int(Id)),
	}
	node.Addr, node.Con = Port(node.Id)

	go node.serverStart()
	time.Sleep(15 * time.Second)
	if node.Id == 0 {
		node.clintStart()
	}
	for {

	}

}

func (s *server) SendMessage(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	log.Printf("*************************************************************************")
	log.Printf("Node " + in.GetName() + " has sent message to Node " + node.Ids)
	go node.clintStart()
	return &pb.Reply{Message: "Node " + node.Ids + " confirmes to node " + in.GetName()}, nil
}

func (node *noden) serverStart() {
	lis, err := net.Listen("tcp", fmt.Sprintf(node.Addr))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (node *noden) clintStart() {
	log.Printf("Node %v has enterede the critical zone", node.Id)
	time.Sleep(time.Second)
	// Set up a connection to the server.
	conn, err := grpc.Dial(node.Con, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	log.Printf("Node %v has left the critical zone", node.Id)
	_, err = c.SendMessage(ctx, &pb.Request{Name: node.Ids})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

}

func Port(NodeId int64) (string, string) {
	file, err := os.Open("ports.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var Port0, Port1 string //Fix cursed line
	for scanner.Scan() {
		IdPort := strings.Split(scanner.Text(), " ")
		Id, _ := strconv.ParseInt(IdPort[0], 10, 64)
		if Id == NodeId {
			Port0 = IdPort[1]
			Port1 = IdPort[2]
		}
	}
	return Port0, Port1

}
