## startup

``` bash
cd pathto/confdemo/src/configdemo
export GOPATH=pathto/confdemo
go build -o confdemo configdemo/main.go
./confdemo -c ../conf/configdemo.yaml
```