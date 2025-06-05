# Notifications Service

A lightweight, scalable proxy service that receives notification requests from various backend services and forwards them to Novu.co with proper payload transformation and workflow mapping.

## 🏗️ Architecture Overview

```text
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   P2P Service   │───▶│ Notifications    │───▶│   Novu.co API   │
│   KYC Service   │    │    Service       │    │                 │
│ Security Service│    │                  │    │                 │
└─────────────────┘    └──────────────────┘    └─────────────────┘
```

The service acts as a secure proxy between internal services and Novu, providing:

- **Authentication** via JWT tokens from AWS Cognito
- **Payload transformation** from internal format to Novu-compatible format
- **Workflow mapping** based on message keys and types
- **Reliability** with retry logic and circuit breakers
- **Observability** with metrics, logging, and health checks

## 🚀 Quick Start

### Prerequisites

- Go 1.24+
- Docker & Docker Compose
- Kubernetes cluster (for production deployment)
- Novu.co account and API key
- AWS Cognito setup for JWT validation

### Local Development Setup

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd notifications-service
   ```

2. **Set up environment variables**

   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Run with Devcontainers**

4. **Test the service**

   ```bash
   curl -X POST http://localhost:8080/v1/notifications \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer <your-jwt-token>" \
     -d '{
       "messageKey": "P2P-ORDER-CREATED",
       "type": "buyer",
       "userId": "user123",
       "payload": {
         "orderId": "order456",
         "amount": 100.50,
         "currency": "USD"
       }
     }'
   ```

## 📁 Project Structure

```text
notifications-service/
├── cmd/server/           # Application entry point
├── internal/
│   ├── config/          # Configuration management
│   ├── handlers/        # HTTP handlers
│   ├── services/        # Business logic
│   ├── clients/         # External service clients (Novu, JWKS)
│   ├── middleware/      # Authentication, logging, metrics
│   └── models/          # Data structures
├── configs/             # Environment-specific config files
├── docker/              # Dockerfiles and compose files
├── scripts/             # Build and deployment scripts
├── tests/               # Integration and e2e tests
└── docs/                # API documentation
```

## 🔧 Configuration

### Environment Variables

```bash
# Required
NOVU_API_KEY=your_novu_api_key
JWKS_URL=https://cognito-idp.region.amazonaws.com/userpoolid/.well-known/jwks.json
JWT_ISSUER=https://cognito-idp.region.amazonaws.com/userpoolid

# Optional
PORT=8080
LOG_LEVEL=info
MAX_RETRIES=3
TIMEOUT_SECONDS=5
```

### Workflow Mappings

Configure notification workflows in `configs/workflow_mappings.yaml`:

```yaml
workflows:
  - messageKey: "P2P-ORDER-CREATED"
    type: "buyer"
    workflowId: "p2p-buy-order-created"
  
  - messageKey: "P2P-ORDER-CREATED" 
    type: "seller"
    workflowId: "p2p-sell-order-created"
    
  - messageKey: "KYC-VERIFICATION-REQUIRED"
    type: "user"
    workflowId: "kyc-verification-needed"
    
  - messageKey: "SECURITY-ALERT"
    type: "urgent"
    workflowId: "security-high-priority"
```

## 📡 API Reference

### POST /v1/notifications

Send a notification request.

**Request Body:**

```json
{
  "messageKey": "P2P-ORDER-CREATED",
  "type": "buyer", 
  "userId": "user123",
  "payload": {
    "orderId": "order456",
    "amount": 100.50,
    "currency": "USD"
  }
}
```

**Response:**

```json
{
  "success": true,
  "requestId": "req_abc123",
  "message": "Notification queued successfully"
}
```

### GET /health

Health check endpoint for liveness probes.

### GET /ready

Readiness check endpoint for readiness probes.

### GET /metrics

Prometheus metrics endpoint.

## 🧪 Testing

### Run Tests

```bash
# Unit tests
make test-unit

# Integration tests with mocks
make test-integration

