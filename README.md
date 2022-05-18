# opencga-cli 
Build
```
git clone git@github.com/phamidko/opencga-cli
cd opencga-cli
go mod init github.com/phamidko/opencga-cli
go install github.com/spf13/cobra-cli@latest
go get -u github.com/spf13/cobra@v1.4.0
go get -u github.com/mvdan/xurls

make build

```

Usage

```
./bin/opencga get version -s iva.mseqdr.org
Fetching from https://iva.mseqdr.org/iva/conf/config.js
        Cellbase Version: 5.0.1                 Git Commit: eaae3a6f7b407c1eebdb1b4bfede941f4b506b30    Git Branch: release-5.0.x
        OpenCGA Version: 2.2.1-SNAPSHOT         Git Commit: 27cf2ae4bb95596daf839f107dac3d8fb63e6715    Git Branch: release-2.2.x
```

Help

```
./bin/opencga -h
./bin/opencga get version -h

```