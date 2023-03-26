To run locally on docker, execute:

```docker run --name mymysql -d -p 6603:3306 -e MYSQL_ROOT_PASSWORD=change_me -v mysql:/var/lib/mysql mysql```

```docker run --name mindmap-be -d -p 3000:3000 -e JWT_SECRET=change_me -e APP_DSN=change_me jedich/mindmap-backend```

```docker run --name mindmap-fe -d -p 80:80 -e VITE_PROXY_IP=ip:port jedich/mindmap-frontend```