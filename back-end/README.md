# exam-system
The goal of this project is to build a web-based system for autolab's instructor building a dynamic exam for different students like GRE exam.

## Front-end
### Tech Stack
*   Language: TypeScript
    *   Pros:
        *   Strong Typed: provide robustness development

*   Framework: React / `create-react-app`
    *   Pros:
        *   Widely-used
        *   Many examples and libraries

*   User Interface: Bootstrap (`React-Bootstrap`)




## Back-end
### Tech Stack
*   Language: golang
    *   Pros:   
        *   Package system 
        *   Interface system
*   Framework:  gin
    *   Pros:
        *   Fast: Routing based on Radix tree, the performance is very powerful.
        *   Support middleware: There are many built-in middleware, such as Logger.
        *   Crash recovery: It can catch program crashes caused by panic, so that Web services can always run.
        *   JSON Validation: You can validate the JSON data format in the request.
        *   Route grouping: Support route grouping (RouteGroup), which can make it easier to organize routes.
        *   Error management mechanism: can collect errors in the program.
*   Database:   MySQL, MongoDB
    *   MySQL: For gin customized structure by gin gorm package
    *   MongoDB: For JSON object storage


### Process
*   Front-end
    *   UI Design
    *   UI Development

*   Back-end    
    | Process                       | Status      |
    | ----------------------------- | ----------- |
    | Framework Structure Design    | Finished    |
    | Authentication                | Integrating |
    | Autolab and Tango Integration | Developing  |
    | Questions Storage             | Developing  |
    | Exam Generator                | Developing  |
