# Setup

For this assignment you'll need a bit of software. We've listed everything you need below.

### git

In order to work locally on the project, you will need to clone your repository onto your machine using git. If you do not have it already, you can download it [here](https://git-scm.com/downloads)

### Go

You will be writing this server in Go. You can download the latest version of Go [here](https://golang.org/dl/).

### Text Editors

We recommend also that you use a very good text editor for this assignment. Here are several options if you're unsure of which to use.

 - [VS Code](https://code.visualstudio.com/download)
 - [Sublime](https://www.sublimetext.com/3)
 - [Atom](https://atom.io/)
 - [Notepad++](https://notepad-plus-plus.org/downloads/)
 - [Vim](https://www.vim.org/download.php)

## Starting

Since this repository acts as a template, you will need to do a bit of extra work before you can clone it to your machine.

1. Click the `Use this template` button at the very top of the repository on GitHub.

2. Name the repository whatever you'd like and give the repository whatever description you'd like. Please also make the repository private.

3. Click `Create repository from template`.

4. From there, you can use `git clone <LINK TO YOUR REPO>` to clone your newly created repository onto your computer and start working!

5. You may also want to set the original repository as a remote in case we make changes to the starter code. You can do that with `git remote add source https://github.com/BearCloud/sp21-assignment-4.git`. If we make changes to the starter code, you can use `git pull source master --allow-unrelated-histories` to integrate the changes with the ones you have.

# Your Task

In this assignment, you will be implementing most of a simple HTTP web server that will perform two primary tasks; echo back parts of an HTTP request and manipulate user credentials. To keep things simple, we will not be worrying about issues of security nor persistent storage of credentials. We will deal with this when we implement Bearchat! All of your work for this assignment will be done in `api/api.go`.

You will need to complete all of the functions in `api/api.go` to complete the assignment. Specifics abotu what each function does is listed in the comments above the functions.

If you need a refresher as to how HTTP works or how to work with HTTP in Go, check out `REFRESHER.md`. You can also ask us any questions on Discord!

# Testing

To test your implementation, we have provided a comprehensive suite of tests in `api/api_test.go`. Simply run `go test -v` in that directory and you should see every test pass if your implementation is correct.

We also encourage you to play around with the server and run it yourself. There are two ways to do this. 

The first way is to open a terminal window and run `go run main.go` in the directory with `main.go`. You should then be able to make requests to the server on port 80. For example, you can try visiting `http://localhost:80/api/getQuery?userID=40` using a web browser or [Insomnia](https://insomnia.rest/products/insomnia) and you should see `40` returned back to you.

The second way is to use [Docker](https://www.docker.com/products/docker-desktop). We have provided a `Dockerfile` you can use to create a container running your server. To run the server in a container, first build the image for the container using the command `docker build -t practice-server .` in the directory with the server files. Then run the server in a container using `docker run -p 80:80 --name practice-server practice-server`. If all went well, you should be able to see the server running if you do `docker ps -a` on another terminal window and you should be able to make requests to it on port 80 as described above. We recommend you use this file as a reference when you make Bearchat!