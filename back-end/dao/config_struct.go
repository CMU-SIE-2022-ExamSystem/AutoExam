package dao

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// @Description assessment
type AutoExam_Assessments struct {
	Id             string     `yaml:"id" json:"id" bson:"id"`                 // name of assessment
	Course         string     `yaml:"course" json:"course" bson:"course"`     // course of assessment
	BaseCourse     string     `json:"-" bson:"base_course"`                   // base course of assessment
	General        General    `yaml:"general" json:"general"`                 // general details of  assessment
	Settings       []Settings `yaml:"settings" json:"settings"`               // questions settings of  assessment
	Draft          bool       `json:"draft" bson:"draft"`                     // whether the assessment could be used for student
	Generated      int        `json:"generated" bson:"generated"`             // whether assessment is generated for all student in the course. 0: not generated, 1: already generated, -1: generated error
	GeneratedError string     `json:"generated_error" bson:"generated_error"` // error message for an error happened when generatings all student's exam
} // @name Assessments

type AutoExam_Assessments_Student struct {
	Description string `yaml:"description" json:"description" bson:"description"`                                                                                 // description of assessment
	Start_at    string `yaml:"start_at" json:"start_at" bson:"start_at" default:"2022-06-15T15:04:05Z" binding:"required,datetime=2006-01-02T15:04:05.000-07:00"` // start time of assessment
	End_at      string `yaml:"end_at" json:"end_at" bson:"end_at" default:"2023-06-15T15:04:05Z" binding:"required,datetime=2006-01-02T15:04:05.000-07:00"`       // end time of assessment
} // @name Assessments

type AutoExam_Assessments_Create struct {
	Name          string `yaml:"name" json:"name" bson:"name" binding:"required"`
	Category_name string `yaml:"category_name" json:"category_name" bson:"category_name" default:"Exam" binding:"required,oneof=Exam Quiz"`
} // @name Assessments

// @Description assessment update structure
type AutoExam_Assessments_Update struct {
	General  General    `yaml:"general" json:"general"`   // general details of the assessment
	Settings []Settings `yaml:"settings" json:"settings"` // questions settings of the assessment
} // @name Assessments

type AutoExam_Assessments_Update_Validate struct {
	BaseCourse string
	General    General    `yaml:"general" json:"general"`   // general details of the assessment
	Settings   []Settings `yaml:"settings" json:"settings"` // questions settings of the assessment
}

// @Description questions general structure
type General struct {
	Name             string `yaml:"name" json:"name" bson:"name" binding:"required"`                                                                                                           // name of assessment
	Description      string `yaml:"description" json:"description" bson:"description"`                                                                                                         // description of assessment
	Category_name    string `yaml:"category_name" json:"category_name" bson:"category_name" default:"Exam" binding:"required,oneof=Exam Quiz"`                                                 // only accept Exam or Quiz
	Start_at         string `yaml:"start_at" json:"start_at" bson:"start_at" default:"2022-06-15T15:04:05Z" binding:"required,datetime=2006-01-02T15:04:05.000-07:00"`                         // start time of assessment
	End_at           string `yaml:"end_at" json:"end_at" bson:"end_at" default:"2023-06-15T15:04:05Z" binding:"required,datetime=2006-01-02T15:04:05.000-07:00"`                               // end time of assessment
	Grading_deadline string `yaml:"grading_deadline" json:"grading_deadline" bson:"grading_deadline" default:"2023-06-16T15:04:05Z" binding:"required,datetime=2006-01-02T15:04:05.000-07:00"` // grading deadline of assessment
	MaxSubmissions   int    `yaml:"max_submissions" json:"max_submissions" bson:"max_submissions" binding:"required,gte=1"`                                                                    // number of submission, Exam category would only accept 1
	Url              string `json:"url" bson:"url"`                                                                                                                                            // assessment url
}

// @Description questions settings structure
type Settings struct {
	Id                []string  `yaml:"id" json:"id" bson:"id"`                                                                       // specify the possible question ids for the assessment, can be empty. If there are multiple ids, the exam would randomly choose from those ids. If there is no id, the exam would randomly select the question based on tag_id and sub_question_number
	Title             string    `yaml:"title" json:"title" bson:"title" default:"title"`                                              // title for the exam of this question
	Tag               string    `yaml:"tag" json:"tag" bson:"tag" binding:"required"`                                                 // tag id of this question
	Max_score         float64   `yaml:"max_score" json:"max_score" bson:"max_score"`                                                  // max score of this question
	Scores            []float64 `yaml:"scores" json:"scores" bson:"scores"`                                                           // detail sub score for each sub_question, can be empty. If this field is empty, the sub score would be divide equally based on max_score and sub_question_number
	SubQuestionNumber int       `yaml:"sub_question_number" json:"sub_question_number" bson:"sub_question_number" binding:"required"` // sub question number of this question, cannot be empty
}

