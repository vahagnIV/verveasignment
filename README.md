# Assignment

## Overview
 Since the records are independent, i.e. we do not need to write joined queries, the sharding pattern can be utilized. If the system were deployed in a cloud environment, a natural choice for the data storage would be the corresponding no-sql database( say,  DynamoDB in AWS) that does the sharding out of the box. However, presumably, the purpose of the assignment was to implement the sharding instead of using the existing solution. The project is split into modules that depend on each other. Below is the description of the modules:
 
 ### libs/datarepository
 This module contains the definition of the Model (DataRow) and an interface `Repo` of the data-repository. The interface represents the storage that can store and retrieve the data recrods. It has 3 methods - Add, Get and BatchInsert. We do not need more functionality as per the assignment. 
 ### libs/sqlrepository
 This module implements the `Repo` interface from the previous model. It depends on `gorm` for further abstraction so that we can use any sql instance as the underlying storage. In the provided example we use `sqlite`.
 ### libs/shardeddatabase
 This module implemets a sharded database, i.e. it receives a list of `Repo` instances and organizes them according to the provided implementation of the `ShardingStrategy` interface. The latter simply takes the id of the data record and returns the index of the shard. Currently, we implement only hashed sharding, however, the code is implementation agnostic. It is assumed, that the shards are independent and insert is done in parallel threads.
 ### libs/factorymethods
 Contains shared methods for initializing the instances of the interfaces.
 
 ### httpserver
 The `httpserver` runs as daemon and serves the records on port 3333 from the existing database. Currently, it assumes that at least one database exists. It has 2 endpoints: 
 - `/promotions/`- according to the assignment returns JSON-serialized DataRow. I assumed that there is a typo in the assignment and the record should be searched by its string uuid and not by an integer id. 
 - `/updateDatabase` - this is for a web hook, through which the server is notified that a new csv file was processed and is ready to be served. 
 
 ### importer
 This is a standalone application that creates databases from the csv file. Currently, we have only one implementation of the data-repository that is sqlite-based. Hence, upon receiving a new csv file, it creates a new folder `/path/to/data/folder/timestamp` and  notifies the server that the database has been updated. The server then, switches the database instance to the new one and starts serving the new items.
 
 

## Build instructions

- Install and configure go version > 1.20 from https://go.dev/dl/
- Open the file `libs/factorymethods/factorymethods.go` and modify the line 17 `const DatabasePath string = ""`. Set the path to a folder where you want the database files to be stored. **I strongly recommend to set absolute path.**  The application should have permissions to read/write to the folder.
- For building the `importer` cd to the corresponding directory and execute `go build`.
- Do the same for the `httpserver`

## Execute instructions
execute
```{bash}
./importer path/to/csv/file.csv
```

After the first database was created we can run the http server in a separate terminal

```{bash}
./httpserver
```

At this point the path `http://127.0.0.1/promotions/uuid` should be available.

Each subsequent run of 
```{bash}
./importer path/to/csv/file{i}.csv
```

will replace the existing database with the new one.

## Thoughts on the deployment

After a slight modification  (mostly in factorymethods), instead of sqlite, we can store our data in the sql instances that support network queries(mysql, mssql, postgresql). Doing this will enable real parallelization as all queries will be executed on a separate machine (instance).
Alternatively, if a sharded database (like DynamoDB) is installed, we can implement another instance `datarepository.Repo`.

Although the load on the server will be substantially decreased, the server itself will remain alone and for a larger number of queries, we will probably need a load balancer with multiple servers. 



## Thought on the code

- This is the first day I saw the golang syntax. I would be very surprised to learn that I did everything properly. In particular, I noticed that there are objects that resemble C++ pointers in golang and in my code everything is passed by copy. At the very least, there is a room for optimization here.
- I did my best to make the code cross-platform, however, the project was tested on Ubuntu 20.04 only.
- At some point I thought to dockerize the solution, however,... :-)
