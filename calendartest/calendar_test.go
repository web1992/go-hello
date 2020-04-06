package calendartest

import (
	"fmt"
	"log"
	"testing"
	"web1992/calendar"
)

func TestCalendar(t *testing.T) {
	date := calendar.Date{}

	err := date.SetYear(2020)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(12)

	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(6)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
}

func TestEvent(t *testing.T) {

	event := calendar.Event{}

	event.Title = "Title"
	event.SetYear(2020)
	event.SetMonth(3)
	event.SetDay(20)

	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())

	date := event.Date
	fmt.Println(date)
	// dot chain
	fmt.Println(event.Date.Year())

}
