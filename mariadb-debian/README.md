
# MariaDB

Taken from [here](https://github.com/docker-library/mariadb/tree/2538af1bad7f05ac2c23dc6eb35e8cba6356fc43)


__VERSION: 10.1__ 


## Setup
- __Export some variables we will use:__
	```
	export DCKR_IMAGE="mariadb"
	export VDC_PGNAME="maria_dbstore"
	export MNT_POINT="/maria-dbdata"
	export MARIA_PORT="3306"
	export DCKR_DBCTNR="mariadb"
	```

- __Build image__
	```
	$ docker build -t $DCKR_IMAGE .
	```

- __Create Volume Data Container__
	If its necessary to keep you database data when container exists. We need to use Volume Data Container concept
	```
	$ docker create -v $MNT_POINT --name $VDC_PGNAME $DCKR_IMAGE /bin/true
	```

- __Run database container in the background__
	```
	# specific port mapping
	$ export DB=$(docker run --detach -p $MARIA_PORT:3306 -e MYSQL_ROOT_PASSWORD=root --volumes-from $VDC_PGNAME --name $DCKR_MARIA_CTNR $DCKR_IMAGE)

	# random port mapping
	$ DB=$(docker run -d -P -e MYSQL_ROOT_PASSWORD=root --name $DCKR_DBCTNR $DCKR_IMAGE )
	```

 - Install sample database
	```
	$ docker exec -ti $DB /bin/bash
	root@696c020fb305:~# mysql --host=localhost --user=root --password=root  <db_create.sql
	```


## Connect
 - Using `mysql`
	```
	$ mysql -h localhost -u root -p 3306 -d testdb
	root@696c020fb305:/# mysql --host=localhost --user=root --password=root -D testdb
	...
	MariaDB [testdb]> select * from COMPANY;
	+----+-------+-----+------------+--------+------------+
	| ID | NAME  | AGE | ADDRESS    | SALARY | JOIN_DATE  |
	+----+-------+-----+------------+--------+------------+
	|  1 | Paul  |  32 | California |  20000 | 2001-07-13 |
	|  2 | Allen |  25 | Texas      |   NULL | 2007-12-13 |
	|  3 | Teddy |  23 | Norway     |  20000 | NULL       |
	|  4 | Mark  |  25 | Rich-Mond  |  65000 | 2007-12-13 |
	|  5 | David |  27 | Texas      |  85000 | 2007-12-13 |
	+----+-------+-----+------------+--------+------------+
	5 rows in set (0.01 sec)

	MariaDB [testdb]>
	```

- Using Go sample
	```
	#TODO: $ go run db_sample.go
	```

- Using python3 sample
	```
	#TODO: $ pip install psycopg2
	$ python3 db_sample.py
	```

## Inspect

- __Check exposed ports__
	```
	$ docker inspect --format='{{range $p, $conf := .NetworkSettings.Ports}}{{(index $conf 0).HostPort}}{{end}}' $DB
	```
- __Check mounted volumes__
	```
	$ docker inspect -f '{{.Config.Volumes}}' $DB
	map[/var/lib/postgresql:{} /var/log/postgresql:{} /etc/

	$docker inspect --format '{{ .HostConfig.VolumesFrom }}' $DB
	[pg_dbstore]
	$ docker inspect --format '{{ .Mounts }}' pg_dbstore
	[{7b047b442212fb4f425a528a377d504c595e555e962f5d765ce81132fc5ce42c....
	```
- Get ip address of the containter:
	```
	$ docker inspect -f '{{ .NetworkSettings.IPAddress }}' $DB
	172.17.0.2
	```

## Manage

- Stop database container
	```
	$ docker stop  $DCKR_DBCTNR && docker rm $DCKR_DBCTNR
	```


- Cleanup database container data
	```
	???
	```
- Backup database files
	```
	???
	```

- get backup data from docker
	```
	???
	```

__TODO:__
- [ ] Check how docker compose can help with it



