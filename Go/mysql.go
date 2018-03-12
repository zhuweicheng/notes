// 针对MySQL数据查询的错误处理方式
package main

import (
	"database/sql"
	"fmt"
	"runtime"
)

// go run test.go
func main() {
	// mysql扩展 https://github.com/go-sql-driver/mysql
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	// db连接一般不主动关闭, 除非确认确实不再需要该连接了, go会进行回收
	defer db.Close()

	// 用户模型, 表结构, 需要一个结构来接收查询结果集
	type User struct {
		Id   int32
		Name string
		Age  int8
	}

	// 保存用户信息列表
	var user User

	// 1、查询一系列值
	// Query结果集需要调用Next()方法进行逐条遍历
	rows, err := db.Query(`
		SELECT id,name,age FROM user
	`)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		// 对于遍历, 只需要判断每次是否有错误产生即可
		// 参数绑定需要数量和位置一一对应
		if err := rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(user.Id, user.Name, user.Age)
	}
	// 处理完毕后, 需要在判断一次遍历过程中是否有错误产生
	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}

	// 2、查询一条记录
	// 查询一条记录时, 不能使用类似if err := db.QueryRow().Scan(&...); err != nil {}的处理方式
	// 因为查询单条数据时, 可能返回var ErrNoRows = errors.New("sql: no rows in result set")该种错误信息
	// 而这属于正常错误
	err = db.QueryRow(`
		SELECT id,name,age WHERE id = ?
	`, 2).Scan(
		&user.Id, &user.Name, &user.Age,
	)
	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		// 使用该方式可以打印出运行时的错误信息, 该种错误是编译时无法确定的
		if _, file, line, ok := runtime.Caller(0); ok {
			fmt.Println(err, file, line)
		}
	}
	fmt.Println(user.Id, user.Name, user.Age)

	// 3、关于NULL
	// 所有查询出来的字段都不允许有NULL, 避免该方式最好的办法就是建表字段时, 不要设置类似DEFAULT NULL属性
	// 还有一些无法避免的情况, 比如下面这个查询
	// 该种查询, 如果不存在, 返回值为NULL, 而非0, 针对该种简单的查询, 直接使用HAVING子句即可
	// 具体的查询, 需要在编码的过程中自行处理
	var age int32
	err = db.QueryRow(`
		SELECT
			SUM(age) age
		FROM user
		WHERE id = ?
		HAVING age <> NULL
	`, 10).Scan(&age)
	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		fmt.Println(err)
	}
	fmt.Println(age)
}
