package controllers_test

import (
	"golang-rest-api-test/src/di"
	"golang-rest-api-test/src/testutils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
)

func Test_CoffeesCreate(t *testing.T) {
	req := testutils.NewRequestBuilder().
		WithMethod(http.MethodPost).
		WithTarget("/api/v1/coffees/").
		WithHeader("Content-Type", "application/json").
		WithBody(map[string]any{
			"name":            "キリマンジャロ",
			"comment":         "美味しいいよ",
			"productionArea":  "india",
			"richScore":       1,
			"bitternessScore": 1,
			"sournessScore":   1,
			"aromaScore":      1,
		}).
		MustBuild()

	recorder := httptest.NewRecorder()
	gc, _ := gin.CreateTestContext(recorder)
	gc.Request = req

	c := di.ProvideCoffeeController()
	c.Create(gc)

	testutils.AssertRequestWithOpenAPISpec(t, req, recorder)
}

func Test_CoffeesGet(t *testing.T) {
	req := testutils.NewRequestBuilder().
		WithMethod(http.MethodGet).
		WithTarget("/api/v1/coffees/id").
		WithHeader("Content-Type", "application/json").
		MustBuild()

	recorder := httptest.NewRecorder()
	gc, _ := gin.CreateTestContext(recorder)
	gc.Request = req
	gc.Request.SetPathValue("id", faker.UUIDHyphenated())

	c := di.ProvideCoffeeController()
	c.Get(gc)

	testutils.AssertRequestWithOpenAPISpec(t, req, recorder)
}
