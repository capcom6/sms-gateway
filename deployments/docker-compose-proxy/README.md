# Quickstart Guide for Docker Compose Single File Deployment

## Prerequisites

Before proceeding with the quickstart guide, make sure you have the following:

- Docker Compose version 2.23.1 or later.
- If you are using Portainer, version 2.20.0 or later is required.
- Your server must have a public IP address.
- A domain name that points to your server's public IP address.

## Quickstart Instructions

Follow these steps to get your service up and running:

### Step 1: Environment Setup

Create a `.env` file in the root directory of your project with the following content, and be sure to replace the `[placeholder]` values with your actual data:

```env
MYSQL_ROOT_PASSWORD=[your_root_password]
MYSQL_PASSWORD=[your_mysql_password]
DOMAIN=sms.example.com
ACME_EMAIL=email@example.com
GATEWAY_PRIVATE_TOKEN=[your_private_token]
```

Make sure that your domain (`DOMAIN`) is properly set up to resolve to your server's public IP address.

### Step 2: Starting the Services

Run the following command to start all services defined in your Docker Compose file:

```sh
docker-compose up -d
```

This command will initiate the download of necessary Docker images, create containers, and start the services in detached mode, allowing them to run in the background.

### Step 3: Accessing the Application

Open a web browser and navigate to the following address:

```
https://sms.example.com
```

Replace `sms.example.com` with the domain you've configured. The application should now be accessible from your browser.
