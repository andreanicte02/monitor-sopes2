package main

import (
	"fmt" //imprimir en conosla
	"net/http" //levantar el server
	"encoding/json" //formato json
	"flag"
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

//Deben de tener permisos 444
//sudo chmod 444 mem_gupo18
var ArchivoCPU = "/proc/cpu_grupo18"
var ArchivoRAM = "/proc/mem_grupo18"
var ArchivoCON = "/proc/pro_grupo18"

type Response struct {
    StatusCode int
    Msg string
}

type conteo struct {	
	Ejecucion int `json:"Ejecucion"`
	Suspendido int `json:"Suspendido"`
	Detenido int `json:"Detenido"`
	Zombie int `json:"Zombie"`
}
var C conteo

//Struct para modulo de CPU
type Procesos []struct {
	PID    int    `json:"PID"`
	Nombre string `json:"Nombre"`
	Estado int    `json:"Estado"`
	UID    int    `json:"uid"`
	Mm     int    `json:"mm"`
	Sub    []struct {
		PID    int    `json:"PID"`
		Nombre string `json:"Nombre"`
		Estado int    `json:"Estado"`
		UID    int    `json:"uid"`
		Mm     int    `json:"mm"`
	} `json:"sub"`
}
var P Procesos

//Struct para modulo de RAM
type ramLectura struct {	
	Total int `json:"total"`
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
      fmt.Println("error1:",err)
	}
	//fmt.Println((data))
	
	err = json.Unmarshal(data, &Ram)
	if err != nil {
        fmt.Println("error2:", err)
	}
	
	structRam := StructRam{
		float64(Ram.Total)/1000, 
		float64(Ram.Total-Ram.Libre)/1000,
		float64((Ram.Total-Ram.Libre) * 100) / float64(Ram.Total),
	}

	//RamAcumalada = append(RamAcumalada, structRam)
	//js, err := json.Marshal(RamAcumalada)
	js, err := json.Marshal(structRam)
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
	err = json.Unmarshal(data, &P)
	if err != nil {
        fmt.Println("error:", err)
	}
	
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(P)		
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

func getConteo(w http.ResponseWriter, r *http.Request){
	data, err := ioutil.ReadFile(ArchivoCON)	
    if err != nil {
      fmt.Println(err)
	}
	err = json.Unmarshal(data, &C)
	if err != nil {
        fmt.Println("error:", err)
	}
	
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(C)
}

// Ruta Raiz
func indexRoute(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	resp := Response{http.StatusOK, "SERVER OK"}	
	json.NewEncoder(w).Encode(resp)		
}

func main()  {
	fmt.Println(">>>> Iniciando Server")

	var dir string
	flag.StringVar(&dir, "dir", "./so-pra-03-client/build/", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	
	router := mux.NewRouter()

	// This will serve files under http://localhost:8000/static/<filename>
	router.PathPrefix("/client/").Handler(http.StripPrefix("/client/", http.FileServer(http.Dir(dir))))

	router.HandleFunc("/",indexRoute)
	router.HandleFunc("/cpu",getCPU).Methods("GET")
	router.HandleFunc("/ram",getRAM).Methods("GET")
	router.HandleFunc("/kill/{id}",kill).Methods("GET")
	router.HandleFunc("/conteo",getConteo).Methods("GET")
	fmt.Println(">>>> Iniciado en puerto 8080 ")
	http.ListenAndServe(":8080", router)

}

