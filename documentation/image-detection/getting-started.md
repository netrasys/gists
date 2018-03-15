<!-- 
To publish to the Azure API Management portal, you must first convert this markdown to raw HTML
Use this website to do so: https://markdowntohtml.com/
-->

<head>
  <title>github-embed demo</title>
  <link rel="stylesheet" type="text/css" href="https://apimgmtsthfps9my3wlnnljm.blob.core.windows.net/content/MediaLibrary/assets/github-embed-1.css">
</head>
<style>
  
  #settings-object {
    margin-bottom: 2rem;
    height: 500px;
  }
</style>

# Getting Started

> The Async Image Detection API allows users to add images to a queue for processing. Once the image has completed processing, a POST callback response will be returned from the Async Image Detection API server to your provided `callback_url` containing any tags found. 

![sequence](https://storage.googleapis.com/netra-public-data/async-image-detection/documentation/api_sequence.png)

# Testing the API (Demo)

> In this demonstration, we'll be making an API call to the `/brands` endpoint.

## Setup a callback endpoint

> To help us get started, we'll be using [webhookinbox](http://webhookinbox.com/) to accept and inspect callbacks.

> Head to [http://webhookinbox.com/](http://webhookinbox.com/) and select "create an inbox". This will setup a temporary inbox with a web address to send our callbacks to. Save the web address of the temporary inbox you just created, we'll use this in our request to the Async Image Detection API. (It should look something like http://api.webhookinbox.com/i/XXXXX/)

## Setting up the request

> We'll make our request to the `/brands` endpoint and use this image below to identify the logo present.

### Test Image

<img src="https://storage.googleapis.com/netra-public-data/async-image-detection/documentation/api-test-image.jpg" width=50%>

### code

> Set the variables `CALLBACK_URL` and `SUBSCRIPTION_KEY`. Make sure to use the callback address that we setup up earlier with Webhook Inbox.

<div id="settings-object"></div>

### Response

> After running the request we should see response a in the console.

```
{
  "message": "Image successfully added to processing queue, results will be sent to callback_url shortly.",
  "image_url": "https://storage.googleapis.com/netra-public-data/async-image-detection/documentation/api-test-image.jpg",
  "callback_url": "http://api.webhookinbox.com/i/yvJlr2fl/in/",
  "request_id": "jas58r9at4jercert1"
}
```

### Callback

> Within a few seconds we should see the callback POST to our Webhook Inbox

<img src="https://storage.googleapis.com/netra-public-data/async-image-detection/documentation/webhook.png" width=75%>

<script src="https://cdnjs.cloudflare.com/ajax/libs/babel-polyfill/6.23.0/polyfill.min.js"></script>
<script src="https://apimgmtsthfps9my3wlnnljm.blob.core.windows.net/content/MediaLibrary/assets/github-embed.min-1.js"></script>
<script>
githubEmbed('#settings-object', 'https://github.com/netrasys/gists/blob/master/documentation/image-detection/.gh-embed.json');
</script>




