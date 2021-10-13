# Fizz-buzz REST API in Go

This project is a technical test assessed by Leboncoin. The main 
objective of this project is to develop a Go REST API based on 
the fizz-buzz logic. It consists in writing all numbers from 1 to 
100, and just replacing all multiples of 3 by "fizz", all 
multiples of 5 by "buzz", and all multiples of 15 by "fizzbuzz".
The output would look like this: "1,2,fizz,4,buzz,fizz,7,8,fizz,
buzz,11,fizz,13,14,fizzbuzz,16,...".

In order to assess my hard skills, Leboncoin asked me to let 
the client choose the numbers, the words and define a limit. As a
bonus challenge, Leboncoin asked me to include an endpoint 
related to statistics.

> **This project is developed to be deployed and tested with 
Heroku services.** If you do not wish to use these services, 
don't forget to check that environment variables `PORT` and
`DATABASE_URL` are set in the environment running the API!


## What is required to make it work?

In order to make the API work, you will need to create a 
PostgreSQL v13 database (local, Docker, Heroku, AWS, Azure, 
etc.), with a table named `fizzbuzz_queries` that includes the 
following columns:

- `id` as "type" `serial` (see doc [here](https://www.postgresql.org/docs/13/datatype-numeric.html#DATATYPE-SERIAL))
- `ip_address` as type `varchar`
- `time` as type `timestamp` and default `NOW()`
- `first_int` as type `int8`
- `second_int` as type `int8`
- `limit` as type `int8`
- `first_string` as type `varchar`
- `second_string` as type `varchar`
- `params_hash` as type `char(32)`

Once the database is created, set an environment variable 
`DATABASE_URL` with your database's URL as the value.


## How to start up and use the API?

After you have created the database, run the following commands 
in your terminal:

    git clone https://github.com/CyrilleGG/leboncoin-fizzbuzz.git
    
    go get

You can run the API locally with this command:

    go run main.go

Once the API started, you will be able to test two endpoints:

- `/fizzbuzz`
- `/fizzbuzz-metrics`

In a first usage case, it is recommended to test the `/fizzbuzz` 
endpoint before testing the `/fizzbuzz-metrics` one. You can 
test the `/fizzbuzz` endpoint with the following parameters:

- `first_int`, `second_int` and `limit` as **integers**
- `first_string` and `second_string` as **strings**

> **Note:** If you don't specify some parameters, the API will 
> use default values for the fizz-buzz function to fill the gaps.


## How to deploy the project with Heroku?

If you planned to deploy this project with Heroku, and if you 
didn't create any Heroku app for this project, open the project 
directory in your terminal and run:

    make setup

> **Note:** this command will work if you installed Heroku CLI 
> and signed in to Heroku. After everything is deployed, don't
> forget to set up the database, as mentioned in **What is 
> required to make it work?**

If you already created the Heroku app and add-on for this 
project, you can run in your terminal:

    make deploy


## Project organization

This project is organized as follows:

- the `main` package is used to start the API
- code related to HTTP operations (as responses, routes, server, 
etc.) is handled by the `server` package
- configuration for database connection is handled by the 
`database` package
- source code related to the `/fizzbuzz` endpoint is in 
`fizzbuzz/fizzbuzz.go` and `fizzbuzz/fizzbuzzmetrics.go`
- source code related to the `/fizzbuzz-metrics` endpoint is in 
`fizzbuzz/fizzbuzzmetrics.go`

## Additional Remarks

#### What is the purpose of hashing queries' parameters in the database?

In the first place, hashing a cryptographic tool used to 
protect data. But, it can also be used as a unique digital 
fingerprint of a dataset (bytes). If you hash the same dataset 
numerous times, you will get the same result. In this project, 
hashing the queries' parameters helps in making the SQL query, 
for the `/fizzbuzz-metrics` endpoint, easier to write because you 
can use `GROUP BY` on a single column, instead of 5.

---

It took me 5 days to develop this project, including researches 
to make it work. I don't consider it finished. Here is a list of 
things I would have enhanced/added:

- use of a database technology developed for recording events, 
like Apache Kafka, instead of PostgreSQL
- more go routines
- better error handling
- better HTTP response handling
- better unit testing (`fizzbuzz_test.go` worked fine before 
adding database related operations)
- get rid of [Julien Schmidt's httprouter](https://github.com/julienschmidt/httprouter) 
because it needs some updates (better query parameters handling, 
unit testing features, etc.)

If you have any feedbacks or questions, send an e-mail to <cyrilleguiot@live.fr>.

Enjoy!