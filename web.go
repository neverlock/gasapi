package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"strings"
	"github.com/gorilla/mux"
	"github.com/PuerkitoBio/goquery"
	"encoding/json"
)

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

	//Uncomment for run on normal server
	//bind := ":8080" 
	//use Getenv cuz dev on Openshif
        bind :=fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	fmt.Printf("listening on %s...", bind)
        http.ListenAndServe(bind, nil)

}

func bcpPrice (w http.ResponseWriter, r *http.Request){
        doc, err := goquery.NewDocument("http://www.bangchak.co.th/oilprice-widget.aspx")
        if err != nil {
                log.Fatal(err)
        }

	lastUpdate := strings.TrimSpace(doc.Find("table tr td tr td.css1").Text())

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

	bcpDiesel := data[3]
	bcpGasoholE85 := data[4]
	bcpGasoholE20 := data[5]
	bcpGasohol91 := data[6]
	bcpGasohol95 := data[7]
	NGV := data[9]


	mapD := map[string]string{"LastUpdate":lastUpdate,"bcpDiesel":bcpDiesel,"bcpGasoholE85":bcpGasoholE85,"bcpGasoholE20":bcpGasoholE20,"bcpGasohol91":bcpGasohol91,"bcpGasohol95":bcpGasohol95,"NGV":NGV}
	js,_ := json.Marshal(mapD)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(js)
}

func pttPrice (w http.ResponseWriter, r *http.Request){
        doc, err := goquery.NewDocument("http://www.pttplc.com/th/getoilprice.aspx")
        if err != nil {
                log.Fatal(err)
        }

	lastUpdate := strings.TrimSpace(doc.Find("div div div span.pttplc-oilpricebanner-row-datetime-format").Text())

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

	blueGasoline95 := data[1]
	blueGasohol91 := data[5]
	blueGasohol95 := data[9]
	blueGasoholE20 := data[13]
	blueGasoholE85 := data[17]
	blueDiesel := data[21]
	hyForceDiesel := data[27]
	NGV := data[31]

	mapD := map[string]string{"LastUpdate":lastUpdate,"blueGasoline95":blueGasoline95,"blueGasohol91":blueGasohol91,"blueGasohol95":blueGasohol95,"blueGasoholE20":blueGasoholE20,"blueGasoholE85":blueGasoholE85,"blueDiesel":blueDiesel,"hyForceDiesel":hyForceDiesel,"NGV":NGV}
	js,_ := json.Marshal(mapD)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(js)
}
