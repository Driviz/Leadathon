package main

import (
	"context"
	"log"

	"github.com/lead/chessgames"
	"github.com/lead/service"
)

func main() {
	ctx := context.Background()
	res, err := chessgames.GetFile(ctx, "https://www.chessgames.com/chessecohelp.html")
	if err != nil {
		log.Println("error getting file", "error", err)
		return
	}

	log.Println("response", res)

	svc := service.NewService(res)
	svc.StartService()
}
