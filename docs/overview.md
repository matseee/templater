# documentation
The application follows the server-client architecture, which, however, remains largely hidden from the user.
The server part of the application is written in golang and the frontend is written in angular.
The angular application is delivered via the golang application.
This in turn communicates via a Restful api with the "server" part of the application.

## client
#### development
```
cd angular-app/
npm install
npm run start
```

## server
#### development
```
// first build the client and after that:
go run .
```

### module overview
Templater is based on the following modules:
| module      | description                                                       |
| ----------- | ----------------------------------------------------------------- |
| templater   |                                                                   |
| keylistener | listens to the keyboard and sends events to the templater package |
| gui         | creates systemtray and webview for the gui                        |


