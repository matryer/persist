# persist

Persist loads and saves Go objects to files.

## Usage

First get it:

```
go get github.com/matryer/persist
```

Then save objects like this:

```
var conf Config
if err := persist.Save("./project.conf", &conf); err != nil {
	log.Fatalln("failed to save config:", err)
}
```

And load them like this:


```
var conf Config
if err := persist.Load("./project.conf", &conf); err != nil {
	log.Fatalln("failed to load config:", err)
}
```