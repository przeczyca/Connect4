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

	if len(position) < 10 {
		var waitGroup sync.WaitGroup
		for i := 0; i < 7; i++ {
			waitGroup.Add(1)
			go runSolver(position, i+1, &scores, &waitGroup)
		}
		waitGroup.Wait()
	} else {
		for i := 0; i < 7; i++ {
			runSolver(position, i+1, &scores, nil)
		}
	}

	for i := 0; i < 7; i++ {
		score := scores[i]
		columnStr, colStrExists := scoreMap[score]
		if !colStrExists {
			scoreMap[score] = strconv.Itoa(i + 1)
		} else {
			scoreMap[score] = columnStr + strconv.Itoa(i + 1)
		}

		if score > maxScore {
			maxScore = score
		}
	}

	return scoreMap[maxScore]
}

func runSolver(position string, newColumn int, scores *[]int, waitGroup *sync.WaitGroup) {
	if waitGroup != nil {
		defer waitGroup.Done()
	}

	var path = getPathToSolver()

	cmd := exec.Command(path, position+strconv.Itoa(newColumn))
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	var err = cmd.Run()
	if err != nil {
		log.Println(fmt.Sprint(err) + ": " + stderr.String())
		log.Fatal(err)
	} else if out.String() == "Invalid Move" {
		return
	}

	score, err := strconv.Atoi(out.String())
	if err != nil {
		log.Fatal(err)
	}

	(*scores)[newColumn-1] = score
}

func getPathToSolver() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for filepath.Base(path) != "Connect4_Solver" && filepath.Base(path) != "app" && path != "/" {
		path = filepath.Dir(path)
	}

	if filepath.Base(path) != "Connect4_Solver" && filepath.Base(path) != "app" {
		log.Fatal("path to /c4solver/ does not exist")
	}

	return path + "/internal/Magic/c4solver"
}