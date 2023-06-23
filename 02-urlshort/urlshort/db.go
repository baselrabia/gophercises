package urlshort

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func LoadFromDB() (map[string]string, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/urls")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`select path,url from urls`)
	if err != nil {
		return nil, err
	}
	list := map[string]string{}
	for rows.Next() {
		var url data
		if err := rows.Scan(&url.Path, &url.Url); err != nil {
			return nil, err
		}
		fmt.Printf("this is a path %v and this is a url %v \n", url, url.Path)
		list[url.Path] = url.Url
	}

	return list, nil
}
