# goweb-structure
goweb-structure is a predefined directory structure for building web applications using the Golang Gin web framework.
This project aims to provide a standardized and organized approach to structuring your Golang web projects, making it easier to maintain and scale your applications.

## Directory Structure
The directory structure follows a modular approach, separating different components of the application into their respective directories.
Here is an overview of the predefined structure:

```
.
├── cmd
│   └── main.go
├── config
│   └── config.go
├── controllers
│   └── ...
├── middlewares
│   └── ...
├── models
│   └── ...
├── services
│   └── ...
├── web
│   ├── static
│   │   ├── css
│   │   │   └── ...
│   │   └── js
│   │       └── ...
│   └── templates
│       └── ...
└── utils
    └── ...
```

### cmd
The cmd directory contains the entry point of the application, main.go.
This is where you initialize the Gin router and any other application-specific configurations.

### config
The config directory is used to store configuration files or modules related to your application's settings.
It typically includes database configurations, environment variables, and other application-specific settings.

### controllers
The controllers directory houses the handlers for different routes in your application.
Each controller file represents a logical grouping of related routes and their corresponding handlers.

### middlewares
The middlewares directory contains custom middleware functions that can be applied to specific routes or groups of routes.
These middleware functions can be used for authentication, logging, error handling, etc.

### models
The models directory is used to define your application's data models or structures.
These models represent the entities in your application and provide methods for interacting with the underlying data storage.

### services
The services directory houses the business logic of your application.
It contains the implementation of various services that your application provides, such as user management, email sending, file handling, etc.

### web
##### static
The static directory is used to store static files like CSS, JavaScript, images, etc., that are served by your application.

##### templates
The templates directory contains the HTML templates used for rendering dynamic content in your application.
These templates can be rendered using a templating engine like html/template or any other engine of your choice.

## Getting Started
To use this predefined directory structure for your Golang Gin web project, follow these steps:
1. Clone this repository or copy the contents of the goweb-structure directory to your project's root directory.
2. Update the cmd/main.go file with your application-specific configurations and routes.
3. Customize and add your controllers, middlewares, models, repositories, services, static files, templates, and utility functions as per your application's requirements.
4. Change module name as per your choice.
5. Run go program:
```bash
go run cmd/main.go
```

Make sure to refer to the official documentation of the Golang Gin web framework (https://github.com/gin-gonic/gin) for more information on how to use Gin and its features.

## Contributing
If you have any suggestions or improvements for this predefined directory structure, feel free to open an issue or submit a pull request. Your contributions are highly appreciated!

## License
This project is licensed under the [MIT License](LICENSE).
