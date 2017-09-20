package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"
	"github.com/velocity-ci/velocity/master/velocity/domain"
	"github.com/velocity-ci/velocity/master/velocity/domain/auth"
	"github.com/velocity-ci/velocity/master/velocity/domain/knownhost"
	"github.com/velocity-ci/velocity/master/velocity/domain/project"
	"github.com/velocity-ci/velocity/master/velocity/domain/user"
	"github.com/velocity-ci/velocity/master/velocity/middlewares"
	"github.com/velocity-ci/velocity/master/velocity/persisters"
	"github.com/velocity-ci/velocity/master/velocity/routers"
)

// VelocityAPI - The Velocity API app
type VelocityAPI struct {
	Router *routers.MuxRouter
	server *http.Server
	db     *gorm.DB
	bolt   *bolt.DB
}

// NewVelocityAPI - Returns a new Velocity API app
func NewVelocityAPI() App {
	velocityAPI := &VelocityAPI{}

	controllerLogger := log.New(os.Stdout, "[controller]", log.Lshortfile)
	dbLogger := log.New(os.Stdout, "[database]", log.Lshortfile)
	renderer := render.New()
	validator, translator := middlewares.NewValidator()

	velocityAPI.db = persisters.NewGORMDB(
		dbLogger,
		&domain.User{},
		&domain.Project{},
	)

	velocityAPI.bolt = persisters.NewBoltDB(dbLogger)

	userDBManager := user.NewDBManager(dbLogger, velocityAPI.db)
	projectDBManager := project.NewDBManager(dbLogger, velocityAPI.db)

	projectBoltManager := project.NewBoltManager(dbLogger, velocityAPI.bolt)

	projectManager := project.NewManager(dbLogger, projectDBManager, projectBoltManager)
	knownhostManager := knownhost.NewManager()

	authController := auth.NewController(controllerLogger, renderer, userDBManager)
	projectController := project.NewController(controllerLogger, renderer, projectManager)
	knownhostController := knownhost.NewController(controllerLogger, renderer, validator, translator, knownhostManager)

	velocityAPI.Router = routers.NewMuxRouter([]routers.Routable{
		authController,
		projectController,
		knownhostController,
	}, true)

	port := os.Getenv("PORT")
	velocityAPI.server = &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        velocityAPI.Router.Negroni,
		ReadTimeout:    1 * time.Hour,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return velocityAPI
}

// Stop - Stops the API
func (v *VelocityAPI) Stop() {
	if err := v.server.Shutdown(nil); err != nil {
		panic(err)
	}
	v.db.Close()
}

// Start - Starts the API
func (v *VelocityAPI) Start() {
	if err := v.server.ListenAndServe(); err != nil {
		log.Printf("error: %v", err)
	}
}
