# GoRedis
Go application using redis
HOw to install redis in windows?
Install wsl and then redis
-------------------

1)wsl --install Ubuntu --web-download

2)wsl -l -v

3)curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg

4)echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list

5)sudo apt-get update

6)sudo apt-get install redis

7)sudo service redis-server start

----------------------------
Redis- Inmemory volatile database ley-value pair database. It is not persistent meanining once the system closes, the data is lost. It is usually used as a cache and can be used on top of other database since complex query can take time, you can store necessary data in redis nd database and when retrieving you can check first in redis and if present, you can fetch from there. Since redis sits in your ram, the data retrieval is very fast.
Everything in redis is stored as string

few commands:
1)redis-cli  -- starts redis cli
2)ping -- return with pong as response
3)quit -- quitting from redis cli
4)set key value -- for example set name shreyas, set name as key nd value as shreyas
5)get key -- return value
6)del key -- return 1 if data presnt else 0. delete key
7)EXISTS key -- returns 1 if key is present else 0
8)keys <pattern> -- returns all keys matching the pattern for example, keys * return all keys
9)flushall---deletes all keys
10) ttl <key> --- example ttl name, returns -1 if no expiration, else return the number of seconds left until expiration, -2 means the key is not present
11)expire <key> <time in sec> -- for example expire name 10 - setsexpiration as 10 seconds for key name
12)setex key time value -- set with expiry time
13)handling list
	lpush listname value -- for example lpush shapes square - adds the data at the beginning of list
	rpush listname value -- adds data at the end of the list
	lrange listname 0 -1 -- get all the data starting from 0 the index till last
	lpop listname -- removes first element from the left and return the removed element
	rpop listname -- removes last eleemnt from the right and return the removed element
14)handling set -- unique data
	sadd setname value -- adds the data in set.
	smembers setname -- get all the values in set
	srem setname value -- removes the value from the list
15)handling hashes -- key value pair
	hset hashsetname key value -- example hset person name shreyas
	hget hashsetname key --gets the value from hash set
	hgetall hashsetname -- gets all key and value pair in this hashset
	hdel hashsetname key -- deletes the key from hashset
	hexists hashsetname key -- returns 1 if key exists else 0


To run application
go run main.go 1 --- 1 is the argument






