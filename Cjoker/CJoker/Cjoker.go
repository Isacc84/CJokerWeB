package CJoker

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/gorilla/websocket"
	"github.com/shirou/gopsutil/disk"
)

var (
	upgrader = websocket.Upgrader{
		// 读取存储空间大小
		ReadBufferSize: 10240,
		// 写入存储空间大小
		WriteBufferSize: 10240,
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	talkArr         chan []uint8
	messageTypeKage chan int
	DisLis          list.List
	autoMoch        string
	te              string
	tmpDiskse       string
	noUseAr         sync.Map
	ms              []uint8
	er              error
	er0             error
	messageTyp0     int
	p0              []byte
	Pussy           *uint64
	se              string
	pin             list.List
)

var DisList = &DisLis
var autoMochi = &autoMoch
var tmpDiskset = &tmpDiskse
var noUseArr = &noUseAr
var msg = &ms
var err = &er
var tem = &te
var messageType = &messageTyp0
var p1 = &p0
var err1 = &er0
var err0 = &er
var pig = &pin

//输出磁盘剩余情况的类型
func DiskFreeMemory() (*[]string, *[]uint64, *sync.Map, error) {
	mapStat, _ := disk.IOCounters()
	for name, _ := range mapStat {
		DisList.PushBack(name)
	}
	// 获取磁盘信息后输入到两个列表中
	var returnLis = make([]string, 0, 10000)
	var returnLisKaGe = make([]uint64, 0, 10000)
	for item := DisList.Front(); item != nil; item = item.Next() {
		info0, _ := disk.Usage(item.Value.(string))
		// data0, _ := json.MarshalIndent(info0, "", "  ")
		// fmt.Println(info0)
		//如果每一项的值不是空那就循环，不然就读取失败
		if reflect.TypeOf(info0) != nil {
			returnLis = append(returnLis, info0.Path)
			returnLisKaGe = append(returnLisKaGe, info0.Free)
		} else {
			err := errors.New("读取磁盘大失败")
			return nil, nil, nil, err
		}
	}
	var lisRan uint64 = uint64(len(returnLis))
	var lisRanKaGe uint64 = uint64(len(returnLisKaGe))
	if lisRan != lisRanKaGe {
		err := errors.New("获取硬盘列表失败")
		return nil, nil, nil, err
	}
	for keys := lisRanKaGe - 1; uint64(keys) > 0; keys-- {
		for key := lisRanKaGe - 1; uint64(key) > 0; key-- {
			if returnLisKaGe[key] >= returnLisKaGe[key-1] {
				temCup := returnLisKaGe[key]
				returnLisKaGe[key] = returnLisKaGe[key-1]
				returnLisKaGe[key-1] = temCup
				temCup0 := returnLis[key]
				returnLis[key] = returnLis[key-1]
				returnLis[key-1] = temCup0
			}
		}
	}
	//对列表内容进行了排序
	var mapDikMes sync.Map
	for keys := 0; uint64(keys) < lisRan; keys++ {
		mapDikMes.Store(returnLis[keys], returnLisKaGe[keys])
	}
	return &returnLis, &returnLisKaGe, &mapDikMes, nil
}

func OptionFileUse() { //读取设置
	fileInfo, _ := os.Lstat("./option/CJokerSet/CDFSsetfile.txt")
	if fileInfo == nil { //判断是否文件存在，fileInfo是不存在就创建该设置文件
		file, err := os.OpenFile("./option/CJokerSet/CDFSsetfile.txt", os.O_CREATE|os.O_WRONLY, 0777)
		defer file.Close()
		if err != nil {
			fmt.Print("创建设置文件大失败")
			return
		} else {
			fmt.Println("是否要是用自动设置硬盘顺序安磁盘剩余空间排列./option/CJokerSet/CDFSsetfile.txt YESs&No 或附加不加入的磁盘C:")
			fmt.Scan(&autoMochi)
			if *autoMochi == "Y" || *autoMochi == "y" || *autoMochi == "Yes" || *autoMochi == "yes" {
				for {
					fmt.Println("请输入不要加入设置的磁盘,不设置了就说no")
					fmt.Scan(&tem)
					if *tem == "no" || *tem == "No" || *tem == "n" || *tem == "N" {
						break
					} else {
						noUseArr.Store(*tem, true)
					}
				}
				var returnLisIndica, returnLisKaGeIndica, mapDikMeIndica, _ = DiskFreeMemory()
				fmt.Print(returnLisIndica, returnLisKaGeIndica, mapDikMeIndica)
				*tmpDiskset = "UseDiskLis"
				for _, val := range *returnLisIndica {
					vue, staut := noUseArr.Load(val)
					if vue == nil || staut == false {
						*tmpDiskset += "," + val
					}
				}
				*tmpDiskset += "\n"
				writer := bufio.NewWriter(file)
				writer.WriteString(*tmpDiskset)
				writer.Flush()
			} else {
				fmt.Println("请根据,来分割磁盘设置，C:,D:,E:,F:,G:")
				fmt.Scan(&tem)
				*tmpDiskset = "UseDiskLis" + *tem
				writer := bufio.NewWriter(file)
				writer.WriteString(*tmpDiskset)
				writer.Flush()
			}
			fileInfo0, _ := os.Lstat("./option/CDFSOptionPussy/masterSev.csv")
			fileInfo1, _ := os.Lstat("./option/CDFSOptionPussy/masterArr.csv")
			if fileInfo0 == nil {
				fmt.Println("你没有提供主机列表")
				return
			}
			if fileInfo1 == nil {
				fmt.Println("未设置主机，自动配色会根据ping进行Y手动输入N")
				fmt.Scan(&se)
				if se == "Y" || se == "y" || se == "yes" || se == "Yes" {
					file, err := os.OpenFile("./option/CDFSOptionPussy/masterSev.csv", os.O_RDONLY, 0777)
					if err != nil {
						fmt.Println(err)
						return
					} else {
						defer file.Close()
						rd := bufio.NewReader(file)
						for {
							line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
							buf := make([]byte, 1024)
							liste, err := net.Listen("tcp", "0.0.0.0:1810")
							if err != nil {
								fmt.Println(err)
								return
							}
							defer liste.Close()
							diaCon, err := net.Dial("tcp", line+":1812")
							defer diaCon.Close()
							starttime := time.Now().UnixNano()
							if err != nil {
								fmt.Println(err)
								return
							}
							_, err = diaCon.Write([]byte(line + ",txt"))
							if err != nil {
								fmt.Println(err)
								return
							}
							WriCon, err := liste.Accept()
							if err != nil {
								fmt.Println(err)
							}
							n, err := WriCon.Read(buf)
							if err != nil {
								fmt.Println(err)
								return
							}
							if string(buf[:n]) == (line + ",txt") {
								edtime := time.Now().UnixNano()
								fmt.Println((edtime - starttime))
							}
							if err != nil || io.EOF == err {
								break
							}
						}
					}
				}
			}
		}
	}
	// var initSetingArr = [1]string{"UseDiskLis"}
	// file, err := os.Open("./option/CJokerSet/CDFSsetfile.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()

	// rd := bufio.NewReader(file)
	// for {
	// 	line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
	// 	var splitLines = strings.Split(line, ",")
	// 	fmt.Print(splitLines[0])
	// 	var isExist bool = false
	// 	for _, value := range initSetingArr { //查看是否有这个设置，如果有就进行设置,没有就自动写入设置
	// 		if value == splitLines[0] {
	// 			isExist = true
	// 		} /* else {
	// 			SetOptionAutomatch()
	// 		}*/
	// 	}
	// 	fmt.Print(isExist)
	// 	if err != nil || io.EOF == err {
	// 		break
	// 	}
	// }
}

func hander(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("./ChemmsGo/index.html"))
	*err0 = page.Execute(w, "")
	if *err0 != nil {
		log.Println(*err0)
		return
	}
}

