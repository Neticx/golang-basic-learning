package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"runtime"
	"sync"
)

func handleRequests()  {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",HelloWorld).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081",router))
}
func init()  {
	//fmt.Println("process started")
	runtime.GOMAXPROCS(5)
	//fmt.Println("start web server on port 8080")
	//http.HandleFunc("/employe",employe)
	//http.HandleFunc("/employes",employes)
	//http.HandleFunc("/fetch",fetchData)
	//http.HandleFunc("/post",httpFormLearning)
	//http.ListenAndServe(":8080",nil)
}
func main()  {
	//fmt.Println("go run in port 8081")
	//handleRequests()

	//pointerLearning()
	//structLearning()
	//accessLevelLearning()
	//interfaceLearning()
	//emptyInterfaceLearning()
	//reflectLearning()
	//goRoutineLearning()
	//channelLearning()
	//loopChannelLearning()
	//channelTimeout()
	//deferLearning()
	//exitLearning()
	//errorLearning()
	//panicLearning()
	//timeLearning()
	//timeRoutineLearning()
	//encodeLearning()
	//hashLearning()
	//saltingHashLearning()
	//flagArgLearning()
	//execLearning()
	//httpListenerLearning()
	//jsonLearning()
	//jsonAPILearning()
	//httpLearning()
	//sqlLearning()
	//prepareSqlLerning()
	//dbexecLearning()
	//nosqlLearning()
	//insertData()
	//waitGroupLearning()
	//mutexLearning()
	webHttpLearning()
}

func HelloWorld(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"Hello World gaysss")
}



//func main() {
//var number int
//number = 2
//fmt.Println(number)
//
//address := "Tangerang"
//fmt.Println(address)
//
//age := 19
//if age < 17 {
//	fmt.Println("tidak cukup")
//}else if age > 17 && age < 25{
//	fmt.Println("time to study and work")
//}else{
//	fmt.Println("too old")
//}
//
//switch age{
//	case 10:
//		fmt.Println("masih 10")
//	case 17:
//		fmt.Println("17 gan")
//	case 19:
//		fmt.Println("time to works")
//	default:
//		fmt.Println("hmm")
//}
//// switch gaya ifelse
//switch {
//	case age <= 17:
//		fmt.Println("umur kurang 17")
//	case age > 17:
//		fmt.Println("umur lebih dari 17")
//		fallthrough //digunakan untuk melanjutkan pengecekkan terhadap case, karena di go jika case terpenuhi maka akan berhenti mengecek
//	case age < 20:
//		fmt.Println("kurang dari 20")
//}
//
////array
//
//var arr[4]int
//arr[1] = 1
//arr[3] = 3
//fmt.Println(arr)
//
//arrr := [3]string{"fahmie","ulul","azmie"}
//fmt.Println(arrr)
//
//unlimited := [...][1]string{{"fahmi"},{"ulul"},{"azmi"}} // ... menandakan array yang tak ditentukan
//fmt.Println(unlimited)
//
//for i := 0; i < len(arrr) ;i++  {
//	fmt.Println(arrr[i])
//}
//
//for _, data := range arrr{ // _ sebagai anonim variabel, karena variabel jika tidak digunakan akan menghasilkan error di go
//	println(data)
//}
//
//// slice
//var slc = []int{1,2,3}
//fmt.Println(slc)
//
//slc[0] = 2
//slc[1] = 3
//fmt.Println(slc)
//
//slices := []string{"fahmi","ulul","azmi"}
//fmt.Println(slices)
//
//slc1 := slices[0:1]
//slc2 := slices[0:3]
//slc3 := slices[0:3:3] //parameter teakhir menandakan batasan kapasitas, yaitu kapasitasnya hanya boleh 3
//
//fmt.Println(slc1)
//
//slices[0]= "fahmie"
//
//fmt.Println(slc2)
//fmt.Println(slc3)
//
//fmt.Println(cap(slc1)) // capacity 3 karena mereference dari slices yang memiliki 3 array
//fmt.Println(len(slc1)) // len 1 karena hanya memiliki 1 array meskipun memiliki 3 capacity
//
////kenapa slc2 indeks ke 0 ikut berganti value? karena slice bukan hasil copy dari elemen array, tapi referensi dari elemen array.
//// memiliki alamat memori yang sama
//
////map adalah jenis array yang memiliki key dan value
//
//maps := map[int]string{}
//
//maps[1] = "satu"
//maps[2] = "dua"
//
//var map2 = map[string]int{"satu":1,"dua":2}
//
//fmt.Println(maps[1])
//fmt.Println(map2)
//delete(map2,"satu") // menghapus map dengan key 2. sama dengan fungsi unset pada php
//fmt.Println(map2)
//
//for key, data := range maps {
//	fmt.Println(key,data)
//}
//var value, isExist = maps[0]// value akan berisi value dari maps[0] sementara isExist akan berisi boolean apakah data dengan key itu ada atau tidak
//
////slice map
//var newslice = make([]map[string]string,3)
//
//var slicemap = []map[string]string{}
//
//newslice[0] = map[string]string{"tes":"tes"}
//
//
//fmt.Println(newslice)
//fmt.Println(slicemap)
//fmt.Println(isAssert(2,3))
////variadic function manually not using slice
//fmt.Println(variadicTest(1,2,3,4,5,4,3,2,1))
////variadic function automation with slice
//var slices = []int{1,2,3,4,5,4,3,2,1}
//fmt.Println(variadicTest(slices ...))
//fmt.Println(sliceRand(slices))
//
//var count = func(x int, y int)(result float64) {
//	var sum = x+y
//	result = float64(sum)
//	return
//}
//fmt.Println(count(2,3))
//
//var data = []string{"a","b","b","b"}
//countRow(data, func(data []string) int {
//	length := len(data)
//	return int(length)
//})
//	pointerLearning()
//
//}

