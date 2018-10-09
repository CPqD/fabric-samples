# Meat Trade Web Application
This is a simple web application to demonstrate the integration of a REST client with the middleware developed in Go.

## Requirements
* [Node.js](https://nodejs.org/en/download/) version 8.11.x
* NPM version 5.6.0

## Node modules
This Web App includes the following modules.

* [express](https://www.npmjs.com/package/express): one of the most used web frameworks for node. 
* [ejs](https://www.npmjs.com/package/ejs): embedded Javascript templates. Is the view engine of the web site.
* [consign](https://www.npmjs.com/package/consign): successor of express-load. It is used to load scripts separated according to MVC architecture (for example, those in models, controllers, routes, etc).

## Other modules
This Web App also includes the version 4 of [Bootstrap](https://getbootstrap.com/docs/4.0/getting-started/introduction/) and its starter template.

## Usage
1. Navigate to `web-application` folder:
    
2. Start Docker container using:
    ```
    ./startApp.sh
    ```
3. If everything worked, you should be able to reach the page in [http://localhost:8080](http://localhost:8080).

The script [`../runAll.sh`](../runAll.sh) also starts this application.

To stop the application run `./stopApp.sh`.
