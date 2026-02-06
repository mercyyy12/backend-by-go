package main

import (
	"fmt"
	"time"
)

type data interface {
	showData() string
	worker() (int, string)
}

// email task
type structEmail struct {
	emailType string
	emailId   int
	emailTime time.Time
	Payload   string
}

// image task
type structImage struct {
	imageType string
	imageId   int
	imageTime time.Time
	Payload   string
}

// report task
type structReport struct {
	reportType string
	reportId   int
	reportTime time.Time
	Payload    string
}

// creates a new email task and stored globally
func (e *structEmail) addEmail(payload string) int {
	e.emailId = c()          // assign new unique id
	e.emailTime = time.Now() // timestamp
	e.emailType = "email"
	e.Payload = payload
	storedata[e.emailId] = e // store in global map
	return e.emailId
}

// creates a new image task and stored globally
func (i *structImage) addImage(payload string) int {
	i.Payload = payload
	i.imageTime = time.Now()
	i.imageType = "image"
	i.imageId = c()
	storedata[i.imageId] = i
	return i.imageId
}

// creates a new report task and stored globally
func (r *structReport) addReport(payload string) int {
	r.Payload = payload
	r.reportTime = time.Now()
	r.reportType = "report"
	r.reportId = c()
	storedata[r.reportId] = r
	return r.reportId
}

// returns formatted string for email
func (e *structEmail) showData() string {
	return fmt.Sprintf("ID: %d | Type: Email | Time: %v | Payload: %s\n",
		e.emailId, e.emailTime.Format("15:04:05"), e.Payload)
}

// returns formatted string for image
func (e *structImage) showData() string {
	return fmt.Sprintf("ID: %d | Type: Image | Time: %v | Payload: %s\n",
		e.imageId, e.imageTime.Format("15:04:05"), e.Payload)
}

// returns formatted string for report
func (e *structReport) showData() string {
	return fmt.Sprintf("ID: %d | Type: Report | Time: %v | Payload: %s\n",
		e.reportId, e.reportTime.Format("15:04:05"), e.Payload)
}

// returns id and type for email
func (s *structEmail) worker() (int, string) {
	return s.emailId, s.emailType
}

// returns id and type for image
func (s *structImage) worker() (int, string) {
	return s.imageId, s.imageType
}

// returns id and type for report
func (s *structReport) worker() (int, string) {
	return s.reportId, s.reportType
}
