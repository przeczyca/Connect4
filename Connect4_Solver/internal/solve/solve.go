package solve

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"
)

func GetBestColumns(position string) string {
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

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for filepath.Base(path) != "Connect4_Solver" && len(path) != 0 {
		path = filepath.Dir(path)
	}

	if filepath.Base(path) != "Connect4_Solver" {
		log.Fatal("Path does not exist")
	}

	cmd := exec.Command(path+"/internal/Magic/c4solver", position+strconv.Itoa(newColumn))
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
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
