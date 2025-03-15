package main

import (
	"context"
	"log"

	"github.com/cerbos/cerbos-sdk-go/cerbos"
)

func main() {
	c, err := cerbos.New("localhost:3593", cerbos.WithPlaintext())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	principal := cerbos.NewPrincipal("seller", "user")
	principal.WithAttr("regular", true)

	kind := "product:object"
	actions := []string{"view:public", "create", "edit"}

	r1 := cerbos.NewResource(kind, "Private_Product_Seller")
	r1.WithAttributes(map[string]any{
		"owner":   "seller",
		"public":  false,
		"flagged": true,
	})

	r2 := cerbos.NewResource(kind, "Public_Product")
	r2.WithAttributes(map[string]any{
		"owner":   "another_seller",
		"public":  true,
		"flagged": true,
	})

	batch := cerbos.NewResourceBatch()
	batch.Add(r1, actions...)
	batch.Add(r2, actions...)

	resp, err := c.CheckResources(context.Background(), principal, batch)
	if err != nil {
		log.Fatalf("Failed to check resources: %v", err)
	}
	for _, result := range resp.Results {
		for action, status := range result.Actions {
			log.Printf("Resource: '%s' -> Actions: '%s' -> Status: '%s'", result.Resource.Id, action, status.Enum())
		}
	}

}
