package database

type FoodComment struct{
	ID int `json:"id"`
	FoodID int `json:"food_id"`
	UserID int `json:"user_id"`
	UserName string `json:"user_name"`
	Comment string `json:"comment"`
	RecordTime string `json:"record_time"`
}


