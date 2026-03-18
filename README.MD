# HttpDo
A simple HTTP client wrapper around the standard library for golang, because said standard library felt clunky to me.
May contain bugs or other issues. This is just a simple side-project without any priority.
### Install
```
go get github.com/RegenJoho/HttpDo
```
and import with
```golang
import "github.com/RegenJoho/HttpDo"
```
### Examples

##### Run a simple GET Request, for example a webhook:
```golang
err = HttpDo.New().GET().Url("http://example.com/your-url").Run().Err()
```
##### POST request with json body:
```golang
err = HttpDo.New().POST().MarshalJSON(yourStruct).Run().Err()
```
Please note that this also works with xml, in that case use MarshalXML instead.
Behind the scenes the encoding/json and encoding/xml is used, so keep in mind that your structs need to have the proper tags.

##### Receive data from a request:
```golang
err := request.GET().Url("http://example.com/your-url").Run().Err()
if err != nil {
	panic(err)
}
data := YourStruct{}
err = request.UnmarshalJSONResponse(&data)
```
As above, the same is possible with xml using UnmarshalXMLResponse.

##### Custom HTTP request method
```golang
err = HttpDo.New().Method("your-method").Url("http://example.com/your-url").Run().Err()
```
The standard HTTP request headers are also provided on their own, but using this anything is possible.
 ### Limitations
- If one part errors out, everything will error. This is simply a limitation of only checking for errors at the end.
- If you forget an essential part, like for example the URL, it will error out at runtime, but your compiler or editor will not scream at you.
