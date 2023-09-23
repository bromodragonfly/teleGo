package telegram

// TODO: create stucrure "basePath" and add information to req separately
type UpdateSResponse struct{
	Ok bool `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID int `json:"update_id"`
	Message string `json:"message"`

}