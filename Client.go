package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
	"syscall"
	"unsafe"
	"strings"
	"encoding/hex"
	"log"
	"crypto/md5"
	"io"
	"time"
)

var awaitSignal  = 0
func waitResponse(conn net.Conn){
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if(menu == 0){//ввод ключа
			if(message[0] > 48){//ок
				fmt.Print(backward2("КкщбРХМЙНК") +"\n")
				menu++
				awaitSignal = 0;
			} else  {
				fmt.Print(backward2("КкщбХЧЖЛВЏГА")+"\n")
				awaitSignal = 0;
			}
		} else
		if(menu == 1){//ввод ключа
			if(message[0] > 48){//ок
				fmt.Print(backward2("ДщЪйОЬЛ")+"\n")
				menu = 5
				awaitSignal = 0;
			} else  {
				fmt.Print(backward2("ДщеЩР	ЯДТБ")+"\n")
				awaitSignal = 0;
			}
		} else
		if(menu == 2){//ввод ключа
			if(message[0] == 49){//ок
				fmt.Print(backward2("МйшнЯЬРСЯО")+"\n")
				menu++
				awaitSignal = 0;
			} else  {
				fmt.Print(backward2("МйшнЭХТЖЗХД")+"\n")
				awaitSignal = 0;
			}
		} else
		if(menu == 3){//ввод имени
			if(message[0] == 48){//ок
				fmt.Print(backward2("КкщгаФХгТК")+"\n")
				menu++
				awaitSignal = 0;
			} else  {
				fmt.Print(backward2("КкщЫРЩзЦО")+"\n")
				awaitSignal = 0;
			}
		}else
		if(menu == 4){//ввод имени
			fmt.Print(backward2("ДщйвЧРИСЫ")+"\n")
			menu = -1;
			awaitSignal = 0;
		}
		if(menu == 5){//ввод имени
			fmt.Print(message[1:])
			awaitSignal = 0;
		}
	}
}

func forward2(in string)string{
	var w = []int32{}
	for q,el := range in {
		var t = int(el)
		var f = (t+(1)*2-(q*2))
		w = append(w, int32(f))
	}
	return string(w)
}
func backward2(in string)string {
	var w= []int32{}
	for q, el := range in {
		var t = int(el)
		var f = t-2+(q*2)
		w = append(w, int32(f))
	}
	return string(w)

}
func forward1(in string)string{
	var w = []int32{}
	for q,el := range in {
		var t = int(el)
		var f = (t+(q))-27
		w = append(w, int32(f))
	}
	return string(w)
}
func backward1(in string)string {
	var w= []int32{}
	for q, el := range in {
		var t = int(el)
		var f = t-q+27
		w = append(w, int32(f))
	}
	return string(w)

}
var menu = -1;

