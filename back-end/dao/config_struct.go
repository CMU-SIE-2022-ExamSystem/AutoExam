package dao

import (
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Student struct {
	StudentId string      `json:"studentId" bson:"studentId"`
	ExamId    string      `json:"examId" bson:"examId"`
	CourseId  string      `json:"course" bson:"course"`
	Settings  map[int]int `json:"settings" bson:"settings"`
}

type AutoExam_Assessments struct {
	Id       string     `yaml:"id" json:"id" bson:"id"`
	Course   string     `yaml:"course" json:"course" bson:"course"`
	General  General    `yaml:"general" json:"general"`
	Settings []Settings `yaml:"settings" json:"settings"`
	Draft    bool       `json:"draft" bson:"draft"`
}

type AutoExam_Assessments_Create struct {
	Name          string `yaml:"name" json:"name" bson:"name" binding:"required"`
	Category_name string `yaml:"category_name" json:"category_name" bson:"category_name" default:"Exam" binding:"required,oneof=Exam Quiz"`
}

type AutoExam_Assessments_Update struct {
	General  General    `yaml:"general" json:"general"`
	Settings []Settings `yaml:"settings" json:"settings"`
}

type General struct {
	Name             string `yaml:"name" json:"name" bson:"name" binding:"required"`
	Description      string `yaml:"description" json:"description" bson:"description"`
	Category_name    string `yaml:"category_name" json:"category_name" bson:"category_name" default:"Exam" binding:"required,oneof=Exam Quiz"`
	Start_at         string `yaml:"start_at" json:"start_at" bson:"start_at" default:"2022-06-15T15:04:05Z" binding:"required,datetime=2006-01-02T15:04:05.000-07:00"`
	End_at           string `yaml:"end_at" json:"end_at" bson:"end_at" default:"2023-06-15T15:04:05Z" binding:"required,datetime=2006-01-02T15:04:05.000-07:00"`
	Grading_deadline string `yaml:"grading_deadline" json:"grading_deadline" bson:"grading_deadline" default:"2023-06-15T15:04:05Z" binding:"required,datetime=2006-01-02T15:04:05.000-07:00"`
	MaxSubmissions   int    `yaml:"max_submissions" json:"max_submissions" bson:"max_submissions" binding:"required,gte=1"`
}

type Settings struct {
	Id        int    `yaml:"id" json:"id" bson:"id"`
	Title     string `yaml:"title" json:"title" bson:"title" default:"title"`
	Tag       string `yaml:"tag" json:"tag" bson:"tag" binding:"required"`
	Max_score int    `yaml:"max_score" json:"max_score" bson:"max_score"`
	Score     []int  `yaml:"score" json:"score" bson:"score"`
}

type Problems struct {
	Name        string        `yaml:"name" json:"name" bson:"name"`
	Type        string        `yaml:"type" json:"type" bson:"type"`
	Description []Description `yaml:"description" json:"description" bson:"description"`
	MaxScore    int           `yaml:"max_score" json:"max_score" bson:"max_score"`
	Optional    bool          `yaml:"optional" json:"optional" bson:"optional"`
}

type Description struct {
	Name string `yaml:"name" json:"name" bson:"name"`
	//Answer string `yaml:"answer" json:"answer" bson:"answer"`
	Score int `yaml:"score" json:"score" bson:"score"`
}

type Categories_Return struct {
	Categories []string `yaml:"categories" json:"categories"`
}

type Draft struct {
	Draft bool `yaml:"draft" json:"draft" bson:"draft"`
}

var (
	Assessment_Catergories []string = []string{"Exam", "Quiz"}
)

func (assessment *AutoExam_Assessments_Create) ToAutoExamAssessments(course string) AutoExam_Assessments {
	start_at, end_at := LocalTime_to_EST()

	general := General{
		Name:             assessment.Name,
		Description:      "",
		Category_name:    assessment.Category_name,
		Start_at:         start_at,
		End_at:           end_at,
		Grading_deadline: end_at,
		MaxSubmissions:   1,
	}

	autoexam := AutoExam_Assessments{
		Id:       assessment.Name,
		Course:   course,
		General:  general,
		Settings: []Settings{},
		Draft:    true,
	}

	return autoexam
}

func (assessment *AutoExam_Assessments_Update) ToAutoExamAssessments(course string) AutoExam_Assessments {
	assessment.General.Start_at = Time_to_EST(assessment.General.Start_at)
	assessment.General.End_at = Time_to_EST(assessment.General.End_at)
	assessment.General.Grading_deadline = Time_to_EST(assessment.General.Grading_deadline)

	autoexam := AutoExam_Assessments{
		Id:       assessment.General.Name,
		Course:   course,
		General:  assessment.General,
		Settings: assessment.Settings,
		Draft:    true,
	}

	return autoexam
}

func (assessment *AutoExam_Assessments) ToAutolabAssessments() models.Download_Assessments {
	general := models.General{}
	general.Default()
	AutoExamToDownloadAssessmentsGeneral(*assessment, &general)

	autograder := models.Autograder{}
	autograder.Default()

	ass := models.Download_Assessments{
		General: general,
		// Problems:   assessment.Problems,
		Autograder: autograder,
	}
	return ass
}

func AutoExamToDownloadAssessmentsGeneral(assessment AutoExam_Assessments, general *models.General) {
	general.Name = assessment.General.Name
	general.Description = assessment.General.Description
	general.Display_name = cases.Title(language.Und).String(assessment.General.Name)
	general.Category_name = assessment.General.Category_name
	general.Max_submissions = assessment.General.MaxSubmissions

	start_at := Time_str_convert(assessment.General.Start_at)
	end_at := Time_str_convert(assessment.General.End_at)
	start_at.Add(-time.Minute * models.TimeAhead)
	end_at.Add(time.Minute * models.TimeDelay)
	general.Start_at = Time_to_Download_Assessment(start_at)
	general.End_at = Time_to_Download_Assessment(end_at)
	general.Due_at = general.End_at
	general.Grading_deadline = Time_to_Download_Assessment_str(assessment.General.Grading_deadline)

}

func (assessment *AutoExam_Assessments) ToAssessments() models.Assessments {
	ass := models.Assessments{
		Name:             assessment.General.Name,
		Display_name:     assessment.General.Name,
		Start_at:         assessment.General.Start_at,
		Due_at:           assessment.General.End_at,
		End_at:           assessment.General.End_at,
		Category_name:    assessment.General.Category_name,
		Grading_deadline: assessment.General.Grading_deadline,
		Autolab:          false,
		AutoExam:         true,
		Draft:            assessment.Draft,
	}
	return ass
}

func Time_to_EST(t string) string {
	loc, _ := time.LoadLocation(models.TimeLoc)
	tm, _ := time.Parse(models.TimeFormat, t)
	return tm.In(loc).Format(models.TimeFormat)
}

func Time_to_UTC(t string) string {
	// loc, _ := time.LoadLocation(models.TimeLoc)
	tm, _ := time.Parse(models.TimeFormat, t)
	// return tm.In(loc).Format(models.TimeFormat)
	return tm.In(time.UTC).Format(models.TimeFormat)
}

func Time_to_Download_Assessment(t time.Time) string {
	// loc, _ := time.LoadLocation(models.TimeLoc)
	// tm, _ := time.Parse(models.TimeFormat, t)
	// return tm.In(loc).Format(models.TimeFormat)
	return t.In(time.UTC).Format(models.DownloadTimeFormat)
}

func Time_to_Download_Assessment_str(t string) string {
	// loc, _ := time.LoadLocation(models.TimeLoc)
	tm, _ := time.Parse(models.TimeFormat, t)
	// return tm.In(loc).Format(models.TimeFormat)
	return tm.In(time.UTC).Format(models.DownloadTimeFormat)
}

func Time_str_convert(t string) time.Time {
	tm, _ := time.Parse(models.TimeFormat, t)
	return tm
}

func LocalTime_to_EST() (start_at, end_at string) {
	loc, _ := time.LoadLocation(models.TimeLoc)
	tm := time.Now()
	start_at = tm.In(loc).Format(models.TimeFormat)
	end_at = tm.AddDate(0, 0, 1).In(loc).Format(models.TimeFormat)
	return
}
