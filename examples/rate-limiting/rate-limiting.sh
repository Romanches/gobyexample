# Запускаючи нашої программу ми побачимо першу
# партія запитів що обслуговуватимуться кожні 200
# мілісекунд.
$ go run rate-limiting.go
request 1 2012-10-19 00:38:18.687438 +0000 UTC
request 2 2012-10-19 00:38:18.887471 +0000 UTC
request 3 2012-10-19 00:38:19.087238 +0000 UTC
request 4 2012-10-19 00:38:19.287338 +0000 UTC
request 5 2012-10-19 00:38:19.487331 +0000 UTC

# Друга ж партія запитів що ми обслуговувємо
# представляє собою 3 запити що не чекають
# лімітування і представляють собою "наплив"
# - дозволену кулькіть запитів без обмеження та
# ще 2 обмежені 200 мілісекундним інтервалом
# запити.
request 1 2012-10-19 00:38:20.487578 +0000 UTC
request 2 2012-10-19 00:38:20.487645 +0000 UTC
request 3 2012-10-19 00:38:20.487676 +0000 UTC
request 4 2012-10-19 00:38:20.687483 +0000 UTC
request 5 2012-10-19 00:38:20.887542 +0000 UTC