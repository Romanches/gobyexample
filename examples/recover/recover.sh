# Запуск программи покаже як вона опрацювала
# сигнал _panic_'и та відновала свою роботу.
$ go run recover.go
Опрацювання паніки: Паніка без причини
Повернення з функції normalFunc.

# Власне _panic_'а була перехоплена в момент коли вона мала б
# рівнем вище в main, і єдиний шлях зробити це є використання
# відкладення виклику функції що буде перевіряти чи потрібно
# функції робитит відновлення.

