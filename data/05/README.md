
```bash
$ docker run -d --name postgres -p 5432:5432 postgres:10
$ apt install postgresql-client-10
$ psql -U postgres -h localhost
```

```sql
create database gwp;
create user gwp with password 'gwp';
```