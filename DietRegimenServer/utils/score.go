package utils

import "time"

const(
	ScoreFoodLook = 1
	ScoreFoodComment = 5
	ScoreFoodCollect = 20

	ScoreDiseaseLook = 1
	ScoreDiseaseComment = 5
	ScoreDiseaseCollect = 15


	ScoreUserLogin = 10
	ScoreUserLook = 1
	ScoreUserComment = 3

)


const(
	LimitUserLoginAwardDay = 1
	LimitUserLookAwardDay = 30
	LimitUserCommentAwardDay = 15
)

const(
	DurationUserScore = time.Hour * 24
)
