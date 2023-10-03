## Sample using auth_request

## Start the application

Start the development environment and nginx with `docker-compose up -d`.

nginx exposes the `8888` port.

Enter the development environment via devcontainer and execute the following command

`go run main.go app1`.

Login and main application server.

The main application displays the HTML of `app2` in an iframe

`go run main.go app2`.

Application server that displays session information.

## URLs

- `app1`
  - `/login` login
  - `/app` application
  - `/auth` Authentication endpoint.
- `app2`
  - `/app` application

## login information

Login user: `user@example.com`  
Password: `password`

## Verification with no login.

`http://localhost:8888/app`

Accessing `/auth_request`.

Use `auth_request` to verify the session against `/auth`.

Since you are not logged in, set `/auth` to redirect to `Location` header

Set the contents of the `Location` header to a cookie in the nginx configuration and render `unauthorized.html

`unauthorized.html` gets the redirect from the cookie and sets the redirect to the parent location.href

Succeeds if redirected to `/login
