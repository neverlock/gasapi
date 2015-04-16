package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"time"
	"html"
	"strings"
	"github.com/gorilla/mux"
	"github.com/PuerkitoBio/goquery"
	"encoding/json"
)

	var bcplastUpdate,bcpDiesel,bcpGasoholE85,bcpGasoholE20,bcpGasohol91,bcpGasohol95,bcpNGV string
	var bluelastUpdate,blueGasoline95,blueGasohol91,blueGasohol95,blueGasoholE20,blueGasoholE85,blueDiesel,hyForceDiesel,blueNGV string
	var current string
	var currentHR,getdataHR int

func main() {
/*
	http.HandleFunc("/", hello)
	bind := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	fmt.Printf("listening on %s...", bind)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		panic(err)
	}
*/
        rtr := mux.NewRouter()
        rtr.HandleFunc("/ptt",pttPrice).Methods("GET")
        rtr.HandleFunc("/bcp",bcpPrice).Methods("GET")
        http.Handle("/", rtr)
//	bind := ":8080"
        bind :=fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	fmt.Printf("listening on %s...", bind)
        http.ListenAndServe(bind, nil)

}

func httplog(r *http.Request){
        //ip,_,_ := net.SplitHostPort(r.RemoteAddr)
        log.Printf("%s - %s - %s - %s - %q",
                r.RemoteAddr,
                r.Proto,//
                r.Method,//
                r.UserAgent(),
                html.EscapeString(r.URL.Path),
        )
}

func bcpPrice (w http.ResponseWriter, r *http.Request){
	httplog(r)
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now:=time.Now().In(loc)
	now1:=fmt.Sprint(now.Year(),now.Month(),now.Day())
	currentHR=now.Hour()
	log.Println(now1)
	log.Println(currentHR)
	if (current == "")||(current != now1) {
		if (getdataHR < 5) && (currentHR >= 5) {
		loc, _ := time.LoadLocation("Asia/Bangkok")
		now:=time.Now().In(loc)
		current=fmt.Sprint(now.Year(),now.Month(),now.Day())
		getdataHR=currentHR
		log.Println(current)
		log.Println(getdataHR)
		getbcpPrice()
		getpttPrice()
		}
	}
	mapD := map[string]string{"LastUpdate":bcplastUpdate,"bcpDiesel":bcpDiesel,"bcpGasoholE85":bcpGasoholE85,"bcpGasoholE20":bcpGasoholE20,"bcpGasohol91":bcpGasohol91,"bcpGasohol95":bcpGasohol95,"NGV":bcpNGV}
	js,_ := json.Marshal(mapD)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(js)
}

func getbcpPrice (){
        doc, err := goquery.NewDocument("http://www.bangchak.co.th/oilprice-widget.aspx")
        if err != nil {
                log.Fatal(err)
        }

	bcplastUpdate = strings.TrimSpace(doc.Find("table tr td tr td.css1").Text())

	var data [50]string
	doc.Find("tr td tr").Each(func(i int,s *goquery.Selection) {
		data[i] = strings.TrimSpace(s.Find("td.css2").Text())
		if (data[i] != ""){
			fmt.Printf("Review %d: %s\n", i, data[i])
		}
	})

/*
Review 3: 25.09
Review 4: 21.98
Review 5: 24.68
Review 6: 26.08
Review 7: 27.40
Review 9: 13.00
*/

	bcpDiesel = data[3]
	bcpGasoholE85 = data[4]
	bcpGasoholE20 = data[5]
	bcpGasohol91 = data[6]
	bcpGasohol95 = data[7]
	bcpNGV = data[9]


}

func pttPrice (w http.ResponseWriter, r *http.Request){
	httplog(r)
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now:=time.Now().In(loc)
	now1:=fmt.Sprint(now.Year(),now.Month(),now.Day())
	currentHR=now.Hour()
	log.Println(now1)
	log.Println(currentHR)

	if (current == "")||(current != now1) {
		if (getdataHR < 5) && (currentHR >= 5) {
		loc, _ := time.LoadLocation("Asia/Bangkok")
		now:=time.Now().In(loc)
		current=fmt.Sprint(now.Year(),now.Month(),now.Day())
		getdataHR=currentHR
		log.Println(current)
		log.Println(getdataHR)
		getbcpPrice()
		getpttPrice()
		}
	}
	mapD := map[string]string{"LastUpdate":bluelastUpdate,"blueGasoline95":blueGasoline95,"blueGasohol91":blueGasohol91,"blueGasohol95":blueGasohol95,"blueGasoholE20":blueGasoholE20,"blueGasoholE85":blueGasoholE85,"blueDiesel":blueDiesel,"hyForceDiesel":hyForceDiesel,"NGV":blueNGV}
	js,_ := json.Marshal(mapD)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(js)
}

func getpttPrice (){
        doc, err := goquery.NewDocument("http://www.pttplc.com/th/getoilprice.aspx")
        if err != nil {
                log.Fatal(err)
        }

	bluelastUpdate = strings.TrimSpace(doc.Find("div div div span.pttplc-oilpricebanner-row-datetime-format").Text())

	var data [50]string
	doc.Find("div div div div div div").Each(func(i int,s *goquery.Selection) {
		data[i] = strings.TrimSpace(s.Find("div.pttplc-oilpricebanner-row-oilprice-price").Text())
		if (data[i] != ""){
			fmt.Printf("Review %d: %s\n", i, data[i])
		}
	})

/*
Review 1: 33.96
Review 5: 26.08
Review 9: 27.40
Review 13: 24.68
Review 17: 21.98
Review 21: 25.09
Review 27: 28.09
Review 31: 13.00
*/

	blueGasoline95 = data[1]
	blueGasohol91 = data[5]
	blueGasohol95 = data[9]
	blueGasoholE20 = data[13]
	blueGasoholE85 = data[17]
	blueDiesel = data[21]
	hyForceDiesel = data[27]
	blueNGV = data[31]

}
