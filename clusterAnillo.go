package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var localhost string
var remotehost string

//El canal es un mecanismo de sincronizacion
// Asignar el contador y leer el contador

func main() {
	//configuracion de nodo local
	bufferIn := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese el puerto local: ")
	puerto, _ := bufferIn.ReadString('\n')

	//limpiar de espacios
	puerto = strings.TrimSpace(puerto)
	localhost = fmt.Sprintf("localhost:%s", puerto)

	//configuracion del nodo remoto
	fmt.Print("Ingrese el puerto del nodo remoto")
	puerto, _ = bufferIn.ReadString('\n')

	puerto = strings.TrimSpace(puerto)
	remotehost = fmt.Sprintf("localhost:%s", puerto)


	//Establezco el modo escucha el nodo actual
	ln, _ := net.Listen("tcp", localhost)

	defer ln.Close()

	for {
		con, _ := ln.Accept()
		go manejadorConexion(con)
		//manejo de N conexiones


	}
}

func manejadorConexion(con net.Conn) {
	//es el que implementa la logica del servicio
	defer con.Close()

	//leer los datos enviados
	bufferIn := bufio.NewReader(con)
	num, _ := bufferIn.ReadString('\n')

	num = strings.TrimSpace(num)

	numero, _ := strconv.Atoi(num) 		//TODO: Reemplazar con JSON

	fmt.Println("Lllego el numero", numero)

	//Logica
	if numero == 0 {
		//finaliza el flujo 
		fmt.Println("Fin del flujo de anillo!!")
	} else {
		enviarNumero(numero - 1)
	}
}

func enviarNumero(num int) {
	con, _ := net.Dial("tcp", remotehost)
	defer con.Close()
	fmt.Fprintln(con, num)
}
