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
-   🚀 Blazing fast and lightweight, thanks to Go and Fibe.
-   ⚙️ Optimized for high-concurrency environments with Fiber's `Prefork` feature.
-   🌐 Listens on both IPv4 and IPv6.
-   🔒 Reliably detects the real client IP behind proxies.

---
## Performance

![MyIP_wrk](https://github.com/user-attachments/assets/14370449-a1e4-4b89-91d6-70e2779d42c3)

The `MyIP` service demonstrates high-performance capabilities for handling concurrent HTTP requests. As shown in the benchmark above, the application is capable of processing over **60,000 requests per second** with low latency (~21 ms average) under heavy load.

This test was performed from an **8-core server** to a **2-core server** (host: `myip.dynx.pro`) using `wrk` with **8 threads and 1000 concurrent connections** over a **15-second duration**. The results highlight the efficiency of the server in returning simple plaintext responses, making it ideal for low-latency, high-throughput scenarios.

## Installation & Usage

### Prerequisites

You need to have **Go** installed on your server.

### Steps

1.  **Clone the repository**
    ```bash
    git clone https://github.com/MrDevAnony/MyIP.git
    cd MyIP
    ```

2.  **Edit .env file**
    This command opens the `.env` configuration file in the nano text editor, allowing you to modify its contents.
    ```bash
    nano .env
    ```

3.  **Build the application**
    This command compiles the source code into a single executable binary named `MyIP`.
    ```bash
    go mod init MyIP && go mod tidy && go build -o MyIP main.go
    ```

4.  **Run the application**
    The server will start and listen on port `3000`.
    ```bash
    ./myip
    ```

5.  **Test it**
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
