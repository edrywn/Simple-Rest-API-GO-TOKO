package seller

type ItemResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender int    `json:"gender"`
	ItemID int    `json:"item_id"`
}
