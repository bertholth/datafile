package datafile
import (
	"bufio"
	"os"
	"strconv"

)

func GetFloats(fileName string) ([3]float64, error)
	var numbers [3]float64 // обьявляем массим 
	file, err := os.Open(fileName) // открываем файл
	// проверяем на ошибки
	if err !== nil {
		retur numbers, err
	}
	i:=0
	scanner:=bufio.NewScanner(file)
	for scanner.Scan {
		numbers[i], err = strconv.ParseFloat(scanner.Text(),64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err=file.Close()
	if err != nil {
		return numbers,err
	}
	return numbers, nil