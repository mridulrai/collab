# user

Backend service for managing users

## Overview
**User** is a backend micro service for managing following information.
* Manage users - user CRUD operations

## Architecture
TODO

## Local Development
To run this service on your local machine make sure that you have following items installed by using your terminal :
* Docker engine ([Read more](https://www.docker.com/))
* **make** tool ([Read more](https://man7.org/linux/man-pages/man1/make.1.html))
* You will need to install `traefik` and `dnsmasq`. Instructions can be found at [local-traefik](https://gitlab.com/cilalabs/global/common-utils/local-traefik) user.

## Service Account Key

In order to authenticate with Firebase Auth for local development, you will need a `serviceAccountKey.json` file.

You can follow these steps to get your service account key:
* Follow the steps mentioned in this doc ([Click here](https://cloud.google.com/docs/authentication/production))
* Copy paste the serviceAccountKey.json in user root directory.

## Run Application

Create a new `.env` file on root directory of the user and copy the content of `.sample.env` in it. Make modifications in environment variables as per your need in the new file.

Once you have checked items in the above list then open your terminal and change to root of user directory on your local machine. There are 2 ways to run the app ie.,

**Option 1**
* Run `docker-compose up`

**Option 2**
* Type the following command to run the service using Makefile.
>make run

Inorder to enter inside docker container type the following command at the root of user directory on your local machine.
>make inside

Inorder to stop containers in docker for the service type the following command at the root of user directory on your local machine.
>make stop

To clean data in database use following command on your terminal
>make clean

## user Management
* [Slack](https://cilalabs.slack.com/)
* [Issues](https://gitlab.com/groups/cilalabs/my-sales-champion/service/-/issues)

## Maintainers
* Mani Soundararajan
* Abhay Joshi
* Saraswati Kalla
