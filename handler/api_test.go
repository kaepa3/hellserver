package handler

import (
	"fmt"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"

	"github.com/kaepa3/hellserver/model"
	"github.com/labstack/echo/v4"
)

func TestAddTrain(t *testing.T) {
	e := echo.New()
	e.POST("/train", AddTrain)
	server := httptest.NewServer(e.Server.Handler)
	t.Cleanup(func() { server.Close() })

	train := model.Train{}
	restyClient := resty.New().SetBaseURL(server.URL)
	restyClient.R().SetHeaders(map[string]string{"Content-Type": "application/json"})

	res, err := restyClient.NewRequest().SetResult(&train).Get(fmt.Sprintf("/train"))
	if err != nil {
		t.Error(err)
	}
	log.Println(res)
}
