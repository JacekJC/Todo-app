package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strings"
)

type Response struct{
	Status string
	StatusCode int
	Data string
}


type Login struct{
	Username string
	Password string
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}




type TaskBody struct{
	Username string
	Taskname string
}

func add_new_task(w *http.ResponseWriter, r *http.Request, d *Database) bool{
	var t TaskBody
	err := json.NewDecoder(r.Body).Decode(&t)
	if(err!=nil){
		fmt.Println("json decode error")
		return false
	}
	fmt.Println(t)
	d.add_task(&t)
	d.save_user_file(t.Username)
	return true
}


func main(){
	
	con_man := new_connmanager()
	db := NewDatabase() 
	
	//db.add_user("Chuube")
	//db.update_user_name_image("Chuube", Fsload_file("Chuube"))

	go con_man.run()

	fmt.Println("test world!!")

	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request){
		enableCors(&w)
		id := strings.TrimPrefix(r.URL.Path, "/user/")
		fmt.Println(id);
		json.NewEncoder(w).Encode(&Response{"success", 200, db.get_user_tasks(id),})
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request){
		fmt.Printf("is making new user!\n")
		enableCors(&w)
		var t Login
		jerr := json.NewDecoder(r.Body).Decode(&t)
		if(jerr!=nil){
			fmt.Fprintf(w, "json error")
		}
		fmt.Println("json is ", t.Username);
		if(!db.contains_name(t.Username)){
			fmt.Println("user did not exist!")
			db.add_new_user(t.Username)
		}else{
			fmt.Println("user existed!")
			db.add_user(t.Username)
		}
		json.NewEncoder(w).Encode(&Response{"success", 200, 
			db.get_user_profile(t.Username),})
	})

	http.HandleFunc("/profile/user/", func(w http.ResponseWriter, r *http.Request){
		enableCors(&w)
		id := strings.TrimPrefix(r.URL.Path, "/profile/user/")
		fmt.Println(id);
		json.NewEncoder(w).Encode(&Response{"success", 200, 
			db.get_user_profile(id),
		})
	})

	http.HandleFunc("/online_users", func(w http.ResponseWriter, r *http.Request){
		enableCors(&w)
		json.NewEncoder(w).Encode(&Response{"success", 200, 
			db.loaded_users(),
		})
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(con_man, w,r)
	})

	http.HandleFunc("/add_task", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		add_new_task(&w, r, db)
	})







	log.Fatal(http.ListenAndServe(":8081", nil))
}
