# Model Update + New Features 2.2.0

**Schedule Deployment:** Friday Jan 19, 2019 11:00 AM EST

## Changes

- Hierarchy for Context model (see reference below)
- 2784 Model Update
    - Lotus Crackers
    - Make Up For Ever
    - Call of Duty World League
    - Albert Heijn
    - Shreiber
    - Pacific Union Financial
    - Trebor Confectionery
    - Columbus Crew
    - Bryant HVAC
    - GK Elite Sportswear
    - Columbus Blue Jacket
    - Yogi Tea
    - HEMA Coffee
    - Seagrams
    - Cadbury's Starbar
    - USTA
    - ULTA Salon
    - Zesta Tea
    - Dodge SRT
    - Absinthe Brevans H.R. Giger
    - Douwe Egberts
    - Land O'Lakes
    - Yelp
    - Jewel Osco
    - Go Naturally Organic
    - Meridian Vineyards
    - Russia Fifa World Cup 2019
    - McDonald's Vers Gemalen Koffie

## Reference

### Context Hierarchy Example

```
{
   "context_hierarchy":{
      "People & Relationships":{
         "People":{
            "Athlete":[
               "tennis player",
               "athlete",
               "football player"
            ],
            "Person":[
               "person"
            ]
         }
      },
      "Sports & Recreation":{
         "Sports":{
            "Sports":[
               "sport"
            ],
            "Tennis":[
               "tennis"
            ]
         }
      }
   },
   "context":[
      {
         "confidence":86,
         "label":"tennis player"
      },
      {
         "confidence":77,
         "label":"athlete"
      },
      {
         "confidence":72,
         "label":"sports"
      },
      {
         "confidence":39,
         "label":"person"
      },
      {
         "confidence":38,
         "label":"tennis"
      },
      {
         "confidence":34,
         "label":"football player"
      }
   ],
   "image_url":"https://pbs.twimg.com/media/DSTDOACW0AE9uXo.jpg",
   "errors":[

   ],
   "request_id":"9ljkgtgjc24co2i"
}
```

# Test Results

## Brands

<table class="tap" width="100%"><tbody><tr><th></th><th>Number</th><th>Description</th><th>Directive</th></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">1</td><td width="65%">Basic Performance:  random VALID images</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">2</td><td width="65%">Basic Performance:  images that contain brands</td><td width="25%"></td></tr><tr><td></td><td colspan="3">
# tests 2<br></td></tr><tr><td></td><td colspan="3">
# pass 2<br></td></tr><tr><td></td><td colspan="3">
# fail 0<br></td></tr><tr><td></td><td colspan="3">
# skip 0<br></td></tr></tbody></table>

<table class="tap" width="100%"><tbody><tr><th></th><th>Number</th><th>Description</th><th>Directive</th></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">1</td><td width="65%">Basic Functionality:  Send 100 images random INVALID images</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">2</td><td width="65%">Basic Functionality:  Send 100 images random VALID images</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">3</td><td width="65%">Basic Functionality:  Send 100 images images that contain brands</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">4</td><td width="65%">Basic Functionality:  Send 100 images random VALID images from client w/ isolated GPU enabled</td><td width="25%"></td></tr><tr><td></td><td colspan="3">
# tests 4<br></td></tr><tr><td></td><td colspan="3">
# pass 4<br></td></tr><tr><td></td><td colspan="3">
# fail 0<br></td></tr><tr><td></td><td colspan="3">
# skip 0<br></td></tr></tbody></table>

<table class="tap" width="100%"><tbody><tr><th></th><th>Number</th><th>Description</th><th>Directive</th></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">1</td><td width="65%">Dedicated Functionality:  Send 100 images random VALID images from client w/ isolated GPU enabled</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">2</td><td width="65%">Dedicated Functionality:  Send 100 images random VALID images from normie client, should not return</td><td width="25%"></td></tr><tr><td></td><td colspan="3">
# tests 2<br></td></tr><tr><td></td><td colspan="3">
# pass 2<br></td></tr><tr><td></td><td colspan="3">
# fail 0<br></td></tr><tr><td></td><td colspan="3">
# skip 0<br></td></tr></tbody></table>

<table class="tap" width="100%"><tbody><tr><th></th><th>Number</th><th>Description</th><th>Directive</th></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">1</td><td width="65%">Priority Functionality:  Send 100 images checks to see that default priority is working</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">2</td><td width="65%">Priority Functionality:  Send 100 images checks to see that HIGH priority is working</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">3</td><td width="65%">Priority Functionality:  Send 100 images checks to see that LOW priority is working</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">4</td><td width="65%">Priority Functionality:  Send 100 images checks to see that DEFAULT, HIGH, and LOW priority is working</td><td width="25%"></td></tr><tr><td></td><td colspan="3">
# tests 4<br></td></tr><tr><td></td><td colspan="3">
# pass 4<br></td></tr><tr><td></td><td colspan="3">
# fail 0<br></td></tr><tr><td></td><td colspan="3">
# skip 0<br></td></tr></tbody></table>

## Humans

<table class="tap" width="100%"><tbody><tr><th></th><th>Number</th><th>Description</th><th>Directive</th></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">1</td><td width="65%">Basic Performance:  random VALID images</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">2</td><td width="65%">Basic Performance:  images that contain humans</td><td width="25%"></td></tr><tr><td></td><td colspan="3">
	# tests 2<br></td></tr><tr><td></td><td colspan="3">
	# pass 2<br></td></tr><tr><td></td><td colspan="3">
	# fail 0<br></td></tr><tr><td></td><td colspan="3">
	# skip 0<br></td></tr></tbody></table>

