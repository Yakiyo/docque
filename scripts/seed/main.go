package main

import (
	"context"
	"math/rand"

	"github.com/Yakiyo/docque/db"
	"github.com/charmbracelet/log"
	"github.com/go-faker/faker/v4"
)

func main() {
	db.Init()
	defer db.Close()
	// random int between 4 and 9
	n := rand.Intn(5) + 4
	log.Info("Generating doctors", "n", n)
	for i := 0; i <= n; i++ {
		genDoc()
	}
}

func genDoc() error {
	ctx := context.Background()
	docName := faker.Name()

	_, err := db.Client.Doctor.CreateOne(
		db.Doctor.Name.Set(docName),
	).Exec(ctx)

	if err != nil {
		return err
	}
	// random int between 5 to 20
	_ = rand.Intn(15) + 5
	return nil
}