//for testing is var x and var y has same value
//func isAssert(x int,y int)(bool) { // bisa juga paramteter ditulis x,y int karena bertipe data sama, jadi di declare hanya sekali
//	fmt.Println(time.Now())
//	return x == y
//}
//
//func variadicTest(data ...int)(result []int){
//	var temp = []int{}
//	for _, res := range data{
//		temp = append(temp,res)
//	}
//	result = temp
//	return
//}
//
//func sliceRand(data []int)[]int{
//	return data
//}
//
//type CountCallback func([]string)(int)
//
//func countRow(data []string, callback CountCallback)  {
//	var length = callback(data)
//	fmt.Println(length)
//}
//
//func pointerLearning()  {
//	var strings string = "tes"
//	var tes *string = &strings
//
//	fmt.Println(*tes)
//	*tes = "done"
//
//	fmt.Println(*tes)
//	fmt.Println(strings)
//
//}
//
//func structLearning()  {
//	var des = desa{}
//	des.Nama = "kampungku ini"
//	des.Kecamatan = "kresek"
//	fmt.Println(des.countLength())
//	des.changeDesa("mukidi")
//	fmt.Println(des.countLength())
//
//}
//
//type (
//	desa struct {
//		Nama string
//		Kecamatan string
//	}
//
//	kecamatan struct {
//		data int
//	}
//)
////tipe data interface berarti bisa mengembalikan tipe data apapun.
//func (des desa) countLength() (interface{}) {
//	var length = len(des.Nama)
//	return length
//}
//
//func (des *desa) changeDesa(nama string)  {
//	des.Nama = nama
//}
////public and private determine in character case, if starting with uppercase like Data, its public. if data, its private
//func accessLevelLearning()  {
//	var datas = Data{}
//	datas.name = "pilonopila"
//	fmt.Println(datas.CountDataLength())
//}
//
//func (bagi *kecamatan) pembagian() int {
//	bagi.data = bagi.data /2
//	return bagi.data
//
//}
//
//
//func (bagi *kecamatan) pengurangan() int {
//	return  bagi.data - 2
//}
//
//type wilayah interface {
//	pembagian() int
//	pengurangan() int
//}
//
//func interfaceLearning()  {
//	var kec wilayah = &kecamatan{10}
//	fmt.Println(kec.pembagian())
//	fmt.Println(kec.pengurangan())
//}
//
//func emptyInterfaceLearning()  {
//	var inter interface{} = &kecamatan{2}
//	var secret = inter.(*kecamatan).data //casting to pointer, because default type of struct is string. so it should cast into each of type
//	var pon *int = &secret
//	fmt.Println(secret)
//	fmt.Println(*pon)
//}
//
//func reflectLearning()  {
//	var number int
//	number = 2
//
//	var name string
//	name = "fahmi"
//
//	var kec = kecamatan{2}
//	var point *int = &number
//
//	var reflectValue = reflect.ValueOf(point)
//
//	fmt.Println(reflect.ValueOf(number))
//	fmt.Println(reflect.ValueOf(name))
//	fmt.Println(reflect.ValueOf(kec))
//
//	fmt.Println(reflectValue.Kind())
//	fmt.Println(reflectValue.Type())
//	fmt.Println(reflectValue.Interface())
//	var tes = reflectValue.Elem() // Elem mengambil elemen asli jika tipe reflect nya adalah pointer
//	fmt.Println(tes.Kind()) // struct, pointer, int, string dll
//
//	var des = &desa{Nama: "pasir",Kecamatan:"kresek"}
//	des.getInfoProperty()
//}
//
//func (d *desa) getInfoProperty()  {
//	value := reflect.ValueOf(d)
//
//	if value.Kind() == reflect.Ptr {
//		value = value.Elem()
//	}
//	//fmt.Println(value)
//
//	var vtype = value.Type()
//	fmt.Println(value)
//	fmt.Println(value.NumField())
//	for i := 0; i < vtype.NumField() ; i++  {
//		fmt.Println("Name: ",vtype.Field(i).Name)
//		fmt.Println("Type: ",vtype.Field(i).Type)
//		fmt.Println("Value: ",value.Field(i).Interface()) // untuk mengakses interface, modifier atau field dari struct harus public
//	}
//}
//
//func goRoutineLearning()  {
//	go show("INI")
//	show("BUDI")
//	fmt.Scanln()
//}
//
//func show(message string)  {
//	for i := 0; i < 50; i++  {
//		fmt.Println((i + 1),message)
//	}
//}
//
//func channelLearning() {
//	var message= make(chan int, 2)
//	//var sayHello = func(who string) {
//	//	var result = fmt.Sprintf("hello %s", who)
//	//	message <- result
//	//}
//	//
//	//go sayHello("fahmi")
//	//go sayHello("ulul")
//	//go sayHello("azmi")
//	//
//	//var person1 = <-message
//	//fmt.Println(person1)
//	//
//	//var person2 = <-message
//	//fmt.Println(person2)
//	//
//	//var person3 = <-message
//	//fmt.Println(person3)
//
//	//channel as param
//	//var dummy = []string{"hello","world","hehe","this","is","sparta"}
//	//
//	//for _, data := range dummy {
//	//	go func(each string) {
//	//		message <- each
//	//	}(data)
//	//}
//	//
//	//for i := 0; i < len(dummy) ;i++  {
//	//	printChan(message)
//	//}
//	//buffered channel
//
//	go func() {
//		for {
//			i := <-message
//			fmt.Println("data received", i)
//		}
//	}()
//
//	for i := 0;i < 10 ;i++  {
//		fmt.Println("send data", i)
//		message <- i
//	}
//}
//
//func printChan(message chan string)  {
//	fmt.Println(<-message)
//}
//
//func loopChannelLearning()  {
//	var data = make(chan int)
//	go setData(data)
//	printData(data)
//}
//
//func setData(ch chan<- int)  {
//	for i := 0; i <= 100 ;i++  {
//		ch<- i
//	}
//	close(ch)
//}
//
//func printData(ch<- chan int)  {
//	for data := range ch {
//		fmt.Println(data)
//	}
//}
//
////channel direction :
//// var chan<- = channel yang hanya bisa menerima data
//// var <-chan = channel yang hanya bisa mengirim data
//
//func channelTimeout()  {
//	rand.Seed(time.Now().Unix())
//	fmt.Println(time.Duration(rand.Int()%10+1))
//	var data = make(chan int)
//	go sendingData(data)
//	retrieveData(data)
//}
//
//func sendingData(ch chan<- int)  {
//	for i:= 0;true ;i++  {
//		ch<- i
//		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
//	}
//}
//
//func retrieveData(ch <-chan int)  {
//// initialize loop label
//loop:
//	for {
//		select {
//			case data := <-ch:
//				fmt.Println("receive data:", data)
//			case <-time.After(time.Second * 5):
//				fmt.Println("timeout guys!!")
//				break loop
//		}
//	}
//
//}
//// defer berfungsi untuk meng-akhirkan eksekusi di paling akhir block code sekalipun setelah di return
//func deferLearning()  {
//	somedata := []string{"pizza","kebab","macaroni"}
//	for _,data := range somedata{
//		reasonChoise(data)
//	}
//}
//
//func reasonChoise(data string)  {
//	defer fmt.Println("<------------->")
//	defer fmt.Println("thank you!")
//
//	if data == "pizza" {
//		fmt.Println("great! you choose", data)
//		return
//	}else if data == "macaroni"{
//		fmt.Println("the",data,"is so nice!")
//		return
//	}
//
//	fmt.Println("you choose",data)
//}
//
////exit statement will stop entire programs. either you was written defer statement, but it will not shown because exit will stop entire program and show exit status code
//func exitLearning()  {
//	var panic bool
//
//	panic = false
//	defer fmt.Println("sipp")
//
//	if panic {
//		os.Exit(1)
//	}else {
//		fmt.Println("works")
//	}
//}
//
//func errorLearning()  {
//	var input string
//
//	fmt.Scanln(&input)
//
//	//var result int
//	//var err error
//
//	//result, err = strconv.Atoi(input)
//	//
//	//if err != nil {
//	//	fmt.Println(err.Error())
//	//}else {
//	//	fmt.Println(result,"success")
//	//}
//
//	if valid, err := validate(input); valid {
//		fmt.Println("works", input)
//	}else {
//		fmt.Println(err.Error())
//	}
//
//}
//
//func validate(input string)(bool,error)  {
//	if strings.TrimSpace(input) == "" {
//		return false,errors.New("cannot accept empty input")
//	}
//	return true, nil
//}
//
//func panicLearning()  {
//	defer catch()
//	var input string
//	fmt.Scanln(&input)
//
//	if valid, err := validate(input); valid {
//		fmt.Println("works")
//	}else {
//		panic(err.Error())
//		fmt.Println("end")
//	}
//}
//
//func catch()  {
//	if r := recover(); r != nil {
//		fmt.Println("error occured : ",r)
//	}else {
//		fmt.Println("works as well")
//	}
//}
//
//func timeLearning()  {
//	var now = time.Now()
//	fmt.Println(now)
//	fmt.Println(now.Year())
//	var date,err = time.Parse("2006-01-02","2018-09-12");
//	fmt.Println(date)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//}
//
//func timeRoutineLearning()  {
//	var input int
//	var ch = make(chan bool)
//	var timeout int = 5
//	go timer(timeout,ch)
//	go watcher(timeout,ch)
//
//	fmt.Println("2x2 = ?")
//	fmt.Scanln(&input)
//
//	if input == 4 {
//		fmt.Println("benar")
//	}else {
//		fmt.Println("salah")
//	}
//
//}
//
//func timer(timeout int, ch chan <- bool)  {
//	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
//		ch <- true
//	})
//}
//
//func watcher(timeout int, ch <- chan bool)  {
//	<-ch
//	fmt.Println("timeout!!!", timeout,"second")
//	os.Exit(0)
//}
//
//func encodeLearning()  {
//	var data = "hello world"
//	fmt.Println([]byte(data))
//	var encoded = base64.StdEncoding.EncodeToString([]byte(data))
//	fmt.Println(encoded)
//	var decoded,_ = base64.StdEncoding.DecodeString(encoded)
//	fmt.Println(string(decoded))
//}
//
//func hashLearning()  {
//	var data = "secret"
//	var sha = sha1.New()
//	sha.Write([]byte(data))
//	var hashed = fmt.Sprintf("%x",sha.Sum(nil))
//	fmt.Println(hashed)
//}
//
//func saltingHashLearning()  {
//	var data = "this is secret"
//	var salt =  fmt.Sprintf("%d",time.Now().UnixNano())
//	fmt.Println(reflect.ValueOf(salt).Kind())
//	var sha = sha1.New()
//	var salted = fmt.Sprintf("text: %s salt: %s",data,salt)
//	fmt.Println(salted)
//	sha.Write([]byte(salted))
//	var hashed = fmt.Sprintf("%x",sha.Sum(nil))
//	fmt.Println(hashed)
//
//}
//
//func flagArgLearning()  {
//	var data = os.Args
//	fmt.Printf("%#v \n", data)
//	var filter = data[1:]
//	fmt.Printf("%#v \n",filter)
//
//	var name = flag.String("name","Anonymous","user name")
//	var age = flag.Int64("age",18,"your age")
//	flag.Parse()
//
//	fmt.Printf("name: %s \n", *name)
//	fmt.Printf("age : %d \n", *age)
//
//}
//
//func execLearning(){
//	var outputs,_ = exec.Command("ls").Output()
//	fmt.Println(string(outputs))
//
//	var output2,_ = exec.Command("git","config","user.name").Output()
//	fmt.Println(string(output2))
//}
//
//func httpListenerLearning()  {
//	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
//		fmt.Fprintln(writer, "hello world")
//	})
//
//	http.HandleFunc("/index",index)
//
//	fmt.Println("starting web on port 8080")
//	http.ListenAndServe(":8080",nil)
//
//}
//
//func index(writer http.ResponseWriter, request *http.Request)  {
//	fmt.Fprintln(writer, "index page")
//}
//
//type datajson struct {
//	FullName string `json:"Name"` //tag ini digunakan yang berarti "data json yang bernama Name itu masuk ke data struct FullName"
//	Roles string
//}
//
//
//func jsonLearning()  {
//	var data datajson
//
//	var jsonString = `{"Name": "Fahmi","Roles": "admin"}`
//	var jsonData = []byte(jsonString)
//
//	var err = json.Unmarshal(jsonData,&data)
//	if err != nil {
//		fmt.Println("error")
//		fmt.Println(err.Error())
//		return
//	}
//
//	fmt.Println(data.FullName)
//	fmt.Println(data.Roles)
//}
//
//type employee struct {
//	Id int
//	Name string
//	Roles string
//}
//var data = []employee{
//	employee{1,"Fahmi","admin"},
//	employee{2,"Ulul","user"},
//	employee{3,"Azmi","editor"},
//}
//
//func jsonAPILearning()  {
//
//	http.HandleFunc("/",index)
//	http.HandleFunc("/employes",employes )
//	http.HandleFunc("/employe",employe )
//	fmt.Println("starting go apps on port 8080")
//	http.ListenAndServe(":8080",nil)
//}
//
//func employes(writer http.ResponseWriter, request *http.Request) {
//	writer.Header().Set("Content-Type","application/json")
//	if request.Method == "POST" {
//		var res,err = json.Marshal(data)
//		if err != nil {
//			http.Error(writer,err.Error(),http.StatusInternalServerError)
//		}
//
//		writer.Write(res)
//		return
//
//	}else if request.Method == "GET" {
//
//	}
//}
//
//func employe(writer http.ResponseWriter, request *http.Request)  {
//	writer.Header().Set("Content-Type","application/json")
//
//	if request.Method == "POST" {
//		var id,errors = strconv.Atoi( request.FormValue("id"))
//		if errors != nil {
//			http.Error(writer, errors.Error(),http.StatusInternalServerError)
//			return
//		}
//
//		var result []byte
//		var err error
//
//		for _,res := range data {
//			if id == res.Id {
//				result,err = json.Marshal(res)
//				if err != nil {
//					http.Error(writer, errors.Error(),http.StatusInternalServerError)
//					return
//				}
//				writer.Write(result)
//				return
//			}
//		}
//
//		http.Error(writer,"user not found guys",404)
//		return
//	}
//}
//
//const BaseUrl  = "http://localhost:8080"
//
//func httpLearning()  {
//	fmt.Println("start web server on port 8080")
//	http.HandleFunc("/employe",employe)
//	http.HandleFunc("/employes",employes)
//	http.HandleFunc("/fetch",fetchData)
//	http.ListenAndServe(":8080",nil)
//}
//
//func fetchData(writer http.ResponseWriter, r *http.Request)  {
//	var result,err = fetchEmployee()
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	writer.Header().Set("Content-Type","application/json")
//	data,err := json.Marshal(result)
//	if err != nil {
//		http.Error(writer,err.Error(),http.StatusInternalServerError)
//	}
//	writer.Write(data)
//
//	for _,each := range result  {
//		fmt.Println(each.Id)
//	}
//}
//
//func fetchEmployee()([]employee, error)  {
//	var client = http.Client{}
//	var err error
//	var data []employee
//
//	request,err := http.NewRequest("POST",BaseUrl+"/employes",nil)
//	if err != nil {
//		return nil,err
//	}
//
//	response,err := client.Do(request)
//	if err != nil {
//		return nil,err
//	}
//
//	defer response.Body.Close()
//	err = json.NewDecoder(response.Body).Decode(&data)
//	if err != nil {
//		return nil,err
//	}
//	return data,nil
//}
//
//func httpFormLearning(w http.ResponseWriter,r *http.Request)  {
//	var err error
//	var client = &http.Client{}
//	var data employee
//
//	var params = url.Values{}
//
//	params.Set("id","2")
//	fmt.Println(params)
//	var payload = bytes.NewBufferString(params.Encode())
//	fmt.Println(payload)
//
//	req,err := http.NewRequest("POST",BaseUrl+"/employe",payload)
//
//	if err != nil {
//		http.Error(w,err.Error(),http.StatusInternalServerError)
//		return
//	}
//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//
//	response,err := client.Do(req)
//
//	if err != nil {
//		http.Error(w,err.Error(),http.StatusInternalServerError)
//		return
//	}
//	defer response.Body.Close()
//
//	err = json.NewDecoder(response.Body).Decode(&data)
//	if err != nil {
//		http.Error(w,err.Error(),http.StatusInternalServerError)
//		return
//	}
//
//	fmt.Println(data)
//	return
//
//}
//
//func sqlLearning()  {
//	db, err := connectSql()
//	if err != nil{
//		fmt.Println(err.Error())
//		return
//	}
//	defer db.Close()
//
//	var query = "SELECT * FROM employee where id = ? OR name like ?"
//	rows ,err := db.Query(query,3,"%J%")
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//
//	defer rows.Close()
//
//	var result []employee
//
//	for rows.Next()  {
//		var each = employee{}
//		var err = rows.Scan(&each.Id,&each.Name,&each.Roles)
//
//		if err != nil {
//			fmt.Println(err.Error())
//			return
//		}
//
//		result = append(result,each)
//	}
//
//	if err = rows.Err(); err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//
//	fmt.Println(result)
//
//	//single query
//	var single = "SELECT id, name, roles FROM employee where id = ?"
//	var res  = employee{}
//	err = db.QueryRow(single,1).
//		Scan(&res.Id,&res.Name,&res.Roles)
//
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//
//	fmt.Println(res)
//
//}
//
//func connectSql()(*sql.DB,error)  {
//	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/golang")
//	if err != nil{
//		return nil,err
//	}
//
//	return db,nil
//}
//
//func prepareSqlLerning()  {
//	db, err := connectSql()
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	defer db.Close()
//
//	stmt, err := db.Prepare("SELECT id,name,roles FROM employee where id = ?")
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	defer stmt.Close() // *stmt ini harus di close juga
//
//	var res1 = employee{}
//	stmt.QueryRow(1).Scan(&res1.Id,&res1.Name,&res1.Roles)
//	fmt.Println(res1)
//	var res2 = employee{}
//	stmt.QueryRow(2).Scan(&res2.Id,&res2.Name,&res2.Roles)
//	fmt.Println(res2)
//	var res3 = employee{}
//	stmt.QueryRow(3).Scan(&res3.Id,&res3.Name,&res3.Roles)
//	fmt.Println(res3)
//}
//
//func dbexecLearning()  {
//	db,err := connectSql()
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	defer db.Close()
//	insertStmt ,err := db.Prepare("INSERT INTO employee (name,roles) values (?,?)")
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	defer insertStmt.Close()
//	data,_ := insertStmt.Exec("last","tes")
//	lookData(db)
//	id,_ := data.LastInsertId()
//	fmt.Println(id)
//
//	updateStmt,err := db.Prepare("update employee set name = ?, roles = ? where id=?")
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	defer updateStmt.Close()
//
//	updateStmt.Exec("final","developer test",id)
//	lookData(db)
//
//	deleteStmt, err := db.Prepare("delete from employee where id = ?")
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	defer deleteStmt.Close() //dont forget to close it
//
//	deleteStmt.Exec(id)
//	lookData(db)
//
//}
//
//func lookData(db *sql.DB){
//	rows , err := db.Query("select * from employee")
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	defer rows.Close()
//	var final []employee
//	for rows.Next()  {
//		var result = employee{}
//
//		rows.Scan(&result.Id,&result.Name,&result.Roles)
//
//		final = append(final,result)
//	}
//	fmt.Println(final)
//}
//
//type student struct {
//	Name string `bson:"name"`
//	Class string `bson:"class"`
//}
func connectNoSql()(*mgo.Session, error)  {
	var session, err = mgo.Dial("localhost")
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}

	return session,nil
}
//
//func nosqlLearning()  {
//	defer catch()
//	var db, err = connectNoSql()
//	if err != nil {
//		panic(err.Error())
//		fmt.Println(err.Error())
//		return
//	}
//	defer db.Close()
//
//	//view data
//	var collection = db.DB("golang").C("student")
//	lookMongo(collection)
//
//	//insert data
//	err = collection.Insert(&student{"this","dua"})
//	if err != nil{
//		panic(err.Error())
//		fmt.Println(err.Error())
//		return
//	}
//	lookMongo(collection)
//
//	var updated = make(chan bool)
//
//	//update data
//	var selector = bson.M{"name":"this"}
//	go updateData(collection,selector,bson.M{"$set": bson.M{"name": "that"}},updated)
//
//	lookMongo(collection)
//	//delete data
//	<-updated
//	err = collection.Remove(bson.M{"name":"that"})
//	if err != nil {
//		panic(err.Error())
//		return
//	}
//	lookMongo(collection)
//
//}
//
//func lookMongo(collection *mgo.Collection)  {
//	var res []student
//	err := collection.Find(nil).All(&res)
//	if err != nil {
//		panic(err.Error())
//		fmt.Println(err.Error())
//		return
//	}
//	fmt.Println(res)
//}
//
//func updateData(collection *mgo.Collection,selector interface{},changes interface{},updated chan <- bool )(error)  {
//	err := collection.Update(selector,changes)
//	if err != nil {
//		panic(err.Error())
//		return err
//	}
//	updated <- true
//	return nil
//}

