package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"io"
	"log"

	pb "github.com/limmiehoang/grpcdemo/student"
)

var (
	serverAddr  = flag.String("s", "localhost:10000", "The server address in the format of host:port")
	function = flag.String("func", "", "Choose a function to execute: find, create, or get")
	id = flag.String("id", "", "Enter student id")
	name = flag.String("name", "", "Enter student name")
	keyword = flag.String("keyword", "", "Get students with this keyword")
)

func createStudent(client pb.StudentClient, createRequest *pb.CreateRequest)  {
	resp, err := client.CreateStudent(context.Background(), createRequest)
	if err != nil {
		log.Fatalf("Could not create student: %v", err)
	}
	if !resp.Success {
		log.Printf("Student id %s already exists.", createRequest.Student.Id)
		return
	}
	log.Printf("A new student has been added with id: %s", resp.Id)
}

func findStudent(client pb.StudentClient, findRequest *pb.FindRequest)  {
	resp, err := client.FindStudent(context.Background(), findRequest)
	if err != nil {
		log.Fatalf("Error on find student: %v", err)
	}
	if resp.Student == nil {
		log.Printf("Could not find student with id: %s", findRequest.Id)
		return
	}
	log.Printf("Found student: %v", resp.Student)
}

func getStudents(client pb.StudentClient, filter *pb.GetFilter) {
	// calling the streaming API
	stream, err := client.GetStudents(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get students: %v", err)
	}
	for {
		customer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.Students(_) = _, %v", client, err)
		}
		log.Printf("Customer: %v", customer)
	}
}

func main()  {
	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewStudentClient(conn)

	if *function == "" {
		log.Fatalf("Please choose a function (find, create, or get) with -func flag.")
	}

	switch *function {
	case "find":
		if *id == "" {
			log.Fatalf("Please enter student id with -id flag to find a student")

		}
		findRequest := &pb.FindRequest{Id: *id}
		findStudent(client, findRequest)
	case "create":
		if *id == "" || *name == "" {
			log.Fatalf("Please give both student id and student name with -id and -name flag to create a student")

		}
		createRequest := &pb.CreateRequest{
			Student: &pb.Model{
				Id:   *id,
				Name: *name,
			},
		}
		createStudent(client, createRequest)
	case "get":
		getFilter := &pb.GetFilter{Keyword: *keyword}
		getStudents(client, getFilter)
	default:
		log.Fatalf("Sorry, function %s doesn't exist.", *function)
	}
}