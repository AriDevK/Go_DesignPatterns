# 🏭 Factory Method

---

Este es un patrón de diseño creacional que proporciona una interfaz para crear objetos en una superclase, mientras permite a las subclases alterar el tipo de objetos que se crearán.

En esta practica realizamos la implementacion de este patron mediante la interfaz ```IProduct``` la cual contiene
todos los metodos necesarios para estandarizar los productos que se manejan durante la aplicacion.

De ahi se genero la struct ```Computer``` la cual implementa la interfaza ```IProduct``` y funciona como clase/struct padre
del resto de Structs (Productos) que se manejaran en la aplicacion.

Ademas de esto, se generaron los **Metodos de Fabrica** los cuales nos permiten crear objetos de distintas clases pero
permitiendonos crearlos y trabajar con ellos de una manera mas generica debido a que heredan de ```Computer```.


Diagrama de Clases para llevar a cabo este patron:
```mermaid
classDiagram
class IProduct
<<interface>> IProduct

IProduct <|-- Computer : Implements
IProduct : int GetStock()
IProduct : string GetName()
IProduct : void SetStock(int) 
IProduct : void SetName(string)

Computer : +int GetStock()
Computer : +string GetName()
Computer : +void SetStock(int) 
Computer : +void SetName(string)


Computer <|-- Laptop : Extends
Laptop : +Computer
Laptop : +IProduct NewLaptop()

Computer <|-- Desktop : Extends
Desktop : +Computer
Desktop : +IProduct NewLaptop()
```

El uso de este patron de diseño nos permite estandarizar el manejo de los objetos de nuestra aplicacion aun y cuando estos
implementen la logica de distinta manera, podremos utilizar metodos genericos que podran manejarlos sin problema.

En caso de requerir agregar un nuevo objeto unicamente tendra que implementar la interfaz de ```IProduct``` y todos los metodos
de nuestra logica de negocio seran capaz de trabajar con estos nuevos objetos.

p.Ej.
En este diagrama aun y cuando tenemos varios tipos de productos completamente diferentes, todos implementan la misma interfaz por lo que 
su logica de negocio pueda ser diferente; en nuestra aplicacion todas se podran manejar bajo el mismo metodo ya que cada una
de las estructuras se encarga de implementar su logica y nosotros no nos preocuparemos si el stock se mide por caja, botella, paquete, etc.Ñ
```mermaid
classDiagram
class IProduct
<<interface>> IProduct
IProduct <|-- Computer : Implements
IProduct <|-- Drinks : Implements
IProduct <|-- CellPhone : Implements

IProduct : int GetStock()
IProduct : string GetName()
IProduct : void SetStock(int) 
IProduct : void SetName(string)


CellPhone : +int GetStock()
CellPhone : +string GetName()
CellPhone : +void SetStock(int) 
CellPhone : +void SetName(string)

CellPhone <|-- Android : Extends
Android : +Computer
Android : +IProduct NewLaptop()

CellPhone <|-- iOS : Extends
iOS : +Computer
iOS : +IProduct NewLaptop()



Drinks : +int GetStock()
Drinks : +string GetName()
Drinks : +void SetStock(int) 
Drinks : +void SetName(string)

Drinks <|-- Soda : Extends
Soda : +Computer
Soda : +IProduct NewLaptop()

Drinks <|-- Water : Extends
Water : +Computer
Water : +IProduct NewLaptop()



Computer : +int GetStock()
Computer : +string GetName()
Computer : +void SetStock(int) 
Computer : +void SetName(string)

Computer <|-- Laptop : Extends
Laptop : +Computer
Laptop : +IProduct NewLaptop()

Computer <|-- Desktop : Extends
Desktop : +Computer
Desktop : +IProduct NewLaptop()


```