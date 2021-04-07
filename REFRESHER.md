# Introduction

This file is for students who need a refresher on the relevant concepts on HTTP or how you use Go to make an HTTP server. We recommend you use it as a reference if you get stuck implementing something and ask us questions on Discord if you still have some confusion. This file has been designed so you can skip between each of the paragraphs you already understand so we encourage you to do that since this is rather long. ***You don't need to read this file to do the assignment***.

This file is divided into two sections; an HTTP Refresher which talks about the HTTP protocol as a whole and the Go Refresher which has practical examples of how to deal with HTTP in Go. You can read one section without having read the other, so if you're stuck with syntax, but understand how HTTP works, feel free to skip to the Go section.

# HTTP Refresher

## Introduction

In order for computers to communicate with each other over a network, they must follow a protocol so all communication between them is understood. The most common protocol for this is the HyperText Transfer Protocol (**HTTP**). 

Note: If you'd prefer an explanation with more diagrams, the *Further Reading* section includes an excellent overview of HTTP with some.

## Client and Server Protocol

HTTP is a client and server protocol. What this means is that in any communication using HTTP, there is always an entity called the **client** that sends **requests** to another entity called the **server**. In **response** to the requests from a client, the server sends back a message that tells the client if their request was handled by the server or not as well as any additional information the client may need (such as information they may have requested from the server).

## Stateless Property and Cookies

HTTP was originally designed for scientists to share documents with one another and, therefore, did not require the client to save any state once a communication was completed. And so, the original protocol was often called **stateless**.

Nowadays, HTTP has grown to include communications that *do* require that the state of the communication be saved across multiple requests. Think of something like social media that keeps you logged in even after you close your web browser and reopen it or online shopping where your cart is saved even if you visit multiple different pages on the website. How is this state saved across requests?

The answer is **cookies**. Cookies are key/value pairs that are stored on the client to save state across multiple HTTP requests. If a server wishes, they can put cookies in a response for the client to save. Then, when the client returns to the server, the client sends the cookies back to the server in their request so the server can remember who they are.

There is a lot more to say about cookies (such as security concerns), but you won't have to deal with them for this assignment. If you'd like to learn more, take a look at the *Further Reading* section or refer to the webcasts on networking and web security. For an actual example of a cookie, check out the *HTTP Request and Response Syntax* section below.

## URLs, API Endpoints, and HTTP Methods

Requests over the internet can broadly be described as clients wishing to perform some action on some resource. For example, a common request on the internet is to post a message onto a social media platform. The resource in this situation is the message you want to post and the action is to create it and store it onto the server so other people can see it. 

Resources in HTTP can be uniquely identified with a **URL** (also called a URI in some contexts). You are most likely familiar with URLs as they are so central to describing things on the internet, so we won't go over them. However, when writing APIs, it is often very cumbersome to write the entire URL for every resource a user can access, so we often abbreviate it to just the end. For example, if we had the URL, `http://api.messageserver.com/posts`, we may remove the base part of the URL (the `http://api.messageserver.com` part) and just say `/posts` if the base part is already known. The last part of the URL is often called an **API Endpoint**.

While the API Endpoint describes the resource being acted on, we still need a way of specifying what to do to the resource being described by the endpoint. For example, if we had the endpoint `/posts/123` this may indicate that we want to do something with the post with ID 123. But what exactly? Do we want to read the post? Do we want to edit the post? Do we want to delete it? The way clients often express the action they want to do on a resource is through **HTTP Methods**. Broadly speaking, there are four things that could be done to a resource; **C**reate, **R**ead, **U**pdate, or **D**elete it. These operations are often called CRUD operations and HTTP has methods for each of these operations.

| HTTP Method | CRUD Operation |
|     :-:     |      :-:       |
|    `POST`   |     Create     |
|    `GET`    |     Read       |
|    `PUT`    |     Update     |
|   `DELETE`  |     Delete     |

While it isn't strictly enforced for a server to follow this convention (as it is certainly possible for an entire server to use `GET` for reading and `POST` for everything else as an example) the choice of what methods to use for certain actions can affect how users use your application. As an example, think of a web browser. When you type in a URL and press enter, your web browser sends a `GET` request and users often have no way of changing it to something like `POST`. So if you have a webpage that can only be accessed through `PUT` requests, your page will be completely inaccessable to web browser users. There are other practical reasons why choosing the right method for each endpoint matters, but for this assignment you won't have to worry about that. There are also other HTTP methods, but we won't concern ourselves with them.

