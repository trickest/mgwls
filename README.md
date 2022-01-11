# mgwls

Merge wordlists

Combines all single words from two wordlist files and concatenates them with each other, with optional delimiter

```
Usage of mgwls:
  -delimiter string
        String delimiter to place between words
  -l string
        Wordlist file (left side)
  -left
        Flag that determines the side of the single word (default true)
  -o string
        Output file (optional)
  -r string
        Wordlist file (right side)
  -w string
        Single word to use for concatenation
```

### Examples

##### left.txt
```
1
2
3
```

##### right.txt
```
a
b
c
```

>Use " " to quote delimiters to avoid command line parsing errors when using reserved characters

```shell script
> go run main.go -l left.txt -r right.txt -delimiter "_"
1_a
1_b
1_c
2_a
2_b
2_c
3_a
3_b
3_c
```

```shell script
>  go run main.go -l left.txt -w foo -left=false
1foo
2foo
3foo
```