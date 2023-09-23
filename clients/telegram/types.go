package telegram


// TODO: сделать общую  структуру "BaseResponse" и добавлять информацию для запроса отдельно
type UpdateSResponse struct{
	Ok bool `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID int `json:"update_id"`
	Message string `json:"message"`

}