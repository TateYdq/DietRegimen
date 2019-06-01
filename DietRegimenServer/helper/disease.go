package helper

import (
	"github.com/gin-gonic/gin"
	"errors"
	"strconv"
)

func GetDiseaseID(c *gin.Context)(diseaseID int,err error){
	dID := c.Query("disease_id")
	if dID == "" {
		return 0,errors.New("get disease_id error")
	}
	diseaseID,err = strconv.Atoi(dID)
	return diseaseID,err
}