type User struct {
	ID bson.ObjectId `bson:"_id"`
	Username string `bson:"username"`
	Email string `bson:"email"`
}

type Method interface {
	insert() error
	getData() (User,error)
}

func (user *User) insert() error{
	db,err := connectNoSql()
	if err != nil {
		log.Fatal("cannot dial mongo",err)
		return err
	}
	defer db.Close()

	collection := db.DB("golang").C("users")

	err = collection.Insert(user)
	if err != nil {
		log.Fatal("cannot insert data", err)
		return err
	}

	return nil
}

func(user *User) getData() (User,error)  {
	var result = User{}

	db,err := connectNoSql()
	if err != nil {
		log.Fatal(err)
		return  result,err
	}
	defer db.Close()

	var selector = bson.M{"_id": user.ID}

	var collection = db.DB("golang").C("users")

	err = collection.Find(selector).One(&result)
	if err != nil {
		log.Fatal(err)
		return result,err
	}
	return result,nil

}
func insertData()  {
	var id = bson.NewObjectId()
	fmt.Println(id)
	var user Method = &User{id,"tes","es@gmail.com"}
	err := user.insert()
	if err != nil {
		log.Fatal(err)
	}
	 data,err := user.getData()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(data)

}

func doCountAndPrint(wg *sync.WaitGroup, x int, y int)  {
	var result = x+y
	wg.Done()
	fmt.Println(result)
}

