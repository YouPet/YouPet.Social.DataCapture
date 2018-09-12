package main

import (
	"context"
	"fmt"

	proto "bwkernel/youpet/youpet.social.datacapture/proto"
	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("Social.DataCapture.Client"),
	)

	service.Init()

	userService := proto.NewUserService("Social.DataCapture", service.Client())

	rsp, err := userService.GetUser(context.TODO(), &proto.GetUserRequest{UserID: 1})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.NickName)

}
