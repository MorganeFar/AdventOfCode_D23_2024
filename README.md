Prerequisites
------------
* Go compiler

Getting started
------------
* <i>main.go </i> : code file
* <i>datas_tests</i> : directory with some examples of sets used in this problem. You can add your own !
----
1. To run the code with the default options :
```sh
go run .
```
2. Test both part with a specified set
```sh 
go run . PATH/TO/INPUT/FILE.txt
```
3. Test only part 1 or part 2 with a specified set
```sh 
go run . PATH/TO/INPUT/FILE.txt 2
```


Parallélisation
PS C:\Users\Morgane FAREZ\Documents\USMB\L3\S6\INFO601_GO\projet\paral> go run .
** Part 1 **
Number of sets of 3 interconnected computers : 7
** Part 2 **
Code : co,de,ka,ta
Time  614.2µs
PS C:\Users\Morgane FAREZ\Documents\USMB\L3\S6\INFO601_GO\projet\paral> go run .
** Part 1 **
Number of sets of 3 interconnected computers : 7
** Part 2 **
Code : co,de,ka,ta
Time  526.9µs

Sans parallélisation
PS C:\Users\Morgane FAREZ\Documents\USMB\L3\S6\INFO601_GO\projet\paral> go run .
** Part 1 **
Number of sets of 3 interconnected computers : 7
** Part 2 **
Code : co,de,ka,ta
Time  621.2µs
PS C:\Users\Morgane FAREZ\Documents\USMB\L3\S6\INFO601_GO\projet\paral> go run .
** Part 1 **
Number of sets of 3 interconnected computers : 7
** Part 2 **
Code : co,de,ka,ta
Time  542µs