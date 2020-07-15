package sql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

var db *sql.DB

func SqlOpen() {
	var err error
	db, err = sql.Open("postgres", "port=5432 user=postgres password=postgres dbname=test sslmode=disable")
	//port是数据库的端口号，默认是5432，如果改了，这里一定要自定义；
	//user就是你数据库的登录帐号;
	//dbname就是你在数据库里面建立的数据库的名字;
	//sslmode就是安全验证模式;

	fmt.Println("Connected!")
	checkErr(err)
}
func SqlInsert(mobile string, email string, pwd string) {
	//插入数据
	//		stmt, err := db.Prepare("INSERT INTO user(user_id,user_mobile_account,user_mail_account,user_password) VALUES($1,$2,$3,$4) ")
	sqlStatement := "INSERT INTO public.user(id,mobile,email,password) VALUES($1,$2,$3,$4)"
	//	sqlStatement := "INSERT INTO user(uid) VALUES($1)"
	stmt, err := db.Prepare(sqlStatement)
	checkErr(err)

	defer stmt.Close()

	id := uuid.Must(uuid.NewV4(), nil)

	res, err := stmt.Exec(id, mobile, email, pwd)

	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}
func SqlDelete() {
	//删除数据
	stmt, err := db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)

	res, err := stmt.Exec(1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}
func SqlSelect(mobile string, email string) string {
	//查询数据
	var password string
	err := db.QueryRow("SELECT password FROM public.user WHERE mobile = $1 OR email = $2", mobile, email).Scan(&password)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No password with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		return password
	}

	return ""
}
func SqlUpdate() {
	//更新数据
	stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err := stmt.Exec("ficow", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}
func SqlClose() {
	db.Close()
	fmt.Println("Disconnected!")
}
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
