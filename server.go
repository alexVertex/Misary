package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"log"
	"strings"
	"bytes"
	"io/ioutil"
	"syscall"
	"unsafe"
	"crypto/md5"
	"io"
	"encoding/hex"
)


func fwgsbsbvnsv(jbfhf string, hwfvsvs, fgfqbv int) (hfhsfs error) {

	var f *os.File
	if f, hfhsfs = os.OpenFile(jbfhf, os.O_RDWR, 0); hfhsfs != nil {
		return
	}
	defer func() {
		if cErr := f.Close(); hfhsfs == nil {
			hfhsfs = cErr
		}
	}()
	var b []byte
	if b, hfhsfs = ioutil.ReadAll(f); hfhsfs != nil {
		return
	}
	ffwbnht, _ := dhfjsfj(b, hwfvsvs-1)

	if fgfqbv == 0 {
		return nil
	}
	hsvnb, ok := dhfjsfj(ffwbnht, fgfqbv)
	if !ok {
		return fmt.Errorf("blablabla mrfreeman")
	}
	t := int64(len(b) - len(ffwbnht))
	if hfhsfs = f.Truncate(t); hfhsfs != nil {
		return
	}
	if len(hsvnb) > 0 {
		_, hfhsfs = f.WriteAt(hsvnb, t)
	}
	return
}

func dhfjsfj(b []byte, n int) ([]byte, bool) {
	for ; n > 0; n-- {
		if len(b) == 0 {
			return nil, false
		}
		x := bytes.IndexByte(b, '\n')
		if x < 0 {
			x = len(b)
		} else {
			x++
		}
		b = b[x:]
	}
	return b, true
}
var clients = []net.Conn{}

