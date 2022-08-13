# exam-system
The goal of this project is to build a web-based system for autolab's instructor building a dynamic exam for different students like GRE exam.

# Prerequisite
*   Install Autolab + Tango in one server    
    *  Please follow [Autolab Documentation](https://docs.autolabproject.com/installation/docker-compose/) to install the system
*   Modify Tango's `Dockefile`  file for installing Python3.9 (used for Graders)
    *   Replace original code 
        ```
        RUN apt-get update && apt-get install -y \
        build-essential \
        gcc \
        git \
        make \
        sudo \
        && rm -rf /var/lib/apt/lists/*
        ```
        with 
        ```
        ENV DEBIAN_FRONTEND noninteractive

        RUN apt-get update && apt-get install -y software-properties-common && \
            add-apt-reposiory -y ppa:deadsnakes/ppa
        RUN apt-get update && apt-get install -y \  
        build-essential \
        gcc \
        git \
        make \
        sudo \
        python3.9 \
        python3-pip \
        python3.9-venv \
        && rm -rf /var/lib/apt/lists/*
        ```
    *   Build the image again       
        ```
        docker build -t autograding_image Tango/vmms/
        ```

# Hardware Suggestion
*  4 Core
*  More than 16 GiB RAM
*  50 GB of hard-drive space
*  Linux OS

# Installation Steps
1. Install Docker and Docker-Compose
   *    Follow [Docker Offical Website](https://docs.docker.com/engine/install/) to install Docker and Docker-Compose
2. Clone this repository    
   ```
   git clone https://github.com/CMU-SIE-2022-ExamSystem/exam-system.git
   ```
3. Preprare a specific domain name for tls security
4. Setup back-end server settings
   1. Enter back-end folder     
        ```
        cd back-end
        ```
   2. Copy settings template
        ```
        cp default/settings-dev-default.yaml settings-dev.yaml
        ```
    3. Modify `settings-dev.yaml` file
       *    ip with specific domain name in `Step 3`
       *    autolab's ip with the desired Autolab system's domain name and specify Autolab system's protocol
       *    autolab's OAuth2 four information for authentication
5. Setup Caddy settings
   1. Go back to root folder
   2. Copy .env template
        ```
        cp template.env autoexam.env
        ```
    3. Modify `autoexam.env`
        * Domain with specific domain name in `Step 3`
    4. Modify `Caddfile` to disable SSL
6. Start server
   ```
   ./start.sh
   ```
7.  Stop server
   ```
   ./stop.sh
   ```


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
*   Database:   MySQL, MongoDB, Redis
    *   MySQL: For gin customized structure by gin gorm package
    *   MongoDB: For JSON object storage
    *   Redis: Distrubuted task queue for time-consuming APIs


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
