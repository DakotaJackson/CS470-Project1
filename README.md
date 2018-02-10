# Project 1 - Path Finding
* Use the following Algorithms
    - Breadth First
    - Lowest Cost
    - Greedy Best First
    - A* With 2+ Different Heuristics

## End Goal
* Completly analyze each of the 4 algorithms with the following criteria (For Write-up)
    - Draw the path on the map.
    - Denote all of the explored squares on the map.
    - Denote the current open/frige list.
    - Report the length of the path found.
    - Report the cost of the path found.

## Compiling/Running
* Compile with (Windows)
```
go build -o ./bin/pathfind ./src/project1/
```

* Run with the following flags to denote whitch algorithms to run
    - breadth
    - lowCost
    - gBestFirst
    - aEuclidean
    - aManhattan
    - all
* Use the `-map` flag with a string denoting the map.txt file to use
* No need to pipe output, program will auto-create and generate it for you
```
./bin/pathfind -map p1Map.txt -[desiredflag]
```
* Or you can simply use the pre-built binary in the same way like so
```
./djProject1 -map p1Map.txt -[desiredflag]
```
* Notes 
    - TODO: the `-all` flag will run and create output for each of the flags before
    - output will be generated in the `tmp/` directory
    - only allowed to move NSEW