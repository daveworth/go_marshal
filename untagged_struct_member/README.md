# JSON Unmarshalling Drama

## Untagged Struct Member

Even untagged struct members can be unmarshalled. To run the `semgrep` check
from Trail of Bits use the following command:

```sh
semgrep --config=./semgrep_rules.yaml --error main.go
```
