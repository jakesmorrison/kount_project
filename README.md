# Coding Project for Kount

This is a coding project for Kount. The requirements for this project were:

1) Install Go. 
2) Use resty with error handling to access database through api. 
3) Write unit test. 
4) Write benchmark functions. 
5) Submit code via github. 

I wrote 3 simple Gp functions to access a rest api I setup on my personal server. The values in the database are vocabulary words (word, type, definition). The code has the ability to retrieve a word from the database, add a word to the database, and delete a word from the database. 

## Files

- access_api.go: Main class that has api calling functions. 
- test_api.go: Testing class that test each function in the main class individually. 
- benchmark_api.go: I ran out of time, so this class is only returning the api response times. 
- results.txt: This file will contain all testing and benchmark results. 
- debug.txt: This file has all debug information from the api calls. 
- README.md

## Running Test and Benchmarks

- Open command line from folder
- Run, "go build access_api.go test_api.go benchmark_api.go". 
- Then run, "./access_api Test_and_Benchmark"
