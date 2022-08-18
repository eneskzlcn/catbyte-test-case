### About Project
The project is like a backend-service of a message program.
It serves an endpoint to accept messages then it produces
a rabbitmq message. Another service that consumes rabbitmq
catches that message and write the message to the redis in a
format that each sender-receiver grouped messages are sorted
in chronological descending order. The last service is
the reporting service which serves an endpoint to list
all messages in chonological descending order with corresponding
sender-receiver pair.

### Local Setup

Firstly, you need to up the redis and rabbitmq containers
which are configured in docker-compose.yml

```shell
   docker compose up -d
```

Then Execute:

```shell
    go run main.go
```


### Endpoints

`GET: /message:` When you hit that endpoint with the request
has body like `{sender:"", receiver:"", message:""}`, it takes
that message and publishes that message to the rabbitmq.

**GET: /message/list?sender=x&receiver=y** When you hit that
endpoint with query params **sender:string**, **received: string**,It
reads the redis for senders message sent to receiver and returns a list of object
array in a format like
`[{sender:"", receiver:"", message:""},
{sender:"", receiver:"", message:""}
]` and in a chronological descending order.