### Compiling and Running ###
- Install go: http://golang.org/doc/install
- git clone https://github.com/jadekler/git-misc.git
- cd git-misc/spiral
- go build main.go
- ./main \<some number\>

### Solution Afterthoughts ###
This solution runs with O(n) space and time complexity, although there are admittedly a fair amount of operations per step. If this had space / time constraints, I would probably try to formulize the solution based on coordinates relative to the 0 position, or look into using a recursive solution (once that builds inner spirals as it goes - although I'm not sure how I'd do that without once again storing the inner data as we go).
