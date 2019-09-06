<img src="../../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# Errors Lab...

## Mission

> Enhance your FizzBuzz game with errors

* Clone github.com/gopherland/labs
* Cd to the errors lab
* The mission here is to surface out of range wrapped errors as follows:
  * If the number <=0 then issue an under range error wrapping it with the current number.
  * If the number is >20 then issue an over range error wrapping the error with the current number.
* Write a program that iterates from 0 to 21 (included) printing the corresponding FizzBuzz number.
  * If we get an UnderRange error, unwrap it and print the original error
  * Otherwise print the wrapped error.
* BONUS!! For comparison use the [Dave Cheney's error package](github.com/pkg/errors)

## Up And Running

```shell
go run game.go
```

## Sample output...

```text
Boom!! number is under range (<=0)
01 1
02 2
fizzbuzz (solution) δ >>> go run game.go
Boom!! number is under range (<=0)
01 1
02 2
03 Fizz
04 4
05 Buzz
06 Fizz
07 7
08 8
09 Fizz
10 Buzz
11 11
12 Fizz
13 13
14 14
15 FizzBuzz
16 16
17 17
18 Fizz
19 19
20 Buzz
Bam!! Invalid FizzBuzz# (21) -- number is over range (>20)
```

---
<img src="../../assets/imhotep_logo.png" width="32" height="auto"/> © 2019 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)