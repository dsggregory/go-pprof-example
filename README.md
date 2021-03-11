# Golang Process Profile (pprof)
This example shows how to integrate pprof into your Go web app.

Go pprof is made available from an internal API server. You may interrogate performance profiles of a running instance using Go pprof and the internal API server. For further reference on pprof, see https://golang.org/pkg/net/http/pprof/.

In order to use the Go pprof tools, you must first create a k8s port forward on the API port:
```bash
kubectl port-forward lrxd-lrxd-worker-0 8080:8080
```
You may then browse to `http://localhost:8080/debug/pprof/` to see what services are available. An example of using this to interrogate memory of the running process would be to then run the following which provides an interactive shell to generate memory reports, etc.
```bash
go tool pprof http://localhost:8080/debug/pprof/heap
```
From the resulting `pprof` shell, try the following subcommands:
* web - creates a graphviz (required) chart of functions and memory used

You may grab periodic samples to compare later by running:
```bash
#!/bin/bash

k8spod=$1

i=3
until [ $i -gt 50 ]
do
  ((i=i+1))
  curl  http://localhost:8081/debug/pprof/heap > heap.pprof.$i
  if [ $? -ne 0 -a -n "$k8spod" ]; then
    kubectl port-forward $k8spod 8081:80 &
    sleep 5
    curl  http://localhost:8081/debug/pprof/heap > heap.pprof.$i
  fi
  sleep 60
done
```

Then get top10 of all the samples:
```bash
#!/bin/bash

for f in `ls -tr heap*`; do
	ls -l $f
	go tool pprof -top $f|head -16
	echo
done > top10.out
```