func waitGroupLearning()  {
	var wg sync.WaitGroup

	for i := 0; i < 10 ;i++  {
		wg.Add(1)
		go doCountAndPrint(&wg,1,i)
	}
	wg.Wait()

	fmt.Println("done")
}

func mutexLearning()  {
	var wg sync.WaitGroup
	var meter counter

	for i := 0; i < 1000 ; i++  {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000 ; j++  {
				meter.add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(meter.Value())
}

type counter struct {
	sync.Mutex
	val int
}

func (counter *counter) add( x int)  {
	counter.Lock()
	counter.val++
	counter.Unlock()
}

func (counter *counter) Value() (x int) {
	return counter.val
}

func webHttpLearning()  {
	http.HandleFunc("/",handlerIndex)
	http.HandleFunc("/blog",handlerBlog)
	http.HandleFunc("/close", func(writer http.ResponseWriter, request *http.Request) {
		var message = "using closure statement"
		writer.Write([]byte(message))
	})
	var err error

	//1 using http.ListenAndServe
	//err = http.ListenAndServe(":8080",nil)
	//2 using http.Server
	var server = new(http.Server)
	server.Addr = ":8080"
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request)  {
	var message = "hello guys"
	w.Write([]byte(message))
}

func handlerBlog(w http.ResponseWriter, r *http.Request)  {
	var message = "blog page"
	w.Write([]byte(message))
}