package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func isExpressionValid(expr string) (isValid bool, err error) {
	// pattern := "/(sqrt)*[0-9]*[\\+\\-\\*\\/\\^\\(\\)]*/g"
	pattern := "/(sqrt)*[0-9]+((\\+|\\-|\\*|\\/|\\^)[0-9])*/g"
	match, _ := regexp.MatchString(pattern, expr)
	if !match {
		return false, errors.New("GoCalc: expression does not match")
	}

	// se le parentesi aperte sono quante quelle chiuse
	if strings.Count(expr, "(") != strings.Count(expr, ")") {
		return false, errors.New("GoCalc: discordant parenthesis number")
	}

	//TODO:
	return true, nil
}

func getExpression() string {
	var expression string
	for {
		fmt.Printf("Inserisci l'espressione: ")

		n, err := fmt.Scanf("%s\n", &expression)
		if err != nil {
			// handle invalid input
			fmt.Println(n, err)
			continue
		}

		isValid, err := isExpressionValid(expression)
		if isValid {
			// handle invalid expression
			fmt.Println("Espressione non valida")
			continue
		}
		break
	}
	return expression
}

func getResult(expr string) (result int, err error) {
	return 0, nil
}

func startCalc() {
	expression := getExpression()
	result, err := getResult(expression)
	if err != nil {
		fmt.Printf("%d\n", result)
		log.Fatal(err)
	}
}

func main() {
	fmt.Printf("##############################################\n")
	fmt.Printf("# Ciao, benvenuto in GoCalc by Damiano Imola #\n")
	fmt.Printf("# Informazioni generali                      #\n")
	fmt.Printf("# Operatori consentiti: '+', '-', '*' e '/'  #\n")
	fmt.Printf("# Radice quadrata: sqrt(expression)          #\n")
	fmt.Printf("# Elevazione a potenza: '^'+esponente (n^2)  #\n")
	fmt.Printf("##############################################\n\n\n")
	startCalc()
}
