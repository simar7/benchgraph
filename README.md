# benchgraph
Visualization of Golang benchmark output using Google charts

## Prefix - what's special about this fork?

The original `benchgraph` is not maintained anymore -- useful PRs like [#4](https://github.com/codingberg/benchgraph/pull/4) and [#5](https://github.com/codingberg/benchgraph/pull/5) have been sitting there for over a year without being picked. Thus people just forked away without bothering to send PRs back.

This fork tries to gather PRs from all different forks and consolidate them into a central location/tool.

![Network graph](network_graph.png "Network graph")


Notable features available in this fork are:

- Able to support sub benchmarks (by [**@miry**](https://github.com/miry))
- Added go.mod (by [**@simar7**](https://github.com/simar7))
- Able to generate graphs locally (by [**@tkanos**](https://github.com/tkanos))
- Makes the default behavior to generate graph locally, instead of sending to some server and publish them publicly. More explanations/reasons on this:
  * Many people don't want to send their results somewhere else and publish them publicly.
  * Even for those who do, it might take well over several attempts to get a all-satisfying result. Publishing those intermediate test results will be a waste of the efforts and being force to do so doesn't make much sense.
  * The original tool has been out of maintenance for a while, and the publishing server might be gone in any minutes. So getting used to not using it might not be a bad idea.
  * Most importantly, the site said that it is very easy to embed the generated charts elsewhere, but I tried and was unsuccessful. Having a local pure html file makes porting the charts elsewhere much easier.
- Allow the generated html file world-readable.

### Download binaries

- The latest binary executables are available under  
https://bintray.com/antoniosun/bin/benchgraph,   
as the result of the Continuous-Integration process.
- I.e., they are built during every git push, automatically by [travis-ci](https://travis-ci.org/), right from the source code, truly WYSIWYG.
- Pick & choose the binary executable that suits your OS and its architecture. E.g., for Linux, it would most probably be the `benchgraph-linux-amd64` file. If your OS and its architecture is not available in the download list, please let me know and I'll add it.
- You may want to rename it to a shorter name instead, e.g., `benchgraph`, after downloading it. 


### Debian package

Available at https://bintray.com/antoniosun/deb/benchgraph,  
or directly at  https://dl.bintray.com/antoniosun/deb:

```
echo "deb [trusted=yes] https://dl.bintray.com/antoniosun/deb all main" | sudo tee /etc/apt/sources.list.d/antoniosun-debs.list
sudo apt-get update

sudo chmod 644 /etc/apt/sources.list.d/antoniosun-debs.list
apt-cache policy benchgraph

sudo apt-get install -y benchgraph
```



### Install Source

To install the source code instead:

```
go get github.com/AntonioSun/benchgraph
```


### Author(s) & Contributor(s)

- [Antonio SUN](https://github.com/AntonioSun)
- [**@codingberg**](https://github.com/codingberg), the original author, and [**@miry**](https://github.com/miry), [**@tkanos**](https://github.com/tkanos), [**@simar7**](https://github.com/simar7) as listed above

_Powered by_ [**WireFrame**](https://github.com/go-easygen/wireframe),  [![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-Y.svg)](http://godoc.org/github.com/go-easygen/wireframe), the _one-stop wire-framing solution_ for Go cli based projects, from start to deploy.

All patches welcome. 

## Introduction
In Golang we can analyze algorithm efficiency by writing benchmark functions and looking at execution time in ns/op. This task might become significantly hindered by increasing number of benchmark tests. One way to handle this is to visualize multiple benchmark results and track the function curve on a graph. The `benchgraph` reads benchmark output lines, prepare data for the graph, and upload data to remote server, which enables online view and html embedding. Graph turns out to be very handy in case of many algorithms that are tested against many arguments, especially if you are studing internal algorithm design.

## Installation

```bash
go get -v github.com/AntonioSun/benchgraph
```

## Naming convention
In order for `benchgraph` to work a coder is required to follow the **naming convention** when coding benchmark functions:
```go
// Naming convention
func Benchmark[Function_name]_[Function_argument](b *testing.B){
...
}
```
For example, if we take one line from the benchmark output,
```bash
BenchmarkF1_F-4       	30000000	        53.7 ns/op
```
it will be parsed and plotted on graph as function `F1(F)=53.7`, taking `F` as an argument and `53.7` as function result. 
In short, X-axis shows function arguments, while Y-axis shows function execution time in ns/op.

## Usage
The output of benchmark is piped through `benchgraph`:

```bash
go test -bench .|benchgraph -title="Graph: F(x) in ns/op"
testing: warning: no tests to run
? PASS
√ BenchmarkF1_F-4       	30000000	        53.7 ns/op
√ BenchmarkF1_FF-4      	20000000	        62.9 ns/op
√ BenchmarkF1_FFF-4     	20000000	        70.0 ns/op
√ BenchmarkF1_FFFF-4    	20000000	        80.3 ns/op
√ BenchmarkF1_FFFFF-4   	20000000	        90.8 ns/op
√ BenchmarkF1_FFFFFF-4  	20000000	        99.5 ns/op
...
Waiting for server response ...
=========================================

http://benchgraph.codingberg.com/1

=========================================
```

In front of every line `benchgraph` places indicator whether line is parsed correctly, or not.
When you see red marks `-` or `?`, it means, either you do not follow the **naming convention** from above, or the line doesn't contain benchmark test at all. At the end, `benchgraph` returns URL to the graph. From there, follow instructions how to embed graph into custom HTML page. Also, you can just share the graph link.

## Help

```bash
benchgraph -help
Usage of benchgraph:
  -apiurl string
    	url to server api (default "http://benchgraph.codingberg.com")
  -oba value
    	comma-separated list of benchmark arguments (default [])
  -obn value
    	comma-separated list of benchmark names (default [])
  -publish
        publish the response publicly
  -title string
    	title of a graph (default "Graph: Benchmark results in ns/op")
```

You can filter out which functions and against which arguments you want to display on graph by passing `-obn` and `-oba` arguments. This can be very handy in case when performing many benchmark tests.

```bash
go test -bench .|benchgraph -title="Graph1: Benchmark F(x) in ns/op" -obn="F2,F3,F4" -oba="F,FF,FFF,FFFF,FFFFF,FFFFFF,FFFFFFF,FFFFFFFF"
```

## Hints on productivity

You can first save benchmark output and then use it later for drawing graphs. This is very handy if your benchmark tests take some time to complete.

```bash
go test -bench . > out

cat out|benchgraph -title="Graph1: Benchmark F(x) in ns/op" -obn="F2,F3,F4" -oba="F,FF,FFF,FFFF,FFFFF,FFFFFF,FFFFFFF,FFFFFFFF"
cat out|benchgraph -title="Graph2: Benchmark F(x) in ns/op" -obn="F2,F3,F4" -oba="0F,F0,F00,F000,F0000,F00000,F000000,F0000000"
```

To have all in local, you can also use the **-local** option :

```bash
go test -bench . > out

cat out|benchgraph -title="Graph1: Benchmark F(x) in ns/op" -obn="F2,F3,F4" -oba="F,FF,FFF,FFFF,FFFFF,FFFFFF,FFFFFFF,FFFFFFFF" -local
cat out|benchgraph -title="Graph2: Benchmark F(x) in ns/op" -obn="F2,F3,F4" -oba="0F,F0,F00,F000,F0000,F00000,F000000,F0000000" -local
```

It will generates on your temp folder, a local html file.

## Online Demo

Here we analyze efficiency of different algorithms for computing parity of uint64 numbers:

http://codingberg.com/golang/interview/compute_parity_of_64_bit_unsigned_integer

There are two graphs embedded into page behind above link:

http://benchgraph.codingberg.com/1

http://benchgraph.codingberg.com/2

*Both above links can be also shared without emebeding into HTML page.*

