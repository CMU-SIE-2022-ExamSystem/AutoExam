package course

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

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
	modules, err := copy_autograders(exam_path, base_course, assessment_name)
	if err != nil {
		response.ErrAssessmentGenerateResponse(c, course, err.Error())
	}
	add_modules(c, exam_path, assessment_name, modules)
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
	url = strings.ReplaceAll(url, "/", "\\/")
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

	// fmt.Println("data written")
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
		response.ErrAssessmentGenerateResponse(c, arg1, err.Error())
		return
	}
}

func copy_autograders(path, base_course, assessment_name string) ([]string, error) {
	path = filepath.Join(path, assessment_name)
	path = filepath.Join(path, "autograder")
	path = filepath.Join(path, "autograders")

	// copy basic grader
	for _, grader := range global.Settings.Basic_Grader {
		utils.Copy_file(grader+".py", Basicgrade_path, path)
	}

	db_path := filepath.Join(dao.DBgrade_path, base_course)
	graders, _ := dao.ReadAllGraders(base_course)

	var modules []string
	for _, grader := range graders {
		file_name := grader.Name + ".py"
		file_path := filepath.Join(db_path, file_name)
		if _, err := os.Stat(file_path); os.IsNotExist(err) {
			err := dao.Storegrader(grader, base_course)
			if err != nil {
				return []string{}, err
			}
		}
		modules = append(modules, grader.Modules...)
		utils.Copy_file(file_name, db_path, path)
	}
	return modules, nil
}

func add_modules(c *gin.Context, path, assessment_name string, modules []string) {
	autograder_path := filepath.Join(path, assessment_name)
	autograder_path = filepath.Join(autograder_path, "autograder")
	requirement_path := filepath.Join(autograder_path, "requirements.txt")
	utils.FileCheckWithC(c, requirement_path)
	f, err := os.OpenFile(requirement_path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		response.ErrFileNotValidResponse(c)
	}
	defer f.Close()
	for _, module := range modules {
		f.WriteString("\n" + module)
	}
}

func GetAssessment(c *gin.Context, course_name, assessment_name string) dao.AutoExam_Assessments {
	// read certain assessment
	assessment, err := dao.ReadExam(course_name, assessment_name)

	// check mongo error
	if err != nil {
		response.ErrMongoDBReadResponse(c, Student_Model)
	}
	return assessment
}

func CheckAssessmentTime(c *gin.Context) {
	course_name, assessment_name := GetCourseAssessment(c)
	assessment := GetAssessment(c, course_name, assessment_name)
	now := time.Now()
	if !now.After(dao.Time_str_convert(assessment.General.Start_at)) {
		response.ErrAssessmentBeforeStartAtResponse(c, assessment_name)
	}
}
