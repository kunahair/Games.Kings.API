package viewmodel

type PlayerViewModel struct {
	Name	string	`json:"name" binding:"required"`
}
