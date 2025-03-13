package main

import (
	"context"
	"fmt"
	"github.com/orbcorp/orb-go"
)

func main() {
	client := orb.NewClient()
	iter := client.Subscriptions.ListAutoPaging(context.TODO(), orb.SubscriptionListParams{
		ExternalCustomerID: orb.F([]string{"Stainless", "OpenAI"}),
	})
	// Automatically fetches more pages as needed.
	for iter.Next() {
		subscription := iter.Current()
		fmt.Printf("%+v\n", subscription)
	}
	if err := iter.Err(); err != nil {
		panic(err.Error())
	}
}
