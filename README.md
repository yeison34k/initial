air same to nodemon 
https://techinscribed.com/5-ways-to-live-reloading-go-applications/
go get -u github.com/cosmtrek/air

docker build -t nombreProyecto:V1 .
◦ {V1 puede ser cualquier tag, :v2. :c876, etc}
• docker images 
◦ {lista las imagenes}
• docker image rm id
◦ {elimina una imagen por su id  agregando: “-f” se hace --force}
• docker ps -a 
◦ {lista contenedores en ejecucion}
docker run nombre-imagen
◦ {ejecuta una imagen}
• docker image ls
◦ {lista las imagenes}
• docker stop [container-name]
◦ {detiene un contenedor}
• docker stop $(docker ps -a -q) 
◦ {detiene todos los contenedores}
• docker logs [container-name]
◦ {logs}
• docker run -it image-name sh
◦ {ver contenido de la imagen}
• docker run --name nombreContainer -d -p 8081:8081 NombreImagen:tag
◦ {con esto ejecuto una imagen como container,  puertoLocal:puertoImagen}