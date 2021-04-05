# Introduction

This file is for students who need a refresher on the relevant concepts on HTTP or how you use Go to make an HTTP server. We recommend you use it as a reference if you get stuck implementing something and ask us questions on Discord if you still have some confusion. This file has been designed so you can skip between each of the paragraphs you already understand so we encourage you to do that since this is rather long. If you already know how everything works, **you don't need to read this file to do the assignment**.

# HTTP Refresher

## Introduction

In order for computers to communicate with each other over a network, they must follow a protocol so all communication between them is understood. The most common protocol for this is the HyperText Transfer Protocol (**HTTP**).

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

As a last point to this section, it is worth noting that the combination of the API endpoint and the HTTP method uniquely defines an action to take on a specific resource. Often programmers call this combination an **HTTP Route** (or just a route). Typically these routes are what get mapped to actual functions in code instead of just the HTTP method or just the API endpoint as you will see in this assignment.

## HTTP Request and Response Syntax

Here is an example HTTP request made to a site called [`jsonplaceholder.typicode.com`](jsonplaceholder.typicode.com):

```
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

Another important header we can see is the `cookie` header which the client sends if they need to send cookies to the server. In this case, we have one cookie with key `__cfduid` and value `daef275dc229d6ef75dcaf29628543c9f1617581200`. This is an example of a cookie that holds session state and you will see more of those when you build Bearchat.

We can also see a header that tells the server we have a JSON in the **body** of our request and that it is 20 bytes long. The JSON object here has only one key (called `"Key"`) with value `"Value"`.

In response, the server may send this message back:

```
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

While we have presented raw HTTP messages here, you will almost never actually write these yourself. As you will see in this assignment (and maybe in the examples below if you read them) that we rely on libraries built into programming languages to do this for us.

# Go HTTP Library Refresher

# Further Reading / Additional Resources

### HTTP
- Mozilla has an excellent overview of how the entire HTTP protocol works. It goes into much more depth than is needed for our assignments, but it has excellent explanations. [Link to that here](https://developer.mozilla.org/en-US/docs/Web/HTTP/Overview).
- For more information on the security aspect of HTTP and cookies, we recommend using the [CS 161](https://cs161.org/assets/notes/web.pdf) notes and lectures.
- For more on HTTP status codes, Mozilla also has an explanation of that. [Link to that here](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status).

### Go
- For more information about what gorilla/mux does, we recommend checking out their `README.md` on [their GitHub](https://www.gorillatoolkit.org/pkg/mux).
