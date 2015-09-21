# Contributions for the Sept 20th WWG_NYC Hackathon 
This includes projects for two different chat apps, taking two different approaches using Go as backend. 

### Team 1 

Tammy and Hana's app leveraged Revel (http://revel.github.io/)  and their existing chat example, looking for ways to make it better. 
They forked the samples from https://github.com/revel/samples
and used the websockets package here: golang.org/x/net/websocket

Rooz, one of our mentors, also suggested bolt storage api to add persistence to the app between browser sessions 
https://github.com/boltdb/bolt

Andrey added a time parser.

Tammy lent her very own DO droplet to do the infra and deployed for all of us to see.  She added some INFRACODE (Tammy I ops heart you!) 
http://45.55.91.238:9000/


### Team 2

Redis as a message broker & db + Go + Redisgo (client lib)
Redis client: https://github.com/garyburd/redigo

Kevin (lead architect at Timehop) was an enormous help here. He & his team wrote a redis client (wrapper) for redigo! 
https://github.com/timehop/jimmy
Huge thank you to Timehop for the infra, There is a working version of this app at  http://wwgo.timehop.com/chat/wwgo (will stay up indefinitely as demo) 

Additional work was done, please add it to the repo via pull request! 



 
