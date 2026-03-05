package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	//"utf8"
)

//эта программа поможет вам сочетать элементы, получая молекулы.

const H = 1.0079 //  Водород 1.0079 а. е. м. (атомных единиц массы).
const C = 12.0107 //  Углерод 12,0107 а.е.м.
const He = 4.002602 //  Гелий 4,002602 а.е.м.
const O = 15.9994 //  Кислород 15,9994 а. е. м.
const AllAtomsPrint = [4]string{"1  -  H, *H\n",
								"2  -  C, *C\n",
								"3  -  He, *H\n",
								"4  -  O, *O\n",
H, *H, C, *C, He, *He, O, *O}

type MoleculesСombinator interface{
	*Molecule // отправлять дальше на комбинацию молекул
}

type Molecule struct {
	MoleculeName string
	Atoms []string
	MMass float64
	//ChemicalProperties string
} 


func (m *Molecule) MoleculaName(name string){
	return Molecule{MoleculeName: name}
}

func (m *Molecule) Atoms(combination string){
	choosenAtoms := combination []string
	return Molecule{Atoms: }
}

// func (m *Molecule) MMass(combination []string) float64 {
// 	for i := in range
// }

func (m *Molecule) ChoosesAtoms(const AllAtoms, Atoms[]string, c combination []string) c {
	fmt.Println("Выберите набор атомов: | Для выхода нажмите х")
	fmt.Println(const AllAtoms)
	var answer int

	_, _ = fmt.Scan(&answer)
	_, _ fmt.Scanln()
	for answerR != "x" {

point0:
var answer string
	_, err := fmt.Scanln(&answer)
	if err !=0 {
		fmt.Println("Ошибка")
		goto point0
		}
	if okAnswer <= 5 && okAnswer <= 0 { 
		append answer(answer, answer []int)
	}
	}

//
//
//


//
func startCapture()err Error{
	fmt.Println("Напишите перевод на русской язык и докажите, что Вы не робот: . или нажмите '2' - я не знаю")
	fmt.Sprintf("bynary", "0000010000111111\n", "00100000\n",
				"0000010001000000\n", "0000010000111110\n", 
				"0000010000110001\n",  "0000010000111110\n",
	 			"0000010001000010\n",
				"Ваш ответ:\n")
	var answer string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	words := strings.Split(line, " ")
	fmt.Println("Вы сказали:", words)
	if words... != "Я робот" {
		fmt.Pintln("Вы человек")
		retutn nil
	} else {
		return errors.New("Нет прав доступа Института")
		
	}
}

func main(err error) {
	defer "Выход"
	err := startCapture()
	if err != nil {
		fmt.Println("Вы робот")
		return
	} else {
		OpenMenu()
	}
}

func OpenMenu() []string {
	
	fmt.Println("---Главное меню---")
	fmt.Println("%-15s, %-5s, %-5s", "1 -- Смоделировать молекулу",
	"2 -- Отчёт последних моделирований", "0 -- Выход")
	for answer != 0 {
	_, err := fmt.Scanln(&answer)
	switch answer {
	case 1:
		choosedAtoms = choosesAtoms(c.combination []string)
	}
	return nil
} 
}