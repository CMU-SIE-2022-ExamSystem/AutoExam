package models

const (
	TimeDelay = 15 // in minutes
	TimeAhead = 15 // in minutes
)

type Download_Assessments struct {
	General    General    `mapstructure:"general"`
	Problems   []Problems `mapstructure:"problems"`
	Autograder Autograder `mapstructure:"autograder"`
}

type General struct {
	Name             string `json:"name" mapstructure:"name"`
	Description      string `json:"description" mapstructure:"description"`
	Display_name     string `json:"display_name" mapstructure:"display_name"`
	Handin_filename  string `json:"handin_filename" mapstructure:"handin_filename"`
	Handin_directory string `json:"handin_directory" mapstructure:"handin_directory"`
	Max_grace_days   int    `json:"max_grace_days" mapstructure:"max_grace_days"`
	Handout          string `json:"handout" mapstructure:"handout"`
	Writeup          string `json:"writeup" mapstructure:"writeup"`
	Max_submissions  int    `json:"max_submissions" mapstructure:"max_submissions"`
	Disable_handins  bool   `json:"disable_handins" mapstructure:"disable_handins"`
	Max_size         int    `json:"max_size" mapstructure:"max_size"`
	Has_svn          bool   `json:"has_svn" mapstructure:"has_svn"`
	Category_name    string `json:"category_name" mapstructure:"category_name"`
	Start_at         string `json:"start_at"  mapstructure:"start_at"`
	Due_at           string `json:"due_at"  mapstructure:"due_at"`
	End_at           string `json:"end_at"  mapstructure:"end_at"`
	Grading_deadline string `json:"grading_deadline"  mapstructure:"grading_deadline"`
}

type Problems struct {
	Name        string `json:"name" mapstructure:"name"`
	Description string `json:"description" mapstructure:"description"`
	Max_score   int    `json:"max_score" mapstructure:"max_score"`
	Optional    bool   `json:"optional" mapstructure:"optional"`
}

type Autograder struct {
	Autograde_timeout int    `mapstructure:"autograde_timeout"`
	Autograde_image   string `mapstructure:"autograde_image"`
	Release_score     bool   ` mapstructure:"release_score"`
}

func (autograder *Autograder) Default() {
	autograder.Autograde_timeout = 180
	autograder.Autograde_image = "autograding_image"
	autograder.Release_score = true
}

func (general *General) Default() {
	general.Handin_filename = "handin.tar"
	general.Handin_directory = "handin"
	general.Max_grace_days = 0
	general.Handout = ""
	general.Writeup = "writeup/exam.html"
	general.Disable_handins = false
	general.Max_size = 2
	general.Has_svn = false
}
