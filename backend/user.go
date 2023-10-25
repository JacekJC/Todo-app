package main

import(
	"strings"
	"fmt"
	"encoding/json"
)

type User struct{
	name string
	img_url string
	data string
	tasks []*Task
}

type Task struct{
	Name string
	State string
}


type Database struct{
	users map[string]*User
}

func (d *Database )contains_name(name string) bool{
	if(fread_file(name) == "read error"){
		return false
	}
	return true
}

func NewDatabase() *Database {
	return &Database{
		users: make(map[string]*User),
	}
}

func (d *Database) get_data(user string) *User{
	if val, ok := d.users[user]; ok{
		return val;
	}
	return nil;
}

func (d *Database) add_task(task *TaskBody){
	d.users[task.Username].tasks = append(
			d.users[task.Username].tasks, 
			&Task{task.Taskname, "0"})
}

func (d *Database) set_data(user string, data string) *User{
	if val, ok := d.users[user]; ok{
		val.data = data;
		return val;
	}
	return nil;
}

func (d *Database) add_new_user(name string) *User{
	create_user_file(name)
	d.users[name] = &User{name, "empty", fread_file(name), []*Task{}}
	return d.users[name]
}

func (d *Database) loaded_users() string{
	str:= "{users:["
	i:=0;
	for _, v := range d.users{
		if(i>0){
			str+=",";
		}
		str+="{\"name\":\""+v.name+"\", \"img\":"+v.img_url+"\"}"
	}
	str+="]"
	return str
}

func (d *Database) unload_user(name string){
	if _, ok := d.users[name]; ok{
		delete(d.users, name)
	}
}

func (d *Database) add_user(name string) *User{
	if val, ok := d.users[name]; ok{
		return val;
	}
	d.users[name] = user_from_file(name)
	d.update_user_name_image(name, fread_file(name))
	return d.users[name]
}

func user_from_file(name string) *User{
	file := fread_file(name);
	cleaned := strings.Replace(file, " ", "", -1)
	cleaned = strings.Replace(file, "\r", "", -1)
	split := strings.Split(cleaned, "\n")
	return &User{split[0], split[1], file, get_tasks_from_file(split)}
}

func get_tasks_from_file(file []string) []*Task {

	res := []*Task{}
	//name := ""
	fmt.Println("getting tasks from file")
	for i := 2; i < len(file); i++{
		t := &Task{"null_task", "0"}
		line := strings.Replace(file[i], "\r", "", -1)
		line = strings.Replace(file[i], "\n", "", -1)
		json.Unmarshal([]byte(line), &t)
		fmt.Println("task: ", t)
		if(t.Name!="null_task"){
			//res[t.Name] = t
			res = append(res, t)
		}
	}
	fmt.Println(res)
	return res
}

func create_user_file(name string) string {
	str := name+"\n"+"null\n"+"tasks\n"
	fwrite_file(name, str)
	return str
}

func (d *Database) update_user_name_image(name string, file string){
	rep := strings.Replace(file, "\r", "", -1);
	split := strings.Split(rep, "\n")
	d.users[name].name = split[0]
	d.users[name].img_url = split[1]
}

func (d* Database) get_user_profile(name string) string{
	return "{\"name\":\""+d.users[name].name + "\", \"img\":\""+ d.users[name].img_url + "\"}"
}

func sub_string(str string, start int, end int) string{
	res := ""
	for(start < end){
		start++;
		res+=string(str[start])
	}
	return res
}

func (t* Task) get_str() string {
	return "{\"name\":\""+t.Name+"\", \"state\":\""+t.State+"\"}"
}

func (u* User) create_file_text() string {
	cstring := u.name+"\n"+u.img_url;
	for _, v := range u.tasks{
		cstring+="\n"+v.get_str()
	}
	return cstring
}

func (d* Database) save_user_file(name string){
	fmt.Println("saving file for ", name)
	user := d.users[name];
	fwrite_and_replace(name, user.create_file_text())
}

func (d* Database) get_user_tasks(name string) string{
	cstring := "{\"Data\":["
	ind:=0;
	for _, v := range d.users[name].tasks{
		if(ind>0){
			cstring+=","
		}
		cstring+=v.get_str()
		ind+=1;
	}
	cstring+="]}"
	fmt.Println(cstring)
	return cstring
}

