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
	MentorId           *primitive.ObjectID `json:"mentor_id,omitempty"` // Linked mentor ID, null until assigned
	College_Details    ClgDetails          `json:"collegeDetails"`
}

type ClgDetails struct {
	CollegeName     string              `json:"college_name"`
	CollegeId       *primitive.ObjectID `json:"college_id,omitempty"` // Optional if colleges are predefined in DB
	Degree          string              `json:"degree"`               // e.g., B.Tech, B.Sc
	Branch          string              `json:"branch"`               // e.g., Computer Science
	YearOfStudy     int                 `json:"year_of_study"`        // 1, 2, 3, 4, etc.
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
	CompanyId       *primitive.ObjectID `json:"company_id,omitempty"` // Optional for predefined companies
	Position        string              `json:"position"`             // e.g., "Software Engineer", "Manager"
	Department      string              `json:"department"`           // e.g., "IT", "HR"
	YearsInCompany  int                 `json:"years_in_company"`     // Years of experience in the current company
	CompanyLocation string              `json:"company_location"`     // e.g., "Mumbai, India"
}
