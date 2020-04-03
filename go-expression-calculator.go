package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

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

func startCalc() {
	expression := getExpression()
	result, err := getResult(expression)
	if err != nil {
		fmt.Println(0, err)
		return
	}
	fmt.Printf("%v", result)
}

func getExpression() string {
	var expression string
	for {
		fmt.Printf("Inserisci l'espressione: ")

		n, err := fmt.Scanf("%s\n", &expression)
		if err != nil {
			fmt.Println(n, err)
			continue
		}

		isValid, err := isExpressionValid(expression)
		if !isValid && err != nil {
			fmt.Println(0, err)
			continue
		}
		break
	}
	return expression
}

func isExpressionValid(expr string) (isValid bool, err error) {
	// da finire, accetta -16(5), invece dovrebbe essere -16+(5)
	pattern := "(sqrt)?([\\(]?[0-9]{1,}[\\+\\-\\*\\/\\^0-9]?([\\)]+[\\+\\-\\*\\/\\^\\)]+)?)"
	match, err := regexp.MatchString(pattern, expr)

	// se il match genera errore
	if err != nil {
		return false, err
	}

	// se il pattern non è rispettato
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

func getResult(expr string) (result int, err error) {
	// conterrà l'indice dell'ultima parentesi aperta
	lastOpenBracketIndex := 0

	for i := 0; i < len(expr); i++ {

		if expr[i] == '(' {
			lastOpenBracketIndex = i
		}

		if expr[i] == ')' {

			subExpression := expr[lastOpenBracketIndex+1 : i]

			result, err := compute(subExpression)
			if err != nil {
				fmt.Println(0, err)
				break
			}

			expr = strings.Replace(expr, expr[lastOpenBracketIndex:i+1], result, 1)
		}
	}
	return 100, nil
}

// ((2+4)+(2+5))

func compute(expr string) (result string, err error) {

	// checks for exponent
	for {
		if strings.ContainsAny(expr, "^") {
			index := strings.Index(expr, "^")
			var result int
			for i := 0; i < int(expr[index+1]); i++ {
				result += int(expr[index-1]) * int(expr[index-1])
			}
			expr = strings.Replace(expr, expr[index-1:index+1], string(result), 1)
		} else {
			break
		}
	}

	// compute basics calcs
	for {
		if strings.ContainsAny(expr, "+-*/") {
			plus := strings.Index(expr, "+")
			minus := strings.Index(expr, "-")
			times := strings.Index(expr, "*")
			on := strings.Index(expr, "/")

			if plus != -1 {
				return string(int(expr[plus-1]) + int(expr[plus+1])), nil
			}

			if minus != -1 {
				return string(int(expr[minus-1]) - int(expr[minus+1])), nil
			}

			if times != -1 {
				return string(int(expr[times-1]) * int(expr[times+1])), nil
			}

			if on != -1 {
				return string(int(expr[on-1]) / int(expr[on+1])), nil
			}
		} else {
			return "ERRORE", errors.New("Nessun operazione disponibile")
		}
	}

}

func findOperator(s string) (index int, err error) {
	if strings.ContainsAny(s, "+-*/") {
		plus := strings.Index(s, "+")
		minus := strings.Index(s, "-")
		times := strings.Index(s, "*")
		on := strings.Index(s, "/")

		if plus != -1 {
			return int(s[plus-1]) + int(s[plus+1]), nil
		}

		if minus != -1 {
			return int(s[minus-1]) - int(s[minus+1]), nil
		}

		if times != -1 {
			return int(s[times-1]) * int(s[times+1]), nil
		}

		if on != -1 {
			return int(s[on-1]) / +int(s[on+1]), nil
		}

		return -1, errors.New("GoCalc: errore interno findOperator")

	}

	return -1, errors.New("GoCalc: operatore non trovato")
}