type Problems struct {
	Name        string        `yaml:"name" json:"name" bson:"name"`
	Type        string        `yaml:"type" json:"type" bson:"type"`
	Description []Description `yaml:"description" json:"description" bson:"description"`
	MaxScore    float64       `yaml:"max_score" json:"max_score" bson:"max_score"`
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

type Statistic struct {
	Number  int     `json:"number" bson:"number"`
	Highest float64 `json:"highest" bson:"highest"`
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

func (assessment *AutoExam_Assessments_Update_Validate) ToAutoExamAssessments(course string) AutoExam_Assessments {
	assessment.General.Start_at = Time_to_EST(assessment.General.Start_at)
	assessment.General.End_at = Time_to_EST(assessment.General.End_at)
	assessment.General.Grading_deadline = Time_to_EST(assessment.General.Grading_deadline)

	autoexam := AutoExam_Assessments{
		Id:         assessment.General.Name,
		Course:     course,
		BaseCourse: assessment.BaseCourse,
		General:    assessment.General,
		Settings:   assessment.Settings,
		Draft:      true,
	}

	return autoexam
}

func (assessment *AutoExam_Assessments) ToDownloadAssessments() models.Download_Assessments {
	general := models.General{}
	general.Default()
	AutoExamToDownloadAssessmentsGeneral(*assessment, &general)

	autograder := models.Autograder{}
	autograder.Default()

	ass := models.Download_Assessments{
		General:    general,
		Problems:   AutoExamToDownloadAssessmentsProblems(*assessment),
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

func AutoExamToDownloadAssessmentsProblems(assessment AutoExam_Assessments) []models.Problems {
	problems := []models.Problems{}

	for i, setting := range assessment.Settings {
		for j := 0; j < setting.SubQuestionNumber; j++ {
			problems = append(problems, models.Problems{
				Name:        ToSubQuestionName(i+1, j+1),
				Description: "",
				Max_score:   setting.Scores[j],
				Optional:    false,
			})
		}
	}
	return problems
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

func (assessment *AutoExam_Assessments) ToAssessmentsStudent() AutoExam_Assessments_Student {
	ass := AutoExam_Assessments_Student{
		Description: assessment.General.Description,
		Start_at:    assessment.General.Start_at,
		End_at:      assessment.General.End_at,
	}
	return ass
}

func (assessment *AutoExam_Assessments) GenerateAssessmentStudent(email, course_name, assessment_name string) Assessment_Student {
	var questions []string
	var problems []Student_Problems
	solutions := make(Student_Questions)
	for i, ass := range assessment.Settings {
		// get specified question_id id
		var question_id string
		if len(ass.Id) != 0 { // limit id
			number := len(ass.Id)
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			question_id = ass.Id[r1.Intn(number)]
		} else { // not limit id
			ids := GetAllQuestionIDBySubQuestionNumber(assessment.BaseCourse, ass.Tag, ass.SubQuestionNumber)
			number := len(ids)
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			question_id = ids[r1.Intn(number)]
		}

		question, _ := ReadQuestionById(question_id)

		sub_solutions := make(Student_Sub_Questions)
		for j, sub_quest := range question.SubQuestions {
			// create problems
			problem := Student_Problems{
				Name:     ToSubQuestionName(i+1, j+1),
				Grader:   sub_quest.Grader,
				MaxScore: ass.Scores[j],
			}
			problems = append(problems, problem)
			// create solutions
			sub_sub_solutions := make(Student_Sub_Sub_Questions)
			for k, sol := range sub_quest.Solutions {
				sub_sub_solutions[ToSubSubQuestionName(i+1, j+1, k+1)] = sol
			}
			sub_solutions[ToSubQuestionName(i+1, j+1)] = sub_sub_solutions
		}

		questions = append(questions, question_id)
		solutions["Q"+strconv.Itoa(i+1)] = sub_solutions
	}

	student := Assessment_Student{
		Email:      email,
		Course:     course_name,
		Assessment: assessment_name,
		Questions:  questions,
		Problems:   problems,
		Solutions:  solutions,
	}

	return student
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

func ToSubQuestionName(question_index, sub_question_index int) string {
	return "Q" + strconv.Itoa(question_index) + "_" + "sub" + strconv.Itoa(sub_question_index)
}
func ToSubSubQuestionName(question_index, sub_question_index, sub_sub_question_index int) string {
	return "Q" + strconv.Itoa(question_index) + "_" + "sub" + strconv.Itoa(sub_question_index) + "_sub" + strconv.Itoa(sub_sub_question_index)
}
