# Dockerfile y `docker build`
Suele pasar que muchas veces las imagenes de Docker que encontramos en el Docker Registry no cumplen con toda la funcionalidad que necesitamos y es por ello que deberemos de crear imagenes personalizadas de Docker que si cumplan con nuestros requerimientos funcionales. Para la creación de estas imágenes personalizadas se utiliza el comando `docker build` que a partir de las instrucciones de un archivo Dockerfile creara una nueva imagen personalizada

## 01 - Dockerfile
Es un archivo de texto simple, con la peculiaridad que no posee extensión, en el cual se escriben las instrucciones que se deben realizar para crear una nueva imagen. Estas instrucciones se ejecutan bajo el formato Top-Down por lo cual se debe prestar especial atención al orden en que se escriban para que no haya problemas con la construcción.
```
FROM golang:alpine
WORKDIR /go/src/app
COPY main.go .
RUN go mod init
RUN go mod tidy
RUN go build 

CMD ["./app"]
```
Donde la linea con la instrucción `FROM` indica que como base de la nueva imagen se utilizará la imagen de golang con el tag alpine. La instrucción `WORKDIR` permite definir el directorio de entrada del nuevo contenedor, todos los demás comandos a partir de este serán ejecutados en el directorio que se especifique.

El comando `COPY` permite copiar el archivo _main.go_ ubicada en la ruta actual del host a el directorio especificado en el segundo parámetro, en caso de que el segundo parámetro sea un punto (.) querrá decir que los archivos se copiaran en el directorio actual dentro del contenedor. Las instrucciones `RUN` simplemente ejecutan los comandos que tienen a su derecha dentro del contenedor, para ello utilizan el sistema operativo y otros recursos propios del contenedor.

Por ultimo podemos ver que la instrucción `CMD` permite definir el comando que se ejecutará dentro del contenedor al momento en que éste sea creado.

<br>

## 02 - docker build
Para poder construir la image a partir del Dockerfile primero debemos de estar ubicados en la ruta o carpeta donde éste último se encuentra, una vez realizado esto se utiliza el comando `docker build` para realizar la construcción de la imagen. La forma estandar de utilizar este comando es la siguiente:
```
sudo docker build -t nombre:tag .
```
Donde con el argumento `-t` le indicamos al demonio de docker que la imagen que cree a partir del Dockerfile debe tener el nombre y tag especificado en la instrucción. A parte del nombre y tag se debe indicar la ruta donde se encuentra el Dockerfile, en este caso se usa el punto (.) ya que estamos en la misma ruta donde se encuentra dicho archivo.

<br>

## 03 - Doker Hub
En caso de no tener una cuenta de Docker Hub primero debe [crear una](https://hub.docker.com/signup), una vez creamos la sesion tenemos disponibles una cantidad ilimitada de repositorios publicos gratuitos, aunque la version estandar permite solamente uno privado, donde podremos crear y guardar en la nube nuestras imagenes.

Para acceder a Docker Hub desde nuestra consola utilizamos el siguiente comando, donde nos pedira ingresar nuestras credenciales.
```
sudo docker login
```

Una vez nos hemos logueado a Docker Hub tenemos la posibilidad de subir imagenes a nuestra cuenta, para hacerlo primero debemos de haber creado una imagen y utilizar el comando `docker tag` para taguear la imagen local con nuestro usuario y el repositorio en la nube.
```
sudo docker tag imagenOrigen usuarioDockerHub/nombreImagenDestino:tag
```

Una vez hecho esto podemos subirla al repositorio con el comando:
```
sudo docker push usuarioDockerHub/nombreImagen:tag
```