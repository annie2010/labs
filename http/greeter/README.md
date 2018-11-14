<img src="../../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# WebService Lab

---
## <img src="../../assets/lab.png" width="auto" height="32"/> Mission

> And that's a Wrap!! Decorate your user greeter as a Web Service...

* Ensure your web server is listening on port 4500
* Create an HTTP service with two endpoints ie / and /greet
* If the user hits / the service should produce a no route match error with a 400 status code
* /greet serves out new greetings
  * The request should take query parameters for the user name and age
  * You should issue a JSON response!
* Make sure the proper JSON headers are set
* Write the necessary tests to shake out your endpoints!
* Run your web service using cmd/svc/main.go
* Now run your cli to call your web service cmd/cli/main.go
* [BONUS] Hit a classmate greeting service if you share the same network

## Expected Output

```text
{
    "Age": 42,
    "Greeting": "Greetings Fernand! You are now 42 years old...",
    "Name": "Fernand"
}
```

## OSX Brew Installs (Totally Optional!!)

```shell
# Install httpie
brew install httpie
```

## Commands

```shell
# Start your web server
go run cmd/srv/main.go
# Hit the new_word url classic
curl -XGET 'http://localhost:4500/greet?name=Fernand&age=42'
# Or...
http :4500/greet name==Fernand age==42
# Run the cli
go run cmd/cli/main.go -age 10 -name fred -url localhost:4500
```

---
<img src="../../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)