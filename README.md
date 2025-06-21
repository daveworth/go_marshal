# JSON Unmarshalling Drama

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
