'use strict'

const request = require('request');
const SUBSCRIPTION_KEY = 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXX';
const CALLBACK_URL = 'http://api.webhookinbox.com/i/XXXXXX/'

const options = {
	url: 'https://api.getnetra.com/image-detection/process/brands',
	method: 'POST',
	headers: {
	  'Ocp-Apim-Subscription-Key': SUBSCRIPTION_KEY,
	  'content-type' : 'application/json'
	},
	body: JSON.stringify({
		"image_url":"https://storage.googleapis.com/netra-public-data/async-image-detection/documentation/api-test-image.jpg",
		"callback_url": CALLBACK_URL
	})
};

function response(error, res, body) {
	if (!error) {
		var info = JSON.parse(body);
		console.log(info);
	}else{
		console.log(error);
	}
}

request(options, response);