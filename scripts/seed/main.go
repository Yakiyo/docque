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
	log.SetReportTimestamp(false)
	// random int between 4 and 9
	n := rand.Intn(5) + 4
	log.Info("generating doctors", "n", n)
	for i := 1; i <= n; i++ {
		genDoc()
	}
	ctx := context.Background()
	docs, _ := db.Client.Doctor.FindMany().Exec(ctx)
	aps, _ := db.Client.Appointment.FindMany().Exec(ctx)
	log.Info("generated doctors", "count", len(docs))
	log.Info("generated appointments", "count", len(aps))
}

// generate a random doctor instance
func genDoc() {
	ctx := context.Background()
	docName := faker.Name()

	doc, err := db.Client.Doctor.CreateOne(
		db.Doctor.Name.Set(docName),
	).Exec(ctx)

	if err != nil {
		panic(err)
	}
	log.Info("created doctor instance", "val", doc)
	// random int between 5 to 20
	n := rand.Intn(15) + 5
	log.Info("generating appointments", "n", n)
	for i := 1; i <= n; i++ {
		genAppointment(doc)
	}
}

// generate a random appointment instance for a doctor
func genAppointment(doc *db.DoctorModel) {
	ctx := context.Background()
	t1, err := time.Parse(faker.TimeFormat, faker.TimeString())
	if err != nil {
		panic(err)
	}
	d, _ := time.ParseDuration("30m")
	t2 := t1.Add(d)
	a, err := db.Client.Appointment.CreateOne(
		db.Appointment.Patient.Set(faker.Name()),
		db.Appointment.Start.Set(t1),
		db.Appointment.End.Set(t2),
		db.Appointment.DoctorID.Set(doc.ID),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}
	log.Info("created appointment instance", "val", a)
}
