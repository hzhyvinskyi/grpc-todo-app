package main

import (
	"context"
	"fmt"
	todopb "github.com/hzhyvinskyi/grpc-todo-app/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
)

var (
	port = ":8080"
	db *mongo.Client
	tododb *mongo.Collection
	mongoCtx context.Context
)

type TodoServiceServer struct {}

type TodoItem struct {
	ID			primitive.ObjectID	`bson:"_id,omitempty"`
	UserID		string				`bson:"user_id"`
	Title		string				`bson:"title"`
	Description	string				`bson:"description"`
}

func (s *TodoServiceServer) CreateTodo(ctx context.Context, req *todopb.CreateTodoReq) (*todopb.CreateTodoRes, error) {
	// Extract Todo message from request
	todo := req.GetTodo()

	// Convert  it to the TodoItem to convert it into BSON
	data := TodoItem{
		UserID:      todo.GetUserId(),
		Title:       todo.GetTitle(),
		Description: todo.GetDescription(),
	}

	result, err := tododb.InsertOne(mongoCtx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid := result.InsertedID.(primitive.ObjectID)

	todo.ID = oid.Hex()

	return &todopb.CreateTodoRes{Todo: todo}, nil
}

func (s TodoServiceServer) ReadTodo(context.Context, *todopb.ReadTodoReq) (*todopb.ReadTodoRes, error) {
	panic("implement me")
}

func (s TodoServiceServer) UpdateTodo(context.Context, *todopb.UpdateTodoReq) (*todopb.UpdateTodoRes, error) {
	panic("implement me")
}

func (s TodoServiceServer) DeleteTodo(context.Context, *todopb.DeleteTodoReq) (*todopb.DeleteTodoRes, error) {
	panic("implement me")
}

func (s TodoServiceServer) ListTodo(*todopb.ListTodoReq, todopb.TodoService_ListTodoServer) error {
	panic("implement me")
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Printf("Server is listening on port %s\n", port)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Unable to listen port %s:\n%v\n", port, err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	srv := &TodoServiceServer{}
	todopb.RegisterTodoServiceServer(s, srv)

	fmt.Println("Connection to MongoDB...")
	mongoCtx = context.Background()
	db, err = mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Couldn't connect to MongoDB: %v\n", err)
	}

	fmt.Println("Connected to MongoDB")

	tododb = db.Database("tdappdb").Collection("todos")

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v\n", err)
		}
	}()
	fmt.Printf("Server successfully started on port %s\n", port)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Println("Stopping the server...")
	s.Stop()
	if err = listener.Close(); err != nil {
		log.Fatalf("Can't close Listener: %v\n", err)
	}
	fmt.Println("Closing MongoDB connection...")

	if err = db.Disconnect(mongoCtx); err != nil {
		log.Fatalf("Can't disconnect from MongoDB: %v\n", err)
	}
	fmt.Println("Done")
}
