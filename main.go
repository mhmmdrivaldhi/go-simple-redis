package main

import (
	"encoding/json"
	"fmt"
	"go-redis/db"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ResponJSON struct {
	Data interface{} `json:"data"`
	Status string	 `json:"status"`
}

type RequestRedis struct{
	Name string `json:"name"`
	Age string	`json:"age"`
}

var key = "gored_intro_app"

func main() {
	// initialize redis
	db.RedisInit()

	// initialize echo web service
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/insert", insert)
	e.GET("/get-redis", getRedis)

	e.Logger.Fatal(e.Start(":8080"))
}

func insert(c echo.Context) error {
	id := c.QueryParam("id")
	name := c.QueryParam("name")
	age := c.QueryParam("age")

	if id == "" || name == "" || age == "" {
		return c.JSON(http.StatusBadRequest, ResponJSON{
			Data: nil,
			Status: "missing parameter(s)",
		})
	}

	rdb := db.RedisConnect()

	reqRedis := RequestRedis{
		Name: name,
		Age: age,
	}
	req, _ := json.Marshal(reqRedis)

	err := rdb.HSet(c.Request().Context(), key, id, req).Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponJSON{
			Data: nil,
			Status: fmt.Sprintf("Error set redis %s", err),
		})
	}
	return c.JSON(http.StatusOK, ResponJSON{
		Data: id,
		Status: "Successfully Inserted Data to redis",
	})
}

func getRedis(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, ResponJSON{
			Data: nil,
			Status: "Missing id parameter",
		})
	}

	rdb := db.RedisConnect()

	val, err := rdb.HGet(c.Request().Context(), key, id).Result()
	if err != nil {
		return c.JSON(http.StatusNotFound, ResponJSON{
			Data: nil,
			Status: fmt.Sprintf("Data Not Found or Error %s", err),
		})
	}

	var request RequestRedis
	err = json.Unmarshal([]byte(val), &request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponJSON{
			Data: nil,
			Status: "failed to parse data from Redis",
		})
	}

	return c.JSON(http.StatusOK, ResponJSON{
		Data: request,
		Status: "Successfully retrieved data from redis",
	})
}