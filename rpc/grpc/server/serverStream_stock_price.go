package server

import (
	"fmt"
	"net/http"
	"rpctest/grpc/protocol/pb"
)

type StockServer struct {
	pb.UnimplementedStockMarketServiceServer
}

// get alphaVantageAPIKey from url: https://www.alphavantage.co/support/#api-key
func fetchStockPrice(symbol string) (float64, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%s&interval=5min&apikey=%s", symbol, alphaVantageAPIKey)
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch stock price: %v", err)
	}
	defer resp.Body.Close()

}
