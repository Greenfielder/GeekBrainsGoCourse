package main

// Home work 6

// 1. Проанализируйте задания предыдущих уроков.
// a. В каких случаях необходима была явная передача указателя в качестве входных параметров и возвращаемых результатов или в качестве приёмника в методах?
// Ответ:
// На пример в случае fmt.Scan(&x), в этом случае нам необходимо явно получить значение в переменную, которая вводилась с клавиатуры.

// b. В каких случаях мы фактически имеем дело с указателями при передаче параметров, хотя явно их не указываем?
// Ответ:
// Рискну предположить,что используя слайсы в качестве переменных при передаже параметров.

// 2. Для арифметического умножения и разыменования указателей в Go используется один и тот же символ — оператор (*).
// Как вы думаете, как компилятор Go понимает, в каких случаях в выражении имеется в виду умножение, а в каких — разыменование указателя?

// Ответ:
// Задание для серфинга интернета )) У операторов есть свой порядок приоритетов.
// У унарных операторов приоритет высокий, у бинарные по ниже приоритет.
// Оператор разыменивания - унарный (один аргумент)
// Оператор умножение - бинарный (два аргумента)
// Скорее всего компилятор сначала проверяет * на количество аргументов стоящих рядом с ним.
