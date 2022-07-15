package seller

import "encoding/json"

type SellerRequest struct {
	Name   string      `json:"name" form:"name" binding:"required"`
	Gander json.Number `json:"gender" form:"gender" binding:"required,number"`
	ItemID json.Number `json:"itemid" form:"item_id" binding:"required,number"`
}
