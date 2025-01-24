package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id         primitive.ObjectID `json:"user_id"`
	Name       string             `json:"name"`
	Profession string             `json:"profession"`
	Email      string             `json:"email"`
	Password   string             `json:"password"`
	Age        int64              `json:"age"`
}

type Student struct {
	Id                 primitive.ObjectID  `json:"student_id"`
	UserId             primitive.ObjectID  `json:"user_id"`
	SubjectsOfInterest []string            `json:"subjects_of_interest"`
	MentorId           *primitive.ObjectID `json:"mentor_id,omitempty"` 
	College_Details    ClgDetails          `json:"collegeDetails"`
}

type ClgDetails struct {
	CollegeName     string              `json:"college_name"`
	CollegeId       *primitive.ObjectID `json:"college_id,omitempty"` 
	Degree          string              `json:"degree"`               
	Branch          string              `json:"branch"`               
	YearOfStudy     int                 `json:"year_of_study"`      
	RollNumber      string              `json:"roll_number"`
	CollegeLocation string              `json:"college_location"`
	GraduationYear  int                 `json:"graduation_year"`
	Grade           string              `json:"grade"` // CGPA
}

type Mentor struct {
	Id             primitive.ObjectID   `json:"mentor_id"`
	UserId         primitive.ObjectID   `json:"user_id"`
	Specialization []string             `json:"specialization"`
	Experience     int                  `json:"experience"`
	Qualification  string               `json:"qualification"`
	Students       []primitive.ObjectID `json:"students"`
	Availability   map[string]string    `json:"availability"`
	Achievements   []string             `json:"achievements"`
	Status         string               `json:"status,omitempty"`
	ApprovalDate   *time.Time           `json:"approval_date,omitempty"`
	AdminComments  string               `json:"admin_comments,omitempty"`
	CompanyDetails CompanyDetails       `json:"company_details"`
}

type CompanyDetails struct {
	CompanyName     string              `json:"company_name"`
	CompanyId       *primitive.ObjectID `json:"company_id,omitempty"` 
	Position        string              `json:"position"`            
	Department      string              `json:"department"`         
	YearsInCompany  int                 `json:"years_in_company"`   
	CompanyLocation string              `json:"company_location"`    
}
