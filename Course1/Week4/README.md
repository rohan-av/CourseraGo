# Week 4

## RFCs: Requests for Comments

**RFCs**: Definitions of Internet protocols and formats
E.g.

- HTML
- URI (Uniform Resource Identifier)
- HTTP
- JSON

Golang has protocol pakcages to decode and encode these protocols, e.g.

- `"net/http"` : `http.Get(www.uci.edu)`
- `"net"`: TCP/IP and socket programming (`net.Dial("tcp","uci.edu:80")`)

JSON

- JavaScript Object Notation, RFC 7159
- Format to represent structured information
- Attribute-value pairs
  - struct/map

## JSON

Properties

- All Unicode
- Human-readable
- Fairly compact representation
- Types can be combined recursively

JSON Marshalling: generating JSON representation from an object

```
p1 := Person(name: "Joe", addr: "A St.", phone: "123")

barr, err := json.Marshal(p1)
```

`Marshal()` returns JSON representation as `[]byte`.

```
var p2 Person

err := json.Unmarshal(barr, &p2)
```

`Unmarshal()` converts a JSON `[]byte` into a Go object.

## File Access (`ioutil`)

- Linear access, not random access
  - Mechanical delay
- Basic operations
  - Open
  - Read
  - Write
  - Close
  - Seek (move read/write head)

**`ioutil` File Read**

```
dat, e := ioutil.ReadFile("test.txt")
```

- `dat` is `[]byte` filled with contents of entire file
- Explicit open/close not needed
- Large files cause a problem

**`ioutil` File Write**

```
dat = "Hello, world"

err := ioutil.WriteFile("outfile.txt", dat, 0777)
```

- Creates a file
- Third argument is permission (Unix-style permission bytes)
  - `0777`: universal permission for read/write

## File Access (`os`)

`os.Open()`: opens a file (returns a file descriptor)
`os.Close()`: closes a file
`Read()`: reads from a file into a `[]byte`

- Controls the amount read (fills the `[]byte`)
- returns no. of bytes read

```
f, err := os.Open("dt.txt")
barr := make([]byte, 10)
nb, err := f.Read(barr)
f.close()
```

`Write()`: writes a `[]byte`
`WriteString()`: writes a string (any Unicode sequence)

```
f, err := os.Create("outfile.txt")

barr := []byte{1, 2, 3}
nb, err := f.Write(barr)
nb, err := f.WriteString("Hi")
```
