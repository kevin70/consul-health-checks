package main

import (
	"fmt"
	"os/exec"
	"testing"
	//"time"
)

func Test_HealthCheck(t *testing.T) {
	cmd := exec.Command("cmd", "/C", "consul-health-checks grpc --address localhost:50051 --service konan")
	err := cmd.Start()

	//errCh := make(chan error)
	//go func() {
	//	errCh <- cmd.Wait()
	//}()

	//err = <-errCh
	cmd.Wait()
	fmt.Println("finished", err)
}
