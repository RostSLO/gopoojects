package main

import (
	"fmt"
	"os"

	"github.com/rostslo/goprojects/playground/calcproject/calc"

	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {

	//ad := calc.Addition{1, 4}
	//calc.ExecuteOperation(&ad)
	//sub := calc.Substraction{1, 4}
	//calc.ExecuteOperation(&sub)

	http.HandleFunc("/", handler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/hijack/", func(w http.ResponseWriter, r *http.Request) {
		hj, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
			return
		}
		conn, bufrw, err := hj.Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Don't forget to close the connection:
		defer conn.Close()
		bufrw.WriteString("Now we're speaking raw TCP. Say hi: ")
		bufrw.Flush()
		s, err := bufrw.ReadString('\n')
		if err != nil {
			log.Printf("error reading string: %v", err)
			os.Exit(0)
		}
		fmt.Fprintf(bufrw, "You said: %q\nBye.\n", s)
		bufrw.Flush()
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println("Done")

}

type Page struct {
	Title  string
	Val1   float32
	Val2   float32
	Result string
}

func handler(w http.ResponseWriter, r *http.Request) {
	var p Page
	p.Title = "Simple Calculator"
	p.Result = ""
	p.Val1 = 0
	p.Val2 = 0
	t, _ := template.ParseFiles("calcpage.html")
	t.Execute(w, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	var ad float32
	p.Title = "Simple Calculator"
	value, _ := strconv.ParseFloat(r.FormValue("val1"), 32)
	p.Val1 = float32(value)
	value, _ = strconv.ParseFloat(r.FormValue("val2"), 32)
	p.Val2 = float32(value)
	op := r.FormValue("op")
	fmt.Println(p.Val1, p.Val2, op)
	if op == "+" {
		ad = calc.ExecuteOperation(&calc.Addition{p.Val1, p.Val2})
	} else if op == "-" {
		ad = calc.ExecuteOperation(&calc.Substraction{p.Val1, p.Val2})
	} else if op == "*" {
		ad = calc.ExecuteOperation(&calc.Multiplication{p.Val1, p.Val2})
	} else {
		ad = calc.ExecuteOperation(&calc.Division{p.Val1, p.Val2})
	}
	p.Result = fmt.Sprintf("%v", p.Val1) + " " + op + " " + fmt.Sprintf("%v", p.Val2) + " = " + fmt.Sprintf("%v", ad)
	t, _ := template.ParseFiles("calcpage.html")
	t.Execute(w, p)

}
