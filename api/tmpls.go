package api

import (
	_ "embed"
	"text/template"
	"time"

	"github.com/Yakiyo/docque/db"
)

//go:embed templates/queue.html
var queue string

var queueTmpl = template.Must(template.New("queue").Parse(queue))

// the `Doctor` type used in queue.html template
type QueueDoc struct {
	Name         string
	Id           int
	Appointments []QueueAppointment
}

type QueueAppointment struct {
	Patient  string
	Start    time.Time
	Duration int
}

func atoqa(a db.AppointmentModel) QueueAppointment {
	return QueueAppointment{
		Patient:  a.Patient,
		Start:    a.Start.UTC(),
		Duration: int(a.End.UTC().Sub(a.Start.UTC())),
	}
}

// iterate a slice `slice` of elements of type T, and apply
// the function `f` which returns a value of type `U`
func iter_map[T any, U any](slice []T, f func(T) U) []U {
	res := []U{}
	for _, v := range slice {
		res = append(res, f(v))
	}
	return res
}
