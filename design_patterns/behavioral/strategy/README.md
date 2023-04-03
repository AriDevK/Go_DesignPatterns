# 🏹 Strategy

---

Strategy es un patrón de diseño de comportamiento que te permite definir una familia de algoritmos, colocar cada uno de ellos en una clase separada y hacer sus objetos intercambiables.

En este caso se implemento una clase que permite intercambiar el algoritmo utilizado para encriptar una contraseña, esto se logra
debido a que todas las demas clases de encriptacion implementan la misma interfaz donde el metodo ```Hash()``` es implementado de 
manera independiente por cada una.

De modo que en nuestro codigo no llamaremos a ```HashMD5```, ```HashSHA```, ```HashXYZ``` sino que mediante el mismo metodo se manejaran 
todas las variantes.

Ademas es posible intercambiar el tipo de algoritmo de manera sencilla ya que unicamente habria que reasignar la clase del algoritmo
que se estara utilizando de la siguiente manera:

```go
import "project/algorithms"

func main(){
    pwdProtector.SetHashAlgorithm(&algorithms.XYZ{})
}
```


<br>
Diagrama UML:

```mermaid
classDiagram
class PasswordProtector
PasswordProtector : -string user
PasswordProtector : -string passwordName
PasswordProtector : -IHashAlgorithm hashAlgorithm
PasswordProtector : +PasswordProtector NewPasswordProtector(string, string, IHashAlgorithm)
PasswordProtector : +void SetPasswordName(string)
PasswordProtector : +void SetHashAlgorithm(IHashAlgorithm)
PasswordProtector : +Hash()

class IHashAlgorithm
<<interface>> IHashAlgorithm
IHashAlgorithm : void Hash(PasswordProtector)

IHashAlgorithm <|-- MD5 : Implements
IHashAlgorithm <|-- SHA : Implements
IHashAlgorithm <|-- SHA256 : Implements

MD5 : +void Hash(PasswordProtector)
SHA : +void Hash(PasswordProtector)
SHA256 : +void Hash(PasswordProtector)


```