func processAss(conn net.Conn) {
	defer conn.Close()
	fmt.Println(conn.RemoteAddr())
	// diaCon, err := net.Dial("tcp", conn.RemoteAddr())
	// defer diaCon.Close()
	// for {
	// 	buf := make([]byte, 1024)
	// 	n, err := conn.Read(buf)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 		return
	// 	}
	// 	_, err = diaCon.Write(buf[:n])
	// }
}

func runLis() {
	lis, err := net.Listen("tcp", "0.0.0.0:1812")
	if err != nil {
		return
	}
	defer lis.Close()
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go processAss(conn)
	}
}

func WriteMessage(conn *websocket.Conn) {
	if len(talkArr) > 0 && len(talkArr) > 0 {
		temWS := <-talkArr
		messageType := <-messageTypeKage
		if err := conn.WriteMessage(int(messageType), temWS); err != nil {
			//if err := conn.WriteMessage(1, []byte("今天。。。"));err != nil {
			log.Println("Writeing error...", err)
			return
		}
	}
}

func noChoose(msg *[]uint8, messageType *int, Beatuetpussy *websocket.Conn) {
	if strings.Split(strings.Split(string(*msg), ",")[0], "=")[0] == "userName" {
		SAOOnlie := [2]string{strings.Split(strings.Split(string(*msg), ",")[0], "=")[1], strings.Split(strings.Split(string(*msg), ",")[1], "=")[1]}
		fileInfo, _ := os.Lstat("./option/CJokerSet/CJokerUserPussy/" + SAOOnlie[0] + ".csv")
		if fileInfo == nil {
			messageTypeKage <- *messageType
			talkArr <- []uint8("userName,false")
		} else {
			file, err := os.OpenFile("./option/CJokerSet/CJokerUserPussy/"+SAOOnlie[0]+".csv", os.O_RDONLY, 0777)
			defer file.Close()
			if err != nil {
				fmt.Println(err)
				messageTypeKage <- *messageType
				talkArr <- []uint8("./option/CJokerSet/CJokerUserPussy/" + SAOOnlie[0] + ".csv err")
			} else {
				rd := bufio.NewReader(file)
				for {
					line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
					fmt.Println(strings.Split(line, ",")[1], SAOOnlie[1])
					fmt.Println((strings.Split(line, ",")[1] == SAOOnlie[1]))
					if strings.Split(line, ",")[0] == SAOOnlie[0] {
						if SAOOnlie[1] == strings.Split(line, ",")[1] {
							messageTypeKage <- *messageType
							talkArr <- []uint8(SAOOnlie[0] + ",oNtrue")
						} else {
							messageTypeKage <- *messageType
							talkArr <- []uint8(SAOOnlie[0] + ",oNfalse")
						}
						break
					}
					fmt.Println(len(talkArr), len(talkArr))
					if err != nil || io.EOF == err {
						break
					}
				}
			}
		}
	}
	WriteMessage(Beatuetpussy)
}

