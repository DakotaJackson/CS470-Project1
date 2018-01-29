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

* Run with the following flags to denote whitch paths to run
    - breadth
    - lowCost
    - gBestFirst
    - aEuclidean
    - aOctile
    - all
* No need to pipe output, program will auto-create and generate it for you
```
./bin/pathfind -[desiredflag]
```
* Notes 
    - the `-all` flag will run and create output for each of the flags before
    - output will be generated in the `tmp/` directory
    - only allowed to move NSEW