package main

import (
	"golang.org/x/net/context"
	"encoding/json"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"strings"

	pb "github.com/limmiehoang/grpc-demo/student"
)

var (
	port = flag.Int("p", 10000, "The server port")
	jsonDBFile = flag.String("db", "data/student_db.json", "A json file containing a list of features")
)

type server struct {
	savedStudents map[string]*pb.Model
}

func (s *server) FindStudent(ctx context.Context, in *pb.FindRequest) (*pb.FindResponse, error) {
	student, exists := s.savedStudents[in.Id]
	if !exists {
		return &pb.FindResponse{Student: nil}, nil
	}
	return &pb.FindResponse{Student: student}, nil
}

func (s *server) CreateStudent(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	if _, exists := s.savedStudents[in.Student.Id]; exists {
		return &pb.CreateResponse{
			Id: "",
			Success: false,
		}, nil
	}
	s.savedStudents[in.Student.Id] = in.Student
	return &pb.CreateResponse{
		Id: in.Student.Id,
		Success: true,
	}, nil
}

func (s *server) GetStudents(filter *pb.GetFilter, stream pb.Student_GetStudentsServer) error {
	for _, student := range s.savedStudents {
		if filter.Keyword != "" {
			if !strings.Contains(student.Name, filter.Keyword) {
				continue
			}
		}
		if err := stream.Send(student); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) loadStudents(filePath string)  {
	var data []byte
	var err error
	data, err = ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to load default students: %v", err)
	}
	var studentArr []*pb.Model
	if err := json.Unmarshal(data, &studentArr); err != nil {
		log.Fatalf("Failed to load default students: %v", err)
	}
	for _, student := range studentArr {
		s.savedStudents[student.Id] = student
	}
}

func newServer() *server {
	s := &server{savedStudents: make(map[string]*pb.Model)}
	s.loadStudents(*jsonDBFile)
	return s
}

func main()  {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterStudentServer(grpcServer, newServer())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}