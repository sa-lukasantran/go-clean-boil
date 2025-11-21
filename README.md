# Source Code Structure
- Domain (Entities): Pure go structs or interfaces - /entity
- Use case (Application): Application-Specific business logic - /usecase
- Persentation /presentation
    - Controller: HTTP request handlers - /controller
    - GrapQL: Grapql request handlers - /grapql
    - Queue: Handle queue message - /queue
- Util: Shared command logic - /util
- Cmd: Entrypoint - cmd
- Adapter: DB, Redis, Connection - /adapters

# Features
## Core / Architecture
- [ ] Clean Architecture layering (entities, repository interfaces, use cases, adapters, infrastructure)
- [ ] Dependency inversion via interfaces
- [ ] Domain event bus (publish / subscribe, async handlers)
- [ ] Transaction manager abstraction
- [ ] ID generation (UUID)
- [ ] Clock abstraction (time provider)

## Configuration & Environment
- [ ] Central config loader (env-based)
- [ ] Feature flags system (in-memory, pluggable source)

## Logging & Error Handling
- [ ] Structured JSON logging (Zap)
- [ ] Central error constants
- [ ] Panic recovery middleware

## HTTP / API Layer
- [ ] HTTP server (Chi)
- [ ] Health check endpoint
- [ ] Example user creation endpoint (use case wiring)
- [ ] Request ID middleware
- [ ] Correlation ID middleware
- [ ] Request logging middleware
- [ ] Rate limiting (token bucket)
- [ ] Timeout middleware
- [ ] Max body size enforcement
- [ ] Content negotiation (Accept header)
- [ ] JSON content-type enforcement
- [ ] Cache-control (no-cache / immutable variants)
- [ ] Security headers (CSP, X-Frame-Options, etc.)
- [ ] Gzip compression
- [ ] ETag generation
- [ ] Idempotency key (stub)
- [ ] API version header middleware
- [ ] JWT auth middleware
- [ ] Role-based authorization middleware
- [ ] Optional vs required auth variants
- [ ] Path params extraction (simple matcher)
- [ ] Query param helper utilities
- [ ] Language detection (Accept-Language)
- [ ] Geo-IP stub middleware
- [ ] Response caching middleware (pluggable store)
- [ ] Worker pool / concurrency limiter
- [ ] Post-response async hook
- [ ] Denylist path blocking
- [ ] Request counter (atomic)
- [ ] Request origin tracking (referer / user-agent)
- [ ] CSRF protection stub
- [ ] CORS middleware
- [ ] Session cookie (signed)
- [ ] OpenTelemetry trace propagation
- [ ] Combined metrics + tracing middleware
- [ ] Debug headers middleware
- [ ] Raw request dump middleware
- [ ] Retry middleware (transient 5xx)
- [ ] Language value injection into context
- [ ] Context value injection helper
- [ ] JSON body size & type enforcement
- [ ] Immutable cache header middleware
- [ ] Response ETag + conditional GET handling
- [ ] Response cache with TTL
- [ ] Idempotent POST placeholder
- [ ] API version tagging
- [ ] Body recording for caching/ETag
- [ ] Request value map utilities

## Authentication & Authorization
- [ ] JWT issuance (HS256)
- [ ] JWT verification
- [ ] Password hashing (bcrypt)
- [ ] Password comparison
- [ ] RBAC role checks

## Persistence & Data
- [ ] Postgres connection setup (sql.DB)
- [ ] User repository interface
- [ ] Transaction manager interface
- [ ] Redis client wrapper (JSON helpers)

## Caching & Rate Limiting
- [ ] Redis-based caching utilities
- [ ] In-memory rate limiter middleware
- [ ] Response caching layer (pluggable)

## Messaging / Async
- [ ] NATS client (publish / subscribe / request)
- [ ] Kafka producer & consumer
- [ ] Domain events bus (in-memory async)

## Observability
- [ ] Prometheus metrics (request count & latency)
- [ ] OpenTelemetry tracer initialization
- [ ] Trace context propagation
- [ ] Structured logging with correlation IDs

## Validation & Utilities
- [ ] Field validation example (user name)
- [ ] Pagination helper (offset calculation)
- [ ] UUID-based ID generator
- [ ] Clock abstraction

## File & Object Storage
- [ ] Local filesystem wrapper
- [ ] S3-compatible object storage (MinIO integration)

## Email / Notification
- [ ] SMTP email sender

## Search
- [ ] OpenSearch client (index documents)

## Scheduling / Background
- [ ] Simple cron-like scheduler (ticker based)
- [ ] Post-response async hook (fire-and-forget tasks)
- [ ] Worker pool request throttling

## Security
- [ ] CSRF token generator (stub)
- [ ] CORS policy middleware
- [ ] Security headers middleware
- [ ] Session cookie signing
- [ ] Request origin tracking
- [ ] Path denylist
- [ ] Rate limiting
- [ ] Role-based access control
- [ ] Idempotency stub
- [ ] Geo-IP placeholder

## Performance & Reliability
- [ ] Gzip compression
- [ ] Retry middleware for transient failures
- [ ] Concurrency limiting (worker pool)
- [ ] Response caching
- [ ] ETag conditional responses
- [ ] Immutable caching headers

## Internationalization (i18n)
- [ ] Translation bundle (lang/key map)
- [ ] Language detection middleware

## Shutdown & Lifecycle
- [ ] Graceful shutdown orchestration
- [ ] Signal handling in main
- [ ] Cleanup hooks

## Misc / Advanced
- [ ] Feature flags (in-memory)
- [ ] Request context value map
- [ ] Debug / diagnostic middlewares
- [ ] After-response hook
- [ ] API version tagging
- [ ] Correlation ID utilities