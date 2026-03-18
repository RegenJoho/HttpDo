# HttpDo
A simple HTTP client wrapper around the standard library for golang, because said standard library felt clunky to me.
May contain bugs or other issues. This is just a simple side-project without any priority.
### Examples

Run a simple GET Request, for example a webhook
```
err = HttpDo.New().GET().Url("http://example.com/your-url").Run().Err()
```
