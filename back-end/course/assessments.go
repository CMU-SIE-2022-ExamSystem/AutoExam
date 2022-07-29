package course

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/gin-gonic/gin"
	cp "github.com/otiai10/copy"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var (
	Template_path   = "./source/template"
	Basicgrade_path = "./source/autograders"
)

func Build_Assessment(c *gin.Context, course, base_course string, autoexam dao.AutoExam_Assessments) (tar_path string) {
	assessment_name := autoexam.General.Name
	exam_path := utils.Find_assessment_folder(c, "exam", course, assessment_name)

	// copy template assessment project and modify information
	copy_template(c, exam_path)
	replace_template(c, exam_path, assessment_name, autoexam.General.Url)
	modify_yml(c, exam_path, autoexam.ToDownloadAssessments())
	copy_autograders(exam_path, base_course, assessment_name)
	make_tar(c, exam_path, assessment_name)

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

// replace template file and writeup
func replace_template(c *gin.Context, path, name, url string) {
	prog := filepath.Join(path, "replace.sh")
	run_exec(c, prog, name, url)
}

func modify_yml(c *gin.Context, pro_path string, assessment models.Download_Assessments) {
	assessment_name := assessment.General.Name
	yml_path := filepath.Join(pro_path, assessment_name+"/"+assessment_name+".yml")
	utils.FileCheck(yml_path)

	// ass := models.Download_Assessments{}
	v := viper.New()
	v.SetConfigFile(yml_path)

	data, err := yaml.Marshal(&assessment)
	if err != nil {
		response.ErrAssessmentInternaldResponse(c, "There is an error when building assessment!")
	}

	err2 := ioutil.WriteFile(yml_path, data, 0)
	if err2 != nil {
		response.ErrAssessmentInternaldResponse(c, "There is an error when building assessment!")
	}

	fmt.Println("data written")
}

func make_tar(c *gin.Context, path, name string) {
	prog := filepath.Join(path, "make.sh")
	run_exec(c, prog, name, "")
}

func run_exec(c *gin.Context, prog, arg1, arg2 string) {
	var cmd *exec.Cmd
	if arg2 == "" {
		cmd = exec.Command(prog, arg1)
	} else {
		cmd = exec.Command(prog, arg1, arg2)
	}

	_, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func copy_autograders(path, base_course, assessment_name string) {
	path = filepath.Join(path, assessment_name)
	path = filepath.Join(path, "autograder")
	path = filepath.Join(path, "autograders")

	// copy basic grader
	for _, grader := range global.Settings.Basic_Grader {
		utils.Copy_file(grader+".py", Basicgrade_path, path)
	}

	db_path := filepath.Join(dao.DBgrade_path, base_course)
	graders, _ := dao.ReadAllGraders(base_course)
	for _, grader := range graders {
		file_name := grader.Name + ".py"
		file_path := filepath.Join(db_path, file_name)
		if _, err := os.Stat(file_path); errors.Is(err, os.ErrNotExist) {
			dao.Storegrader(grader, base_course)
		}
		utils.Copy_file(file_name, db_path, path)
	}
}
