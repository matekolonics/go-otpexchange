# go-otpexchange
Exchange rates API for Hungarian bank OTP written in Go.

## Installation:
`go get https://github.com/mattee12/go-otpexchange`

## Usage:
Let's convert 10 EUR to HUF.
```
buy, sell := otpexchange.convert(otpexchange.EUR, otpexchange.HUF, 10.0)
```
The function returns two float64 values (buy and sell).
