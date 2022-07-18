package course

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/gin-gonic/gin"
	cp "github.com/otiai10/copy"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var (
	Template_path  = "./source/template"
	Autograde_path = "./source/autograders"
)

func Build_Assessment(c *gin.Context, course string, assessment models.Download_Assessments) (tar_path string) {
	assessment_name := assessment.General.Name
	exam_path := utils.Find_folder(c, "exam", course, assessment_name)

	// copy template assessment project and modify information
	copy_template(c, exam_path)
	replace_template(c, exam_path, assessment_name, assessment_name)
	modify_yml(c, exam_path, assessment)
	copy_autograders(exam_path, assessment_name)
	make_tar(c, exam_path, assessment_name, assessment_name)

	tar_path = filepath.Join(exam_path, assessment_name+".tar")
	utils.FileCheck(tar_path)

	return
}

func copy_template(c *gin.Context, path string) {
	// delete current folder
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		os.RemoveAll(path)
	}

	// copy a new folder
	err := cp.Copy(Template_path, path)

	if err != nil {
		response.ErrAssessmentInternaldResponse(c, "There is an error when building assessment!")
	}
}

func replace_template(c *gin.Context, path, name, display_name string) {
	prog := filepath.Join(path, "replace.sh")
	run_exec(c, prog, name)
}

func modify_yml(c *gin.Context, pro_path string, assessment models.Download_Assessments) {
	assessment_name := assessment.General.Name
	yml_path := filepath.Join(pro_path, assessment_name+"/"+assessment_name+".yml")
	utils.FileCheck(yml_path)

	ass := models.Download_Assessments{}
	v := viper.New()
	v.SetConfigFile(yml_path)
	if err := v.ReadInConfig(); err != nil {
		response.ErrAssessmentInternaldResponse(c, "There is an error when building assessment!")
	}
	if err := v.Unmarshal(&ass); err != nil {
		response.ErrAssessmentInternaldResponse(c, "There is an error when building assessment!")
	}
	ass.General.Name = assessment_name
	ass.General.Display_name = assessment_name

	data, err := yaml.Marshal(&ass)
	if err != nil {
		response.ErrAssessmentInternaldResponse(c, "There is an error when building assessment!")
	}

	err2 := ioutil.WriteFile(yml_path, data, 0)
	if err2 != nil {
		response.ErrAssessmentInternaldResponse(c, "There is an error when building assessment!")
	}

	fmt.Println("data written")
}

func make_tar(c *gin.Context, path, name, display_name string) {
	prog := filepath.Join(path, "make.sh")
	run_exec(c, prog, name)
}

func run_exec(c *gin.Context, prog, arg1 string) {
	cmd := exec.Command(prog, arg1)
	_, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func copy_autograders(path, assessment string) {
	path = filepath.Join(path, assessment)
	path = filepath.Join(path, "autograder")
	path = filepath.Join(path, "autograders")

	// TODO should copy autograders based on configurations
	utils.Copy_file("multiple_blank.py", Autograde_path, path)
	utils.Copy_file("multiple_choice.py", Autograde_path, path)
	utils.Copy_file("single_blank.py", Autograde_path, path)
	utils.Copy_file("single_choice.py", Autograde_path, path)

}
