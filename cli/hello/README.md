<img src="../../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# CLI Flags Lab

---
## <img src="../../assets/lab.png" width="auto" height="32"/> Mission


> Leveraging the code from the greeter lab, create a program that takes in parameters
> for the user name and age and display the message.

* Clone the Git repository from the template link into your Go workspace (Skip if already cloned!)
  * For Example: cd $GOPATH/src/github.com/YOUR_GIT_USERNAME
  * git clone https://github.com/gopherland/labs.git
* cd cli/hello
* Edit the file *main.go* and add your code
  * Note: Default user: *No one*
  * Note: Default age: *42* of course!
* Verify your application is running correctly with/without args and CLI help is displaying the correct information
<br/>
* BONUS! Print out the program help. Can you explain what's going on with Usage of? and how to fix it?

### Setup

```shell
# Navigate to your own Go workspace
$ cd $HOME/gopherland
# Clone Labs Repo if not done already
$ git clone https://github.com/gopherland/labs.git
# Lab dir
cd cli
# Get cracking!!
```

### Sample Output

```shell
go run main.go               # => Hello, my name is No One. I am 42 years old!"
go run main.go -u Fernand    # => Hello, my name is Fernand. I am 42 years old!
go run main.go -a 21         # => Hello, my name is No One. I am 21 years old!
go run -help
# Usage of hello:
# -a int
#       Specify a user age (default 42)
# -u string
#       Specify a user name (default "No One")
# Build your application and publish it to $GOBIN
go install
# In a separate terminal run your new application
hello -u Fred -a 10
```

---
<img src="../../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)