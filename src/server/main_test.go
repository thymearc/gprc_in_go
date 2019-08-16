package tests
import (
	"context"
	"log"
	"os"
	"time"
	"testing"
	"google.golang.org/grpc"
	pb "../server/CalcService"
)

const (
	address     = "localhost:50051"
	defaultName = "calc"
)

func TestService(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)
	a := "2"
	b := "2"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CalcAdd(ctx, &pb.Request{A: a, B: b})
	if err != nil {
		log.Fatalf("something went wrong: %v", err)
		t.Error()
	}
	log.Printf("calc result: %s", r.Result)
	os.Stdout.WriteString("Result")
	os.Stderr.WriteString(r.Result)
	if r.Result != "4" { 
		t.Error() }
}
