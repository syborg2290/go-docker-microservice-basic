package main

import (
	"flag"
)

func main() {
	// client := client.New("http://localhost:3000")

	// price, err := client.FetchPrice(context.Background(), "ETH")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", price)
	// return
	
	listenAddr := flag.String("listenAddr", ":3000", "listen address the service is running")
	flag.Parse()
	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJsonAPIServer(*listenAddr, svc)
	server.Run()

	// price, err := svc.FetchPrice(context.Background(), "ETH")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(price)
}
