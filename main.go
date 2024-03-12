package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// greetUser возвращает приветствие в зависимости от времени суток в таймзоне пользователя
func greetUser(params RequestParams) (string, error) {
	location, err := time.LoadLocation(params.Timezone)
	if err != nil {
		return "Ошибка при загрузке таймзоне", err
	}

	now := time.Now().In(location)

	var greeting string
	switch hour := now.Hour(); {
	case hour < 12:
		greeting = "Доброе утро"
	case hour < 18:
		greeting = "Добрый день"
	default:
		greeting = "Добрый вечер"
	}

	return fmt.Sprintf("%s, %s!", greeting, params.Name), nil
}

type RequestParams struct {
	Name     string `form:"name" binding:"required"`
	Timezone string `form:"timezone" binding:"required"`
}

func main() {
	router := gin.Default()

	router.GET("/greet", func(c *gin.Context) {
		var params RequestParams
		if err := c.BindQuery(&params); err != nil {
			c.String(400, "Некорректные параметры запроса")
			return
		}

		greeting, err := greetUser(params)
		if err != nil {
			c.String(400, greeting)
			return
		}
		c.String(200, greeting)
	})

	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
}
