import requests

SUBSCRIPTION_KEY = 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXX'
CALLBACK_URL = 'http://api.webhookinbox.com/i/XXXXXX/'

headers = {
    'Ocp-Apim-Subscription-Key': SUBSCRIPTION_KEY,
    'Content-Type': 'application/json'
}

request = {
    'image_url': 'https://storage.googleapis.com/netra-public-data/async-image-detection/documentation/api-test-image.jpg',
    'callback_url': CALLBACK_URL
}

response = requests.post('https://api.getnetra.com/image-detection/process/brands',
                         headers=headers,
                         json=request)
print 'Sent request:', response.status_code, response.text
