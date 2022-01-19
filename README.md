# hil8
hilight lines matching patterns in a stream

I needed to highlight lines from the output of a program,
had not idea how to do that and no time to investigate more than a few minutes,
so here comes `hil8` which highlights lines matching a pattern in a stream.

```
# highlights lines containing the string foobar in file
$ hil8 foobar < file
```

```
# highlights lines NOT containing the string foobar in file
$ hil8 -reverse foobar < file
```

```
# highlights lines matching the regular expression ^foobar in file
$ hil8 -regexp '^foobar' < file
```

```
# highlights in red lines matching the regular expression ^foobar in file
$ hil8 -color red -regexp '^foobar' < file
```

```
# highlights in red lines matching the regular expressions ^foobar or barbaz$ in file
$ hil8 -color red -regexp '^foobar' 'barbaz$' < file
```
