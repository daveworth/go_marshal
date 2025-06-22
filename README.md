# JSON Unmarshalling Drama

Toy examples while working through Trail of Bits' blog post about Golang parsers
in marshaling and unmarshaling Go Structs:
https://blog.trailofbits.com/2025/06/17/unexpected-security-footguns-in-gos-parsers/

This also gave me an opportunity to experiment with
[`semgrep`](https://semgrep.dev/), a tool I've been intending to use and
evaluate for quite some time. While the Semgrep Registry already had 2/3 rules
for the examples presented I chose to use the 1/3 not in the Registry (but in
the post) and modify it to match the others for purposes of learning.

## Examples with Semgrep

### Untagged Struct Member

Even untagged struct members can be unmarshalled. To run the `semgrep` check
from Trail of Bits use the following command:

```sh
semgrep --config=untagged_struct_member/semgrep_rules.yaml --error untagged_struct_member/main.go
```

### Incorrect ignore directive in JSON tags

The JSON package treats first positional parameter to struct tag as the name of
the struct member if there are additional tags. As a side effect we can still
unmarshall in badly configured structs.

```sh
semgrep --config=incorrect_ignore_struct_member/semgrep_rules.yaml --error incorrect_ignore_struct_member/main.go
```

### Incorrect `omitempty` directive in JSON Tags

This is almost identical to the Incorrect Ignore Directive in JSON tags bug
above as the first positional parameter is the name the unmarshaller looks for
but instead of `-` as the JSON key name `omitempty` will fire.

```sh
semgrep --config=incorrect_omitempty_struct_member/semgrep_rules.yaml --error incorrect_omitempty_struct_member/main.go
```

## Examples without semgrep

While the use of semgrep was very educational the following examples less
clearly mapped to the tool but were equally ~educational~ shocking.

### Unmarshal "case insensitivity"

This is an incredibly shocking example in a few ways: Not only do [the docs](https://pkg.go.dev/encoding/json#Unmarshal)
*strongly* imply that while field matching is case-insensitive, matching case is
preferred.

> To unmarshal JSON into a struct, Unmarshal matches incoming object keys to the
> keys used by Marshal (either the struct field name or its tag), preferring an
> exact match but also accepting a case-insensitive match.

What's even wilder is how it's not about case-insensitivity, but unicode
code-point "folding" meaning that "actions" and "aKtionſ" are considered the
same!!! This was a major "wat?!" moment.

### Parser Differentials

This example is surprising in other ways but relies on architectural constraints
(and argues for both technical standards *and*, perhaps, a monorepo) in that
inter-, and intra-, ecosystems interoperability requires true standards or we
find differing, equally valid but conflicting, behaviors leading to
inconsistencies with how payloads are interpreted..

### Polyglots

This was a fun comparison of JSON, XML, and YAML parsing of a single payload
finding that parser decisions lead to internally inconsistent unmarshaling of
the same struct member.
