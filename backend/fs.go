package main

import(
	"fmt"
	"os"
)

func IndexOfSubstring(str, Substr string) int{
	for i := 0; i < len(str); i++{
		if str[i:i+len(Substr)] == Substr{
			return i;
		}
	}

	return -1
}

func fwrite_file(file_name string, data string) bool {
	f, err := os.Open("data/user/"+file_name)
	if(err!=nil){
		fmt.Println("file does not exist.. creating file.")	
		f, err := os.Create("data/user/"+file_name)
		if(err!=nil){
			fmt.Println("create failed!!!")	
			return false
		}
		f.Write([]byte(data))
		f.Close()
		return true
	}
	fmt.Println("\nfile saved!\n ", file_name, " :\n", data)
	_, w_err := f.Write([]byte(data))
	if(w_err!=nil){
		fmt.Println("write error!")
		f.Close()
		return false
	}
	f.Close()
	return true
}

func fwrite_and_replace(name string, data string) bool {
	err := fwrite_file(name+"_temp", data)
	if(err==false){
		return false
	}
	r_err := os.Rename("data/user/"+name+"_temp","data/user/"+name)
	if(r_err != nil){
		return false
	}
	return true
}

func fread_file(file_name string) string {

	i := IndexOfSubstring(file_name, "/")
	if(i!=-1){
		return "illegal read"
	}

	i = IndexOfSubstring(file_name, "\\")
	if(i!=-1){
		return "illegal read"
	}

	dat, err := os.ReadFile("data/user/"+file_name)
	if(err!=nil){
		fmt.Println("failed to read file ", file_name);
		return "read error"
	}
	fmt.Println(string(dat))
	return string(dat)

}