var temSMS string
var SMS = &temSMS

func dilen() {
	a, b, _, _ := DiskFreeMemory()
	for {
		*SMS = "slave:"
		file, err := os.OpenFile("./option/CJokerSet/SlaveServer.csv", os.O_RDONLY, 0777)
		if err != nil {
			fmt.Println("你这瓜不熟怎么办", err)
			continue
		}
		br := bufio.NewReader(file)
		for {
			lin, erl := br.ReadString('\n')
			*SMS += lin
			if erl != nil || io.EOF == erl {
				break
			}
		}
		file.Close()
		*SMS += "?"
		*SMS += "mater:"
		file, err = os.OpenFile("./option/CJokerSet/MaseterServer.csv", os.O_RDONLY, 0777)
		if err != nil {
			fmt.Println("你这瓜不熟怎么办", err)
			continue
		}
		br = bufio.NewReader(file)
		for {
			lin, erl := br.ReadString('\n')
			*SMS += lin
			if erl != nil || io.EOF == erl {
				break
			}
		}
		file.Close()
		*SMS += "?DickLen:"
		for k, v := range *a {
			bb := *b
			*SMS += v + "." + strconv.FormatInt(int64(bb[int64(k)]), 10) + ","
		}
		*SMS += "?upload:"
		temMes, _ := ioutil.ReadDir("./upload/")
		for _, file := range temMes {
			*SMS += file.Name() + "|" + strconv.FormatInt(file.Size(), 10) + ","
		}
		*SMS += "?down:"
		temMes, _ = ioutil.ReadDir("./down/")
		for _, file := range temMes {
			*SMS += file.Name() + "|" + strconv.FormatInt(file.Size(), 10) + ","
		}
		*SMS += "?emo:"
		temMes, _ = ioutil.ReadDir("./temDir/")
		for _, file := range temMes {
			*SMS += file.Name() + "|" + strconv.FormatInt(file.Size(), 10) + ","
		}
		*SMS += "?"
		time.Sleep(time.Second)
		messageTypeKage <- 1
		talkArr <- []uint8(*SMS)
		WriteMessage(azhe)
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//   完成握手 升级为 WebSocket长连接，使用conn发送和接收消息。
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer conn.Close()
	azhe = conn
	messageTypeKage <- 1
	talkArr <- []uint8("周幽王,烽火戏诸侯")
	go dilen()
	for {
		fmt.Println("Pussy", *Pussy)
		*messageType, *msg, *err0 = conn.ReadMessage()
		if *err0 != nil {
			log.Println("Reading error...", *err0)
			conn.Close()
			break
			return
		}
		if string(*msg) == "" {
			continue
		}
		log.Printf("Read from client msg:%v \n", string(*msg))
		// log.Printf("zhuan%T", conn)
		go noChoose(msg, messageType, conn)
	}
	//调用连接的WriteMessage和ReadMessage方法以一片字节发送和接收消息。实现如何回显消息：
	//p是一个[]字节，messageType是一个值为websocket.BinaryMessage或websocket.TextMessage的int。
}

func litPag() {
	http.HandleFunc("/", hander)
	err := http.ListenAndServe(":80", http.FileServer(http.Dir("./ChemmsGO")))
	if err != nil {
		log.Fatal("ListenAndServe", err.Error())
		return
	}
}

const (
	upload_path string = "./upload/"
)

func uploadHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	if r.Method != "GET" {
		//获取文件内容 要这样获取
		file, head, err := r.FormFile("upfile")
		fmt.Println(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		//创建文件
		fW, err := os.Create(upload_path + head.Filename)
		if err != nil {
			fmt.Println("文件创建失败")
			return
		}
		defer fW.Close()
		_, err = io.Copy(fW, file)
		if err != nil {
			fmt.Println("文件保存失败")
			return
		}
		fmt.Println(w, head.Filename+" 保存成功")
		//io.WriteString(w, head.Filename+" 保存成功")
		//http.Redirect(w, r, "/success", http.StatusFound)
		//io.WriteString(w, head.Filename)
	}
}
func stpck() {
	http.HandleFunc("/stock", wsHandler)
	http.HandleFunc("/upload", uploadHandle)
	err := http.ListenAndServe(":1919", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err.Error())
		return
	}
}

var azhe *websocket.Conn

func Main(pussy *uint64) {
	Pussy = pussy
	talkArr = make(chan []uint8, 100)
	messageTypeKage = make(chan int, 100)
	go runLis()
	// app := iris.New()
	// app.RegisterView(iris.HTML("./ChemmsGo", ".html"))
	// app.HandleDir("/", "./ChemmsGo")
	// app.Get("/", func(ctx iris.Context) {
	// 	ctx.View("index.html")
	// })
	// app.Run(iris.Addr(":80"))
	// 监听 地址 端口
	go litPag()
	go stpck()
	fmt.Print("摩罗摩那一")
}

// func SetOptionAutomatch( /*setTar string*/ ) {
// 	file, err := os.OpenFile("./option/CJokerSet/CDFSsetfile.txt", O_WRONLY|O_CREAT, 0777)
// 	if err != nil {
// 		fmt.Printf(err)
// 	}
// }
