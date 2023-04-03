# 💾 Cache

---

En esta práctica se implementó un sistema de caché parecido a Redis, el cual para este caso nos
permitiera calcular las operaciones de la una implementación recursiva de la serie de fibonacci, guardarla
y leerla cuando vuelva a ser necesario conocer su valor.

Se utilizó conceptos de:
- WaitGroups
- Locks (Mutex)


 La aplcación implementa una struct para manejar la logica de la Memoria Cache la cual cuenta con atributo
 del tipo *Function* el cual es una Función que recive un entero y regresa una interface (Generic) asi como su error

Maneja una propieda llamada cache basa en un map de Int-FunctionResult, siendo la key (int) el valor previamente
calculado y el value (FunctionResult) una struct que nos permite obtener un interface (Generic) y un error en caso de
se presente alguno.

 ```go
import "sync"

type MemoryCache struct {
	f     Function
	cache map[int]FunctionResult
	lock  sync.Mutex
}

type Function func(key int) (interface{}, error)
type FunctionResult struct {
value interface{}
err   error
}
```

Esta implementacion cuenta además con un metodo constructor para inicializar el map y recibir la funcion que estara
manejando los datos, asi como también una funcion para obtener los datos de nuestro map segundo su key.

Esta última funcion hace uso del objeto Mutex para detener y continuar el flujo del proceso mediante los locks, notese
que se bloqueó el flujo de manera explícita en ciertos puntos del proceso, esto para poder dar paso a realizar ciertas acciones
al resto de goroutines cuando uno de estos termine el primero proceso.