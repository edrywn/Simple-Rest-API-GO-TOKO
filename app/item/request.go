package item

import "encoding/json"

type ItemRequest struct {
	Name  string      `json:"name" form:"name" binding:"required"`
	Price json.Number `json:"price" form:"price" binding:"required,number"`
	Stock json.Number `json:"stock" form:"stock" binding:"required,number"`
}
