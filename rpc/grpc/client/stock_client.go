package client

import (
	"context"
	"fmt"
	"log"
	"rpctest/grpc/protocol/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartStockClient() {

	conn, err := grpc.NewClient(*address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewStockServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// send stock symbol to server
	symbol := &pb.StockSymbol{Symbol: "IBM"}
	stream, err := c.GetStockPrices(ctx, symbol)
	if err != nil {
		log.Fatal("Could not get stock prices: %v", err)
	}

	for {
		data, err := stream.Recv()
		if err != nil {
			log.Fatal("Failed to receive stock price: %v", err)
		}

		fmt.Printf("Received stock data: Symbol: %s, Open: %s, High: %s, Low: %s, Price: %s, Volume: %s, Latest Trading Day: %s, Previous Close: %s, Change: %s, Change Percent: %s, Timestamp: %s\n",
			data.GetSymbol(), data.GetOpen(), data.GetHigh(), data.GetLow(), data.GetPrice(), data.GetVolume(), data.GetLatestTradingDay(), data.GetPreviousClose(), data.GetChange(), data.GetChangePercent(), data.GetTimestamp())

	}
}
