package main

import (
	"errors"
)

// Основаня глобальная переменная, содержашая число
var number int=0

// Максимальное значение текущего числа.
// По умолчанию максимум -- максимальное значение int.
var globalMaximumValue=2147483647

// Возвращает текущее число. В самом начале это ноль.
// * Аргументы- нет
// * Возвращает- текущее значение number. Тип int
func getNumber() int{
	return number
}

// Увеличивает текущее число на один. После каждого вызова этого
// метода getNumber() будет возвращать число на один больше.
// * Аргументы- нет
// * Возвращает-  нет
func incrementNumber(){
	number+=1;

	//Если текущее число достигает максимального значения, оно обнуляется
	//начинает снова возвращать ноль, и снова один после следующего вызова
	if number>=globalMaximumValue{
		number=0
	}

}

//Устанавливает максимальное значение текущего числа.
func setMaximumValue(maximumValue int) error{

	//Нельзя позволять установить тут число больше лимита для  int.
	if maximumValue>2147483647{
		return errors.New("Wrong Maximum value. Cant be more than 2147483647")
	}

	//Нельзя позволять установить тут число меньше нуля.
	if maximumValue<=0{
		return errors.New("Wrong Maximum value. Cant be less than 1")
	}

	//Если при смене максимального значения число резко начинает
	//превышать максимальное значение, то число надо обнулить.
	if maximumValue<number {
		number = 0
	}

	//Устанавливаем максимальное значение
	globalMaximumValue=maximumValue

	return nil
}