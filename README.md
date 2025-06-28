# Credit Card Generator

This is a test harness that will generate valid credit card numbers for test purposes.

Currently supported card types are: 

- Visa
- Mastercard
- American Express.

To just generate cards go to [https://ccgenerator.alwaysdata.net/](https://ccgenerator.alwaysdata.net/)

## Development

For local development you will need to run this codebase in a server because it uses ES6 modules.  To make it easy to run locally, there is a go webserver included, just run the following:  

```golang
go run main.go
```

This will start up a local web-server hosting the card generation HTML page, you then just need to navigate to [http://localhost:3000](http://localhost:3000).

## Future

This is designed to be an easy way to generate test card numbers when testing apps that expect a valid credit card number.  All numbers are randomly generated on the fly.

The intent is to render something that is close enough to a real credit card to enable card scanners to scan the card displayed in the web page.  This is really aimed at people writing apps that require a credit card and allow you to use the camera to scan card details (although since the numbers are a valid format it will work for anything that requires a credit card number in a valid format).  

Unfortunately; it would seem that a lot of libraries are quite clever and will only scan embossed credit card numbers.  This means that at the moment it doesn't really work for intelligent libraries like this at the moment.