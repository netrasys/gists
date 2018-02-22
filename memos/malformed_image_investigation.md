# Malformed Image Investigation

It seems that the issue regarding Netra API callbacks returning "Malformed Image" is revolving around pbs.twimg.com (Twitter's CDN Service). An image can become available for a moment, and then become unavailable shortly after. The issue was brought up within the Twitter Development Community (See: [pbs-twimg-com-cdn-problems](https://twittercommunity.com/t/pbs-twimg-com-cdn-problems/92623)). After a number of tests, Netra can also confirm that the issue persists today.

## Simple Browser Download

Using a browser, we tried to retrieve an image and it was not available.

![https://i.imgur.com/jeyamgM.png](https://i.imgur.com/jeyamgM.png)

Moments later, the same image was available.

![https://i.imgur.com/cNuR7t4.png](https://i.imgur.com/cNuR7t4.png)

## Download Image Programatically

To reproduce the Netra API, we expiremented programatically.

### Script

This scripts downloads, and writes the file from an array of `urls`

```python
import requests

urls = ["https://pbs.twimg.com/media/DSfK-tXX4AcxLOs.jpg"]

for i, url in enumerate(urls):
	file_name = url.split('/')[len(url.split('/'))-1]
	print 'downloading {}...'.format(file_name)
	response = requests.get(url)
	print response.status_code
	if response.status_code == 200:
		print 'writing file {}...'.format(file_name)
		with open(file_name, 'wb') as f:
			f.write(response.content)
```

### Running the Script

The script did work, and successfully downloaded the image.

```
| ~/Repositories/tests/malformed_image @ Dominics-MacBook-Pro (dominiccabral)
| => python script.py
downloading DSfK-tXX4AcxLOs.jpg...
200
writing file DSfK-tXX4AcxLOs.jpg...
```

Moments later, the resource was unavailable.

```
________________________________________________________________________________
| ~/Repositories/tests/malformed_image @ Dominics-MacBook-Pro (dominiccabral)
| => python script.py
downloading DSfK-tXX4AcxLOs.jpg...
404
```

## Simple nslookup

nslookup is a network administration command-line tool for querying the Domain Name System to obtain domain name or IP address mapping or for any other specific DNS record.

```
________________________________________________________________________________
| ~ @ Dominics-MacBook-Pro (dominiccabral)
| => date
Tue Jan  2 13:18:37 EST 2018
________________________________________________________________________________
| ~ @ Dominics-MacBook-Pro (dominiccabral)
| => nslookup pbs.twimg.com
Server:		192.168.0.1
Address:	192.168.0.1#53

Non-authoritative answer:
pbs.twimg.com	canonical name = wildcard.twimg.com.
Name:	wildcard.twimg.com
Address: 104.244.46.71
```
Moments later...
```
________________________________________________________________________________
| ~ @ Dominics-MacBook-Pro (dominiccabral)
| => date
Tue Jan  2 13:20:07 EST 2018
________________________________________________________________________________
| ~ @ Dominics-MacBook-Pro (dominiccabral)
| => nslookup pbs.twimg.com
Server:		192.168.0.1
Address:	192.168.0.1#53

Non-authoritative answer:
pbs.twimg.com	canonical name = cs196.wac.edgecastcdn.net.
cs196.wac.edgecastcdn.net	canonical name = cs2-wac.apr-8315.edgecastdns.net.
cs2-wac.apr-8315.edgecastdns.net	canonical name = cs2-wac-us.8315.ecdns.net.
cs2-wac-us.8315.ecdns.net	canonical name = cs45.wac.edgecastcdn.net.
Name:	cs45.wac.edgecastcdn.net
Address: 72.21.91.70
