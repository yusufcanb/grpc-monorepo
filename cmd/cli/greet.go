package cli

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/yusufcanb/grpc-monorepo/pkg/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

// greetCmd represents the apply command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greet the GRPC server",
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := protocol.NewGreeterClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		nameToGreet, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("name should be provided")
		}
		r, err := c.SayHello(ctx, &protocol.HelloRequest{Name: nameToGreet})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		log.Println("Greeting: %s", r.GetMessage())
	},
}

func init() {
	greetCmd.Flags().StringP("name", "n", "", "Name to be greeted.")
	rootCmd.AddCommand(greetCmd)
}
