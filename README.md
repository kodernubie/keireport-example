# KeiReport Example

To use Keireport

1. install keireport package

```
go get github.com/kodernubie/keireport
```

2. Import package

```
import "github.com/kodernubie/keireport"
```

3. Load tempalate from file / string

```
rpt, err := keireport.LoadFromFile("simple.krpt")
```

4. Generate to your intended format

```
rpt.GenToFile("simple.pdf")
```