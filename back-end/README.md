# exam-system


# Development

## Error handling
*   func in [`middlewares/recovery.go/GinRecovery`](middlewares/recovery.go)
    *   Case 1 (Customized error message with [ErrorResponse](error/error.go))
        *   Status:     ErrorResponse.Status
        *   Message:    ```json
                        {    
                            "error": {
                                ErrorResponse.Scope: ErrorResponse.Message
                            }
                        }
                        ```
        *   Receive
            *   [`error/ErrorResponse`](error/error.go)
                *   ``` go
                        type ErrorResponse struct {
                            Status  int
                            Scope   string  // eg: Authenticaion
                            Message string  // eg: Token error  
                        }
                    ```
        *   Usage
            *   add `panic(error.ErrorResponse{...})` in the code
        *   Example
            *   `panic` api in [test.go](router/test.go)
    *   Case 2 (other internal error)
        *   Would return internal error with no message
        
## Hot reload development
*   run `./bin/air`
*   Description
    *   Reload go server for any chages in the project files
*   Reference
    *   https://github.com/cosmtrek/air


## Swagger usage
*   Example
    *   [autheticaion/Auth function](authentication/authentication.go)
*   Swagger UI
    *   http://localhost:8080/swagger/index.html
*   Reference
    *   https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format

## .env file
*   Reference
    *   https://github.com/joho/godotenv

## Develop reference
*   https://juejin.cn/column/6968662583138238478