Compilar y montar modulos de kernel

# Cpu (procesos)

```
>cd Modulos/cpu
>sudo make
```

Montar modulo

```
>sudo insmod cpu_grupo18.ko
```

Mostra mensaje

```
>sudo dmesg
```

Limpiar la carpeta donde se compilo el modulo

```
>sudo make clean
```

Desmontar el modulo

```
>sudo rmmod cpu_grupo18.ko
```



# Ram 

```
>cd Modulos/ram
>sudo make
```

Montar modulo

```
>sudo insmod mem_grupo18.ko
```

Mostra mensaje

```
>sudo dmesg
```

Limpiar la carpeta donde se compilo el modulo

```
>sudo make clean
```

Desmontar el modulo

```
>sudo rmmod mem_grupo18.ko
```

# API
Se necesita tener instalado go
cd /API
go get -u github.com/gorilla/mux
go run main.go

# Funcionamiento API

* Ir a http://localhost:3000/
```json
{
	"StatusCode":200,
	"Msg":"SERVER OK"
}
```


* Formato para Procesos CPU http://localhost:3000/cpu
```json
[
	{
		"PID":int,
		"Nombre":"string",
		"Estado":int,
		"uid":int,
		"mm":int,
		"sub":[
		{
			"PID":int,
			"Nombre":"string",
			"Estado":int,
			"uid":int,
			"mm":int
		},
		{
			"PID":int,
			"Nombre":"string",
			"Estado":int,
			"uid":int,
			"mm":int
		}
		]
	}
]

```