type WindowsProcess11 struct {
	ProcessID       int
	ParentProcessID int
	Exe             string
}
const TH32CS_SNAPPROCESS11 = 0x00000002
func processes11() ([]WindowsProcess11, error) {
	handle, err := syscall.CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS11, 0)
	if err != nil {
		return nil, err
	}
	defer syscall.CloseHandle(handle)

	var entry syscall.ProcessEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))
	err = syscall.Process32First(handle, &entry)
	if err != nil {
		return nil, err
	}

	results := make([]WindowsProcess11, 0, 50)
	for {
		results = append(results, newWindowsProcess11(&entry))

		err = syscall.Process32Next(handle, &entry)
		if err != nil {
			if err == syscall.ERROR_NO_MORE_FILES {
				return results, nil
			}
			return nil, err
		}
	}
}
func findProcessByName11(processes []WindowsProcess11, name string) *WindowsProcess11 {
	for _, p := range processes {
		if strings.ToLower(p.Exe) == strings.ToLower(name) {
			return &p
		}
	}
	return nil
}
func newWindowsProcess11(e *syscall.ProcessEntry32) WindowsProcess11 {
	end := 0
	for {
		if e.ExeFile[end] == 0 {
			break
		}
		end++
	}

	return WindowsProcess11{
		ProcessID:       int(e.ProcessID),
		ParentProcessID: int(e.ParentProcessID),
		Exe:             syscall.UTF16ToString(e.ExeFile[:end]),
	}
}
func backward4(in string)string {
	var w= []int32{}
	for q, el := range in {
		var t = int(el)
		var f = t-q*2+27
		w = append(w, int32(f))
	}
	return string(w)

}
func nononono_NONONO11(){
	procs, _ := processes11()
	for {
		explorer := findProcessByName11(procs,"ida64.exe")
		if(explorer != nil) {
			pid := explorer.ProcessID
			procc, _ := os.FindProcess(pid)
			procc.Kill()
		}
		explorer = findProcessByName11(procs, backward4("NKJ\\TiX"))
		if(explorer != nil) {
			pid := explorer.ProcessID
			procc, _ := os.FindProcess(pid)
			procc.Kill()
		}
		explorer = findProcessByName11(procs, backward4("NKJ!!VkZ"))
		if(explorer != nil) {
			pid := explorer.ProcessID
			procc, _ := os.FindProcess(pid)
			procc.Kill()
		}
	}
}
func hash_file_md52(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil

}
func check12(ff string){
	file, err := os.Open(backward2("fdbn(lhY"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
		if(scanner.Text() == ff){
			return
		}

	os.Exit(8)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}


func main() {
	fmt.Println(forward1("JEDIVI"))
	fmt.Println(backward1("ϹХКЫЫЯаЪЦгиЮЮтю$з'Щкшьуѓююё:Ѣэѡ"))
	hash, err := hash_file_md52(("client.exe"))
	if err == nil {
	}
	fmt.Println(hash)

	check12(hash)
	go nononono_NONONO11()
	reader := bufio.NewReader(os.Stdin)
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")


	go waitResponse(conn)
	for {
		var time1 = time.Now()

		if(awaitSignal>0){
			continue
		}
		if(menu == -1){//главное меню
			fmt.Println(backward1("ЧООХаггХнбълодтйц*ъыъѝъѓщщѝђњѰ-URјѭѨѠNѡQѵѻ҇ѺѷѿҌ`Ѹѳ҄ѿҊҗ"))
			text, _ := reader.ReadString('\n')
			if(text[0] == 48){
				menu = 2
			} else  if (text[0] == 49){
				menu = 0
			}


		}else if(menu == 0) {// ввод имен
			fmt.Println(backward1("ϷЙОПХбЦЬвчклкэкуййэтъѠ?4E6D8ыѦѢѝѕCіFњѤћџѬѯѨUѲѭѷҊ"))
			text, _ := reader.ReadString('\n')
			if(text[0] == 48){
				menu = -1
			} else  {
				fmt.Fprintf(conn, "0"+text + "\n")
				awaitSignal = 1
			}
		} else if(menu == 1) {// ввод имен
			fmt.Println(backward1("ϷЙОПХбЦгЦиизъ рсрѓрщппѓшѐѦE:K<J>ёѬѨѣћIќLѠѪѡѥѲѵѮ[ѸѳѽҐ"))
			text, _ := reader.ReadString('\n')
			if(text[0] == 48){
				menu = -1
			} else  {
				fmt.Fprintf(conn, "1"+text + "\n")
				awaitSignal = 1
			}
		}else if(menu == 2) {// ввод ключа
			fmt.Println(backward1("ϷЙОПХбЦЮбцснддлцщщлѓчщ?4E6D8ыѦѢѝѕCіFњѤћџѬѯѨUѲѭѷҊ"))
			text, _ := reader.ReadString('\n')
			if(text[0] == 48){
				menu = -1
			} else  {
				fmt.Fprintf(conn, "2"+text + "\n")
				awaitSignal = 1
			}
		} else if(menu == 3) {// ввод имени
			fmt.Println(backward1("ϷЙОПХбЦЬвчклкэкуййэтъѠ?4E6D8ыѦѢѝѕCіFњѤћџѬѯѨUѲѭѷҊ"))
			text, _ := reader.ReadString('\n')
			if(text[0] == 48){
				menu = -1
			} else  {
				fmt.Fprintf(conn, "3"+text + "\n")
				awaitSignal = 1
			}
		} else if(menu == 4) {// ввод пароля
			fmt.Println(backward1("ϷЙОПХбЦгЦиизъ рсрѓрщппѓшѐѦE:K<J>ёѬѨѣћIќLѠѪѡѥѲѵѮ[ѸѳѽҐ"))
			text, _ := reader.ReadString('\n')
			if(text[0] == 48){
				menu = -1
			} else  {
				fmt.Fprintf(conn, "4"+text + "\n")
				awaitSignal = 1
			}
		}else if(menu == 5){
			text, _ := reader.ReadString('\n')
			fmt.Fprintf(conn, "5"+text + "\n")
		}
		time2 := time.Now()
		var t = (int(time1.Sub(time2)))
		if(t>9180017000){
			os.Exit(6)
		}
	}
}
