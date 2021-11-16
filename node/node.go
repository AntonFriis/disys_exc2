package main

import (
	"bufio"
	pb "disys_exc2/p2p"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type node struct {
	client pb.P2PServiceClient
	server pb.UnimplementedP2PServiceServer
	Addr string
	Id int64
	IsFirst bool
}

func (node node) mustEmbedUnimplementedP2PServiceServer() {
	panic("implement me")
}

func (node node) Disconnect(ctx context.Context, send *pb.Send) (*pb.Response, error) {
	return &pb.Response{Message: "Node "+ strconv.FormatInt(node.Id, 10) + " has disconnected!"}, nil
}

var conn *grpc.ClientConn
var ctx context.Context
var LastId int64

func NewNode(Id int64) *node {
	node := node{
		client: nil,
		Addr: Port(Id),
		Id: Id,
		IsFirst: true,
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
		if Id == 0 {Port0 = IdPort[1]} //Fix cursed line
		if Id == NodeId {return IdPort[1]}
	}
	return Port0 //Fix cursed line
}

func (node *node) Connect(ctx context.Context, send *pb.Send) (*pb.Response, error) {
	var err error
	ConPort := Port(send.Port)
	conn, err = grpc.Dial(ConPort, grpc.WithInsecure())
	return &pb.Response{Message: "Node "+ strconv.FormatInt(node.Id, 10) + " has connected on port: " + ConPort}, err
}

func (node *node) listen() {
	log.Printf("Node %v: listening on port: %v\n", node.Id, node.Addr)
	lis, err := net.Listen("tcp", node.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer() // n is for serving purpose

	pb.RegisterP2PServiceServer(server, node.server)
	// Register reflection service on gRPC server.
	reflection.Register(server)

	// start listening
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	// pass the port as an argument and also the port of the other node
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Arguments required: <id> as int 0-4")
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

	LastId := pb.LatestId{Id: LastId}

	//Create node
	node := NewNode(Id)

	//Increment LastId by 1
	LastId.Id = LastId.Id+1

	//var opts []grpc.DialOption
	// Set up a connection to the server.
	msg, err := node.Connect(ctx, &pb.Send{Port: 0})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Printf(msg.Message)
	defer conn.Close()

	node.client = pb.NewP2PServiceClient(conn)

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	go node.listen()

	for {
		time.Sleep(60 * time.Second)
	}
}

