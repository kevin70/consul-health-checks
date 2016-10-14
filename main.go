package main

import (
	"fmt"
	"github.com/urfave/cli"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "grpc",
			Usage: "grpc 健康检查",
			Action: func(c *cli.Context) error {
				fmt.Println("=============================")
				addr := c.Args().First()
				fmt.Println(addr)

				conn, err := grpc.Dial(addr)
				if err != nil {
					fmt.Println("EEEEEEEEEEEEEE")
					fmt.Errorf("连接grpc服务[%v]失败\n", addr)
					return err
				}
				defer conn.Close()

				fmt.Println("11111..............................")
				hc := healthpb.NewHealthClient(conn)
				fmt.Println("..............................")

				resp, err := hc.Check(context.Background(), &healthpb.HealthCheckRequest{})
				if err != nil {
					fmt.Errorf("grpc 服务[%v]健康检查失败\n", addr)
					return err
				}

				if resp.Status == healthpb.HealthCheckResponse_UNKNOWN {
					fmt.Println("grpc 服务 UNKNOWN")
				} else if resp.Status == healthpb.HealthCheckResponse_NOT_SERVING {
					fmt.Println("grpc 服务 NOT_SERVING")
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
