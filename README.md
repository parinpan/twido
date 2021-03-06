# Twido
Twido is a twitter bot library built from scratch that can notify twitter users a twitter video download link. It's also pretty modifiable, extendable and configurable, so you can add your own layer on it. [Let's Take a Look at TwidoPlease Download the Video Twitter Account for Example](https://twitter.com/twidopls).

!["Twido Interaction to User when request a twitter video download link"](https://lh3.googleusercontent.com/-IrhrDWM8iJx8tb9mU867QEaXUK7oZAQ_nwPUNl2oyvphmXMaXUIebSfH6KbYQn5Ho1jfUwaBDvW)

## **Twido's Dependencies**
There are some dependencies that Twido needs to work with. Twido uses Redis and communicate with it using Go Redis library. To start exploring it deeper, you need to install them.

 1. Install Golang on your machine. [Click for Tutorial](https://golang.org/doc/install).
 2. Install Redis Server/Client on your machine. [Click for Tutorial](https://www.digitalocean.com/community/tutorials/how-to-install-and-secure-redis-on-ubuntu-18-04).
 3. Install Go Redis by typing this in your terminal ``go get -u github.com/go-redis/redis``

## **How To Setup Twido**

 1. Clone this repository into your *$GOPATH*
 2. Make sure Twido's Dependencies are already installed on your machine
 3. Rename the **production.json.sample** config file in project *config* directory to **production.json**
 4. Setup the **production.json** configuration
 
	 - Setup your twitter credential keys and put them in these lines (required)
	 
	 ```json
	 "twitterApiKey": {
        "consumer": "<YOUR_CONSUMER_KEY>",
        "consumerSecret": "<YOUR_CONSUMER_SECRET_KEY>",
        "access": "<YOUR_ACCESS_TOKEN>",
        "accessSecret": "<YOUR_ACCESS_TOKEN_SECRET>"
    }
    ```
    - Setup your Rebrandly URL Shortener service credential key and put them in these lines, setup the active field to *false* when you don't need this service (optional)
    
    ```json
    "urlShortener": {
        "rebrandly": {
            "active": true,
            "baseUrlApi": "https://api.rebrandly.com/v1/links",
            "apiKey": "<YOUR_API_KEY>",
            "domain": "<YOUR_CUSTOM_DOMAIN>"
        }
     }
    ```
    - Setup your twitter (bot) account username and max count per search, modify them in these lines
    ```json
    "twitterObservation": {
		"keyword": "@yourtwitterusername",
		"maxSearchCount": "100"
	}
    ``` 
    - Setup your forward link and notification string format, modify them in these lines (you can arrange reserved template tags such as *{video_url} {username} {tweet_id}* to your need)
    
    ```json
    "stringFormat": {
        "forwardLink": "http://<YOUR_DOMAIN_HERE>/?video={video_url}&requested_by={username}&tweet_id={tweet_id}",
        "notification": "I ask the almighty God about the video link for you and it's accessible on: {video_url} \n\n-Don't hesitate to come back buddy, @{username}"
    }
    ```
    - Setup the configuration absolute path to where it's actually located by opening **init.go** file in project *config* directory, find these lines and change it to the right absolute path (required)
    ```go
    var TwidoConfig, TwidoConfigErr = NewConfiguration(ConfigurationOption{
		    Environment: "production",
		    BasePath: 	 "/absolute/path/go/src/twido/config", // change it to yours
	})
    ```
5. Hooray, the setup is done! If you still feel roaming about Twitter & Rebrandly Credential Keys, it may help you:
	- [Twitter Developer Docs](https://developer.twitter.com/en.html)
	- [Rebrandly Developer Docs](https://developers.rebrandly.com/docs/get-started)

## **How To Run Twido After Setting Up**
This is the most exciting part. You're gonna run twido on twitter as a bot that will help users to find their twitter video download link. Basically, you just need five steps below:
1. Build the **twido.go** file in the project directory into an executable file, type it in your terminal:

```
go build /path/to/the/project/directory/twido.go
```

2. After building **twido.go** file, now you have a file named **twido** without *.go*. It's an executable file you need to run the bot. Move it to your */usr/local/bin*, type it in your terminal:
```
mv twido /usr/local/bin/twido
```

3. Now, you need to execute */usr/local/bin/twido* executable file every **N** minute(s), so it can check the users' new requests continuously. Here, you will need the cronjob service. Type it in your terminal:
```
crontab -e
```
4. Add the snippet below to the end of crontab file line. So, the executable file will be executed by the operating system every 1 minute.
```
*/1 * * * * /usr/local/bin/twido >> /your/path/to/the/custom.log 2>&1
```
5. Done

## **Contact Me** 
If you want to offer opportunities or facing difficulties setting up this twitter bot. You can catch me up on:
- [Linkedin](https://linkedin.com/in/fachrinfan)
- [Twitter](https://twitter.com/fachrinfan)
- [Instagram](https://instagram.com/fachrinfan)
- [Email](mailto:fachrinnn@gmail.com)
