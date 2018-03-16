# Security Guide

> Due to the asynchronous nature of Netra’s API, it is required that any user of the Image Detection API expose a server endpoint for callback data. To ensure security, please follow the best practices listed below.

## Use HTTPs

> Netra only accepts API requests to be made over the `https://`
 protocol. To ensure end-to-end encryption,  although not required, it is suggested that the client’s “callback_url”/exposed endpoint is also `https://`
.

## Secure Callbacks

> A request to any of the “/process” endpoints allows the client to submit their own query string that will later be returned with their callback as query parameters. 

*Example of an HTTP request using query string to pass in custom parameters. Setting `foo` to `bar` and `token` to `123abc`*

```
POST https://api.getnetra.com/image-detection/process/brands?foo=bar&token=123abc HTTP/1.1
Host: api.getnetra.com
Content-Type: application/json
Ocp-Apim-Subscription-Key: ••••••••••••••••••••••••••••••••

{
	"image_url":"https://storage.googleapis.com/netra-public-data/async-image-detection/documentation/api-test-image.jpg",
	"callback_url":"https://my.webapp.com/callback",
	"threshold":80
}
```

> To prevent unwanted access to the client’s endpoint, It is recommended the client use this feature to add a “token” or “key” to in the query string of their request to then be validated by the client’s exposed endpoint upon callback. 

__Note: If the client’s callback is not `https://`
, the query string would technically be exposed.__
