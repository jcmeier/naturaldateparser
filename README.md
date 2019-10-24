# Natural date parser written in Go

This package helps to parse dates which are written in a natural language like for example:

- 5 days ago
- 1 week ago
- 10 minutes ago
- 5 seconds ago

## Examples
The API methods return an time object:

```go
    p := CreateNaturalDateParser()
    res, err := p.Parse("10 hours ago")
```