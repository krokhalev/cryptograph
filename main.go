package main

import (
	"fmt"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"time"
)

func main() {
	// No keys required for crypto data
	client := marketdata.NewClient(marketdata.ClientOpts{})

	request := marketdata.GetCryptoBarsRequest{
		TimeFrame: marketdata.OneMin,
		Start:     time.Now().UTC().Add(time.Duration(-1) * time.Hour),
		End:       time.Now().UTC(),
	}

	bars, err := client.GetCryptoBars("BTC/USD", request)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(bars))

	lowestPrice := bars[0].Low
	highestPrice := bars[0].High
	for _, bar := range bars {
		fmt.Printf("%+v\n", bar)
		if bar.Low < lowestPrice {
			lowestPrice = bar.Low
		}
		if bar.High > highestPrice {
			highestPrice = bar.High
		}
	}
	fmt.Println(lowestPrice)
	fmt.Println(highestPrice)
}
