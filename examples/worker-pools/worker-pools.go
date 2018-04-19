// У цьому прикладі ми розглянемо як реалізовувати
// _пул робітників_ використовуючи горутини та канали.

package main

import "fmt"
import "time"

// Ось приклад, робітника якого ми будемо використовувати
// щоб запустити кілька різних екземплярів. Ці робітники будуть
// отримувати роботу на каналі `jobs` та надсилати сповіщення
// про завершення роботи до каналу `results`. Роботою нашого
// "працівника" буде сон довжиною в секунду (така собі симуляція
// реальної роботи).
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("працівник", id,
            "почав проботу над завданням", j)
        time.Sleep(time.Second)
        fmt.Println("працівник", id,
            "закінчив роботу над завданням", j)
        results <- j * 2
    }
}

func main() {

    // Щоб використовувати "пул працівників" нам необхідно
    // створити два канали: один з яких буде передавати завдання
    // робітнику, а інший отримавути результати роботи.
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Запускаємо трох працівників, але спочатку вони блокуються
    // тому що жодної роботи для них ще немає.
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Надсилаємо 5 завдань о каналу роботи, після чого
    // закриваємо канал для індикації що це вся робота
    // що наразі доступна.
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    // Нарешті підбираємо результати роботи наших прцівників.
    for a := 1; a <= 5; a++ {
        <-results
    }
}