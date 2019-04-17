
# Twido
Twido is a twitter bot library built from scratch that can notify twitter users a twitter video download link. It's also pretty modifiable, extendable and configurable, so you can add your own layer on it.

## **Twido's Dependencies**
There are some dependencies that Twido needs to work with. Twido uses Redis and communicate with it using Go Redis library. To start exploring it deeper, you need to install them.

 1. Install Golang on your machine. [Click for Tutorial](https://golang.org/doc/install).
 2. Install Redis Server/Client on your machine. [Click for Tutorial](https://www.digitalocean.com/community/tutorials/how-to-install-and-secure-redis-on-ubuntu-18-04).
 3. Install Go Redis by typing this in your terminal ``go get -u github.com/go-redis/redis``

## **How To Setup Twido**

 1. Clone this repository into your *$GOPATH*
 2. Make sure Twido's Dependencies are already installed on your machine
 3. Rename the **production.json.sample** config file in *config* directory to **production.json**
 4. Setup the **production.json** configuration
	 - Setup your twitter credentials key and put them in these lines (required)
	 
	 ```
	 "twitterApiKey": {
        "consumer": "<YOUR_CONSUMER_KEY>",
        "consumerSecret": "<YOUR_CONSUMER_SECRET_KEY>",
        "access": "<YOUR_ACCESS_TOKEN>",
        "accessSecret": "<YOUR_ACCESS_TOKEN_SECRET>"
    }
    ```
    - Setup your Rebrandly URL Shortener service credentials key and put them in these lines (optional)
    
    ```
    "urlShortener": {
        "rebrandly": {
            "active": true,
            "baseUrlApi": "https://api.rebrandly.com/v1/links",
            "apiKey": "<YOUR_API_KEY>",
            "domain": "<YOUR_CUSTOM_DOMAIN>"
        }
     }
    ```
    - Setup your forward link and notification string format, modify them in these lines (you can arrange reserved template tags such as *{video_url} {username} {tweet_id}* to your need)
    
    ```
    "stringFormat": {
        "forwardLink": "http://<YOUR_DOMAIN_HERE>/?video={video_url}&requested_by={username}&tweet_id={tweet_id}",
        "notification": "I ask the almighty God about the video link for you and it's accessible on: {video_url} \n\n-Don't hesitate to come back buddy, @{username}"
    }
    ```
