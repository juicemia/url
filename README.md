# url

_A little utility for manipulating URLs_

For those wondering what URL-encode means, try playing around with [this online
converter](http://string-functions.com/urlencode.aspx).

## Install

```bash
go get github.com/juicemia/url
go install github.com/juicemia/url
```

## Usage

From command:

```bash
Usage:
  url [command]

  Available Commands:
    encode      URL-encode an input string
      help        Help about any command

      Flags:
        -h, --help   help for url
```

## Example

When running this command:

```bash
url encode 'I like writing README examples & encoding strings'
```

`url` will output `I+like+writing+README+examples+%26+encoding+strings`

fax:
