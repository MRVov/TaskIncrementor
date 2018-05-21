Тестовое задание  на Golang

Состоит из 3-х функций

# Глобальные переменные
number - текущее значение числа
По умолчанию- 0

globalMaximumValue - максимальное значение числа

По умолчанию- максимально допустимое для int- 2147483647

Назначение- Если текущее число достигает этого значения, оно обнуляется начинает снова возвращать ноль, и снова один после следующего вызова incrementNumber()

# getNumber
 Аргументы- нет

 Возвращает-  текущее значение number. Тип int


# incrementNumber
 Увеличивает текущее число на один. После каждого вызова этого метода getNumber() будет возвращать число на один больше.
 Аргументы- нет

 Возвращает-  нет


# setMaximumValue
Аргумент-maximumValue максимальное значение. Тип int


Возвращает-  может возвращать ошибку

Устанавливает максимальное значение текущего числа.  Перезаписывая значение глобальной переменной globalMaximumValue
Не может быть меньше нуля и больше 2147483647
Если при смене максимального значения  number  его превышает, то число надо обнулить.

