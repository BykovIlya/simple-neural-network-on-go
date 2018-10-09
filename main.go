package main

import (
	"github.com/fxsjy/gonn/gonn"
	"fmt"
)

func CreateNN()  {
	//создаем нейронную сеть с тремя входными нейронами
	// 16 скрытых нейронов
	// 4 выходных нейрона
	neuralNetwork := gonn.DefaultNetwork(3, 16, 4, false)

	// массив входящих параметров
	// 1 параметр: количество здоровья (0.1 - 1.0)
	// 2 параметр: наличие оружия
	// 3 параметр: количество врагов
	input := [][]float64 {
		[]float64{0.5, 1, 1}, []float64{0.9, 1, 2}, []float64{0.8, 0, 1},
		[]float64{0.3, 1, 1}, []float64{0.6, 1, 2}, []float64{0.4, 0, 1},
		[]float64{0.9, 1, 7}, []float64{0.6, 1, 4}, []float64{0.1, 0, 1},
		[]float64{0.6, 1, 0}, []float64{1, 0, 0} }

	// создаем цели
	target := [][]float64 {
		[]float64{1, 0, 0, 0}, []float64{1, 0, 0, 0}, []float64{1, 0, 0, 0},
		[]float64{0, 1, 0, 0}, []float64{0, 1, 0, 0}, []float64{0, 1, 0, 0},
		[]float64{0, 0, 1, 0}, []float64{0, 0, 1, 0}, []float64{0, 0, 1, 0},
		[]float64{0, 0, 0, 1}, []float64{0, 0, 0, 1} }

	// обучение. Кол-во итераций: 10000
	neuralNetwork.Train(input,target,100000)

	//  сохраняем результат в файл
	gonn.DumpNN("gonn", neuralNetwork)
}

func GetResults(out []float64) string {
	max := -99999.0
	pos := -1

	//  поиск позиции нейрона с наибольшим весом
	for i, val := range out {
		if (val > max) {
			max = val
			pos = i
		}
	}

	switch pos {
	case 0: return "attack"
	case 1: return "steal"
	case 2: return "run"
	case 3: return "to do nothing"
	}
	return ""
}

func main()  {
	CreateNN()
	nn := gonn.LoadNN("gonn")

	var hp float64 = 0.7
	var weapon float64 = 1.0
	var enemyCount float64 = 1.0

	out := nn.Forward([]float64{hp, weapon, enemyCount})
	fmt.Println(GetResults(out))
}