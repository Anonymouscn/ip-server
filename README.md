# IP Server

## Usage

1. clone this project
    ```shell
    git clone https://github.com/Anonymouscn/ip-server.git
    ```
   
2. check the config in `application.yml`

    ```shell
    server:
    name: $service_name
    port: $service_port
    ```

3. build project with docker

    ```shell
    ./build.sh # on unix/linux
    ```

    ```shell
    ./build.bat # on windows
    ```
4. save docker image in `bin` folder


### Pull from docker hub

1. pull images

   ```shell
   docker pull pgl888999/ip-server
   ```

2. run a container

   ```shell
   docker run --name $container_name -p $prot:8080/tcp -p $port:8080/udp -itd pgl888999/ip-server
   ```

### Tips

* check logs on docker container
   ```shell
   docker logs $container_name
   ```
