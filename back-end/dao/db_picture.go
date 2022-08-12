package dao

import (
	"math/rand"
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"gorm.io/gorm"
)

var (
	maxIDLength = 32
)

type PictureFile struct {
	gorm.Model
	ID         string `gorm:"primaryKey, type:varchar(32)"`
	BaseCourse string `gorm:"index:base_course, type:varchar(64)"`
	Type       string `gorm:"type:varchar(32)"` // eg "jpg" "png"
	Content    []byte `gorm:"type:longblob"`
}

func getRandStr() string {

	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	// length is [1, 32]
	length := rand.Intn(maxIDLength) + 1
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func InsertOnePicture(base_course string, pictureType string, content []byte) (string, error) {
	var instanceId string
	// make sure new id doesn't exist in the table
	for {
		instanceId = getRandStr()
		var instance PictureFile
		rows := global.DB.Where(&PictureFile{ID: instanceId}).Find(&instance)
		if rows.RowsAffected < 1 {
			// this id is not used
			break
		}
	}

	new_instance := PictureFile{ID: instanceId, BaseCourse: base_course, Type: pictureType, Content: content}
	if err := global.DB.Create(&new_instance).Error; err != nil {
		return "", err
	}
	return instanceId, nil
}

func UpadateOnePicture(baseCourse string, pictureID string, pictureType string, content []byte) (bool, error) {
	var new_data PictureFile
	rows := global.DB.Where(&PictureFile{BaseCourse: baseCourse, ID: pictureID}).Find(&new_data)
	if rows.RowsAffected < 1 {
		return false, nil
	}
	if err := global.DB.Model(new(PictureFile)).Where("id=?", pictureID).Updates(PictureFile{Type: pictureType, Content: content}).Error; err != nil {
		return true, err
	}

	return true, nil
}

func DeletePicture(baseCourse string, pictureID string) (bool, error) {
	var new_data PictureFile
	rows := global.DB.Where(&PictureFile{BaseCourse: baseCourse, ID: pictureID}).Find(&new_data)
	if rows.RowsAffected < 1 {
		return false, nil
	}
	if err := global.DB.Where(&PictureFile{ID: pictureID}).Delete(&PictureFile{}).Error; err != nil {
		return true, err
	}
	return true, nil
}

func SearchOnePictureBasedOnID(pictureID string) (bool, string, []byte, error) {
	var new_data PictureFile
	rows := global.DB.Where(&PictureFile{ID: pictureID}).Find(&new_data)
	if rows.RowsAffected < 1 {
		return false, "", nil, rows.Error
	}
	return true, new_data.Type, new_data.Content, rows.Error
}

func SearchCoursePictureIDs(base_course string) (bool, []string, error) {
	var instances []PictureFile
	result := global.DB.Where(&PictureFile{BaseCourse: base_course}).Find(&instances)
	if len(instances) == 0 {
		return false, nil, result.Error
	}
	var ids []string
	for _, instance := range instances {
		ids = append(ids, instance.ID)
	}
	return true, ids, result.Error
}
