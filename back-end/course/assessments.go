package course

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	cp "github.com/otiai10/copy"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type Assessment_body struct {
	Name             string `json:"name" mapstructure:"name"`
	Start_at         string `json:"start_at"  mapstructure:"start_at"`
	Due_at           string `json:"due_at"  mapstructure:"due_at"`
	End_at           string `json:"end_at"  mapstructure:"end_at"`
	Grading_deadline string `json:"grading_deadline"  mapstructure:"grading_deadline"`
	Category_name    string `json:"category_name" mapstructure:"category_name"`
}

type Assessment struct {
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

func (autograder *Autograder) fill_defaults() {

	if autograder.Autograde_timeout == 0 {
		autograder.Autograde_timeout = 180
	}

	if autograder.Autograde_image == "" {
		autograder.Autograde_image = "autograding_image"
	}
}

var (
	Template_path = "./source/template"
)

func Build_Assessment(course, assessment string) (tar_path string) {
	// build course folder
	coures_path := Build_Course(course)

	// create assessment folder
	ass_path := filepath.Join(coures_path, assessment)
	utils.CreateFolder(ass_path)

	pro_path := filepath.Join(ass_path, "template")
	// copy template assessment project and modify information
	copy_template(pro_path)
	replace_template(pro_path, assessment, assessment)
	modify_yml(pro_path, assessment)
	make_tar(pro_path, assessment, assessment)

	tar_path = filepath.Join(pro_path, assessment+".tar")
	utils.FileCheck(tar_path)
	return
}

func copy_template(path string) {
	// delete current folder
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		fmt.Println("===================")
		os.RemoveAll(path)
		fmt.Println("===================")
	}

	// copy a new folder
	err := cp.Copy(Template_path, path)
	if err != nil {
		panic(err)
	}
}

func replace_template(path, name, display_name string) {
	prog := filepath.Join(path, "replace.sh")
	run_exec(prog, name)
}

func modify_yml(pro_path, assessment string) {
	yml_path := filepath.Join(pro_path, assessment+"/"+assessment+".yml")
	utils.FileCheck(yml_path)

	ass := Assessment{}
	v := viper.New()
	v.SetConfigFile(yml_path)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&ass); err != nil {
		panic(err)
	}
	ass.General.Name = assessment
	ass.General.Display_name = assessment

	data, err := yaml.Marshal(&ass)
	if err != nil {
		panic(err)
	}

	err2 := ioutil.WriteFile(yml_path, data, 0)

	if err2 != nil {
		panic(err2)
	}

	fmt.Println("data written")
}

func make_tar(path, name, display_name string) {
	prog := filepath.Join(path, "make.sh")
	run_exec(prog, name)
}

func run_exec(prog, arg1 string) {
	cmd := exec.Command(prog, arg1)
	_, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func Download_Assessment() {

}
