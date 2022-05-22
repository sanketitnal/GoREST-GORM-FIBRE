# GoREST-GORM-FIBRE
REST API using Fibre and GORM

#### Database Configuration
- If using Postgres DB, edit database credentials at [here](https://github.com/sanketitnal/GoREST-GORM-FIBRE/blob/master/database/postgres.go).
- If want to use another database, you may refer [here](https://github.com/sanketitnal/GoREST-GORM-FIBRE/blob/master/database/postgres.go) to create similar database setup file. Add database connection call at [here](https://github.com/sanketitnal/GoREST-GORM-FIBRE/blob/master/database/database.go). On using another database, will have to rewrite references to database (currently database.PostgresDB) in services files. [example at](https://github.com/sanketitnal/GoREST-GORM-FIBRE/blob/master/services/user/userservice.go).

#### Run Postgres container(Example)
- ```docker container run --detach --publish 5432:5432 --name postgres-for-gorest -e POSTGRES_PASSWORD=password postgres:latest```.
- Check default environment variables at [postgres docker hub](https://hub.docker.com/_/postgres).

####  Configuring and running application
- Update port to [app.Listen(":port")](https://github.com/sanketitnal/GoREST-GORM-FIBRE/blob/master/main.go).
- From application directory run ```go mod tidy```.
- From application directory run ```go build```.
- Run the generated executable.
- Instead we can directly run ```air``` from the application directory for live reloading. More about golang live reloading at [air](https://github.com/cosmtrek/air).

#### User Model, Json data format
- uid is unsigned int (uint64).
- first_name is string.
- last_name is string.
- contact_info is unsigned (uint64).
```
{
        "uid": 1,
        "first_name": "FirstName",
        "last_name": "LastName",
        "contact_info": 0000055555
}
```

#### REST endpoints
1. ##### POST "/user/create"
- Creates new user.
```
curl --location --request POST 'http://127.0.0.1/user/create' \
--form 'uid="1"' \
--form 'first_name="Firstname"' \
--form 'last_name="Lastname"' \
--form 'contact_info="1010110101"'
```
- uid must be an integer greater than 0. If uid is invalid, database will create its own uid.

2. ##### GET "/user/getUserById/:uid"
```
curl --location --request GET 'http://localhost/user/getUserById/1'
```

3. ##### GET "/user/getAll"
```
curl --location --request GET 'http://localhost/user/getAll'
```

4. ##### DELETE "/user/delete/:id"
```
curl --location --request DELETE 'http://localhost/user/delete/7'
```

5. ##### PUT "/user/update"
- uid is mandatory.
- Will only update columns provided.
```
curl --location --request PUT 'http://localhost/user/update' \
--form 'uid="3"' \
--form 'first_name="newName"' \
--form 'last_name="newName"' \
--form 'contact_info="7777788888"'
```

##### References
- https://youtu.be/dpx6hpr-wE8
- https://youtu.be/5SeYS2aRF34
