package application

func (w *webApp) mapRoutes() {
	w.webRouter.GET("/users/:id", w.UserService.Recive)
	w.webRouter.GET("/users", w.UserService.List)
	w.webRouter.GET("/user/find/", w.UserService.FindAll)
	w.webRouter.POST("/users", w.UserService.Create)
	w.webRouter.PUT("/users/:id", w.UserService.Update)
	w.webRouter.DELETE("/users/:id", w.UserService.Delete)
}
