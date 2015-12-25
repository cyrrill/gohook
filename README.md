# gohooker
Receive Webhooks from services like Stripe, and forward them from your public server to your Go dev environment

## Server
Start the server by running

```
./gohooker
```

There must be an .htpasswd file in the same directory which contains the HTTP Basic Auth for the webhook. You can use apache2-utils, which installs the htpasswd utility.

A webhook service can be then pointed to:
```
http://<username>:<password>@<yourpublicserver.com>:8080/webhook
```
Anything in the POST body is passed along to the client


