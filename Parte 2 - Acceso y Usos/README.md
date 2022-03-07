# Acceso y Usos
Ya que se tiene conocimiento sobre el funcionamiento y la creación de contenedores a partir de imágenes de Docker Hub, el siguiente paso es poder acceder a los contenedores activos, conceder acceso publico mediante sus puertos y también poder trabajar dentro de ellos.

<br>

## 01 - Acceso Remoto con Apaches
Una vez logeado en el sistema que tiene el Docker Engine instalado crearemos un contenedor en este caso con el sistema operativo de Ubuntu version 18.04 para poder trabajar cómodamente dentro de el.
```
sudo docker run -it -d -p 5000:80 --name apacheserver ubuntu:18.04 /bin/bash
```
Como se puede observar se utilizará como imagen base la del repositorio de Ubuntu en su versión 18.04 para crear un contenedor que correrá en segundo plano gracias al parámetro `-d`, también se puede apreiciar que se expondrá el puerto 80 del contenedor con el 5000 del host, esto debido a que apache funciona nativamente en el puerto 80 y será necesario que se enlace dicho punto de acceso a otro disponible en la máquina anfitriona.

<br>

Luego se accede dentro del contenedor con el comando
```
sudo exec -it apacheserver /bin/bash
```
Una vez dentro del contenedor actualizaremos los repositorios de la imagen de ubuntu y posteriormente se instala Apache ingresando los siguientes comandos en la terminal de Ubuntu
```
$ apt-get update
$ apt-get install apache2
```
Solamente queda iniciar el servidor de apache, y posteriormente acceder a la página por defecto que trae el servidor al ser instalado por primera vez
```
$ /etc/init.d/apache2 start
```
Para poder acceder a dicha página se escribe la dirección IP del host (localhost en caso de que no se esté trabajando en la nube) y acceder al puerto expuesto al contenedor, en nuestro caso sería el puerto `5000`.

<br>

## 02 - Trabajando dentro del contenedor
Como desarrolladores muchas veces no se tiene la disponibilidad de que exista un paquete en los repositorios del sistema operativo que cumpla con las necesidades que tenemos, y en consecuencia debemos de crear nosotros mismos el programa en cuestión. 

Para realizar este trabajo primero debemos de acceder al contenedor y dentro de él escribir a mano el código en un archivo de texto plano mediante la consola integrada, lo cual puede ser muy dificil si en la imagen del contenedor no se encuentran configurados editores de texto o no es posible utilizar los comandos `apt`, o equivalentes de otras distribuciones, para descargarlos de los repositorios oficiales.

Como solución a estos inconvenientes actualmente existen varios clientes que permiten acceder a el sistema de archivos del contenedor y mostrarlo de forma gráfica para que el usuario final lo pueda manejar de una mejor manera, uno de estos clientes es una extensión del conocido editor de código [Visual Studio Code](https://code.visualstudio.com) llamado [Remote-Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) que permite acceder a cualquier contenedor que este corriendo en el mismo host donde esté instalado y ejecutandose VS Code.

Se recomienda consultar la [documentación](https://github.com/microsoft/vscode-remote-release.git) completa de su uso para un mejor funcionamiento.