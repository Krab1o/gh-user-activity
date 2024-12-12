# gh-user-activity

This is an application to check recent activity of github users from 
[Roadmap site](https://roadmap.sh). You can find this particular project
[here](https://roadmap.sh/projects/github-user-activity).

## Build guide

1. Clone repository, using ssh, https or zip.
2. Go to the root folder of the project.
3. Type `go build` to build the project.
4. Type `./gh-user-activity 

## Usage guide

`gh-user-activity` takes 2 arguments separated by whitespace:

1. Username is mandatory.
2. Amount of last actions (varies from 1 to 100) is optional. Default value
is 30.

Example:
```golang
gh-user-activity krab1o 17
```
will show krab1o`s last 17 events. 