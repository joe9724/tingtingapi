package _var

import (
	"fmt"
	"github.com/jinzhu/gorm"

	_"net/url"
	_"log"
	_"io/ioutil"
	_"net/http"
	"encoding/json"
)
type Respoonse struct{
	Code int64 `json:"code,omitempty"`
	// cover
	Msg string `json:"msg,omitempty"`
}
func Response200(code int64,msg string) (string) {
	var response Respoonse
	response.Code = code
	response.Msg = msg

	out, _ := json.Marshal(response)
	fmt.Println("ous is",out)
	return(string(out))
}
func GetResourceDomain(filetype string) (string){
	var val string
	if(filetype == "avatar") {
		val = "http://tingting-resource.bitekun.xin/resource/image/avatar/"
	}
	if(filetype == "icon") {
		val = "http://tingting-resource.bitekun.xin/resource/image/icon/"
	}else if(filetype == "cover"){
		val = "http://tingting-resource.bitekun.xin/resource/image/cover/"
	}else if(filetype == "mp3"){
		val = "http://tingting-resource.bitekun.xin/resource/mp3/"
	}else if(filetype == "amr"){
		val = "http://tingting-resource.bitekun.xin/resource/amr/"
	}else if(filetype == "other"){
		val = "http://tingting-resource.bitekun.xin/resource/other/"
	}
	return val
}
func SendMsg(mobile string,str string,t int64) bool{
	var sendOk bool
	sendOk = true
	var content string
	if t==0 {
		content = "【听听阅读】(5分钟内有效)验证码是:" + str
	}else if t == 1{
		content = "【听听阅读】(5分钟内有效)绑定手机号后初始密码是:" + str
	}else if t == 2{
		content = "【听听阅读】(5分钟内有效)快捷登录初始密码是:" + str
	}
	fmt.Println(content)
	/*var requestUrl string
	requestUrl = "http://mes.sh-hstx.com:8800/sendXSms.do?username=tusheng1234&password=abcd1234&mobile="+mobile+"&content="+content+"&dstime=&productid=100035"
	fmt.Println("send str is",url.QueryEscape(requestUrl))
	u, _ := url.Parse(requestUrl)
	//q := u.Query()
	//u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	fmt.Println("url is",u.String())
	if err != nil {
		log.Fatal(err)
		//return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		sendOk = false
	}else{
		sendOk = true
	}

	fmt.Printf("SP网关回复内容是:%s", result)*/

	return sendOk
}

func OpenConnection() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", "root:Tingtingyuedu654321!!!@tcp(47.104.131.147:3306)/tingting?charset=utf8&parseTime=True&loc=Local")
	return db,err
	/*switch os.Getenv("GORM_DIALECT") {
	case "mysql":
		// CREATE USER 'gorm'@'localhost' IDENTIFIED BY 'gorm';
		// CREATE DATABASE gorm;
		// GRANT ALL ON gorm.* TO 'gorm'@'localhost';
		fmt.Println("testing mysql...")
		dbhost := os.Getenv("GORM_DBADDRESS")
		if dbhost != "" {
			dbhost = fmt.Sprintf("tcp(%v)", dbhost)
		}
		db, err = gorm.Open("mysql", fmt.Sprintf("gorm:gorm@%v/gorm?charset=utf8&parseTime=True", dbhost))
	case "postgres":
		fmt.Println("testing postgres...")
		dbhost := os.Getenv("GORM_DBHOST")
		if dbhost != "" {
			dbhost = fmt.Sprintf("host=%v ", dbhost)
		}
		db, err = gorm.Open("postgres", fmt.Sprintf("%vuser=gorm password=gorm DB.name=gorm sslmode=disable", dbhost))
	case "foundation":
		fmt.Println("testing foundation...")
		db, err = gorm.Open("foundation", "dbname=gorm port=15432 sslmode=disable")
	case "mssql":
		// CREATE LOGIN gorm WITH PASSWORD = 'LoremIpsum86';
		// CREATE DATABASE gorm;
		// USE gorm;
		// CREATE USER gorm FROM LOGIN gorm;
		// sp_changedbowner 'gorm';
		fmt.Println("testing mssql...")
		db, err = gorm.Open("mssql", "sqlserver://gorm:LoremIpsum86@localhost:1433?database=gorm")
	default:
		fmt.Println("testing sqlite3...")
		db, err = gorm.Open("sqlite3", filepath.Join(os.TempDir(), "gorm.db"))
	}
	// db.SetLogger(Logger{log.New(os.Stdout, "\r\n", 0)})
	// db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	if os.Getenv("DEBUG") == "true" {
		db.LogMode(true)
	}
	db.DB().SetMaxIdleConns(10)
	return*/
}
