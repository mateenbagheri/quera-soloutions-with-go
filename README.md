# Quera soloutions with Go

Here I keep my soloutions for [Quera](https://quera.org/) problems written in Golang.
Starting this repository to try out learning Go better and have something to do in 
my free time. 

## Structure
Here is a sample of project dir tree.
```
Project Dir:.
│   go.mod
│   README.md
│
└───contest
    ├───<Problem ID>-<Problem Name>
    │       <Problem ID>.go
    │       problem.pdf
    ...
```
As you can see, each folder/dir contains a problem pdf explaining the problem and having
examples for how to solve it and a solution file written in Go. 
## How to find problem ID
In quera website, each problem is assiciated with a unique problem ID which is embedded in
Its webpage URL.
For example, the problem ID for the problem bellow is 52545
```https://quera.org/problemset/52545/```
## How to run
If you have any specific problem that you want to run and check, you simply just need to 
have Go installed. Clone the project and write the following commad.
```
go run contest/<problem dir>/<problemID>.go
```
Only educational usage of this repository is consented.
