<img src="../../assets/gophernand.png" align="right" width="100" height="auto"/>

<br/>
<br/>
<br/>

# LetsGo! Testing FizzBuzz Lab...

## <img src="../../assets/lab.png" width="auto" height="32"/> Mission

> Remember the FizzBuzz lab in the Control section?
> Write tests to ensure you are generating numbers correctly.

* Clone the Git repository from the template link into your **letsgo** workspace.
* Use the existing solution or copy your own solution from the control section into game.go
* Test your compute function using table testing!
* Ensure you've achieved complete code coverage!

### Setup

```shell
# Navigate to your own Go workspace
$ cd $HOME/letsgo
# Clone Labs Repo
$ git clone https://git@github.com/imhotepio/letsgo_labs.git
# Lab dir
cd $HOME/letsgo/letsgo_labs/testing/fizzbuzz
# Dependencies
$ go get -u github.com/cespare/reflex
# Running tests
$ go test
# Watch your function/test files and run the tests!
$ reflex -g '*.go' go tests
# Check coverage and generate HTML report
$ go test -coverprofile=cov.out && go tool cover -html=cov.out
```

---
<img src="../../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)