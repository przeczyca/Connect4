package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type connect4 struct {
	Column string `json:"column"`
}

type position struct {
	Position string `json:"position"`
}

var oldPositionMap = map[string]string{
	"":   "4",
	"1":  "4",
	"2":  "3",
	"3":  "4",
	"4":  "4",
	"5":  "4",
	"6":  "5",
	"7":  "4",
	"11": "4",
	"12": "4",
	"13": "3",
	"14": "4",
	"15": "4",
	"16": "6",
	"17": "4",
	"21": "2",
	"22": "5",
	"23": "3",
	"24": "4",
	"25": "4",
	"26": "5",
	"27": "4",
	"31": "4",
	"32": "6",
	"33": "4",
	"34": "4",
	"35": "5",
	"36": "4",
	"37": "4",
	"41": "4",
	"42": "2",
	"43": "6",
	"44": "4",
	"45": "2",
	"46": "6",
	"47": "4",
	"51": "5",
	"52": "4",
	"53": "3",
	"54": "4",
	"55": "4",
	"56": "2",
	"57": "4",
	"61": "4",
	"62": "4",
	"63": "4",
	"64": "4",
	"65": "5",
	"66": "3",
	"67": "3",
	"71": "4",
	"72": "2",
	"73": "4",
	"74": "4",
	"75": "5",
	"76": "4",
	"77": "4",
}

var positionMap = map[string]string{
	"": "4",
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
		bestColumns = getBestColumns(position.Position)
	} else {
		bestColumns = mapOutput
	}

	randomCharIndex := rand.Intn(len(bestColumns))
	output.Column = string(bestColumns[randomCharIndex])

	fmt.Println("Best columns: " + bestColumns)
	fmt.Println("Choosing " + output.Column)

	c.IndentedJSON(http.StatusOK, output)
}

func getBestColumns(position string) string {
	fmt.Println("Solving " + position)
	//map; key: score; value: string
	scoreMap := map[int]string{}
	maxScore := -999999
	scores := []int{-999999, -999999, -999999, -999999, -999999, -999999, -999999}
	var waitGroup sync.WaitGroup
	for i := 0; i < 7; i++ {
		waitGroup.Add(1)
		go runSolver(position, i+1, &scores, &waitGroup)
	}
	waitGroup.Wait()
	for i := 0; i < 7; i++ {
		score := scores[i]
		columnStr, colStrExists := scoreMap[score]
		if !colStrExists {
			scoreMap[score] = strconv.Itoa(i + 1)
		} else {
			scoreMap[score] = columnStr + strconv.Itoa(i+1)
		}

		if score > maxScore {
			maxScore = score
		}
	}

	return scoreMap[maxScore]
}

func runSolver(position string, newColumn int, scores *[]int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	cmd := exec.Command("./Magic/c4solver", position+strconv.Itoa(newColumn))
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Println(fmt.Sprint(err) + ": " + stderr.String())
		log.Fatal(err)
	} else if out.String() == "Invalid Move" {
		log.Printf("Invalid Move: " + position + strconv.Itoa(newColumn))
		return
	}

	score, err := strconv.Atoi(out.String())
	if err != nil {
		log.Fatal(err)
	}

	(*scores)[newColumn-1] = score
}

func main() {
	//createPositionsFile(4, "fourthPositions1.txt")
	//newFileBasedOnOldFile("5", "sixthPositions5.txt", "fullFifthPositions.txt")
	setBeginningPositions()
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/getOneBestMove", getOneBestMove)
	router.Run(":8080")
}

func setBeginningPositions() {
	files := [5]string{
		"1stPositions.txt",
		"2ndPositions.txt",
		"3rdPositions.txt",
		"4thPositions.txt",
		"5thPositions.txt"}
	for i := 0; i < len(files); i++ {
		file, err := os.Open("./positionFiles/" + files[i])
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := strings.Split(scanner.Text(), " ")
			positionMap[line[0]] = line[1]
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		file.Close()
	}
}

func createPositionsFile(positionSize int, fileName string) {
	_, err := os.Stat(fileName)
	if !errors.Is(err, os.ErrNotExist) {
		deleteErr := os.Remove(fileName)
		if deleteErr != nil {
			log.Fatal(err)
		}
	}
	file, err := os.Create(fileName)

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	columns := []string{"1", "2", "3", "4", "5", "6", "7"}
	var positions []string
	getCombinations(columns, "1", positionSize, &positions)

	var parentWaitGroup sync.WaitGroup
	for limit := 7; limit < len(positions)+7; limit += 7 {
		for i := limit - 7; i < len(positions) && i < limit; i++ {
			position := positions[i]
			parentWaitGroup.Add(1)
			go addBestColumnsToFile(file, &parentWaitGroup, position)
		}
		parentWaitGroup.Wait()
	}
}

func getCombinations(columns []string, position string, length int, positions *[]string) {
	if len(position) == length {
		if position == "111111" || position == "222222" || position == "333333" || position == "444444" || position == "555555" || position == "666666" {
			return
		}
		*positions = append(*positions, position)
	} else {
		for i := 0; i < len(columns); i++ {
			getCombinations(columns, position+columns[i], length, positions)
		}
	}
}

func addBestColumnsToFile(file *os.File, parentWaitGroup *sync.WaitGroup, position string) {
	defer parentWaitGroup.Done()
	file.WriteString(position + " " + getBestColumns(position) + "\n")
}

func newFileBasedOnOldFile(pre string, newFileName string, oldFileName string) {
	oldFile, err := os.Open(oldFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer oldFile.Close()

	_, err = os.Stat(newFileName)
	if !errors.Is(err, os.ErrNotExist) {
		deleteErr := os.Remove(newFileName)
		if deleteErr != nil {
			log.Fatal(err)
		}
	}

	newFile, err := os.Create(newFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer newFile.Close()

	scanner := bufio.NewScanner(oldFile)
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	var postPositions [16807]string
	index := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		postPositions[index] = line[0]
		index++
	}

	var parentWaitGroup sync.WaitGroup
	for limit := 10; limit < len(postPositions)+10; limit += 10 {
		for i := limit - 10; i < len(postPositions) && i < limit; i++ {
			position := postPositions[i]
			parentWaitGroup.Add(1)
			go addBestColumnsToFile(newFile, &parentWaitGroup, pre+position)
		}
		parentWaitGroup.Wait()
	}
}