<table class="tap" width="100%"><tbody><tr><th></th><th>Number</th><th>Description</th><th>Directive</th></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">1</td><td width="65%">Basic Functionality:  Send 100 images random INVALID images</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">2</td><td width="65%">Basic Functionality:  Send 100 images random VALID images</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">3</td><td width="65%">Basic Functionality:  Send 100 images images that contain humans</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">4</td><td width="65%">Basic Functionality:  Send 100 images random VALID images from client w/ isolated GPU enabled</td><td width="25%"></td></tr><tr><td></td><td colspan="3">
	# tests 4<br></td></tr><tr><td></td><td colspan="3">
	# pass 4<br></td></tr><tr><td></td><td colspan="3">
	# fail 0<br></td></tr><tr><td></td><td colspan="3">
	# skip 0<br></td></tr></tbody></table>

<table class="tap" width="100%"><tbody><tr><th></th><th>Number</th><th>Description</th><th>Directive</th></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">1</td><td width="65%">Dedicated Functionality:  Send 100 images random VALID images from client w/ isolated GPU enabled</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">2</td><td width="65%">Dedicated Functionality:  Send 100 images random VALID images from normie client, should not return</td><td width="25%"></td></tr><tr><td></td><td colspan="3">
	# tests 2<br></td></tr><tr><td></td><td colspan="3">
	# pass 2<br></td></tr><tr><td></td><td colspan="3">
	# fail 0<br></td></tr><tr><td></td><td colspan="3">
	# skip 0<br></td></tr></tbody></table>

<table class="tap" width="100%"><tbody><tr><th></th><th>Number</th><th>Description</th><th>Directive</th></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">1</td><td width="65%">Priority Functionality:  Send 100 images checks to see that default priority is working</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">2</td><td width="65%">Priority Functionality:  Send 100 images checks to see that HIGH priority is working</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">3</td><td width="65%">Priority Functionality:  Send 100 images checks to see that LOW priority is working</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">4</td><td width="65%">Priority Functionality:  Send 100 images checks to see that DEFAULT, HIGH, and LOW priority is working</td><td width="25%"></td></tr><tr><td></td><td colspan="3">
	# tests 4<br></td></tr><tr><td></td><td colspan="3">
	# pass 4<br></td></tr><tr><td></td><td colspan="3">
	# fail 0<br></td></tr><tr><td></td><td colspan="3">
	# skip 0<br></td></tr></tbody></table>

## Context

<table class="tap" width="100%"><tbody><tr><th></th><th>Number</th><th>Description</th><th>Directive</th></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">1</td><td width="65%">Basic Performance:  random VALID images</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">2</td><td width="65%">Basic Performance:  images that contain context</td><td width="25%"></td></tr><tr><td></td><td colspan="3">
	# tests 2<br></td></tr><tr><td></td><td colspan="3">
	# pass 2<br></td></tr><tr><td></td><td colspan="3">
	# fail 0<br></td></tr><tr><td></td><td colspan="3">
	# skip 0<br></td></tr></tbody></table>

<table class="tap" width="100%"><tbody><tr><th></th><th>Number</th><th>Description</th><th>Directive</th></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">1</td><td width="65%">Basic Functionality:  Send 100 images random INVALID images</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">2</td><td width="65%">Basic Functionality:  Send 100 images random VALID images</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">3</td><td width="65%">Basic Functionality:  Send 100 images images that contain context</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">4</td><td width="65%">Basic Functionality:  Send 100 images random VALID images from client w/ isolated GPU enabled</td><td width="25%"></td></tr><tr><td></td><td colspan="3">
	# tests 4<br></td></tr><tr><td></td><td colspan="3">
	# pass 4<br></td></tr><tr><td></td><td colspan="3">
	# fail 0<br></td></tr><tr><td></td><td colspan="3">
	# skip 0<br></td></tr></tbody></table>

<table class="tap" width="100%"><tbody><tr><th></th><th>Number</th><th>Description</th><th>Directive</th></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">1</td><td width="65%">Dedicated Functionality:  Send 100 images random VALID images from client w/ isolated GPU enabled</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">2</td><td width="65%">Dedicated Functionality:  Send 100 images random VALID images from normie client, should not return</td><td width="25%"></td></tr><tr><td></td><td colspan="3">
	# tests 2<br></td></tr><tr><td></td><td colspan="3">
	# pass 2<br></td></tr><tr><td></td><td colspan="3">
	# fail 0<br></td></tr><tr><td></td><td colspan="3">
	# skip 0<br></td></tr></tbody></table>

<table class="tap" width="100%"><tbody><tr><th></th><th>Number</th><th>Description</th><th>Directive</th></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">1</td><td width="65%">Priority Functionality:  Send 100 images checks to see that default priority is working</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">2</td><td width="65%">Priority Functionality:  Send 100 images checks to see that HIGH priority is working</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">3</td><td width="65%">Priority Functionality:  Send 100 images checks to see that LOW priority is working</td><td width="25%"></td></tr><tr><td class="green" width="5%"></td><td width="5%" class="center">4</td><td width="65%">Priority Functionality:  Send 100 images checks to see that DEFAULT, HIGH, and LOW priority is working</td><td width="25%"></td></tr><tr><td></td><td colspan="3">
	# tests 4<br></td></tr><tr><td></td><td colspan="3">
	# pass 4<br></td></tr><tr><td></td><td colspan="3">
	# fail 0<br></td></tr><tr><td></td><td colspan="3">
	# skip 0<br></td></tr></tbody></table>
