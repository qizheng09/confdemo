## startup

``` bash
cd pathto/confdemo/src/configdemo
export GOPATH=pathto/confdemo
go get gopkg.in/yaml.v2
go get github.com/golang/glog
go build -o confdemo configdemo/main.go
./confdemo -c ../conf/configdemo.yaml
```