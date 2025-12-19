package main

import (
	"instargam/protogen/userpb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer conn.Close()

	userClient := userpb.NewUserServiceClient(conn)

	r := gin.Default()

	r.GET("/users", func(ctx *gin.Context) {
		res, err := userClient.GetUsers(ctx, &userpb.GetUsersRequest{})
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": "error lah pkoknyaasa",
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"users": res.Users,
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run HTTP server: %v", err)
	}
}
