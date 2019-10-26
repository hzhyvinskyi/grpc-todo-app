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
	port     = ":8080"
	db       *mongo.Client
	tododb   *mongo.Collection
	mongoCtx context.Context
)

type TodoServiceServer struct{}

type TodoItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      string             `bson:"user_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
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

func (s *TodoServiceServer) ReadTodo(ctx context.Context, req *todopb.ReadTodoReq) (*todopb.ReadTodoRes, error) {
	// String convertation from proto to ObjectID
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Can't convert to ObjectID: %v", err),
		)
	}

	result := tododb.FindOne(ctx, bson.M{"_id": oid})

	// Create empty Todo item to write decoded result into
	data := TodoItem{}

	if err = result.Decode(&data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Can't find Todo with ObjectID %s: %v", req.GetId(), err),
		)
	}

	// Cast to ReadTodoRes type
	response := &todopb.ReadTodoRes{
		Todo: &todopb.Todo{
			Id:          oid.Hex(),
			UserId:      data.UserID,
			Title:       data.Title,
			Description: data.Description,
		},
	}

	return response, nil
}

func (s TodoServiceServer) UpdateTodo(ctx context.Context, req *todopb.UpdateTodoReq) (*todopb.UpdateTodoRes, error) {
	todo := req.GetTodo()
	oid, err := primitive.ObjectIDFromHex(todo.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Can't convert to ObjectID: %v", err),
		)
	}

	update := bson.M{
		"user_id":     todo.GetUserId(),
		"title":       todo.GetTitle(),
		"description": todo.GetDescription(),
	}

	// Convert oid into unordered bson document
	filter := bson.M{"_id": oid}

	// Result will be BSON encoded
	// Added options to return updated BSON document instead of original
	result := tododb.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	decodedTodo := TodoItem{}
	if err := result.Decode(&decodedTodo); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Can't find Todo with ObjectID %s", todo.GetId(), err),
		)
	}

	return &todopb.UpdateTodoRes{
		Todo: &todopb.Todo{
			Id:          decodedTodo.ID.Hex(),
			UserID:      decodedTodo.UserID,
			Title:       decodedTodo.Title,
			Description: decodedTodo.Description,
		},
	}, nil
}

func (s TodoServiceServer) DeleteTodo(ctx context.Context, req *todopb.DeleteTodoReq) (*todopb.DeleteTodoRes, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Can't convert to ObjectID: %v", err),
		)
	}

	if _, err = tododb.DeleteOne(ctx, bson.M{"_id": oid}); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Can't find/delete Todo with id %s: %v", req.GetId(), err),
		)
	}

	return &todopb.DeleteTodoRes{
		Success: true,
	}, nil
}

func (s TodoServiceServer) ListTodo(req *todopb.ListTodoReq, stream todopb.TodoService_ListTodoServer) error {
	data := &TodoItem{}

	cursor, err := tododb.Find(context.Background(), bson.M{})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	defer cursor.Close(context.Background())

	// Iterates through the Cursor. If there are no items, loop will break
	for cursor.Next(context.Background()) {
		err := cursor.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Unavailable,
				fmt.Sprint("Can't decode data: %v", err),
			)
		}

		// If there aren't any errors, send Todo via stream
		err = stream.Send(&todopb.ListTodoRes{
			Todo: &todopb.Todo{
				Id:          data.ID.Hex(),
				UserId:      data.UserID,
				Title:       data.Title,
				Description: data.Description,
			},
		})
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Stream Send Error: %v", err),
			)
		}

		if err = cursor.Err(); err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Unknown Cursor Error: %v", err),
			)
		}
	}

	return nil
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
