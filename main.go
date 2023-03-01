package main

import (
	"helloworld/handler"
	pb "helloworld/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("helloworld"),
		service.Version("latest"),
	)

	//go func() {
	//	for i := 0; i < 10; i++ {
	//		err := events.Publish("test", fmt.Sprintf("hello %d", i))
	//		if err != nil {
	//			logger.Debugf("%v", err)
	//		}
	//	}
	//}()
	//go func() {
	//	c , err :=  events.Consume("test")
	//	if err != nil {
	//		logger.Fatal(err)
	//	}
	//	for {
	//		event, ok := <- c
	//		if !ok {
	//			break
	//		}
	//		var content string
	//		json.Unmarshal(event.Payload, &content)
	//		logger.Debugf("%s", content)
	//
	//	}
	//}()

	// Register handler
	pb.RegisterHelloworldHandler(srv.Server(), new(handler.Helloworld))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
