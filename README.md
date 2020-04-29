# Go-Mysql-API

## Install Docker on Ubuntu Using Default Repositories
#### Step 1: Update Software Repositories
``` sudo apt-get update ```
#### Step 2: Uninstall Old Versions of Docker
``` sudo apt-get remove docker docker-engine docker.io ```
#### Step 3: Install Docker
``` sudo apt install docker.io ```
#### Step 4: Start and Automate Docker
``` sudo systemctl start docker ```
``` sudo systemctl enable docker ```

## Start the application
#### Step 1: Clone the Project Open a terminal and run the following commands
```
git clone https://github.com/ShubhamNatekar/Dockerized-GO-Mysql-API.git
cd Dockerized-GO-Mysql-API
```
#### Step 2:
```
docker build --tag myapp .
```
```
docker run -it myapp 
```
