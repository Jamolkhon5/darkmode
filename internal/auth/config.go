package auth

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net/http"

	"github.com/Jamolkhon5/darkmode/pkg/proto/auth_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

var gClient auth_v1.AuthV1Client

func InitClient(conn *grpc.ClientConn) {
	gClient = auth_v1.NewAuthV1Client(conn)
}

func VerifyToken(r *http.Request) (string, error) {
	// Получаем токен из заголовка
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		return "", fmt.Errorf("no authorization header")
	}

	// Создаем метаданные для gRPC запроса
	md := metadata.New(map[string]string{
		"Authorization": authToken,
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// Делаем запрос к auth сервису
	userInfo, err := gClient.GetUser(ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("Authentication failed: %v", err)
		return "", err
	}

	// Возвращаем ID пользователя
	return userInfo.GetUser().GetId(), nil
}
