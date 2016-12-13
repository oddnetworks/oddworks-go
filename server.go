package main

import (
	"fmt"

	"github.com/oddnetworks/oddworks-go/services"
	"github.com/oddnetworks/oddworks-go/types"
)

func main() {
	secrets := make(map[string]string)
	secrets["vimeo"] = "c65a2d4832c0c2e0f7eb5525e1b5eec7"

	odd := types.Channel{ID: "odd-networks", Secrets: secrets}
	dronetv := types.Channel{ID: "dronetv"}

	catalog := &services.CatalogService{
		Store: &stores["memory"],
	}

	fmt.Println(catalog.GetVideo("1", odd))
	fmt.Println(catalog.GetCollection("1", odd))
	fmt.Println(catalog.GetVideo("1", dronetv))
}
