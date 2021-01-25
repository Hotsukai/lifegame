# The Game of Life

## usage 
```
go run main.go <height :int> <width :int> <initializeLifeRate :float(0~1)> <interval :int>
```
### option
- `-r` , `--routine`
    - use GoRoutine
    - default is false
- `-p <boolean>` , `--print <boolean>`
    - print game Field
    - default is true
##  help
go run main.go -h

## sample 
```sh
go run main.go 10 10 0.4 1
go run main.go 10000 10000 0.4 0 -r -p false
```
