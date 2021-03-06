package main

import(
	"github.com/TrickyGameJolt/GoGameJolt" // Please note, this comes in as package gj and not as package GoGameJolt
	"bufio"
	"strings"
	"fmt"
	"os"
	"time"
)

const gameid="336383"
const gamekey="e8a4b4be97e11da42183a5751cef877b"
var username=""
var token=""


func MyTrim(a string) string{
	return strings.Trim(a," \t\n\r\x00")
}

func RawInput(q string) string{
    buf := bufio.NewReader(os.Stdin)
    fmt.Print(q)
    sentence, err := buf.ReadBytes('\n')
    if err != nil {
        fmt.Println(err)
        return ""
    } else {
		//fmt.Println("*"+string(sentence)+"*")
        return MyTrim(string(sentence))
    }
}

func main(){
	// Authentication Test
	username=RawInput("User Name: ")
	token   =RawInput("    Token: ")
	user:=gj.Login(gameid,gamekey,username,token)
	//fmt.Println("Logged in = ",user.LoggedIn)
	if !user.LoggedIn { os.Exit(0) }
	for{
		fmt.Println("1 = Test session")
		fmt.Println("2 = Test achievement/trophy")
		fmt.Println("3 = Test user score")
		fmt.Println("4 = Test guest score")
		fmt.Println("5 = Test general score fetching")
		fmt.Println("6 = Test user score fetching")
		fmt.Println("Q = Quit")
		kz:=RawInput("Make your choice: ")
		if kz=="q" || kz=="Q" { break }
		switch kz{			
			case "1":
				fmt.Println("Starting session")
				user.StartSession()
				fmt.Println("50 pings")
				for i:=0;i<50;i++{fmt.Print(".")}
				fmt.Print("\r")
				for i:=0;i<50;i++{
					user.Ping()
					fmt.Print("o")
					time.Sleep(1100)
				}
				fmt.Println("\nClosing session")
				user.CloseSession()
				fmt.Println("\n\n")
			case "2":
				trophies:=[]string{"Bronze","Silver","Gold","Platinum"}
				ids:=[]string{"92488","92489","92490","92491"}
				nm:=[]string{"1","2","3","4"}
				fmt.Println("Test Trophies")
				for i,tn := range trophies{ fmt.Println(i+1,"=",tn) }
				tkz:=RawInput("What shall I test? ")
				for i:=0;i<len(nm);i++{
					if nm[i]==tkz {
						if user.AwardTrophy(ids[i]){
							fmt.Println(trophies[i]+" awarded!")
						} else {
							fmt.Println(trophies[i]+" failed!")
						}
					}
				}
				fmt.Println();
			case "3":
				timestamp := int32(time.Now().Unix())
				sts:=fmt.Sprintf("%d",timestamp)
				if user.SubmitScore(sts+" unix time",sts,"341872"){
					fmt.Println("Score ",timestamp," succesfully submitted")
				} else {
					fmt.Println("Too bad. I could not submit ",timestamp)
				}
			case "4":
				timestamp := int32(time.Now().Unix())
				sts:=fmt.Sprintf("%d",timestamp)
				guest:=fmt.Sprintf("Guest %x",timestamp)
				if gj.SubmitGuestScore(guest,gameid,gamekey,sts+" unix time",sts,"341872"){
					fmt.Println("Score ",timestamp," succesfully submitted to a guest")
				} else {
					fmt.Println("Too bad. I could not submit ",timestamp)
				}
			case "5":
				a:=gj.FetchScore(gameid,"","341872",gamekey)
				if a["success"]=="false" { 
					fmt.Println("FAILED!") 
				} else {
					_,ok:=a["score"]
					i:=0
					for ok {
						is:=""
						if i>0 {is=fmt.Sprintf("%d",i)}
						fmt.Println(i+1,a["score"+is],"\tscored by user: ",a["user"+is],"\tscored by guest:",a["guest"+is])
						i++
						_,ok=a[fmt.Sprintf("score%d",i)]
					}
				}
			case "6":
				a:=user.FetchScore("","341872")
				if a["success"]=="false" { 
					fmt.Println("FAILED!") 
				} else {
					_,ok:=a["score"]
					i:=0
					for ok {
						is:=""
						if i>0 {is=fmt.Sprintf("%d",i)}
						fmt.Println(i+1,a["score"+is],"\tscored by user: ",a["user"+is],"\tscored by guest:",a["guest"+is])
						i++
						_,ok=a[fmt.Sprintf("score%d",i)]
					}
				}
			default: 
				fmt.Println("I don't understand. Please try again!")
		}
	}
}
