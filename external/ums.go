package external

import (
	"context"
	"fmt"
	"kautsarhasby/ewallet-transaction/external/proto/tokenvalidation"
	"kautsarhasby/ewallet-transaction/internal/models"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type External struct {

}

func (*External) ValidateToken(ctx context.Context, token string) (models.TokenData,error) {
	var (
		resp models.TokenData
	)
	conn, err := grpc.Dial("localhost:7000",grpc.WithInsecure())
	if err != nil {
		return resp, errors.Wrap(err, "failed to dial ums grpc")
	}
	defer conn.Close()
	
	client := tokenvalidation.NewTokenValidationClient(conn)

	req := tokenvalidation.TokenRequest{
		Token: token,
	}
	response, err := client.ValidateToken(ctx, &req)
	if err != nil {
		return resp, errors.Wrap(err, "failed to validate token")
	}

	if response.Message != "Success" {
		return resp, errors.New("got error from ums" + response.Message)
	}

	resp.UserID = response.Data.UserId
	resp.Username = response.Data.Username
	resp.Fullname = response.Data.FullName

	fmt.Println("RSPONSE =>>>>", resp)

	return resp, nil
}