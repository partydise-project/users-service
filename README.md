# Users-service Microservices

![](https://miro.medium.com/v2/resize:fit:720/format:webp/1*e_EhbsyNw6tJbYjwwbjVVg.png)
Users-service imagen by <a HREF="https://betterprogramming.pub/how-are-you-structuring-your-go-microservices-a355d6293932">Better Programming.</a>.

This repository houses the source code for a microservice dedicated to managing all user-related processes within the Partydise project. The microservice utilizes a relational database for efficient storage and retrieval of information. Furthermore, a security layer has been implemented using Auth0, ensuring a secure and authenticated handling of interactions with users. Dive into the code here to explore the implementation details and functionalities of this crucial component in the Partydise project.

# Getting Started
### Prerequisites
- Go 1.20 (should still be backwards compatible with earlier versions)

# Used Tools
1. <a HREF="https://github.com/gin-gonic/gin">Gin as web framework  </a>
2. <a HREF="https://gorm.io/">GORM</a>
3. <a HREF="https://github.com/auth0/go-jwt-middleware/">Auth0 go-jwt-middleware</a>