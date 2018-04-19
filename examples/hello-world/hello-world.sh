# Задля запуску программи, збережіть її код до файлу
# `hello-world.go` та використайте команду `go run`.
$ go run hello-world.go
hello world

# Інколи, ми можемо забажати скомпілювати нашу программу
# у вигляді двійкового файлу. Це можливо за допомогою
# команди `go build`.
$ go build hello-world.go
$ ls
hello-world	hello-world.go

# Тепер запустіть скомпільований файл напряму.
$ ./hello-world
hello world

# Після того, як ми опанували запуск та компіляцію
# базових Go программи, давайте дізнаємось більше про
# саму мову.