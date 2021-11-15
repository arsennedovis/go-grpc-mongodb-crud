package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/arsennedovis/go-grpc-microservice/car/carpb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var collection *mongo.Collection

type server struct {
}

type carItem struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	OwnerID string             `bson:"Owner_id"`
	GovID   string             `bson:"Gov_id"`
	Brand   string             `bson:"Brand"`
}

//List Of Cars
func (*server) ListCar(req *carpb.ListCarRequest, stream carpb.CarService_ListCarServer) error {
	fmt.Println("List car request")
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		data := &carItem{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Error while decoding: %v", err))
		}
		stream.Send(&carpb.ListCarResponce{Car: dataToCarPb(data)})
	}
	if err := cur.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Ukknown internal error: %v", err))
	}
	return nil
}

//Delete Car
func (*server) DeleteCar(ctx context.Context, req *carpb.DeleteCarRequest) (*carpb.DeleteCarResponce, error) {
	fmt.Println("Delete car request")

	oid, err := primitive.ObjectIDFromHex(req.GetCarId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintln("Cannot parse ID"))
	}

	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot delete object in Mongo: %v", err))
	}
	if res.DeletedCount == 0 {
		if err != nil {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find car in Mongo: %v", err))
		}
	}
	return &carpb.DeleteCarResponce{CarId: req.GetCarId()}, nil
}

// Update Car
func (*server) UpdateCar(ctx context.Context, req *carpb.UpdateCarRequest) (*carpb.UpdateCarResponce, error) {
	fmt.Println("Update car request")

	car := req.GetCar()
	oid, err := primitive.ObjectIDFromHex(car.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintln("Cannot parse ID"))
	}
	data := &carItem{}
	res := collection.FindOne(context.Background(), bson.M{"_id": oid})
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find car with specificed ID: %v", err))
	}

	//we update data
	data.OwnerID = car.GetOwnerId()
	data.GovID = car.GetGovId()
	data.Brand = car.GetBrand()

	_, updateErr := collection.ReplaceOne(context.Background(), bson.M{"_id": oid}, data)
	if updateErr != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot update object in Mongo: %v", updateErr))
	}
	return &carpb.UpdateCarResponce{
		Car: dataToCarPb(data),
	}, nil
}

// fn data to block
func dataToCarPb(data *carItem) *carpb.Car {
	return &carpb.Car{
		Id:      data.ID.Hex(),
		GovId:   data.GovID,
		OwnerId: data.OwnerID,
		Brand:   data.Brand,
	}
}

// Read Car Info
func (*server) ReadCar(ctx context.Context, req *carpb.ReadCarRequest) (*carpb.ReadCarResponce, error) {
	fmt.Println("Read car request")
	carID := req.GetCarId()

	oid, err := primitive.ObjectIDFromHex(carID)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintln("Cannot parse ID"))
	}
	// Create empty struct
	data := &carItem{}
	res := collection.FindOne(context.Background(), bson.M{"_id": oid})
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find car with specificed ID: %v", err))
	}
	return &carpb.ReadCarResponce{
		Car: &carpb.Car{
			Id:      data.ID.Hex(),
			GovId:   data.GovID,
			OwnerId: data.OwnerID,
			Brand:   data.Brand,
		},
	}, nil
}

//CreateCar
func (*server) CreateCar(ctx context.Context, req *carpb.CreateCarRequest) (*carpb.CreateCarResponce, error) {
	fmt.Println("Create car request")
	car := req.GetCar()

	data := carItem{
		OwnerID: car.GetOwnerId(),
		GovID:   car.GetGovId(),
		Brand:   car.GetBrand(),
	}

	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %v", err))
	}
	objectID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot convert to OID %v", err))
	}
	return &carpb.CreateCarResponce{
		Car: &carpb.Car{
			Id:      objectID.Hex(),
			OwnerId: car.GetOwnerId(),
			GovId:   car.GetGovId(),
			Brand:   car.GetBrand(),
		},
	}, nil
}

func main() {
	// if crush the go code, we get the file name and line number. FEATERE
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Connection to mongoDB")

	//MongoDB Connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	fmt.Println("Car Service Started")
	collection = client.Database("mydb").Collection("cars")

	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalln("Disconnected fron server MongoDB")
		}
	}()

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	carpb.RegisterCarServiceServer(s, &server{})
	//Register reflection service on gRPC
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait Ctrl-C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Stoping the listener")
	lis.Close()
	fmt.Println("Closing MongoDB Connection")
	client.Disconnect(context.TODO())
	fmt.Println("End of program")
}
