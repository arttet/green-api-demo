# Green API Proxy Demo

A proxy server to simplify and manage requests to the Green API, with a simple SvelteKit web UI for interaction.

## 🛠 Management & Development

This project uses **Go** for the backend proxy, **SvelteKit** for the frontend, and **Just** for running development tasks.

## 📜 Just Commands

Available commands for managing the project:

```shell
$ just help
Available recipes:
    help   # Show available recipes and their descriptions

    [Development]
    fmt    # Format code
    lint   # Lint code
    build  # Build app
    run    # Run app in development mode

    [Web]
    bundle # Build the site for production
    dev    # Start development server
```

## ✨ Features

- Simplified Green API integration
- CORS-safe proxy
- Centralized logging
- Configurable timeout
- Clean architecture

## 🚀 Getting Started: Local Development

This project consists of a static Svelte frontend and a Proxy server. To bypass browser security restrictions (CORS and Mixed Content) when running the site from a local build, you must run the backend locally.

### Prerequisites

Make sure you have `just`, `bun`, `go` installed on your system. It is used to automate the startup process for both the frontend and the backend.

### Running the Project

To get the full environment up and running, follow these two steps in separate terminal tabs:

#### Start the Go Proxy server

This command starts the local server (typically on port 8080). This server handles API requests that the browser would otherwise block due to CORS policies.

```sh
just run
```

#### Start the Svelte Web page

This command launches the Svelte development server. Running the frontend locally (on localhost) is necessary to communicate with the local backend without triggering "Mixed Content" (HTTPS to HTTP) blocks.

```sh
just dev
```
