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

