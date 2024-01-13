package routes

import (
	"fmt"
	"math/rand"
	"net/http"

	"connect4/Connect4_Solver/internal/positionFiles"
	"connect4/Connect4_Solver/internal/solve"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type connect4 struct {
	Column string `json:"column"`
}

type position struct {
	Position string `json:"position"`
}

var positionMap = map[string]string{
	"": "4",
}

func NewRouter() {
	positionFiles.SetBeginningPositions()
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/getOneBestMove", getOneBestMove)
	router.Run(":8080")
}

func getOneBestMove(c *gin.Context) {
	var position position
	var output connect4

	if apiErr := c.BindJSON(&position); apiErr != nil {
		return
	}

	fmt.Println("Solving " + position.Position)

	mapOutput, exists := positionMap[position.Position]
	bestColumns := ""

	if !exists {
		bestColumns = solve.GetBestColumns(position.Position)
	} else {
		bestColumns = mapOutput
	}

	randomCharIndex := rand.Intn(len(bestColumns))
	output.Column = string(bestColumns[randomCharIndex])

	fmt.Println("Best columns: " + bestColumns)
	fmt.Println("Choosing " + output.Column)

	c.IndentedJSON(http.StatusOK, output)
}
