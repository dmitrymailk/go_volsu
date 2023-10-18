package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Структура для хранения адреса
type Address struct {
	Name      string
	City      string
	Zip       int
	Street    string
	Building  string
	Apartment int

	BuildingType   string
	FloorsCount    int
	ResidentsCount int
	Phone          string
}

// Структура для хранения всех адресов
type Addresses struct {
	List []Address
}

// Метод добавления нового адреса
func (a *Addresses) Add(address Address) {
	a.List = append(a.List, address)
}

// Метод удаления адреса по индексу
func (a *Addresses) Remove(index int) {
	a.List = append(a.List[:index], a.List[index+1:]...)
}

// Метод преобразования адреса в строку
func (a *Address) String() string {
	return fmt.Sprintf("%s, %s, %s", a.City, a.Street, a.Building)
}

// Метод извлечения номера телефона
func (a *Address) GetPhone() string {
	return a.Phone
}

// Подсчет адресов в городе
func (a *Addresses) CountInCity(city string) int {
	count := 0
	for _, addr := range a.List {
		if addr.City == city {
			count++
		}
	}
	return count
}

// Сохранение в файл
func (a *Addresses) SaveToFile(filename string) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// Загрузка из файла
func (a *Addresses) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, a)
}

func main() {
	/*
		2. Задать структуру для хранения почтовых адресов,
		включая
			ФИО,
			город,
			индекс,
			улицу,
			номер дома и квартиры,
			тип строения,
			этажность,
			число проживающих,
			номер телефона .

		Реализовать методы
			ok добавления/удаления адреса,
			преобразование адреса в строку,
			извлечения
				номера телефона,
				города,
				индекса,
			подсчета числа адресов в заданном городе,
			загрузка и выгрузка всей структуры в файл.
	*/

	// Создание структуры с адресами
	addresses := Addresses{
		List: []Address{
			{
				Name:      "Иванов Иван Иванович",
				City:      "Москва",
				Zip:       101000,
				Street:    "Ленина",
				Building:  "10",
				Apartment: 105,
				Phone:     "+7 495 123 45 67",
			},
			{
				Name:      "Петров Петр Петрович",
				City:      "Санкт-Петербург",
				Zip:       191180,
				Street:    "Чайковского",
				Building:  "88",
				Apartment: 412,
				Phone:     "+7 812 345 67 89",
			},
		},
	}

	// Добавление нового адреса
	newAddr := Address{
		Name: "Сидоров Сидр Сидорович",
		City: "Казань",
	}
	addresses.Add(newAddr)

	// Проверка количества адресов после добавления
	count := len(addresses.List)
	fmt.Println("Адресов после добавления:", count)

	// Удаление адреса
	addresses.Remove(1) // удаляем первый адрес

	// Проверка количества адресов после удаления
	count = len(addresses.List)
	fmt.Println("Адресов после удаления:", count)

	// Получение адреса в виде строки
	addr := addresses.List[0]
	addrStr := addr.String()

	fmt.Println("Адрес:", addrStr)

	// Получение номера телефона
	phone := addresses.List[0].GetPhone()
	fmt.Println("Телефон:", phone)

	// Получение Индекс
	index := addresses.List[0].Zip
	fmt.Println("Индекс:", index)

	// Подсчет адресов в городе
	count = addresses.CountInCity("Москва")
	fmt.Println("Адресов в Москве:", count)

	// Сохранение в файл
	err := addresses.SaveToFile("./task_3/addresses.json")
	if err != nil {
		fmt.Println(err)
	}

	// Загрузка из файла
	var loaded Addresses
	err = loaded.LoadFromFile("./task_3/addresses.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Загружено адресов: %d\n", len(loaded.List))
}
