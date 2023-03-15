To run locally on docker, run:

```docker run --name mymysql -d -p 6603:3306 -e MYSQL_ROOT_PASSWORD=change_me -v mysql:/var/lib/mysql mysql```

```docker run --name mindmap-be -d -p 3000:3000 -e JWT_SECRET=change_me jedich/mindmap-backend```

```docker run --name mindmap-fe -d -p 8080:5173 jedich/mindmap-frontend```