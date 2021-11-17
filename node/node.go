package main

import (
	"bufio"
	pb "disys_exc2/p2p"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type node struct {
	Client pb.P2PServiceClient
	ctx    context.Context
	conn   *grpc.ClientConn
	Addr   int64
	Id     int64
	IsLast bool
}

type Server struct {
	pb.UnimplementedP2PServiceServer
	this node
}

func NewNode(Id int64) *node {
	node := node{
		Addr:   0000,
		Id:     0,
		IsLast: true,
	}
	return &node
}

// Port TODO: Fix cursed code
func Port(NodeId int64) string {
	file, err := os.Open("ports.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var Port0 string //Fix cursed line
	for scanner.Scan() {
		IdPort := strings.Split(scanner.Text(), " ")
		Id, _ := strconv.ParseInt(IdPort[0], 10, 64)
		if Id == 0 {
			Port0 = IdPort[1]
		} //Fix cursed line
		if Id == NodeId {
			return IdPort[1]
		}
	}
	return Port0 //Fix cursed line
}

func (node *node) Connect(ctx context.Context, send *pb.Send) (*pb.Response, error) {
	var err error
	ConPort := "localhost:" + Port(send.Port)
	node.conn, err = grpc.Dial(ConPort, grpc.WithInsecure())
	if err != nil {
		log.Printf("Connection error when trying to join on port: " + ConPort)
	}

	node.Client = pb.NewP2PServiceClient(node.conn)

	return &pb.Response{Message: "Node " + strconv.FormatInt(node.Id, 10) + " has connected on port: " + ConPort}, err
}

func (node *node) Disconnect(ctx context.Context, send *pb.Send) (*pb.Response, error) {
	err := node.conn.Close()
	return &pb.Response{Message: "Node " + strconv.FormatInt(node.Id, 10) + " has disconnected!"}, err
}

func (node *node) Reconnect(ctx context.Context, send *pb.Send) (*pb.Response, error) {
	_, err := node.Disconnect(ctx, send)
	if err != nil {
		return nil, err
	}
	_, err = node.Connect(ctx, send)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Message: "Node " + strconv.FormatInt(node.Id, 10) + " has reconnected on port: " + strconv.FormatInt(node.Addr, 10)}, err
}

func (node *node) listen(server Server) {
	log.Printf("Node %v: listening on port: %v\n", node.Id, node.Addr)
	lis, err := net.Listen("tcp", "localhost:"+strconv.FormatInt(node.Addr, 10))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	n := grpc.NewServer() // n is for serving purpose

	pb.RegisterP2PServiceServer(n, server)
	// Register reflection service on gRPC Server.
	reflection.Register(n)

	// start listening
	if err := n.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	//main TODO: remove args requirement and automatically increase id by 1.
	// pass the port as an argument and also the port of the other node
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Arguments required: \"go run node/node.go <id>\" as int 0-4")
		os.Exit(1)
	}

	// args in order
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

	//Create node
	node := NewNode(Id)

	//Create server
	server := Server{this: *node}

	//Reconnect the old nodes

	// Set up a connection to first node.
	msg, err := node.Connect(node.ctx, &pb.Send{Port: node.Addr + 1})
	if err != nil {
		return
	}
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Printf(msg.Message)
	defer node.conn.Close()

	node.Client = pb.NewP2PServiceClient(node.conn)

	var cancel context.CancelFunc
	node.ctx, cancel = context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	go node.listen(server)

	for {
		time.Sleep(60 * time.Second)
	}
}