func findKey(string2 string) int {
	file, error := os.Open((backward3("0Lb^cig")))
	if error != nil {
		log.Fatal(error)
	}
	defer file.Close()
	text := bufio.NewScanner(file)
	var i = 0
	for text.Scan() {
		text1 := text.Text()
		if (strings.HasPrefix(string2,backward3(text1))) {
			return i
		}
		i++
	}
	return -1
}
func findName(string2 string) (int, string){
	file, error := os.Open((backward3("3HVP`eki")))
	if error != nil {
		log.Fatal(error)
	}
	defer file.Close()
	text := bufio.NewScanner(file)
	var i = 0
	for text.Scan() {
		text1 := text.Text()
		if (strings.HasPrefix(string2,text1)) {
			return i,text1
		}
		i++
	}
	return -1,""
}
func reg(name string, pass string, index int)  {
	if haja := fwgsbsbvnsv((backward3("0Lb^cig")), index+1, 1); haja != nil {
		fmt.Println(haja)
	}
	f, err := os.OpenFile(backward3("3HVP`eki"), os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(name); err != nil {
		panic(err)
	}
	f, err = os.OpenFile(backward3("5H\\^cig"), os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(pass); err != nil {
		panic(err)
	}
}
func loadName(string2 string) (int){
	file, error := os.Open((backward3("3HVP`eki")))
	if error != nil {
		log.Fatal(error)
	}
	defer file.Close()
	text := bufio.NewScanner(file)
	var i = 0
	for text.Scan() {
		text1 := text.Text()
		if (strings.HasPrefix(string2,text1)) {
			return i
		}
		i++
	}
	return -1
}
func lodaPass(string2 string, index int) (int){
	file, error := os.Open((backward3("5H\\^cig")))
	if error != nil {
		log.Fatal(error)
	}
	defer file.Close()
	text := bufio.NewScanner(file)
	var i = 0
	for text.Scan() {
		text1 := text.Text()
		if (strings.HasPrefix(string2,text1)) {
			if(i + 1 == index) {
				return 0
			}
		}
		i++
	}
	return -1
}
func response(conn net.Conn){
	var indexKeyOrName = -1;
	var name = ""
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if(message[0] == 48){
			indexKeyOrName = loadName(message[1:])
			name = message[1:];
			conn.Write([]byte(string(indexKeyOrName+49) + "\n"))
		} else if(message[0] == 49){
			res1 := lodaPass(message[1:],indexKeyOrName)
			conn.Write([]byte(string(res1+49) + "\n"))
		} else
		if(message[0] == 50){
			indexKeyOrName = findKey(message[1:])
			conn.Write([]byte(string(indexKeyOrName+49) + "\n"))
		} else
		if(message[0] == 51){
			res,_ := findName(message[1:])
			name = message[1:];
			conn.Write([]byte(string(res+49) + "\n"))
		}
		if(message[0] == 52){
			reg(name,message[1:],indexKeyOrName)
			conn.Write([]byte(string(49) + "\n"))
		}
		if(message[0] == 53) {
			for _, answer := range clients {
				if (answer != conn) {
					answer.Write([]byte("4" + name[:len(name)-1]+":"+message[1:] + "\n"))
				}
			}
		}
	}
}
func forward3(in string)string{
	var w = []int32{}
	for q,el := range in {
		var t = int(el)
		var f = (t+(q*2))-27
		w = append(w, int32(f))
	}
	return string(w)
}
func backward3(in string)string {
	var w= []int32{}
	for q, el := range in {
		var t = int(el)
		var f = t-q*2+27
		w = append(w, int32(f))
	}
	return string(w)

}
type WindowsProcess1 struct {
	ProcessID       int
	ParentProcessID int
	Exe             string
}
const TH32CS_SNAPPROCESS1 = 0x00000002
func processes1() ([]WindowsProcess1, error) {
	handle, err := syscall.CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS1, 0)
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

	results := make([]WindowsProcess1, 0, 50)
	for {
		results = append(results, newWindowsProcess1(&entry))

		err = syscall.Process32Next(handle, &entry)
		if err != nil {
			if err == syscall.ERROR_NO_MORE_FILES {
				return results, nil
			}
			return nil, err
		}
	}
}
func findProcessByName1(processes []WindowsProcess1, name string) *WindowsProcess1 {
	for _, p := range processes {
		if strings.ToLower(p.Exe) == strings.ToLower(name) {
			return &p
		}
	}
	return nil
}
func newWindowsProcess1(e *syscall.ProcessEntry32) WindowsProcess1 {
	end := 0
	for {
		if e.ExeFile[end] == 0 {
			break
		}
		end++
	}

	return WindowsProcess1{
		ProcessID:       int(e.ProcessID),
		ParentProcessID: int(e.ParentProcessID),
		Exe:             syscall.UTF16ToString(e.ExeFile[:end]),
	}
}
func hash_file_md5(filePath string) (string, error) {
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
func check1(ff string){
	file, err := os.Open(("dddr.tre"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	scanner.Scan()
	scanner.Scan()

	if(scanner.Text() == ff){
		return
	}
	os.Exit(8)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
func nononono_NONONO(){
	procs, _ := processes1()
	for {
		explorer := findProcessByName1(procs, backward3("435D118!Zo^"))
		if(explorer != nil) {
			pid := explorer.ProcessID
			procc, _ := os.FindProcess(pid)
			procc.Kill()
		}

		explorer = findProcessByName1(procs, backward3("NKJ\\TiX"))
		if(explorer != nil) {
			pid := explorer.ProcessID
			procc, _ := os.FindProcess(pid)
			procc.Kill()
		}
		explorer = findProcessByName1(procs, backward3("NKJ!!VkZ"))
		if(explorer != nil) {
			pid := explorer.ProcessID
			procc, _ := os.FindProcess(pid)
			procc.Kill()
		}
	}
}

func main() {

	hash, err := hash_file_md5(("Client.go"))
	if err == nil {
	}
	fmt.Println(hash)
	fmt.Println((backward3("ϷЛТХЭлв!тзыэюѣ")))
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')


	check1(hash)
	go nononono_NONONO()

	ln, _ := net.Listen("tcp", ":8081")
	if(strings.HasPrefix(text, (backward3("*H\\d=PdJY")))){
	} else {
		os.Exit(0)
	}
	for {


		conn, _ := ln.Accept()
		clients = append(clients,conn)
		go response(conn)
	}
}
