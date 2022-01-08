/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package controller

import (
	"github.com/gin-gonic/gin"
)

// Controller holds the Controller data.
type Controller struct {
}

// NewController returns new controller.
func NewController() (*Controller, error) {
	return &Controller{}, nil
}

// RegisterRoutes registers all the helath API routes.
func (ctl *Controller) RegisterRoutes() *gin.Engine {
	router := gin.Default()

	sr := router.Group("/api/v1/consumer_report_user_health")

	// Routes.
	sr.POST("/health", ctl.Health)
	sr.POST("/shutdown", ctl.Shutdown)

	return router
}
