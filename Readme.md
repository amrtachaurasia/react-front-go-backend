## What's wrong with this project ?
This is a basic blogging api, with golang backend, sqlite and react frontend. There are a lot of redundant codes, and bad software practices such as server.go contains code for sqlite as well which is not proper. But yeah I was earning when I wrote it.
![Page1](https://i.imgur.com/cDjJMJA.png)
![Page2](https://i.imgur.com/pH1pv4f.png)
## How to run this project ?
### For backend
1. Install go from [Golang](https://golang.org/doc/install?download=go1.11.linux-amd64.tar.gz), Follow instruction properly.
2. cd react-front-go-backend/backend
3. go run server.go
### For frontend
1. The front-end is an experimental build, so you need not install anything.
2. cd react-front-go-backend/frontend
3. Run index.html on a browser.
## What did I learnt while doing this ?
### For backend in go, common pitfalls and flight-rules:
1. w.WriteHeader(310), it is used to return custom response.
2. w.Header().Set("Content-Type", "application/json"), is used to edit default http headers.
3. if you map your request string of json to golang int or vice versa you will get errors containing this keyword "unmarshal"
4. "_" is used when you don't want to do anything with the variable but the function is returning it. So in golang you have to define variables if function is returning two values but if you are not going to use it then you code will give errors.
5. In struct value of parameters should start with capital or you might get undecipherable errors.
6. Ref[https://github.com/campoy/go-web-workshop]
### For frontend in react, common pitfalls and flight-rules:
1. If you want to change DOM the only way you can add is through return in render nothing else.
2. If you try to add jsx somewhere else, you will fail badly.
```
data1={
    Name: this.state.blogid,
    Description: "like"
  };
console.log(data1)
  fetch('http://localhost:8080/like', {
  method: 'post',
  headers: {
    "Content-Type": "application/json"
  }
  body: JSON.stringify(data1)
})

```
3. This is how you do post method properly.
4. If you want to change anything in DOM you have to add code like 

```
<p class="content">{this.state.content.Content}</p>
```
5. Now whenever any event happens, state changes then this will definitely change.
6. If you want to add conditional rendering you have to do it in render function.

```
fetch('http://localhost:8080/blog/'+name)
  .then(response => response.json())
  .then(json => {
    console.log(json)
    this.setState({content: json})
    this.setState({likebutton: 1})
    this.setState({blogid: name})
    console.log("hi there "+this.state.content.Title)

  })
```
7. this is important as fetch passes the return value to then and is passed as response, and response.json() is passed as json in 
.then()

