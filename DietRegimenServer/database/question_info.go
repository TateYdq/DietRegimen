package database

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/sirupsen/logrus"
)

type QuestionInfo struct {
	QuestionID int `json:"question_id" gorm:"column:question_id;primary_key"`
	Info string `json:"info"`
	AnswerA string `json:"answer_a"`
	ResponseA string `json:"response_a"`
	AnswerB string `json:"answer_b"`
	ResponseB string `json:"response_b"`
	AnswerC string `json:"answer_c"`
	ResponseC string `json:"response_c"`
}

type AnswerSheet struct {
	ID int `json:"id" gorm:"column:id;primary_key"`
	QuestionID int `json:"question_id"`
	UserID int `json:"user_id"`
	Answer string `json:"answer"`
	RecordTime string `json:"record_time"`
}


func AddQuestion(question QuestionInfo)(id int,err error){
	err = DrDatabase.Create(&question).Error
	if err != nil{
		logrus.WithError(err).Errorf("AddQuestion err")
	}
	return question.QuestionID,err
}


func SelectQuestionInfo(userID int)(questionLists[] QuestionInfo,err error){
	err = DrDatabase.Model(QuestionInfo{}).Scan(&questionLists).Error
	if err != nil{
		logrus.WithError(err).Errorf("SelectQuestionInfo err")
	}
	return questionLists,err
}

func CreateQuestionSheet(request AnswerSheet)(err error){
	request.RecordTime = utils.GetCurTime()
	err = DrDatabase.Create(&request).Error
	if err != nil{
		logrus.WithError(err).Errorf("CreateQuestionSheet err")
	}else{
		logrus.Infof("CreateQuestionSheet success,id:%v",request.ID)
	}
	return err
}

