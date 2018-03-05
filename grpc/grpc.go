package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Sharykhin/gl-mail-grpc"
	"github.com/Sharykhin/gl-mail-manager/entity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var client api.FailMailClient

func init() {
	cert := os.Getenv("GRPC_PUBLIC_KEY")
	if cert == "" {
		log.Fatal("Env variable GRPC_PUBLIC_KEY is not specified")
	}

	cred, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatalf("Could not load tls cert: %s", err)
	}

	address := os.Getenv("GRPC_SERVER_ADDRESS")
	if address == "" {
		log.Fatal("Env variable GRPC_SERVER_ADDRESS is not specified")
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatalf("Could not connet to a grpc server: %v", err)
	}
	client = api.NewFailMailClient(conn)
}

// CreateFailMail creates a new row of failed mail
func CreateFailMail(mm entity.MailMessage, reason string) (*api.FailMailResponse, error) {
	ctx := context.Background()
	p, err := json.Marshal(mm.Payload)
	if err != nil {
		return nil, fmt.Errorf("could not marshal: %v", err)
	}

	fmr := api.FailMailRequest{
		Action:  mm.Action,
		Payload: p,
		Reason:  reason,
	}
	return client.CreateFailMail(ctx, &fmr)
}
