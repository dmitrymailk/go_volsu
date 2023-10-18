### 1. Упражнения для освоения синтаксиса языка

```bash
go run ./task_1/3.go
```
```bash
go run ./task_1/4.go
```

### 2. Обработка символьных строк. Регулярные выражения.
```bash
go run ./task_2/8.go
```
```bash
go run ./task_2/9.go
```

### 3. Структуры. Методы.
```bash
go run ./task_3/2.go
```

### 4. Пакеты, Интерфейсы
```bash
cd task_4 && mkdir polynom
```
```bash
cd task_4/polynom && go mod init test.com/polynom
```
```bash
cd task_4 && go mod init test.com/myapp
```
Добавить в файл `task_4/go.mod`
```mod
module test.com/myapp

go 1.21.3
// эту строчку
replace test.com/polynom => ./polynom 

require test.com/polynom v0.0.0-00010101000000-000000000000

```
Выполнить эту команду она добавит require.
```bash
go mod tidy
```

```bash
go run 2.go
```

### 5. Параллельное программирование. Горутины. Потоки и файлы
```bash 
go run task_5/2.go
```

### 6. Обработка дерева каталогов файлов. Поиск и преобразование содержимого файлов.
```bash
go run task_6/3.go
```