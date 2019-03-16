package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type panPerson struct {
	ID int `json:id`
	PAN string `json:pan`
	Name string `json:name`
	Fname string `json:fname`
	Dob string `json:dob`
}

type dlPerson struct {
	ID int `json:id`
	DL string `json:dl`
	Name string `json:name`
	Fname string `json:fname`
	IssuedDate string `json:issued_date`
	Address string `json:address`
}

var panPeople []panPerson
var dlPeople []dlPerson

func main() {

	panPeople = append(panPeople,
		panPerson{ID: 1, PAN: "KLDAJ8932" , Name:"johnDi", Fname:"fname-1", Dob:"20/01/1989"},
		panPerson{ID: 2, PAN: "KLDAJ8922" , Name:"mabeee", Fname:"fname-2", Dob:"22/02/1939"},
		panPerson{ID: 3, PAN: "KLDAJ8962" , Name:"louise", Fname:"fname-3", Dob:"12/06/1987"},
		panPerson{ID: 4, PAN: "KLDAJ8912" , Name:"mavarin",Fname:"fname-4", Dob:"21/02/1939"},
		panPerson{ID: 5, PAN: "KLDAJ8902" , Name:"jonesa", Fname:"fname-5", Dob:"09/06/1987"},
		panPerson{ID: 6, PAN: "BPHPT82815", Name:"	KUMMARIKUNTLA TEJA", Fname:"KUMMARIKUNTLA SRINIVASA RAO", Dob:"05/11/1999"})

	dlPeople = append(dlPeople,
		dlPerson{ID: 1, DL:"AP40300323542018", Name:"MANI SAI PRASAD M", Fname:"M ANJI PRASAD", IssuedDate:"09-11-2018", Address:"11-143,EBD COLONY E.B.C.COLONY,CHITTOOR, PULICHERLA CHITTOOR-517172"},
		dlPerson{ID: 2, DL:"AP40300323542028", Name:"RAHUL RAM", Fname:"RAM KUMAR", IssuedDate:"03-12-2008", Address:"11-143,EBD ADS E.ADSF.C.COLONY,DS, PULICHERLA ASD-517172"},
		dlPerson{ID: 1, DL:"AP40300323542048", Name:"KUMAAR SAURAVA", Fname:"DINESH", IssuedDate:"09-09-2018", Address:"11-143,EBD B.C.ADS,CHITTOOR, PULICHERLA CHITTOOR-517172"},
		dlPerson{ID: 1, DL:"AP40300323542068", Name:"ANJALI HANMY", Fname:"RAM PRASAD", IssuedDate:"10-11-2018", Address:"E.B.C.COLONY,CHITTOOR, PULICHERLA CHITTOOR-517172"},
		dlPerson{ID: 1, DL:"AP40300323542078", Name:"RAM KIRAN", Fname:"SHIRAN", IssuedDate:"09-11-2001", Address:"11-143,EBD COLOCHITTOOR, PULICHERLA CHITTOOR-517172"})

	router := mux.NewRouter()

	router.HandleFunc("/pan/People", getPanPeople).Methods("GET")
	router.HandleFunc("/pan/People/{pan}", getPanPerson).Methods("GET")

	router.HandleFunc("/dl/People",getDlPeople	).Methods("GET")
	router.HandleFunc("/dl/People/{dl}",getDlPerson).Methods("GET")

	log.Fatal(http.ListenAndServe(GetPort(), router))
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == ""  {
		port = "4747"
		fmt.Println("INFO: no port env var detected, defatulign to" )
	}
	return ":" + port
}

func getPanPeople(w http.ResponseWriter, r *http.Request)  {
	json.NewEncoder(w).Encode(panPeople)
}

func getPanPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, _ := params["pan"]

	for _, person := range panPeople {
		if person.PAN == i {
			json.NewEncoder(w).Encode(&person)
		}
	}
}

func getDlPeople(w http.ResponseWriter, r *http.Request)  {
	json.NewEncoder(w).Encode(dlPeople)
}

func getDlPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, _ := params["dl"]

	for _, person := range dlPeople {
		if person.DL == i {
			json.NewEncoder(w).Encode(&person)
		}
	}
}