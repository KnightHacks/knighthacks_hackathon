// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Event struct {
	ID        string     `json:"id"`
	Hackathon *Hackathon `json:"hackathon"`
}

func (Event) IsEntity() {}

type Hackathon struct {
	ID        string     `json:"id"`
	Term      *Term      `json:"term"`
	Attendees []*User    `json:"attendees"`
	Sponsors  []*Sponsor `json:"sponsors"`
	Events    []*Event   `json:"events"`
}

type HackathonFilterInput struct {
	Year     int       `json:"year"`
	Semester *Semester `json:"semester"`
}

type Sponsor struct {
	ID         string       `json:"id"`
	Hackathons []*Hackathon `json:"hackathons"`
}

func (Sponsor) IsEntity() {}

type Term struct {
	Year     int      `json:"year"`
	Semester Semester `json:"semester"`
}

type User struct {
	ID         string       `json:"id"`
	Hackathons []*Hackathon `json:"hackathons"`
}

func (User) IsEntity() {}

type Semester string

const (
	SemesterFall   Semester = "FALL"
	SemesterSpring Semester = "SPRING"
	SemesterSummer Semester = "SUMMER"
)

var AllSemester = []Semester{
	SemesterFall,
	SemesterSpring,
	SemesterSummer,
}

func (e Semester) IsValid() bool {
	switch e {
	case SemesterFall, SemesterSpring, SemesterSummer:
		return true
	}
	return false
}

func (e Semester) String() string {
	return string(e)
}

func (e *Semester) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Semester(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Semester", str)
	}
	return nil
}

func (e Semester) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
