package main

import (
	"context"
	"fmt"

	proto "bwkernel/youpet/youpet.social.datacapture/proto"
	"bwkernel/youpet/youpet.social.datacapture/src/spider"
)

type UserService struct{}

func (service *UserService) GetUser(ctx context.Context, req *proto.GetUserRequest, res *proto.GetUserResponse) error {
	res.NickName = "Frank"
	return nil
}

func main() {

	chongbabaSpider, err := spider.NewSpider("chongbaba")
	if err != nil {
		fmt.Println(err)
		return
	}

	blogs, err := chongbabaSpider.SpiderUrl("http://www.chongbaba.com/news/")

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range blogs {
		fmt.Println(v.Title)
	}

	// service := micro.NewService(
	// 	micro.Name("Social.DataCapture"),
	// )

	// service.Init()

	// proto.RegisterUserServiceHandler(service.Server(), new(UserService))

	// if err := service.Run(); err != nil {
	// 	fmt.Println(err)
	// }

}
