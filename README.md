# perspective


## Development

### Best Practices

#### Struct Tagging

All structs should be tagged for json, yaml and xml to support easy marshalling using any of these
formats. This is easily done using the `gomodifytags` tool which can be installed with the following
command.

Install `gomodifytags`
```shell script
go get github.com/fatih/gomodifytags
```

Tag structs
```
gomodifytags -file <filename> -all -add-tags json,yaml,xml -add-options json=omitempty -w
```