package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	// Exit code 0 - Check is passing
	// Exit code 1 - Check is warning
	// Any other code - Check is failing
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "grpc",
			Usage: "gRPC 健康检查",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "address", Usage: "127.0.0.1:2345"},
				cli.StringFlag{Name: "service", Usage: "gRPC service name"},
				cli.BoolFlag{Name: "secure", Usage: "default value is false"},
			},
			Action: func(c *cli.Context) error {
				addr := c.String("address")
				service := c.String("service")
				secure := c.Bool("secure")

				var opts []grpc.DialOption
				if !secure {
					opts = append(opts, grpc.WithInsecure())
				}

				conn, err := grpc.Dial(addr, opts...)
				if err != nil {
					log.Println("连接 gRPC 服务失败", addr, err)
					return cli.NewExitError(err.Error(), 2)
				}
				defer conn.Close()

				hc := healthpb.NewHealthClient(conn)
				resp, err := hc.Check(context.Background(), &healthpb.HealthCheckRequest{Service: service})
				if err != nil {
					log.Println("gRPC 服务健康检查失败", addr, err)
					return cli.NewExitError(err.Error(), 2)
				}

				if resp.Status == healthpb.HealthCheckResponse_UNKNOWN {
					log.Println("gRPC 服务 UNKNOWN")
				} else if resp.Status == healthpb.HealthCheckResponse_NOT_SERVING {
					log.Println("gRPC 服务 NOT_SERVING")
				} else {
					// SERVING
					log.Println("检测成功")
					return nil
				}
				return cli.NewExitError(err.Error(), 2)
			},
		},
	}
	app.Run(os.Args)
}
