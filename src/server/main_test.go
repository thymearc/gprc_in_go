package tests
import (
	"context"
	"log"
	"time"
	"testing"
	"google.golang.org/grpc"
	pb "../server/CalcService"
)

const (
	address     = os.Getenv("APP_ADDR") + ":" + os.Getenv("APP_PORT")
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
	if r.Result != "result: 4" {
		t.Error() }
	a = "5"
	b = "2"
	r, err = c.CalcMultiply(ctx, &pb.Request{A: a, B: b})
	if err != nil {
		log.Fatalf("something went wrong: %v", err)
		t.Error()
	}
	if r.Result != "result: 10" {
		t.Error() }
}
