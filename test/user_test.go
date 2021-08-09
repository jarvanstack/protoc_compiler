package test

import (
	"fmt"
	"protoc_compiler/proto"
	"testing"

	"google.golang.org/grpc"
)

func Test_1(t *testing.T) {
	u := proto.NewUserClient(&grpc.ClientConn{})
	fmt.Printf("u: %v\n", u)
}