As a last point to this section, it is worth noting that the combination of the API endpoint and the HTTP method uniquely defines an action to take on a specific resource. Often programmers call this combination an **API Route** (or just a route). Typically these routes are what get mapped to actual functions in code instead of just the HTTP method or just the API endpoint as you will see in this assignment.

## HTTP Request and Response Syntax

Here is an example HTTP request made to a site called [`jsonplaceholder.typicode.com`](jsonplaceholder.typicode.com):

```http
POST /posts HTTP/2
Host: jsonplaceholder.typicode.com
user-agent: insomnia/2021.2.2
cookie: __cfduid=daef275dc229d6ef75dcaf29628543c9f1617581200
content-type: application/json
content-length: 20

{
	"Key" : "Value"
}
```

The first line (also called the start line) specifies the **HTTP Method** (`POST`) and **API Endpoint** (`/posts`) as well as the version of HTTP (version `2`) the client is using for their request. 

Below the start line we have the **headers**. These headers provide metadata about the request that the server can use. For example, the `user-agent` header tells the server that the client is using `Insomnia` to send the request.

Another important header we can see is the `cookie` header which the client sends if they need to send cookies to the server. In this case, we have one cookie with key `__cfduid` and value `daef275dc229d6ef75dcaf29628543c9f1617581200`. This is an example of a cookie that holds session state. You will see more of those when you build Bearchat.

We can also see a header that tells the server we have a JSON in the **body** of our request and that it is 20 bytes long. The JSON object here has only one key (called `"Key"`) with value `"Value"`.

In response, the server may send this message back:

```http
HTTP/2 201 
date: Mon, 05 Apr 2021 00:07:24 GMT
content-type: application/json; charset=utf-8
content-length: 33

{
  "Key": "Value",
  "id": 101
}
```

As you can see, the structure is almost identical as the request, but the start line is different. Notice that the HTTP method is missing (why would the server not need to send this back?) and that there is a number after the HTTP version. That number is called a **Status Code** and, much like return codes in programs, the code gives information back to the client about whether their request was successful or not (and if not why it wasn't). In this case, `201` corresponds to the status code for `Created` which tells us that some kind of resource was *created* by our post request from earlier. Some common status codes you may have seen are `200` for `OK` and `404` for `Not Found`. There is a very neat system for these status codes and if you'd like to see it, see the *Further Reading* section.

While we have presented raw HTTP messages here, you will almost never actually write these yourself. As you will see in this assignment, and maybe in the examples below if you read them, we rely on libraries built into programming languages to do this for us.

# Go Refresher

## Introduction

Dealing with raw HTTP messages is rather cumbersome. So programmers often rely on abstractions given by Go to simplify the process of processing requests. In this section, we will outline some useful parts of Go's HTTP and JSON library. 

Each example given has a link you can use to run it in the Go playground where you can see the output of the example and modify it to try something new. We encourage you to play around with everything to understand it better!

## `net/http`

Recall that HTTP is a client and server protocol, which means clients send requests to servers that respond to them. Clients communicate to servers by sending requests across routes that the server has opened to them. But how are each of things represented in Go? The built-in Go package `net/http` provides most of the tools we will use in our HTTP servers. Let's go through each one by one.

### Requests and `http.Request`
The first thing we'll talk about are requests. In Go, [**http.Request**](https://golang.org/pkg/net/http/#Request) is a struct that represents everything that is present in some request to a server including things like its body and the HTTP method it is using. Since you are the server programmer, you will usually not be creating these yourself (unless you are testing your code) and these will be given to you when a client makes a request. We encourage you to look at the documentation or [this playground](https://play.golang.org/p/eRVx2c_yAFU) for some examples of common operations using `http.Request`. We've also listed some of the ones used in the playground below:

```Go
// Assuming we have an http.Request called `req`

// This line gets the cookie named "Cloud_Cookie" from the request. Gives an error if there's no such cookie.
cookie, err := req.Cookie("Cloud_Cookie")

// This line gets the query parameter "q1" from the request. Gives an empty string if there isn't one.
param := req.FormValue("q1")

// These lines decode a JSON in the request body into some struct (in this case a WordPair struct). Gives an error if the JSON could not be decoded into the struct.
var wp WordPair
err := json.NewDecoder(req.Body).Decode(&wp)
```
### Cookies and `http.Cookie`

Recall that cookies are key/value pairs that allow us to save state across multiple HTTP requests. `net/http` gives us the `http.Cookie` struct that allows us to make our own cookies and also get certain properties about the cookies such as the time they expire.

We have discussed above how to extract a cookie from a request. In the next section, we'll discuss how to set cookies on the client. Once you have the cookie, here are some things you can get from it:

```Go
// Assuming we have an http.Cookie in `cookie`

// Some properties you can access from a cookie
name := cookie.Name
value := cookie.Value
expireTime := cookie.Expire
```

Here's [a playground](https://play.golang.org/p/5-VLNRWm_fg) you can use to play with certain attributes of a cookie.

### Responses and `http.ResponseWriter`

Lastly we'll talk about how to actually write responses back to the client. The abstraction Go provides for creating responses from a server is the `http.ResponseWriter`. There are two fundamental operations you can do with a `ResponseWriter`; write to the response body and write a status code to the response (you can also write the headers of the response using this, but we won't be using this).

