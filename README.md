## 🔐 Project Name: Keytide

> **Status: scaffold.** What is implemented today is an HTTP server, one
> product endpoint, a logging middleware and an auth middleware that fails
> closed pending a real token verifier. gRPC, Kafka, RBAC, OAuth2, token
> revocation and audit trails are described below as intent, not as status —
> none of them are started. Everything under "Core Features" should be read
> as a roadmap.

### Description:

Keytide is a plug-and-play Identity and Access Management (IAM) microservice designed for modern web applications. Built with Go, gRPC, and Kafka, it provides a high-performance, scalable, and secure platform for managing authentication, authorization, and user identity.

Designed to be modular, lightweight, and easily integrable, Keytide offers a clean gRPC API surface for real-time operations and leverages Kafka for asynchronous, lossless event handling such as auditing, session revocation, and user state changes.

Whether you’re building a SaaS platform or need to offload IAM complexity from your monolith, Keytide empowers your stack with enterprise-grade access control in a cloud-native way.

### 🧩 Core Features

-  🔑 Authentication & Authorization (JWT, OAuth2, RBAC)

-  📦 Modular & Configurable — enable or disable features with ease

-  ⚡ gRPC-first API — strongly typed, fast, and scalable

- 🔄 Kafka Integration — event-sourced, lossless, async handling

-  🔐 Secure by Default — mTLS, token revocation, audit trails

-  📊 Observability — Prometheus metrics, structured logs, tracing

-  🛠️ SDK-Ready — Auto-generated gRPC clients for multiple languages