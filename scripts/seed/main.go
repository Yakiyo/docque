package main

import (
	"context"
	"math/rand"
	"time"

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
	for i := 1; i <= n; i++ {
		genDoc()
	}
}

// generate a random doctor instance
func genDoc() error {
	ctx := context.Background()
	docName := faker.Name()

	doc, err := db.Client.Doctor.CreateOne(
		db.Doctor.Name.Set(docName),
	).Exec(ctx)

	if err != nil {
		return err
	}
	// random int between 5 to 20
	n := rand.Intn(15) + 5
	for i := 1; i <= n; i++ {
		genAppointment(doc)
	}
	return nil
}

// generate a random appointment instance for a doctor
func genAppointment(doc *db.DoctorModel) {
	t1, err := time.Parse(faker.TimeFormat, faker.TimeString())
	if err != nil {
		panic(err)
	}
	d, _ := time.ParseDuration("30m")
	t2 := t1.Add(d)
	db.Client.Appointment.CreateOne(
		db.Appointment.Patient.Set(faker.Name()),
		db.Appointment.Start.Set(t1),
		db.Appointment.End.Set(t2),
		db.Appointment.DoctorID.Set(doc.ID),
	)
}
