// Інколи ми можемо бажати сортувати колекцію данних за
// чомось іншим крім натурального порядоку. Наприклад, нам
// потрібно відсортувати  рядки за довжиною напротивагу
// сортуванню по алфавіту, наступний приклад, як це можна
// зробити в Go.

package main

import "sort"
import "fmt"

// Щоб відсотувати щось за власними правилами в Go, нам
// спочатку треба мати це "шось". Ми створюємо тип `byLength`
// (заДовжиною) який є превдонімом до вбудованого типу -
// зріз рядків `[]string`.
type byLength []string

// Нам необхідно реалізувати `sort.Interface`, тобто створити методи
// `Len`, `Less` та `Swap` - що будуть оперувати на нашому типі данних,
// що дозволить нам використати узагальнену функцію `Sort` пакету `sort`.
// `Len` та `Swap` зазвичай будуть схожими на інші свої аналоги.
// Найбільш відрізнятись буде  `Less`, вона буде реалізовувати власну
// логіку сортування, яка в нашому випадку полягає в порівнянні довжини
// рядків.
func (s byLength) Len() int {
    return len(s)
}
func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

// З готовими, ми нарешті можемо реалізувати наше власне
// сортування за допомогою конвертації початкового зрізу `fruits`
// у зріз `byLength` (заДовжиною), і використання `sort.Sort` до
// цього типу данних.
func main() {
    fruits := []string{"персик", "банан", "ківі"}
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)
}