# E2E tests (requires Novu dev environment)
make test-e2e

# Load testing
make test-load

# All tests with coverage
go test ./... -v -cover
```

### Test Categories

- **Unit Tests**: Mock external dependencies, test business logic
- **Integration Tests**: Test against mock services
- **E2E Tests**: Full flow testing with real Novu development environment
- **Load Tests**: Performance and scalability testing

## 🚢 Deployment

### Docker Build

```bash
docker build -t notifications-service:latest .
```

### Kubernetes Deployment

```bash
# Apply Kubernetes manifests
kubectl apply -f deployments/k8s/

# Rolling update
kubectl set image deployment/notifications-service \
  notifications-service=notifications-service:v1.2.0
```

### Environment Strategy

- **Development**: Local Docker Compose with mocks
- **Staging**: Kubernetes cluster with Novu development instance
- **Production**: Multi-AZ Kubernetes deployment with auto-scaling

## 📊 Monitoring & Observability

### Key Metrics

- `notifications_total{status, messageKey, type}` - Total notifications processed
- `notification_duration_seconds{messageKey, type}` - Processing latency
- `notification_errors_total{error_type, messageKey}` - Error counts
- `novu_api_calls_total{status_code}` - Novu API call metrics
- `jwt_validation_total{status}` - Authentication metrics

### Health Checks

- **Liveness**: Service is running (`/health`)
- **Readiness**: Service can handle traffic (`/ready`)
- **Dependencies**: Novu API and JWKS endpoint connectivity

### Logging

- Structured JSON logging with correlation IDs
- Configurable log levels (DEBUG for dev, INFO for production)
- Request/response logging with performance metrics

## 🔒 Security

### Authentication & Authorization

- JWT token validation using JWKS from AWS Cognito
- Signature verification and claims validation
- Rate limiting to prevent abuse

### Security Best Practices

- API keys stored in environment variables/secrets
- Non-root container execution
- Input validation and sanitization
- CORS configuration
- TLS/HTTPS enforcement

## 🛠️ Development

### Technology Stack

- **Language**: Go 1.24+
- **HTTP Framework**: Standard `net/http` with `gorilla/mux`
- **Configuration**: YAML with `gopkg.in/yaml.v3`
- **JWT**: `golang-jwt/jwt/v5`
- **Metrics**: Prometheus client library
- **Logging**: Standard `log/slog`

### Branch Strategy

- `main` - Production-ready code
- `develop` - Integration branch
- `feature/*` - Feature development
- `hotfix/*` - Production fixes

### Contributing

1. Create feature branch from `develop`
2. Implement changes with tests
3. Run full test suite
4. Submit pull request with description
5. Code review and merge

## 🔄 CI/CD Pipeline

Automated pipeline with GitHub Actions:

- **Test**: Unit, integration, and E2E tests
- **Build**: Docker image creation and registry push
- **Deploy**: Automated deployment to staging/production
- **Rollback**: Quick rollback capability

## 🆘 Troubleshooting

### Common Issues

**Service not starting:**

- Check environment variables are set correctly
- Verify JWKS URL is accessible
- Ensure Novu API key is valid

**Authentication failures:**

- Verify JWT token format and signature
- Check JWKS endpoint connectivity
- Validate token issuer and audience claims

**Novu API errors:**

- Check API key permissions
- Verify workflow IDs exist in Novu
- Review payload format requirements

### Logs and Debugging

```bash
# View service logs
kubectl logs -f deployment/notifications-service

# Check specific error patterns
kubectl logs deployment/notifications-service | grep "ERROR"

# Monitor metrics
curl http://localhost:8080/metrics
```

## 📚 Additional Resources

- [API Documentation](./docs/api.md)
- [Architecture Deep Dive](./docs/architecture.md)
- [Deployment Guide](./docs/deployment.md)
- [Novu.co Documentation](https://docs.novu.co/)

## 📄 License

MIT License

## 🤝 Support

For support and questions:

- Create an issue in this repository
- Contact the development team
- Check the troubleshooting guide above
