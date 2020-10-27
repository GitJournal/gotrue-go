package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"context"
	"fmt"

	"github.com/gobuffalo/uuid"
	"github.com/netlify/gotrue/api"
	"github.com/netlify/gotrue/conf"
	"github.com/netlify/gotrue/storage"
	"github.com/sirupsen/logrus"
)

type GoTrueAPI struct {
	api *api.API
}

func (a *GoTrueAPI) Serve(globalConfig *conf.GlobalConfiguration, config *conf.Configuration) {
	db, err := storage.Dial(globalConfig)
	if err != nil {
		logrus.Fatalf("Error opening database: %+v", err)
	}
	defer db.Close()

	ctx, err := api.WithInstanceConfig(context.Background(), config, uuid.Nil)
	if err != nil {
		logrus.Fatalf("Error loading instance config: %+v", err)
	}
	a.api = api.NewAPIWithVersion(ctx, globalConfig, db, "version")

	l := fmt.Sprintf("%v:%v", globalConfig.API.Host, globalConfig.API.Port)
	logrus.Infof("GoTrue API started on: %s", l)
	a.api.ListenAndServe(l)
}

func (a *GoTrueAPI) Settings(r *http.Request) (*api.Settings, error) {
	reader := strings.NewReader("")
	req := httptest.NewRequest(http.MethodGet, "/settings", reader)
	res := httptest.NewRecorder()

	err := a.api.Settings(res, req)
	if err != nil {
		return nil, err
	}

	var settings api.Settings
	err = json.Unmarshal(res.Body.Bytes(), &settings)
	if err != nil {
		return nil, err
	}

	return &settings, nil
}
