package database


type UserFoodRelation struct {
	UserID int `json:"user_id"`
	FoodID int `json:"food_id"`
	FoodName string `json:"food_name"`
	Score int `json:"score"`
}