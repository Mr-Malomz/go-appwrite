package api

import "github.com/gin-gonic/gin"

type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes() {
	app.Router.POST("/project", app.createdProjectHandler())
	app.Router.GET("/project/:projectId", app.getProjectHandler())
	app.Router.PATCH("/project/:projectId", app.updateProjectHandler())
	app.Router.DELETE("/project/:projectId", app.deleteProjectHandler())
}
