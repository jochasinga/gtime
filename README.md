gtime
=====
Deal with Go time stress-free.

Description
-----------
Wished you could deal with time in Go like you did in Python or Ruby, like

```go

// How much is 0.5 second?
s := gtime.Second(0.5)

```

instead of

```go

// How much is 0.5 second?
var s = time.Duration(500) * time.Millisecond

```

*gtime* is the tool set for doing just that.

Install
-------
```bash

go get github.com/jochasinga/gtime

```

Or better, use a [package manager](https://github.com/golang/go/wiki/PackageManagementTools) like [Godep](https://github.com/tools/godep) or [Glide](https://glide.sh/).

Examples
--------

```go

package main

import "github.com/jochasinga/gtime"        

func main() {
        h := gtime.Hour(24)
        fmt.Printf("%.f hours >> %.f hours\n", h, h.Hours())
	fmt.Printf("%.f hours >> %.f minutes\n", h.Minutes())
	fmt.Printf("%.f hours >> %d nanoseconds\n", h.Nanoseconds())
	fmt.Printf("%.f hours >> %.f seconds\n", h.Seconds())
	fmt.Printf("%.f hours >> %q as string\n", h.String())
	fmt.Printf("%.f hours >> %v as time.Duration\n", h.Duration())

}

```

The above prints:

```bash

24 hours >> 24 hours
24 hours >> 1440 minutes
24 hours >> 86400000000000 nanoseconds
24 hours >> 86400 seconds
24 hours >> "24h0m0s" in string
24 hours >> 24h0m0s in time.Duration

```

Want to find out how much 30 minutes is?

```go

m := gtime.Minute(30.0)
fmt.Println(m.Hours())   // 0.500000
fmt.Println(m.Seconds()) // 1800.000000

```

Or `Second`:

```go

s := gtime.Second(0.3)
fmt.Println(s.Nanoseconds()) // 300000000
fmt.Println(s.String())      // "300ms"

```

gtime.Duration
--------------

Both `gtime` types and `time.Duration` implement `Duration`,
therefore they can be used hand in hand.

```go

var s1 gtime.Second = 0.5
var s2 = time.Duration(1) * time.Second

add := func(t1, t2 gtime.Duration) (s float64){
        s = t1.Seconds() + t2.Seconds()
        return
}

sum := add(s1, s2)  // 1.500000

```

Helper functions
----------------

### Ftos(t float64) time.Duration
Converts a floating point number to a duration in second.

```go

s := gtime.Ftos(0.5)

```

### Stof(s time.Duration) float64
Converts a duration in second to floating point number.

```go

s := time.Duration(32.0) * time.Second
t := gtime.Stof(s)

```
