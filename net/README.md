# 🚢 Net Port Scanner

---

En esta practica realizamos un scanner de puertos en servidores de manera concurrente; en esta practica se hizo uso de 
los siguientes temas:
- Goroutines
- WaitGroups
- Flags
- Net Connections

<br>

### Flags
En cuanto a los flags, estos son una forma de mandar parametros por consola a nuestras aplicaciones en Go, para esto la libreria
estandar nos proporciona el paquete "flag", para poder utilizar esta libreria debemos de declarar los flags como
variables globales y parsearlos antes de utilizarlos, por ejemplo:
```go
import "flag"

var myFlag = flag.String("NAME", "DEFAULT_VALUE", "DESCRIPTION")

func main() {
    flag.Parse()
    // Process
}
```

<br>

### Net Connections
En cuanto a las conexiones para escaneo de puertos la misma libreria estandar nos proporciona el paquete "net" el cual contiene
funciones varias para el acceso a distintos metodos relacionados con las redes informaticas en distintos protocolos.

En esta practica utilizamos la funcion ```net.Dial()``` el cual nos permite conectarnos a una direccion mediante una lista de protocolos
de la siguiente manera:

```go
import "net"

func main(){
    conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
    if err != nil {
        log.Default("Error on connection")
    }
    
    _ = conn.Close()
}
```

Este procedimiento nos permite revisar si hubo una conexion exitosa con el servidor.