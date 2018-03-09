// Go підтримує кодування/декодування
// [base64](https://uk.wikipedia.org/wiki/Base64).

package main

// За допомогою цього синтаксису ми імпортуємо пакет
// `encoding/base64` до якого будемо звертатись, як
// до `b64` замість стандартного звернення `base64`.
// Це дозволить нам зберегти трошка місяця під час
// програмування.
import b64 "encoding/base64"
import "fmt"

func main() {

    // Це рядок який ми будемо кодувати та декодувати.
    data := "abc123!?$*&()'-=@~"

    // Go підтриму як стандартну так і URL-сумісну base64.
    // Спочатку приведемо приклад як використовувати
    // стандартне base64. Механізм кодування потребує зріз
    // байтів, отож ми сконвернуємо наш рядом до потрібного типу.
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    // Розкодування може повертати помилку, яку ви можете перевірити
    // якщо ви не знаєте напевне чи була вхідний (не закодований)
    // рядок привально сформований.
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    // А це ми кодуємо за допомогою URL-сумісного base64.
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)

    // Розкодовуємо, аналогічно тому, як ми використовували
    // стандартне кодування base64, кількома рядками вище.
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}
