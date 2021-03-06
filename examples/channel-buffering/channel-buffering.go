// За налаштування ми замовчуванню канали не _буферизуються_,
// тобто вони такі що тільки прийматимуть надсилання (`chan <-`)
// якщо відповідний отримувач (`<- chan`) готовий прийняти
// надіслане значення. _Буферизовані канали_ приймуть обмежену кількість
// значень, без того щоб отримувач запросив ці значення.

package main

import "fmt"

func main() {

    // Ось ми створюємо канал рядків, що буферизоватиме до 2 значень
    messages := make(chan string, 2)

    // Оскільки канал буферизований, ми можемо надіслати в нього
    // оці значення без супровідного одночасного читання з іншого кінця.
    messages <- "buffered"
    messages <- "channel"

    // І ми можемо отримати ці значення пізніше.
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
