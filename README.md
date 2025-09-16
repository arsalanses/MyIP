# MyIP

A simple, high-performance, and scalable API that returns a client's public IP address. Built with **Go** and the **Fiber** web framework, it's designed to handle extremely high traffic with minimal resource usage.

This service correctly identifies the visitor's real IP address, even when behind a reverse proxy or CDN like Cloudflare.

## Live Version

A public and rate-limit-free version of this service is deployed and maintained by **DynX**.

-   **IPv4 & IPv6:** [`https://myip.dynx.pro`](https://myip.dynx.pro)
-   **IPv4 Only:** [`https://ipv4.myip.dynx.pro`](https://ipv4.myip.dynx.pro)
-   **IPv6 Only:** [`https://ipv6.myip.dynx.pro`](https://ipv6.myip.dynx.pro)

## Features

-   ✅ Returns the visitor's public IP address in plain text.
-   🚀 Blazing fast and lightweight, thanks to Go and Fiber.
-   ⚙️ Optimized for high-concurrency environments with Fiber's `Prefork` feature.
-   🌐 Listens on both IPv4 and IPv6 (`[::]`).
-   🔒 Reliably detects the real client IP behind proxies.

---

## Installation & Usage

### Prerequisites

You need to have **Go** installed on your server.

### Steps

1.  **Clone the repository**
    ```bash
    git clone https://github.com/MrDevAnony/MyIP.git
    cd MyIP
    ```

2.  **Build the application**
    This command compiles the source code into a single executable binary named `MyIP`.
    ```bash
    go mod init MyIP && go mod tidy && go build -o MyIP main.go
    ```

3.  **Run the application**
    The server will start and listen on port `3000`.
    ```bash
    ./myip
    ```

4.  **Test it**
    You can now get your IP by making a request to the server.
    IPv4:
    ```bash
    curl http://(IPv4 or IPv6):3000/
    ```
---

## Running as a Systemd Service

To ensure the application runs continuously in the background, automatically starts at boot, and restarts if it fails, you can set it up as a `systemd` service.

1.  **Create the Service File**
    Open a new service file with a text editor like `nano`:
    ```bash
    sudo nano /etc/systemd/system/MyIP.service
    ```

2.  **Add the Configuration**
    Copy and paste the following content into the file. **Important:** Make sure the `ExecStart` and `WorkingDirectory` paths match the location where you built your application. The example below assumes your project is in `/root/MyIP`.

    ```ini
    [Unit]
    Description=MyIP Service
    After=network.target

    [Service]
    ExecStart=/root/MyIP/MyIP
    WorkingDirectory=/root/MyIP
    Restart=always
    User=root
    Group=root

    [Install]
    WantedBy=multi-user.target
    ```

3.  **Enable and Start the Service**
    Run the following commands to reload the systemd manager, enable your service to start on boot, and start it immediately.

    ```bash
    # Reload the systemd daemon to recognize the new service
    sudo systemctl daemon-reload

    # Enable the service to start automatically on boot
    sudo systemctl enable MyIP.service

    # Start the service right away
    sudo systemctl start MyIP.service
    ```

4.  **Check the Status**
    You can verify that the service is running correctly with the following command:
    ```bash
    sudo systemctl status MyIP.service
    ```

    ## Running with HTTPS (SSL/TLS)

To serve the API over a secure HTTPS connection, you'll need SSL/TLS certificates (e.g., from [Let's Encrypt](https://letsencrypt.org/)) and a small modification to the `main.go` file.

### 1. Obtain SSL Certificates

First, make sure you have your certificate files (`fullchain.pem`) and private key files (`privkey.pem`) on your server.

### 2. Modify `main.go`

You need to switch from `app.Listen` to `app.ListenTLS`.

-   **Find this line** in your `main.go` file:
    ```go
    log.Fatal(app.Listen("[::]:3000"))
    ```

-   **Replace it** with the following, making sure to provide the correct paths to your certificate and private key files. The standard port for HTTPS is `443`.

    ```go
    log.Fatal(app.ListenTLS("[::]:443", "/path/to/your/fullchain.pem", "/path/to/your/privkey.pem"))
    ```

    **Example:**
    Based on the code comments, you would change this:
    ```go
    log.Fatal(app.Listen("[::]:3000"))
    // log.Fatal(app.ListenTLS("[::]:443", "/etc/letsencrypt/live/ipv4.myip.dynx.pro/fullchain.pem", "/etc/letsencrypt/live/ipv4.myip.dynx.pro/privkey.pem"))
    ```
    To this:
    ```go
    // log.Fatal(app.Listen("[::]:3000"))
    log.Fatal(app.ListenTLS("[::]:443", "/etc/letsencrypt/live/ipv4.myip.dynx.pro/fullchain.pem", "/etc/letsencrypt/live/ipv4.myip.dynx.pro/privkey.pem"))
    ```

### 3. Rebuild and Restart

After saving the changes to `main.go`, rebuild your application and restart the service.

```bash
# Rebuild the binary
go build -o myip main.go

# If using systemd, restart the service
sudo systemctl restart myip.service
