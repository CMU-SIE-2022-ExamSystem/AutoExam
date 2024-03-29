# exam-system back-end

## Project Structure
| Folder       | Descriptions                            |
|--------------|-----------------------------------------|
| autolab      | functions for autolab communication     |
| config       | configuration of back-end server        |
| controller   | back-end API handler                    |
| default      | default files of yaml                   |
| global       | global variables of back-end server     |
| initialize   | initialization functions                |
| middleswares | middlewares packages of back-end server |
| models       | models structure packages               |
| response     | customized response structure packages  |
| router       | mapping url to controller               |
| test         | development router                      |
| utils        | utilities functions package             |


## Run server
*   `./run.sh`
    *   This script would install go swagger and air package. Also, it would run build swagger index automatically and run the back-end server

## Development

### Error handling
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
        
### Hot reload development
*   run `./bin/air`
*   Description
    *   Reload go server for any chages in the project files
*   Reference
    *   https://github.com/cosmtrek/air


### Swagger usage
*   Example
    *   [autheticaion/Auth function](authentication/authentication.go)
*   Swagger UI
    *   http://localhost:8080/swagger/index.html
*   Reference
    *   https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format

### .yaml file
*   Usage
    *   Use viper to process yaml configuration files and integrate into gin framework
*   Reference
    *   https://github.com/spf13/viper

### Develop reference
*   https://juejin.cn/column/6968662583138238478
