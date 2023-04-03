# 🔒 Sync Mutex

---

En esta practica se utilizo un objecto de la struct **sync.Mutex** el cual 
nos permite generar una especia de candado para el flujo de una rutina concurrente

<br>

### Usando un Lock Simple
En esta funcion se utilizo un objeto de la struct **sync.WaitGroup** para poder manejar
una gran cantidad de procesos concurrentes; ademas de un objeto **sync.Mutex** que mediante el
metodo *.Lock()* detiene al resto de goroutines hasta que la primera en llegar ejecute el receiver *.Unlock()*.

```go
 import "sync"

 func Method(wg *sync.WaitGroup, lock *sync.Mutex) {
    defer wg.Done()
    lock.Lock() 
	// Process...
    lock.Unlock()
 }
```

<br>

### Usando un Lock de Escritura y Lectura (RW)
Esta funcion realiza la lectura de un recurso compartido en un sistema (en este caso una variable), debido
a la naturaleza del *sync.Mutex* no podemos utilizar la funcion *.Lock()* y *.Unlock()* para los casos de lectura
ya que de esta forma nuestras goroutines de lectura estaria bloqueada por las de escritura.

Debido a esto el paquete **sync** proporciona otra implementacion del la struct **Mutex** que permite manejar de manera
personalizada los procesos de lectura y escritura para que uno no interfiera con el otro, esta implementacion es la struct **sync.RWMutex**

Esta struct por definicion "hereda" de la clare Mutex original (sync.Mutex) y genera metodos especificos para el manejo de las lecturas
debido a esto los objetos de la struct *sync.RWMutex* tienen acceso al receiver *.Lock()* y *.Unlock()* para poder manejar los
procesos relacionados con las escrituras. Y los metodos *.RLock()* y *.RUnlock()* para las lecturas.

>Nota: Aun y cuando la struct de **RWMutex** herede de **Mutex** los metodos que hagan uso de esta ultima implementacion deberan de 
> ser actualizados, cambiando la implementacion base (*sync.Mutex*) por la especifica (*sync.RWMutex*)


```go
var value int = 100

func ReadMethod(lock *sync.RWMutex) int {
    lock.RLock()
    v := value
    lock.RUnlock()
    return v
}
```
   


<br>
<br>

### Definicion de la struct *sync.RWMutex*
```go
type RWMutex struct {
    w           Mutex        // held if there are pending writers
    writerSem   uint32       // semaphore for writers to wait for completing readers
    readerSem   uint32       // semaphore for readers to wait for completing writers
    readerCount atomic.Int32 // number of pending readers
    readerWait  atomic.Int32 // number of departing readers
}
```