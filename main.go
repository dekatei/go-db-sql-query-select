package main

import (
	"database/sql"
	"fmt"

	//"log"

	_ "modernc.org/sqlite"
)

type Sale struct {
	Product int
	Volume  int
	Date    string
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Sale.
// Теперь, если передать объект Sale в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date:%s", s.Product, s.Volume, s.Date)
}

func selectSales(client int) ([]Sale, error) {
	var sales []Sale

	/* напишите код здесь
	В функции надо реализовать:
	Подключение к БД
	ыполнение SELECT-запроса
	Заполнение массива в переменной sales объектами Sale, в которых будут данные из таблицы.
	*/
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		//log.Println(err)
		return sales, err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT product, volume, date FROM sales WHERE client = :client`, sql.Named("client", client))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		sale := Sale{}

		err := rows.Scan(&sale.Product, &sale.Volume, &sale.Date)
		if err != nil {
			//log.Println(err)
			return sales, err
		}

		sales = append(sales, sale)
	}

	if err := rows.Err(); err != nil {
		//log.Println(err)
		return nil, err
	}

	return sales, nil
}

func main() {
	client := 208

	sales, err := selectSales(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sale := range sales {
		fmt.Println(sale)
	}
}
