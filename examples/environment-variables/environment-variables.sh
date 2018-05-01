# Запуск програми показже що `FOO` матиме значення
# (бо ми встановили його), `BAR` пусте бо воно відсутнє
# в середовищі.
$ go run environment-variables.go
FOO: 1
BAR:

# Перелік ключів змінних середовища залежить від машити,
# тобто середовиза машини.
TERM_PROGRAM
PATH
SHELL
...

# якщо ми встановимо `BAR` всередовищі спочатку, програма що
# запуститься слідом матиме змогу скористатись з цієї зміннної.
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
