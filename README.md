# goProject

This is a project for my Emerging Technologies course. It is a Banking app that lets you login in and view your account or create a new account

## Overview
This app has 3 pages:

1. Login: This page allows you to login with an existing account or create a new account
  * For the purposes of demoing it also has a CreateDb button the adds 4 premade accounts with the following usernames and passwords:
  * User: "JoyceL", Pass: "78ad45"
  * User: "YellowSquare", Pass: "DaisyB00k"
  * User: "sheehan87", Pass: "password"
  * User: "LJones", Pass: "050690"
2. Create: This page is used to create a new account. It asks for a name, username and password and saves it to the database.
  * An 8 digit account is randomly generated as well
3. Result: This page is displayed after logging in and displays the details of your account

## Technologies
Go - https://golang.org/

Macaron - https://go-macaron.com/

MongoDB - https://www.mongodb.com/

Html

## Setup (Windows 10 only)

### Setting up Go
1. Download Go from the [main website](https://golang.org/dl/). 
  * Downloading the executeable on Windows will automatically save it to C:\Go in your local directory, unless you specify otherwise. 
2. Open Control panel, go into System and Security, then into System and choose Advance Sytem Settings on the left-sided toolbar (admin rights required). 
3. Select Enviroment Variables. 
4. Under system variables there should be a variable called "GOROOT". This variable should point towards the directory you installed Go in (e.g. C:\Go). 
4. Create a new variable called "GOPATH" and set its value to your Go workspace. 

![alt tag](http://i.imgur.com/CdXfxSl.png "My system variables with GOROOT & GOPATH set")

Download the zip file from this repository and unpack it in your GOPATH so it looks like this .

[GOPATH]/

|------- >bin/

|------- >main/

|------- >package/

|------- >src/

Open up a cmd prompt and enter the following commands: (note: if cmd can't find the Go commands, your GOROOT is set incorrectly).

>go get gopkg.in/macaron.v1

>go get gopkg.in/mgo.v2

macaron is the web framework I used in this project and mgo is a connector to database.

Go is now setup!

###Setting up MongoDB

1. Download MongoDB from [here](https://www.mongodb.com/download-center?jmp=nav#community).

2. Select the option "Windows Server 2008 R2 64-bit and later, with SSL support x64".

3. MongoDB should install (by default) into your program files directory.

4. In enviroment variables edit the path variable

5. Select new and add in the path to the bin folder located in the mongo directory

![alt tag](http://i.imgur.com/TKnnfU8.png "My PATH with mongo's bin folder added")

Test MongoDB by entering "mongod" into a cmd promt. if it works you create your own mongo server which must be on for this app to work.

Mongo is now setup!

## Running the app

1. Startup mongodb by going into a cmd prompt and writing "mongod" into it. A mongo server should now be running.

2. Start main.exe in the main folder of the project, either by clicking on it or through a console.
  * It should display the following message: "[Macaron] listening on 0.0.0.0:4000 (development)
  
3. Open up a web browser and navigate to "localhosthost:4000". The main page should now be displayed.