Typically in programming languages, we're used to doing some work inside of a function and when we've finished all our computation we use a `return` statement to give back a value or, maybe if something went wrong, we throw an exception and exit the function. If you squint your eyes a little, a `ResponseWriter` encapsulates the same concept. Let's see that in an example.

```Go
// Assume we have an http.ResponseWriter `w`.

// doSomeWork() is some function that gives us something to write back to the response and possibly some error if it failed.
body, err := doSomeWorkThatMightError()

if err != nil {
  // If there was an error, let the client know through the status code (and maybe the body). This is like throwing an exception!
  http.Error(w, "Some error happened!", http.StatusInternalServerError) 
} else {
  // If there wasn't an error, we can just write the body to the response. This is like returning from a function!
  w.WriteHeader(http.StatusCreated)
  w.Write(body)
}
```
As you can see in this example, using a `ResponseWriter` is not too much different from what you're already used to. It's just learning a few new functions. Note in the example above, we used `Write()` to write to the body of the response, but if you have a `string` you want to write to it, it is often more convenient to use [`fmt.Fprintf()`](https://golang.org/pkg/fmt/#Fprintf). If you'd like to play around with this example, here is a [link to a playground](https://play.golang.org/p/EOpKHwHFSVl).

The playground also includes an example of setting a cookie. Here it is for your convenience.

```Go
// Assume we have an http.ResponseWriter `w`

// Sets a cookie with Name="Cool_Cookie" and Value="Chocolate_Chip".
http.SetCookie(w, &http.Cookie{
  Name:  "Cool_Cookie",
  Value: "Chocolate_Chip",
})
```

## `encoding/json`

A very common operation for clients of an HTTP server to do is send key/value pairs to a server. Think of logging into a website, as an example. Logging in requires you to type in your `username` and `password`. The `username` and `password` act as keys and your input acts as the values for each of these keys.

Since this is such a common operation, many HTTP requests use **JSON** to encode the body. JSON (JavaScript Object Notation) is a convenient way of writing key value pairs that can easily be interpreted by most programming languages. We won't go over the syntax of a JSON file since that is well documented already (see the *Further Reading* section). Instead we will talk about what makes it convenient in Go.

One of the main things programmers in Go work with are structs. Structs, at their most basic, are collections of variables each with their own name. If we think of the names in a struct as keys and the values of each of the variables as values, then there is a very simple way to map things in a JSON file to structs. Simply use each key in the JSON as a guide for what the value for each variable in the struct should be. The `encoding/json` package simplifies this process (typically called **unmarshalling** or **decoding**) greatly.

Here's an example of it in action:
```Go
// A Pizza struct that holds some information about a pizza. Notice the names on the right map exactly to the keys in the JSON.
type Pizza struct {
  Radius  float64 `json:"radius"`
  Topping string  `json:"topping"`
}

// A sample JSON that holds our pizza. We'd like to convert this from a JSON to an actual Pizza struct.
JSON := []byte(`{
  "radius" : 12.8,
  "topping" : "pepperoni"
}`)

// The variable that will hold our Pizza.
var p Pizza

// json.Unmarshal() takes our JSON and converts it into a Pizza struct with the correct types. Then it is stored into p.
err := json.Unmarshal(JSON, &p)
```
Notice in the example above, the actual conversion of the JSON to a struct only takes a single line. The other lines were just initializing all the variables we needed to do it. That's way more convenient that reading it manually!

There are different functions to call if you want to put a JSON into a response or extract a JSON from a request. For examples with those, see the sections above.

For the rest of this section, we'll describe some caveats you should be aware of when using this library. All these caveats are also shown in [this playground](https://play.golang.org/p/uFGmt9dFVRf).

If you try to unmarshal a JSON with extra keys, the extra keys are ignored by Go by default. In the example above, if we had added a `sauce` key to the JSON but not the struct, the resulting struct would be exactly the same. You can change this behavior using [`DisallowUnknownFields`](https://golang.org/pkg/encoding/json/#Decoder.DisallowUnknownFields)

Also if you unmarshal a JSON into a struct and the JSON is missing keys (say the JSON in our pizza example is missing the `radius`) the struct is initialized to its zero value for that variable. See the playground for an example of that.

## Routes and `gorilla/mux`

We have talked extensively about how to use Go to manipulate requests and responses, but we have yet to talk about how to actually create routes for clients to access our server through.

Recall that a route is a combination of an HTTP method and an API endpoint. In Go (and most other languages), these routes are what get mapped to functions in the language. Go in particular uses something called a **mux** (short for multiplexer). It's the job of a mux to match API endpoints to functions in code that will take in requests from clients and give back responses.

Go by default comes with its own [mux](https://golang.org/pkg/net/http/#ServeMux), however, for our purposes we'd like a mux that offers more convenience. That's where `gorilla/mux` comes in. `gorilla/mux` is a popular Go package that provides a mux (which they call a `Router`) with a whole ton of convenient features for developers. The most relevant for our assignments is the ability to assign endpoints specific HTTP methods so they become full fledged routes. We recommend looking at the *Further Reading* section if you'd like to see the other benefits.

Let's see an example of how to assign a function to a route.

```Go
// Assume we have a mux.Router called router

// Here is an example of an HTTP handler
func handler(w http.ResponseWriter, r *http.Request) {
  // Some code here that does stuff
}

// Registers the API Endpoint "/test/handler" to the function handler. Also makes it so the endpoint only responds to GET requests.
router.HandleFunc("/test/handler", handler).Methods(http.MethodGet)
```

Some important things to notice here are first off the function signature for `handler`. Any function that takes in a `http.ResponseWriter` and a pointer to a `http.Request` and returns nothing is called a `handler` as it handles the processing of an HTTP request. 

The second parameter to `HandleFunc` can be any arbitrary `handler`. This means that pretty much anything that you could code in Go could be used to handle an HTTP request. This is a major departure from the origins of HTTP where typically endpoints only served static files and it forms the basis of dynamic web applications.

Another thing to note is that if you tried to access `"/test/handler"` with something other than a `GET` request, you will be met with a `405 Method Not Allowed` error in response. This is one of the convenient features `gorilla/mux` gives us.

# Further Reading / Additional Resources

### HTTP
- Mozilla has an excellent overview of how the entire HTTP protocol works. It goes into much more depth than is needed for our assignments, but it has excellent explanations. [Link to that here](https://developer.mozilla.org/en-US/docs/Web/HTTP/Overview).
- For more information on the security aspect of HTTP and cookies, we recommend using the [CS 161](https://cs161.org/assets/notes/web.pdf) notes and lectures.
- For more on HTTP status codes, Mozilla also has an explanation of that. [Link to that here](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status).

### Go
- For more information about what gorilla/mux does, we recommend checking out their `README.md` on [their GitHub](https://www.gorillatoolkit.org/pkg/mux).
- For more information about the syntax of JSON, this is a good [resource](https://medium.com/omarelgabrys-blog/json-in-a-nutshell-7d638dfea7cc). You won't need the stuff on AJAX or Javascript.
