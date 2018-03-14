export SUBSCRIPTION_KEY=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
export CALLBACK_URL=http://api.webhookinbox.com/i/XXXXXX/

curl -v -X POST "https://api.getnetra.com/image-detection/process/brands" \
-H "Content-Type: application/json" \
-H "Ocp-Apim-Subscription-Key: $SUBSCRIPTION_KEY" \
-d '{ "image_url":"https://storage.googleapis.com/netra-public-data/async-image-detection/documentation/api-test-image.jpg", "callback_url":"'$CALLBACK_URL'", "threshold":80 }'
