package naturaldateparser

import (
	"testing"
	"time"
)

func equalIgnoreSeconds(a time.Time, b time.Time) bool {
	return a.Minute() == b.Minute() &&
		a.Hour() == b.Hour() &&
		a.Day() == b.Day() &&
		a.Month() == b.Month() &&
		a.Year() == b.Year()
}

func TestParseSecond(t *testing.T) {
	p := CreateNaturalDateParser()
	res, err := p.Parse("1 second ago")
	if err != nil {
		t.Error(err)
	}

	expected := time.Now().Add(time.Second * -1)

	if !(equalIgnoreSeconds(res, expected) && res.Second() == expected.Second()) {
		t.Error("Dates are not equal")
	}
}

func TestParseSeconds(t *testing.T) {
	p := CreateNaturalDateParser()
	res, err := p.Parse("2 seconds ago")
	if err != nil {
		t.Error(err)
	}

	expected := time.Now().Add(time.Second * -2)

	if !(equalIgnoreSeconds(res, expected) && res.Second() == expected.Second()) {
		t.Error("Dates are not equal")
	}
}

func TestParseHour(t *testing.T) {
	p := CreateNaturalDateParser()
	res, err := p.Parse("1 hour ago")

	if err != nil {
		t.Error(err)
	}

	if !equalIgnoreSeconds(res, time.Now().Add(time.Hour*-1)) {
		t.Error("Dates are not equal")
	}
}
func TestParseHours(t *testing.T) {
	p := CreateNaturalDateParser()
	res, err := p.Parse("10 hours ago")

	if err != nil {
		t.Error(err)
	}

	if !equalIgnoreSeconds(res, time.Now().Add(time.Hour*-10)) {
		t.Error("Dates are not equal")
	}
}

func TestParseMinute(t *testing.T) {
	p := CreateNaturalDateParser()
	res, err := p.Parse("1 minute ago")

	if err != nil {
		t.Error(err)
	}

	if !equalIgnoreSeconds(res, time.Now().Add(time.Minute*-1)) {
		t.Error("Dates are not equal")
	}

}

func TestParseMinutes(t *testing.T) {
	p := CreateNaturalDateParser()
	res, err := p.Parse("10 minutes ago")

	if err != nil {
		t.Error(err)
	}

	if !equalIgnoreSeconds(res, time.Now().Add(time.Minute*-10)) {
		t.Error("Dates are not equal")
	}

}

func TestParseDay(t *testing.T) {
	p := CreateNaturalDateParser()
	res, err := p.Parse("1 day ago")

	if err != nil {
		t.Error(err)
	}

	if !equalIgnoreSeconds(res, time.Now().Add(time.Hour*24*-1)) {
		t.Error("Dates are not equal")
	}
}

func TestParseYesterday(t *testing.T) {
	p := CreateNaturalDateParser()
	res, err := p.Parse("yesterday")

	if err != nil {
		t.Error(err)
	}

	if !equalIgnoreSeconds(res, time.Now().Add(time.Hour*24*-1)) {
		t.Error("Dates are not equal")
	}
}

func TestParseDays(t *testing.T) {
	p := CreateNaturalDateParser()
	res, err := p.Parse("2 days ago")

	if err != nil {
		t.Error(err)
	}

	if !equalIgnoreSeconds(res, time.Now().Add(time.Hour*24*-2)) {
		t.Error("Dates are not equal")
	}
}

func TestParseWeek(t *testing.T) {
	p := CreateNaturalDateParser()
	res, err := p.Parse("1 week ago")

	if err != nil {
		t.Error(err)
	}

	if !equalIgnoreSeconds(res, time.Now().Add(time.Hour*24*7*-1)) {
		t.Error("Dates are not equal")
	}
}

func TestParseWeeks(t *testing.T) {
	p := CreateNaturalDateParser()
	res, err := p.Parse("2 weeks ago")

	if err != nil {
		t.Error(err)
	}

	if !equalIgnoreSeconds(res, time.Now().Add(time.Hour*24*7*-2)) {
		t.Error("Dates are not equal")
	}
}

func TestParseMonth(t *testing.T) {
	p := CreateNaturalDateParser()
	res, err := p.Parse("2 months ago")

	if err != nil {
		t.Error(err)
	}

	if !equalIgnoreSeconds(res, time.Now().Add(time.Hour*24*30*-2)) {
		t.Error("Dates are not equal")
	}
}

func TestParseToday(t *testing.T) {
	p := CreateNaturalDateParser()
	res, err := p.Parse("today")

	if err != nil {
		t.Error(err)
	}

	now := time.Now()

	if !(res.Day() == now.Day() &&
		res.Month() == now.Month() &&
		res.Year() == now.Year()) {
		t.Error("Dates are not equal")
	}
}
