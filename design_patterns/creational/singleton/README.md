# 🎼 Singleton

---

Singleton es un patrón de diseño creacional que nos permite asegurarnos de que una clase tenga una única instancia, a la vez que proporciona un punto de acceso global a dicha instancia.

En esta practica realizamos este patron mediante la simulacion de una conexion a una Base de datos,
este patron de diseño dicta que la clase que implemente este patron no debera de crear objetos directamente desde el constructor
sino que debera de tener un punto global (un metodo publico) el cual sea capaz de revisar si actualmente ya existe una instancia
en caso de ser asi nos retornara la instancia, caso contrario la crea.

<br>
En este ejemplo se logra apreciar el metodo encargado de instanciar el objeto con la conexion a la base de datos; se aprecia el como
se genera una variable global privada en el archivo la cual hace referencia al objeto global de la clase, es decir el que se estara regresando
en cada llamada.

Ademas podemos ver como al ejecutar la funcion que genera la instancia esta revisa si el objeto tiene un valor o no y en caso de tenerlo
nos retornada la misma instancia global, caso contrario la instanciara normalmente.

En este caso se hizo uso de un lock de la struct ```sync.Mutex``` para poder manejar el flujo de instanciasion en entornos donde
varios clientes quieran acceder al mismo recurso al mismo tiempo.
```go
import "sync"

var db *Database
var lock sync.Mutex

type Database struct {
}

func GetDataBaseInstance() *Database {
    defer lock.Unlock()
    lock.Lock()
    if db == nil {
        fmt.Println("Creating DB Connection")
        db = &Database{}
        db.CreateSingleConnection()
    } else {
        fmt.Println("DB Already Connected")
    }
    return db
}
```