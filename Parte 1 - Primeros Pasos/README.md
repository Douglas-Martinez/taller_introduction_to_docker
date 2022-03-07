# Primeros Pasos
Tras haber instalado el Docker Engine en nuestro sistema podemos utilizar en todo momento los comandos `sudo docker ps` y `sudo docker images` para ver respectivamente un listado con la informacion de los contenedores activos y las imagenes que actualmente se encuentran en nuestro sistema.

<br>

## 01 - Hello World 
Para verificar que Docker esté correctamente instalado ejecutamos la imagen de `hello-world` oficial que Docker, Inc. tiene disponible para esta verificación.

Si la imagen, en este caso `hello-world`, no se encuentra en el entorno local Docker automáticamente la descarga de Docker Regitry y posteriormente levanta el contenedor.
```
sudo docker run hello-world
```

<br>

## 02 - Comando _pull_
Se puede descargar una imagen desde Docker Hub utilizando el comando `sudo docker pull nombreImagen:tag`, por ejemplo:
```
sudo docker pull python:3.9.10
```
Donde python es el nombre del repositorio de la imagen y en este caso el tag hace referencia a la versión 3.9.10 de este lenguaje de programación. Para mas información se puede acceder a [Docker Hub](https://hub.docker.com), buscar la imagen deseada y revisar todos los tags disponibles para dicha imagen.

<br>

## 03 - Comando _run_ y la creación de Contenedores
La sintáxis básica del comando docker run es `sudo docker run [OPTIONS] IMAGE [COMMAND] [ARG...]`

Como es normal el comando `sudo docker run` admite muchas opciones y banderas o argumentos los cuales pueden modificar el comportamiento y configuración inicial del contenedor. Normalmente los parámetros y argumentos mas utilizados son:
- `nombre:tag`: es un parámetro que indica que imagen y que tag será utilizado para la creación del contenedor.
- `comando`: indica que instrucción debe realizarse al momento que el contenedor sea creado.
- `-i`: indica que se requiere que la consola de entrada se mantenga corriendo incluso si no está adjunta (interactiva).
- `-t`: asigna una consola de tipo TTY
- `-d`: ejecuta el contenedor en segundo plano y regresa a la terminal imprimiendo el ID del contenedor.
- `-p`: enlaza los puertos indicados del contenedor a los del host.
- `-v`: monta y enlaza un volumen.
- `-e`: aloja dentro del contenedor las variables de entorno que se especifiquen como valor
- `--name`: asigna un nombre para el contenedor, este valor debe ser de tipo string.

Para el ejemplo:
```
sudo docker run -it -d --name mipython -p 3010:3000 python:alpine sh
```
Tenemos que con la bandera `-it` le indicamos a docker que queremos que se habilite una consola interactiva, con la bandera `-d` que el contenedor se ejecute en segundo plano, en el caso del flag `--name mipython` le asignamos el nombre de "mipython" al contenedor que se va a crear a continuación. Con `-p 3010:300` indicamos que necesitamos que el puerto 3010 de nuestro host este enlazado con el puerto 3000 del contenedor; en este argumento el numero del lado izquierdo del simbolo de dos puntos indica siempre el puerto del host, mientras que el del lado derecho hace referencia al puerto del contenedor que se va a enlazar. Luego tenemos que la imagen que se va a utilizar es la de `python` con su tag `alpine` y por ultimo le decimos al motor de docker que una vez creado el contenedor ejecute el comando `sh`.

<br>

Con ese comando, tras ser creado el contenedor nos mostrara el identificador único de éste y será con ese id o con el nombre que posteriormente podremos acceder a él.

<br>

## 04 - Comando _exec_ para acceder al contenedor
Para ingresar a un contenedor que esté ejecutandose previamente se utiliza el comando `docker exec [ARGS] nombre/idContainer [COMMAND]` donde, como identificador del contenedor se puede utilizar ya sea el nombre o el id unico del mismo:
```
sudo docker exec -it mipython sh
```
<br>

## 05 - Manejo del estado del contenedor
- Para detener la ejecución de un contenedor en específico se utiliza el comando ```sudo docker stop nombre/idContainer```.
- De la misma forma si se desea iniciar un contenedor que se encuentra en estado de pausa o detenido se utiliza el comando ```sudo docker start nombre/idContainer```.
- Finalmente si se desea remover o eliminar un contenedor pausado se utiliza el comando ```sudo docker rm nombre/idContainer``` o si se desea remover un contenedor activo se le agrega la bandera `-f` al comando anterior para forzar la detención y posteriormente la eliminación.

Cabe resaltar que en los tres casos `nombre/idContainer` hace referencia al nombre con el que fue creado el contenedor o el identificador que Docker Engine le asignó cuando fue creado, por ejemplo:
```
sudo docker stop mipython
```