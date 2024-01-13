package positionFiles

import (
	"bufio"
	"connect4/Connect4_Solver/internal/solve"
	"errors"
	"log"
	"os"
	"strings"
	"sync"
)

func SetBeginningPositions(positionMap map[string]string) {
	files := [5]string{
		"1stPositions.txt",
		"2ndPositions.txt",
		"3rdPositions.txt",
		"4thPositions.txt",
		"5thPositions.txt"}

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(files); i++ {
		file, err := os.Open(path + "/internal/positionFiles/" + files[i])
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
	file.WriteString(position + " " + solve.GetBestColumns(position) + "\n")
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
