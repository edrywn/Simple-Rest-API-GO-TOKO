package app

import (
	"github.com/edrywn/toko-online/app/controllers"
	"github.com/edrywn/toko-online/app/item"
)

func (server *Server) initializeRoutes() {

	itemRepository := item.NewRepository(server.DB)
	itemService := item.NewService(itemRepository)
	itemHandler := controllers.NewItemHandler(itemService)

	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
	server.Router.HandleFunc("/items", itemHandler.GetItems).Methods("GET")
	server.Router.HandleFunc("/item", itemHandler.CreateItem).Methods("POST")
	server.Router.HandleFunc("/item/{ID}", itemHandler.GetItem).Methods("GET")
	server.Router.HandleFunc("/item/{ID}", itemHandler.UpdateItem).Methods("PUT")
	server.Router.HandleFunc("/item/{ID}", itemHandler.DeleteItem).Methods("DELETE")
}
