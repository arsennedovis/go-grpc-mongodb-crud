package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/arsennedovis/go-grpc-microservice/car/carpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Car Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v\n", err)
	}
	defer cc.Close()

	c := carpb.NewCarServiceClient(cc)

	//create Car
	fmt.Println("Create new Car")
	car := carpb.Car{
		GovId:   "AA 1111 BB",
		OwnerId: "Arsen Nedovis",
		Brand:   "BMW",
	}
	createCarRes, err := c.CreateCar(context.Background(), &carpb.CreateCarRequest{Car: &car})
	if err != nil {
		log.Fatalf("Unexpected error %v\n", err)
	}
	fmt.Printf("Car was been created: %v\n", createCarRes)
	carID := createCarRes.GetCar().GetId()

	// read Car
	fmt.Println("Read Car info")

	_, err2 := c.ReadCar(context.Background(), &carpb.ReadCarRequest{CarId: "61907bf24a5c74c5dc44af40"})
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v\n", err2)
	}

	readCarReq := &carpb.ReadCarRequest{CarId: carID}
	readCarRes, readCarErr := c.ReadCar(context.Background(), readCarReq)
	if readCarErr != nil {
		fmt.Printf("Error happened while reading: %v\n", err2)
	}
	fmt.Printf("Car was read: %v\n", readCarRes)

	// update Car
	fmt.Println("Update Car")
	newCar := &carpb.Car{
		Id:      carID,
		GovId:   "DC 2222 BM",
		OwnerId: "Mira Nedovis",
		Brand:   "BMW e46",
	}
	updateRes, updateErr := c.UpdateCar(context.Background(), &carpb.UpdateCarRequest{Car: newCar})

	if updateErr != nil {
		fmt.Printf("Error happened while updating: %v\n", updateErr)
	}
	fmt.Printf("Car was read: %v\n", updateRes)

	//Delete Car
	deleteRes, deleteErr := c.DeleteCar(context.Background(), &carpb.DeleteCarRequest{CarId: carID})
	if deleteErr != nil {
		fmt.Printf("Error happened while deleting: %v\n", deleteErr)
	}
	fmt.Printf("Car was deleted: %v\n", deleteRes)

	// list cars

	fmt.Println("Updating a car")
	stream, err := c.ListCar(context.Background(), &carpb.ListCarRequest{})
	if err != nil {
		log.Fatalf("Error while calling ListCar streaming RPC: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something went wrong while iterating through the stream: %v", err)
		}

		fmt.Println(res.GetCar())
	}
}
