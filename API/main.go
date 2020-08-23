package main

import (
	"fmt" //imprimir en conosla
	"net/http" //levantar el server
	"encoding/json" //formato json
	"github.com/gorilla/mux" //para levantar el router
	"io/ioutil" //Entradas por form
	"strconv" //String to int
	"os"
	"syscall"
)

//archivos creados por los modulos 
// CPU -> /proc/mem_grupo18/mem_grupo18.json
// RAM -> /proc/cpu_grupo18/cpu_grupo18.json
// Estos los deje solo para probar
var ArchivoCPU = "./CPU.json"
var ArchivoRAM = "./ram.json"

type Response struct {
    StatusCode int
    Msg string
}

//Struct para modulo de CPU
type Proceso struct {	
	PID int `json:"PID"`
	Nombre string `json:"Nombre"`
	Estado int `json:"Estado"`
}
var Procesos [] Proceso

//Struct para modulo de RAM
type ramLectura struct {	
	Memoria int `json:"Memoria"`
	Libre int `json:"Libre"`
}
var Ram ramLectura
//Struct para modulo de RAM
type StructRam struct {	
	TotalServer float64
	TotalConsumida float64
	Porcentaje float64
}
var RamAcumalada [] StructRam

// retorna JSON acumulando en cada llamada la anterior
// localhost:3000/RAM
// Metodo GET
func getRAM(w http.ResponseWriter, r *http.Request){
	data, err := ioutil.ReadFile(ArchivoRAM)
    if err != nil {
      fmt.Println(err)
	}
	err = json.Unmarshal(data, &Ram)
	if err != nil {
        fmt.Println("error:", err)
	}
	structRam := StructRam{
		float64(Ram.Memoria) / 1024, 
		float64(Ram.Libre) / 1024,
		float64(Ram.Memoria * 100) / float64(Ram.Libre),
	}

	RamAcumalada = append(RamAcumalada, structRam)
	
	js, err := json.Marshal(RamAcumalada)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(js)
}

// retona JSON con Procesos 
// localhost:3000/CPU 
// Metodo GET
func getCPU(w http.ResponseWriter, r *http.Request){
	data, err := ioutil.ReadFile(ArchivoCPU)
    if err != nil {
      fmt.Println(err)
	}
	err = json.Unmarshal(data, &Procesos)
	if err != nil {
        fmt.Println("error:", err)
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(Procesos)		
}

// kill Proceso envio IDproceso
// localhost:3000/kill/ID 
// Metodo GET 
func kill(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	pid, err := strconv.Atoi(vars["id"]) //Enviar un String para convertir
	//Si existe un error 
	if err != nil{
		fmt.Fprintf(w,"ID invalido en kill")
		return
	}else{
		fmt.Printf("Se procedera a matar al proceso: %d \n", pid);
		process, err := os.FindProcess(pid);
		if err != nil {
			fmt.Println("error:", err)
		}
		err = process.Signal(syscall.Signal(0)) // if nil then is ok to kill
		if err != nil {
			fmt.Println("error:", err)
		}
		err = process.Kill()
		if err != nil {
			fmt.Println("error:", err)
		}
	}
}

// Ruta Raiz
func indexRoute(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	resp := Response{http.StatusOK, "SERVER OK"}	
	json.NewEncoder(w).Encode(resp)		
}

func main()  {
	fmt.Println("Iniciando Server")
	
	router := mux.NewRouter()
	router.HandleFunc("/",indexRoute)
	router.HandleFunc("/cpu",getCPU).Methods("GET")
	router.HandleFunc("/ram",getRAM).Methods("GET")
	router.HandleFunc("/kill/{id}",kill).Methods("GET")
	http.ListenAndServe(":3000", router)

}