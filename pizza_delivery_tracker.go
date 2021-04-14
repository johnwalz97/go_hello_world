package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	X int
	Y int
}

type delivery struct {
	position coordinate
	agent    int
}

func preProcessMoves(movesString string) string {
	movesString = strings.TrimSpace(movesString)
	movesString = strings.Trim(movesString, "\n")

	return movesString
}

func validateMoves(movesString string) {
	for i, move := range movesString {
		if move != 60 && move != 62 && move != 94 && move != 118 {
			err := fmt.Sprintf("Invalid move detected (%s) at position %d in input string!", string(move), i)
			panic(err)
		}
	}
}

func vectorizeMoves(movesString string) []coordinate {
	movesMap := map[int32]coordinate{
		94:  coordinate{0, 1},
		118: coordinate{0, -1},
		62:  coordinate{1, 0},
		60:  coordinate{-1, 0},
	}

	movesList := make([]coordinate, len(movesString))

	for i, move := range movesString {
		movesList[i] = movesMap[move]
	}

	return movesList
}

func trackDeliveries(moves []coordinate, numAgents int) ([]delivery, float64) {
	agentPositions := make([]coordinate, numAgents)
	currentAgent := 0

	deliveries := make([]delivery, len(moves)+1)
	deliveries[0] = delivery{coordinate{0, 0}, 0}

	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	for i, move := range moves {
		agentPositions[currentAgent].X += move.X
		agentPositions[currentAgent].Y += move.Y

		deliveries[i+1] = delivery{agentPositions[currentAgent], currentAgent}

		if agentPositions[currentAgent].X < minX {
			minX = agentPositions[currentAgent].X
		} else if agentPositions[currentAgent].X > maxX {
			maxX = agentPositions[currentAgent].X
		}
		if agentPositions[currentAgent].Y < minY {
			minY = agentPositions[currentAgent].Y
		} else if agentPositions[currentAgent].Y > maxY {
			maxY = agentPositions[currentAgent].Y
		}

		if currentAgent == numAgents-1 {
			currentAgent = 0
		} else {
			currentAgent += 1
		}
	}

	gridSize := math.Sqrt(math.Pow(float64(maxX-minX), 2) + math.Pow(float64(maxY-minY), 2))

	return deliveries, gridSize
}

func getUniqueDeliveries(deliveries []delivery) []delivery {
	var uniqueDeliveries []delivery
	keys := make(map[string]bool)

	for _, delivery := range deliveries {
		deliveryKey := fmt.Sprintf("%d-%d", delivery.position.X, delivery.position.Y)
		if _, value := keys[deliveryKey]; !value {
			keys[deliveryKey] = true
			uniqueDeliveries = append(uniqueDeliveries, delivery)
		}
	}

	return uniqueDeliveries
}

func processMoves(movesString string, numAgents int) {
	movesString = preProcessMoves(movesString)
	validateMoves(movesString)

	moves := vectorizeMoves(movesString)

	deliveries, gridSize := trackDeliveries(moves, numAgents)

	fmt.Printf("\nTotal number of unique deliveries is %d\n", len(getUniqueDeliveries(deliveries)))
	fmt.Printf("\nGrid diagonal length is %f\n", gridSize)
}

func main() {
	filePath := os.Args[1]
	numAgents, _ := strconv.Atoi(os.Args[2])

	data, _ := ioutil.ReadFile(filePath)

	inputString := string(data)

	processMoves(inputString, numAgents)
}
