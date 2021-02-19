# Simple Auth
Simple REST API + bare bones frontend client that allows you to register, log in and request some auth-guarded data.
This is just a practice project.

## Reasoning
This project was created for two reasons:
1. To just practice JWT auth implementation with Go. I don't think it's a generally good idea to implement such crucial piece of app on your own, there are libraries that can do a lot of heavy lifiting for you and will do a much better job. However I think it can be beneficial to understand, in principle, how you authentication system works. 

2. To check the idea of storing a JWT in a cookie set by server. Whenever I had to implement the login / registration feature as a frontend dev, the backend would always return JWT in a login response body, and expect the JWT to be in the `Authorization` header for auth-guarded routes. The problem with this approach is that it's the frontend that is responsible for storing the JWT somehow. This means that:   
      * Any other javascript on your site has access to it. And, let's be real, you are not able to check the source code for every google tag manager and facebook pixel that your client wanted you to include on the page. Same goes for every `left-pad.js` that you have in your `node_modules`   
      * You have to implement the logic around storing / removing the token, and initializing all the services that need it. While it's not terribly complicated, it does add some overhead.   

      Storing JWT in a cookie set by server seems to solve both of those problems. By marking the cookie as `httpOnly` you prevent any JS on the client side from accessing it. And because it's a cookie, the browser does the whole job of storing / removing the cookie.

## Potential problems with this solution
* Some security problems around cookies that I'm not aware of
* Frontend does not have access to JWT payload. This can be solved by splitting JWT into two cookies - one (not `httpOnly`) for header and payload, and one (`httpOnly`) for signature
* If you have SSR with your SPA, your server might need to make some initial requests to the backend, and cookies are only included in server <-> browser communication.
