package main

import (
	CJoker "CJokerTCL/Cjoker/CJoker"
	"bufio"
	"fmt"
	"io"
	_ "net/http"
	"os"
	"strings"
	"time"
)

func CjokerOpt() {
	fmt.Println("请输入需要登陆的用户名,不存在就创建")
	var (
		passwd       string
		passTheword  string
		reZeroKara   string
		rereZeroKara string
		ChongShe     string
		ChongShemi   string
		usePussy     string
		resese       int64
		splitLines   string
		ChongSe      bool
	)
	fmt.Scanln(&usePussy)
	fileInfo, _ := os.Lstat("./option/CJokerSet/CJokerUserPussy/" + usePussy + ".csv")
pasSet:
	if fileInfo == nil || ChongSe == true { //判断是否文件存在，fileInfo是不存在就创建该设置文件
		file, err := os.OpenFile("./option/CJokerSet/CJokerUserPussy/"+usePussy+".csv", os.O_CREATE|os.O_WRONLY, 0777)
		defer file.Close()
		if err != nil {
			fmt.Print("创建设置Cjoker文件大失败")
			return
		} else {
			fmt.Println("接下来请设置用户密码，不要空")
			fmt.Print(usePussy, "密码为:")
			fmt.Scanln(&passwd)
			fmt.Println("")
			fmt.Print(usePussy, "请再次输入:")
			fmt.Scanln(&passTheword)
			fmt.Println("")
			if passTheword != passwd {
				fmt.Println("宁这两次密码不一样啊，重来！！！这么粗心还想开军舰")
				goto pasSet
			}
			fmt.Print(usePussy, "从零开始的异世界设置为，不要空:")
			fmt.Scanln(&reZeroKara)
			fmt.Println("")
			fmt.Print(usePussy, "再从零开始的异世界设置:")
			fmt.Scanln(&rereZeroKara)
			fmt.Println("")
			if reZeroKara != rereZeroKara {
				fmt.Println("宁这两次从零不一样啊，重来！！！这么粗心还想开军舰")
				goto pasSet
			}
			fmt.Println("宁设置的密码为", passTheword)
			fmt.Println("宁设置的再从零开始的异世界设置为", rereZeroKara)
			var tmpPassset string = usePussy + "," + passTheword + ",\n" + "ゼロから始める異世界生活," + rereZeroKara + ",\n" + "pornography,,\n"
			writer := bufio.NewWriter(file)
			writer.WriteString(tmpPassset)
			writer.Flush()
		}
	} else {
		fmt.Println("已有用户设置，是否重设\nY&N")
		file, err := os.OpenFile("./option/CJokerSet/CJokerUserPussy/"+usePussy+".csv", os.O_RDONLY, 0777)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		rd := bufio.NewReader(file)
		for {
			line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
			if err != nil || io.EOF == err {
				if resese == 0 {
					ChongSe = true
					goto pasSet
					break
				}
				break
			}
			if strings.Split(line, ",")[0] == "ゼロから始める異世界生活" {
				splitLines = strings.Split(line, ",")[1]
			}
			resese += 1
		}
		if splitLines == "" {
			fmt.Println("您这个瓜哪够十五斤啊，你这秤有问题，看，吸铁石，你自己去设置罢./option/CJokerSet/CJokerUserPussy/")
		}
		fmt.Scanln(&ChongShe)
		if ChongShe == "Y" || ChongShe == "y" || ChongShe == "Yes" || ChongShe == "yes" {
			fmt.Println("从零开始的异世界生活，请输入之前进入异世界的凭据")
			fmt.Scanln(&ChongShemi)
			if ChongShemi == splitLines {
				fmt.Println(ChongShemi, splitLines)
				ChongSe = true
				goto pasSet
			} else {
				fmt.Println("小朋友们，偷偷看好康的是大人才能做的事，不要为了看好康的横冲直撞，不然就会被大人丢进垃圾桶里，保护环境，从你我坐起")
			}
		}
	}
}

func main() {
	// CJokerDisk.DiskFreeMemory()
	// CJoker.OptionFileUse()
	fmt.Println(`団長、車の用意できました
▏m
█▏　､⺍
█▏ ⺰ʷʷｨ
█◣▄██◣
◥██████▋
　◥████ █▎
　　███▉ █▎
　◢████◣⌠ₘ℩
　　██◥█◣\≫
　　██　◥█◣AAA
　　█▉　　█▊
　　█▊　　█▊
　　█▊　　█▋
　　 █▏　　█▙
　　 █ 
だからよ...止まるじゃねえぞ
欢迎使用Chemms柴犬文件系统，接下来请根据步骤完成配置
​`)
	// CjokerOpt()
	// app := iris.New()
	// app.RegisterView(iris.HTML("./ChemmsGo", ".html"))
	// app.HandleDir("/", "./ChemmsGO")
	// app.Get("/", func(ctx iris.Context) {
	// 	ctx.View("index.html")
	// })
	// go app.Run(iris.Addr(":80"))
	var a uint64
	var Pussy = &a
	CJoker.Main(&a)
	for {
		fmt.Print("卡其脱离太", *Pussy, "分钟")
		time.Sleep(time.Minute)
		*Pussy = *Pussy + 1
	}
}
