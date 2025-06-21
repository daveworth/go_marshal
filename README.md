# JSON Unmarshalling Drama

Toy examples while working through Trail of Bits' blog post about Golang parsers
in marshalling and unmarshalling Go Structs:
https://blog.trailofbits.com/2025/06/17/unexpected-security-footguns-in-gos-parsers/

## Untagged Struct Member

Even untagged struct members can be unmarshalled. To run the `semgrep` check
from Trail of Bits use the following command:

```sh
semgrep --config=untagged_struct_member/semgrep_rules.yaml --error untagged_struct_member/main.go
```

## Incorrect ignore directive in JSON tags

The JSON package treats first positional parameter to struct tag as the name of
the struct member if there are additional tags. As a side effect we can still
unmarshall in badly configured structs.

```sh
semgrep --config=incorrect_ignore_struct_member/semgrep_rules.yaml --error incorrect_ignore_struct_member/main.go
```

## Incorrect `omitempty` directive in JSON Tags

This is almost identical to the Incorrect Ignore Directive in JSON tags bug
above as the first positional parameter is the name the unmarshaller looks for
but instead of `-` as the JSON key name `omitempty` will fire.

```sh
semgrep --config=incorrect_omitempty_struct_member/semgrep_rules.yaml --error incorrect_omitempty_struct_member/main.go
```
