# Postgres DB

Taken from [here](https://docs.docker.com/engine/examples/postgresql_service/)


__WORKS FOR PostgresDB VERSION: 9.6__ 


## Setup
- __Export some variables we will use:__
	```
	export DCKR_IMAGE="eg_postgresql"
	export VDC_PGNAME="pg_dbstore"
	export MNT_POINT="/dbdata"
	export PG_PORT="32770"
	export DCKR_PG_CTNR="postgres"
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
	$ export DB=$(docker run --detach -p $PG_PORT:5432 --volumes-from $VDC_PGNAME --name $DCKR_PG_CTNR $DCKR_IMAGE)

	# random port mapping
	$ DB=$(docker run -d -P --name --name $DCKR_PG_CTNR $DCKR_IMAGE )
	```

 - Install sample database
	```
	$ export PGPASSWORD="docker" && psql -h localhost -p $PG_PORT -d docker -U docker <db_create.sql
	```


## Connect
 - Using `psql`
	```
	$ PGPASSWORD="docker" && psql -h localhost -p $PG_PORT -d docker -U docker 
	
	# test select on imported data
	export PGPASSWORD="docker"; echo -n  "select * from company" |psql -h localhost -p $PG_PORT -d docker -U docker
	 id | name  | age |                      address                       | salary | join_date
	----+-------+-----+----------------------------------------------------+--------+------------
	  1 | Paul  |  32 | California                                         |  20000 | 2001-07-13
	  2 | Allen |  25 | Texas                                              |        | 2007-12-13
	  3 | Teddy |  23 | Norway                                             |  20000 |
	  4 | Mark  |  25 | Rich-Mond                                          |  65000 | 2007-12-13
	  5 | David |  27 | Texas                                              |  85000 | 2007-12-13
	(5 rows)
	```

- Using Go sample
	```
	$ go run db_sample.go
	```

- Using python3 sample
	```
	$ pip install psycopg2
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
	$ docker stop  $DCKR_PG_CTNR && docker rm $DCKR_PG_CTNR
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



