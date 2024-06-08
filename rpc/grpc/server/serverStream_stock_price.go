package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"rpctest/grpc/protocol/pb"
	"time"
)

type StockServer struct {
	pb.UnimplementedStockServiceServer
}

type GlobalQuote struct {
	Symbol           string `json:"01. symbol"`
	Open             string `json:"02. open"`
	High             string `json:"03. high"`
	Low              string `json:"04. low"`
	Price            string `json:"05. price"`
	Volume           string `json:"06. volume"`
	LatestTradingDay string `json:"07. latest trading day"`
	PreviousClose    string `json:"08. previous close"`
	Change           string `json:"09. change"`
	ChangePercent    string `json:"10. change percent"`
}

type AlphaVantageResponse struct {
	GlobalQuote GlobalQuote `json:"Global Quote"`
}

// get alphaVantageAPIKey from url: https://www.alphavantage.co/support/#api-key
func FetchStockPrice(symbol string) (*GlobalQuote, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&interval=1min&apikey=%s", symbol, alphaVantageAPIKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch stock price: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}

	var result AlphaVantageResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &result.GlobalQuote, nil
}

func (s *StockServer) GetStockPrices(req *pb.StockSymbol, stream pb.StockService_GetStockPricesServer) error {
	symbol := req.GetSymbol()
	if symbol == "" {
		return fmt.Errorf("empty symbol")
	}

	for {
		quote, err := FetchStockPrice(symbol)
		if err != nil {
			log.Printf("Error fetching stock price for %s: %v", symbol, err)
			return err
		}
		data := &pb.StockData{
			Symbol:           quote.Symbol,
			Open:             quote.Open,
			High:             quote.High,
			Low:              quote.Low,
			Price:            quote.Price,
			Volume:           quote.Volume,
			LatestTradingDay: quote.LatestTradingDay,
			PreviousClose:    quote.PreviousClose,
			Change:           quote.Change,
			ChangePercent:    quote.ChangePercent,
			Timestamp:        time.Now().Format(time.RFC3339),
		}

		if err := stream.Send(data); err != nil {
			return err
		}

		time.Sleep(10 * time.Second)
	}
}
