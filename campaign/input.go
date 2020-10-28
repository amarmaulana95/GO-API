package campaign

type GetCampainDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
