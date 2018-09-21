<img src="assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

[![Build Status](https://travis-ci.org/YOUR_GIT_HANDLE/picker.svg?branch=master)](https://travis-ci.org/YOUR_GIT_HANDLE/picker)
[![Go Report Card](https://goreportcard.com/badge/github.com/YOUR_GIT_HANDLE/picker)](https://goreportcard.com/report/github.com/YOUR_GIT_HANDLE/picker)
[![GoDoc](https://godoc.org/github.com/YOUR_GIT_HANDLE/picker?status.svg)](http://godoc.org/github.com/YOUR_GIT_HANDLE/picker)

# Package Lab

---
## <img src="assets/lab.png" width="auto" height="32"/> Mission

> Package deal! Implement a picker package to pick a word from a given word dictionary

* The picker package loads words from the given assets directory and randomly pick a new word
* Copy gopherland/labs/packaging/picker to your $HOME/gopherland/picker
* Create a new module file using the following command:
  ```shell
  go mod init github.com/YOUR_GIT_USER_HANDLE/picker
  ```
* Update the test import path to the picker to reflect your own git user handle
* Run the test command to make sure the tests are still passing and coverage is good!
* Ensure all exported items have the correct documentation
* Using your own github account, create a new public repo named **picker**
* Setup your git repo using:
  ```shell
  git init
  git add .
  git commit -m 'Init drop'
  git remote add origin git@github.com:YOUR_GIT_USER_HANDLE/picker.git
  git push --set-upstream origin master
  ```
* Using semantic versioning tag your release
  ```shell
  git tags v1.0.0
  git push --tags
  ```
* In cmd/main.go update the import path to use your new repo
* In the root of your repo run your cli app
  ```shell
  go run cmd/main.go
  go run cmd/main.go -dic musicians -dir assets
  # Clean up your dependencies
  go mod tidy
  ```
* Checkout your docs and score card!
  ```shell
  open https://godoc.org/github.com/YOUR_GIT_USER_HANDLE/picker
  open https://goreportcard.com/report/github.com/YOUR_GIT_USER_HANDLE/picker
  ```
* Update the README.md badge section with your own github handle
* Congratulation!! You've just built and published your very own go package!
*
* [BONUS] Change your CLI to reference your classmate repo and run the CLI using
  their implementation!!
* [BONUS] Use [Travis CI](https://travis-ci.org), add your repo and earn an extra badge!

---
<img src="assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)