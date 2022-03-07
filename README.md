# TALLER: INTRODUCCIÓN A DOCKER
## Instalación en Windows
- Instalar previamente WSL 2 backend
- Descargar el instalador de [Docker Desktop para Windows](https://www.docker.com/products/docker-desktop)
- Doble click en *Docker Desktop Installer.exe* para iniciar la instalación
- Seguir las instrucciones del wizard hasta finalizar el proceso de instalación

<br>

[Link con la guía oficial de Docker para Windows](https://docs.docker.com/desktop/windows/install/)

<br>

## Instalación Linux (Ubuntu para este taller)
### _Versión rápida de instalación_
- Actualizar los repositorios del sistema: 
    ``` sh
    sudo apt-get update
    ```
- Instalar Docker:
    ``` sh
    sudo apt-get install docker.io 
    ```

### _Versión recomendada por Docker, Inc._
1. Configurar el repositorio
    - Actualiar los repositorios: 
        ```
        sudo apt-get update
        ```
    - Habilitar `apt` para usar un repositorio sobre HTTPS: 
        ```
        sudo apt-get install \
            ca-certificates \
            curl \
            gnupg \
            lsb-release
        ```
    - Agregar la llave GPG official de Docker
        ```
        curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
        ```
    - Configurar el repositorio estable
        ```
        echo \
        "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
        $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
        ```
2. Instalar Docker Engine
    - Actualizar los repositorios del sistema: 
        ```
        sudo apt-get update
        ```
    - Instalar la última version estable del motor de Docker
        ```
        sudo apt-get install docker-ce docker-ce-cli containerd.io
        ```