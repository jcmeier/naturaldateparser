package naturaldateparser

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type NaturalDateParser struct {
	secondRegex *regexp.Regexp
	minuteRegex *regexp.Regexp
	hourRegex   *regexp.Regexp
	dayRegex    *regexp.Regexp
	weekRegex   *regexp.Regexp
	monthRegex  *regexp.Regexp
}

func CreateNaturalDateParser() *NaturalDateParser {
	secondRegex, err := regexp.Compile("(\\d\\d*) seconds? ago")
	if err != nil {
		log.WithError(err).Panic("Second regex does not compile")
	}

	minuteRegex, err := regexp.Compile("(\\d\\d*) minutes? ago")
	if err != nil {
		log.WithError(err).Panic("Minute regex does not compile")
	}

	hourRegex, err := regexp.Compile("(\\d\\d*) hours? ago")
	if err != nil {
		log.WithError(err).Panic("Hour regex does not compile")
	}

	dayRegex, err := regexp.Compile("(\\d\\d*) days? ago")
	if err != nil {
		log.WithError(err).Panic("Day regex does not compile")
	}

	weekRegex, err := regexp.Compile("(\\d\\d*) weeks? ago")
	if err != nil {
		log.WithError(err).Panic("Week regex does not compile")
	}

	monthRegex, err := regexp.Compile("(\\d\\d*) months? ago")
	if err != nil {
		log.WithError(err).Panic("Month regex does not compile")
	}

	return &NaturalDateParser{
		secondRegex: secondRegex,
		minuteRegex: minuteRegex,
		hourRegex:   hourRegex,
		dayRegex:    dayRegex,
		weekRegex:   weekRegex,
		monthRegex:  monthRegex,
	}
}

func (p *NaturalDateParser) Parse(naturaldate string) (time.Time, error) {
	secondMatch := p.secondRegex.FindStringSubmatch(naturaldate)
	if secondMatch != nil {
		minutes, err := strconv.Atoi(secondMatch[1])

		if err != nil {
			return time.Now(), err
		}

		return time.Now().Add(time.Second * time.Duration(minutes) * -1), nil

	}

	minuteMatch := p.minuteRegex.FindStringSubmatch(naturaldate)
	if minuteMatch != nil {
		minutes, err := strconv.Atoi(minuteMatch[1])

		if err != nil {
			return time.Now(), err
		}

		return time.Now().Add(time.Minute * time.Duration(minutes) * -1), nil

	}

	hourMatch := p.hourRegex.FindStringSubmatch(naturaldate)
	if hourMatch != nil {
		hours, err := strconv.Atoi(hourMatch[1])

		if err != nil {
			return time.Now(), err
		}

		return time.Now().Add(time.Hour * time.Duration(hours) * -1), nil
	}

	if strings.Contains(naturaldate, "today") {
		return time.Now(), nil
	}

	dayMatch := p.dayRegex.FindStringSubmatch(naturaldate)
	if dayMatch != nil {
		days, err := strconv.Atoi(dayMatch[1])

		if err != nil {
			return time.Now(), err
		}

		return time.Now().Add(time.Hour * 24 * time.Duration(days) * -1), nil
	}

	weekmatch := p.weekRegex.FindStringSubmatch(naturaldate)
	if weekmatch != nil {
		weeks, err := strconv.Atoi(weekmatch[1])

		if err != nil {
			return time.Now(), err
		}

		return time.Now().Add(time.Hour * 24 * 7 * time.Duration(weeks) * -1), nil
	}

	monthMatch := p.monthRegex.FindStringSubmatch(naturaldate)
	if monthMatch != nil {
		months, err := strconv.Atoi(monthMatch[1])

		if err != nil {
			return time.Now(), err
		}

		return time.Now().Add(time.Hour * 24 * 30 * time.Duration(months) * -1), nil
	}

	return time.Now(), errors.New("Could not be parsed")
}
