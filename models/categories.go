package models

import (
	"context"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Category struct{
	Id int `json:"id"`
	Name string `json:"name" form:"name"`
}
func countCategory(src string) int {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT COUNT(*) FROM categories WHERE "name" ILIKE '%' || $1 || '%'`
	row := db.QueryRow(context.Background(), sql, src)

	var num int
	row.Scan(&num)
	return num
}
func ShowAllCategories() []Category  {
	db := lib.DB()
	defer db.Close(context.Background())

	// offset := (page - 1) * limit
	// sql := `SELECT * FROM categories WHERE "name" ILIKE '%' || $1 || '%'
	// 	LIMIT $2
	// 	OFFSET $3
	// 	`
	sql := `SELECT * FROM categories`
	// rows, _ := db.Query(context.Background(), sql, search, limit, offset)
	rows, _ := db.Query(context.Background(), sql)

	data, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[Category])
	
	// count := countCategory(search)

	// return data, count
	return data
}
func ShowCategoryById (id int) Category {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "categories" where id = $1`
	row, _ := db.Query(
		context.Background(),
		sql,
		id,
	)

	category, _ := pgx.CollectRows(row, pgx.RowToStructByPos[Category])
	var result Category
	for _, v := range category {
		if v.Id == id {
			result = v
		}
	}

	return result
}
func CreateNewCategory (data Category) Category {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `insert into "categories" ("name") values ($1) returning "id", "name"`
	rows := db.QueryRow(
		context.Background(),
		sql,
		data.Name,
	)

	var result Category
	rows.Scan(
		&result.Id,
		&result.Name,
	)
	return result
}
func EditCategory (data Category, id int) Category {
	db := lib.DB()
	defer db.Close(context.Background())

	dataSql := `UPDATE categories SET name = $1 WHERE id= $2`

    db.Exec(context.Background(), dataSql, data.Name, id)

	data.Id = id
	return data
}
func RemoveCategory (id int) Category {
	db := lib.DB()
    defer db.Close(context.Background())

	cat := ShowCategoryById(id)

	sql := `delete from "categories" where "id" = $1`
	db.Exec(context.Background(), sql, id)

	return cat
}