# goexpert-challenge-multithreading

Requests CEP data from `https://apicep.com/` and `https://viacep.com.br/` and shows the response from the faster one.  
If none responds in 1 second, shows a timeout error.

You can specify a CEP as an argument:

```
go run . 01489-900
```