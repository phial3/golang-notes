
# Example

## Load

The csv file:

```csv
name, ID, number
name1, 1, 1.2
name2, 2, 2.3
name3, 3, 3.4
```

Your Go code:

```go
type Demo struct {                                // A structure with tags
	Name string  `csv:"name"`
	ID   int     `csv:"ID"`
	Num  float64 `csv:"number"`
}

tab := []Demo{}                                   // Create the slice where to put the content
err  := csvtag.LoadFromPath(
	"file.csv",                                   // Path of the csv file
	&tab,                                         // A pointer to the create slice
	csvtag.CsvOptions{                            // Load your csv with optional options
		Separator: ';',                           // changes the values separator, default to ','
		Header: []string{"name", "ID", "number"}, // specify custom headers
})
```

You can also load the data from an io.Reader with:

```go
csvtag.LoadFromReader(youReader, &tab)
```

Or from a string with:

```go
csvtag.LoadFromString(yourString, &tab)
```

## Dump

Your Go code:

```go
type Demo struct {                         // A structure with tags
	Name string  `csv:"name"`
	ID   int     `csv:"ID"`
	Num  float64 `csv:"number"`
}

tab := []Demo{                             // Create the slice where to put the content
	Demo{
		Name: "some name",
		ID: 1,
		Num: 42.5,
	},
}

err := csvtag.DumpToFile(tab, "csv_file_name.csv")
```

You can also dump the data into an io.Writer with:

```go
err := csvtag.DumpToWriter(tab, yourIOWriter)
```

Or dump to a string with:

```go
str, err := csvtag.DumpToString(tab)
```

The csv file written:

```csv
name,ID,number
some name,1,42.5
```