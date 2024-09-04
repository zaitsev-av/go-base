package main

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"math"
)

var (
	red          = color.New(color.FgRed).SprintFunc()
	redBold      = color.New(color.FgRed, color.Bold).SprintFunc()
	success      = color.New(color.FgGreen).SprintFunc()
	green        = color.New(color.FgGreen, color.BgBlack).Add(color.BlinkSlow).SprintFunc()
	yellow       = color.New(color.FgYellow).SprintFunc()
	yellowBoldUl = color.New(color.FgYellow, color.Bold).Add(color.Underline).SprintFunc()
	fgCyan       = color.New(color.FgCyan, color.Bold).SprintFunc()
)

func main() {
	previewResult := preview()
	if previewResult != 1 {
		fmt.Println(yellowBoldUl("Счетчик колорий пока не готов"))
	}
	for {
		_, err := calculateMass()
		if err != nil {
			continue
		}
		isRepeat := checkRepeat()
		if !isRepeat {
			break
		}
	}
}

func preview() uint8 {
	var startOutput uint8
	fmt.Println(yellowBoldUl("Это консольное приложение для выбора различных калькуляторов"))
	fmt.Println(green("Нажмите 1 для того чтобы выбрать Калькулятора массы тела"))
	fmt.Println(green("Нажмите 2 для того чтобы выбрать Счетчик колорий"))
	fmt.Scan(&startOutput)
	return startOutput
}

func checkRepeat() bool {
	var newCalculation string
	fmt.Println(yellow("Хотите произвести расчет еще раз?(yes/no)"))
	fmt.Scan(&newCalculation)
	return newCalculation == "yes" || newCalculation == "y"
}

func calculateMass() (bool, error) {
	height, weight := getUserOutput()
	if height <= 0 || weight <= 0 {
		fmt.Println(redBold("❗️Критическа ошибка🤬"))
		return false, errors.New("NO_CORRECT_INPUT")
	}
	imt := calculateIMT(height, weight)
	outputResult(imt)
	return true, nil
}

func outputResult(imt float64) {
	fmt.Println(printResultIMT(imt))
}

func calculateIMT(height, weight float64) float64 {
	const IMTPower = 2
	return weight / math.Pow(height/100, IMTPower)
}

func getUserOutput() (float64, float64) {
	var height float64
	var weight float64
	fmt.Println(fgCyan("Калькулятор мыссы тела"))
	fmt.Println("Введите свой рост:")

	fmt.Scan(&height)

	fmt.Println("Введите свой вес:")

	fmt.Scan(&weight)

	return height, weight
}

func printResultIMT(imt float64) string {
	switch {
	case imt <= 16.0:
		return fmt.Sprintf(red("У вас сильный дефицит массы тела, ваш индекс массы тела составляет: %.2f"), imt)
	case imt >= 16.0 && imt <= 18.5:
		return fmt.Sprintf(yellow("У вас дефицит массы тела, ваш индекс массы тела составляет %.2f"), imt)
	case imt >= 18.5 && imt <= 25:
		return fmt.Sprintf(success("Вы в норме, ваш индекс массы тела составляет %.2f"), imt)
	case imt >= 25 && imt <= 30:
		return fmt.Sprintf(yellow("У вас избыточная масса тела, ваш индекс массы тела составляет %.2f"), imt)
	case imt >= 30 && imt <= 35:
		return fmt.Sprintf(yellow("У вас 1-я степень ожирения, ваш индекс массы тела составляет %.2f"), imt)
	case imt >= 35 && imt <= 40:
		return fmt.Sprintf(red("У вас 2-я степень ожирения, ваш индекс массы тела составляет %.2f"), imt)
	case imt >= 40 && imt <= 45:
		return fmt.Sprintf(red("У вас 3-я степень ожирения, ваш индекс массы тела составляет %.2f"), imt)
	default:
		return fmt.Sprintf("Скорее всего вы ввели некоректные данные для расчета, но ваш ИМТ равен %.2f 🤔", imt)
	}
}
