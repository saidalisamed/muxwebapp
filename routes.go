package main

// Setup routes
func (a *App) configureRoutes() {
	// Product routes
	productRoutes := a.Router.PathPrefix("/product").Subrouter()
	productRoutes.HandleFunc("/all", a.getProducts).Methods("GET")
	productRoutes.HandleFunc("/create", a.createProduct).Methods("POST")
	productRoutes.HandleFunc("/{id:[0-9]+}", a.getProduct).Methods("GET")
	productRoutes.HandleFunc("/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	productRoutes.HandleFunc("/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")

	// Session routes
	sessionRoutes := a.Router.PathPrefix("/session").Subrouter()
	sessionRoutes.HandleFunc("/set", a.sessionSet).Methods("GET")
	sessionRoutes.HandleFunc("/get", a.sessionGet).Methods("GET")

	a.Router.HandleFunc("/template", a.templateDemo).Methods("GET")
}
