# appsconcept

This is my presentation of the fizz buzz aplication. I tried to make it as easy and to implement as it was in the real world. 

There were some options that I took that I regret somehow, like having raw SQL to make the queries. I did it like that because I though that implementing a DQL (like gorm) would take quite bit more time and, as time is not my side, I though I was being smart, however, the amount of erros I got due to syntax erros (because of reserved words) took me more time to fix than to actually implemented a gorm from the beginning.

Also, the application does not have a proper migration feature, meaning, in order to create the database, I would need to find a way to do it and I initially started by doing the query on the code, but than changed it to a sql file that runs on the docker-compose. This is simply a workaround that saved me a bit of time.

Beside that, the application is running by simply having a dockerfile creating a executable and running the same one. With that, the application will run and we can now make requests (i used Postman) to check every feature.

The application is using a DDD approach where I have specific folders for specific parts of the application. So:

- "cmd" folder is responsible for having our entrypoint commands.
- "internal" folder is where everything else related to the logic and dependencies is located, like:
    - "domain" folder for the entities model;
    - "repository" folder for the init of database access and it's SQL ;
    - "service" where it will be where the main business logic is located and having access to all other dependencies.

To boot the application a simple __docker compose up --build__ should do the trick and you have now a full functional webservice at your disposal.

The endpoints are the follow:
- http://localhost:8080/fizzbuzz?int1=3&int2=5&limit=15&str1=gohan&str2=goku
- http://localhost:8080/stats



