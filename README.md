# mindmap-go

---
Test coverage:
```
ok      mindmap-go/app/services 0.421s  coverage: 50.0% of statements
ok      mindmap-go/test (cached)        coverage: [no statements]
```
Integration tests available at ```test/integration_test.go```

Service available at: http://ec2-52-58-173-131.eu-central-1.compute.amazonaws.com/

Deployed by GitHub Actions + Terraform, [Terraform project here](https://github.com/Jedich/mindmap-terraform)

To run locally on docker, execute:

```docker run --name mymysql -d -p 6603:3306 -e MYSQL_ROOT_PASSWORD=change_me -v mysql:/var/lib/mysql mysql```

```docker run --name mindmap-be -d -p 3000:3000 -e JWT_SECRET=change_me -e APP_DSN=change_me jedich/mindmap-backend```

```docker run --name mindmap-fe -d -p 80:80 -e VITE_PROXY_IP=ip:port jedich/mindmap-frontend```