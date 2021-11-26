package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var remotehost string


func main() {
	bufferIn := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese el puerto remoto: ")
	puerto,_ := bufferIn.ReadString('\n')
	puerto = strings.TrimSpace(puerto)

	remotehost = fmt.Sprintf("localhost:%s",puerto)

	fmt.Print("Ingrese el valor del numero: ")
	num, _ := bufferIn.ReadString('\n')

	num = strings.TrimSpace(num)
	numero, _ := strconv.Atoi(num)


	enviarNumero(numero)
}

func enviarNumero(numero int) {
	con, _ := net.Dial("tcp",remotehost)
	defer con.Close()

	fmt.Fprintln(con,numero)
	
}