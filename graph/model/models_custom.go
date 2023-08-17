package model

type HackathonApplication struct {
	ID                    string            `json:"id"`
	Status                ApplicationStatus `json:"status"`
	User                  *User             `json:"user"`
	UserID                string            `json:"userID"`
	Hackathon             *Hackathon        `json:"hackathon"`
	HackathonID           string            `json:"hackathonID"`
	WhyAttend             []string          `json:"whyAttend"`
	WhatDoYouWantToLearn  []string          `json:"whatDoYouWantToLearn"`
	ShareInfoWithSponsors bool              `json:"shareInfoWithSponsors"`
	ResumeURL             string            `json:"resumeURL"`
}

func (HackathonApplication) IsEntity() {}
