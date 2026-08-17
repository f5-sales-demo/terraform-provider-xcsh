---
page_title: "Guide: Authentication Methods"
subcategory: "Guides"
description: |-
  Comprehensive guide to authenticating the F5 Distributed Cloud Terraform
  provider using API tokens, P12 certificates, or PEM certificates.
---

# Authentication Methods

This guide covers authentication configuration for the F5 Distributed Cloud Terraform provider.

## Quick Reference

| Method          | Complexity | Best For                         | Security               |
| --------------- | ---------- | -------------------------------- | ---------------------- |
| API Token       | Simplest   | Development, quick testing       | Bearer token over TLS  |
| P12 Certificate | Moderate   | Production, CI/CD                | Mutual TLS (mTLS)      |
| PEM Certificate | Advanced   | When tooling requires PEM format | Derived from P12, mTLS |

## Prerequisites

- **Terraform >= 1.0** — Download from <https://www.terraform.io/downloads>
- **F5 Distributed Cloud Account** — Sign up at <https://www.f5.com/cloud/products/distributed-cloud-console>
- **Console Access** — Tenant credentials with permissions to generate API tokens or certificates

## Creating Credentials in F5 Distributed Cloud

### Personal Credentials

Personal credentials are tied to your user account and suitable for development environments.

#### Creating an API Token

1. Open the F5 Distributed Cloud Console.
2. Navigate to **Administration** → **Personal Management** → **Credentials**.
3. Select **+ Add Credentials**.
4. Enter a name (for example, `terraform-dev-token`).
5. In the **Credential Type** list, select **API Token**.
6. Choose an expiration date.
7. Select **Generate**, and then copy the token value.

~> **Warning:** Copy and store your token immediately. You cannot retrieve the token value after closing the dialog.

#### Creating a P12 Certificate

1. Navigate to **Administration** → **Personal Management** → **Credentials**.
2. Select **+ Add Credentials**.
3. Enter a name (for example, `terraform-dev-cert`).
4. In the **Credential Type** list, select **API Certificate**.
5. Enter and confirm a password.
6. Select an expiration date.
7. Select **Download** to save the `.p12` certificate file.

-> **Tip:** Store the P12 certificate securely and never commit it to version control.

### Service Credentials (IAM)

Service credentials are managed through IAM and recommended for CI/CD pipelines and production automation. You can scope service credentials to specific roles and namespaces.

#### Creating a Service API Token

1. Navigate to **Administration** → **IAM** → **Service Credentials**.
2. Select **+ Add Service Credentials**.
3. Enter a name (for example, `terraform-cicd-token`).
4. In the **Credential Type** list, select **API Token**.
5. Assign roles and namespaces to restrict credential scope.
6. Choose an expiration date.
7. Select **Generate**, and then copy the token value.

#### Creating a Service P12 Certificate

1. Navigate to **Administration** → **IAM** → **Service Credentials**.
2. Select **+ Add Service Credentials**.
3. Enter a name (for example, `terraform-cicd-cert`).
4. In the **Credential Type** list, select **API Certificate**.
5. Assign roles and namespaces to restrict scope.
6. Enter and confirm a password.
7. Select an expiration date.
8. Select **Download** to save the `.p12` certificate file.

## Provider Configuration

### Method 1: API Token Authentication (Simplest)

API tokens provide bearer token authentication over TLS. This is the quickest way to get started.

**Using Environment Variables:**

```bash
export XCSH_API_URL="https://<XC_TENANT>.console.ves.volterra.io"
export XCSH_API_TOKEN="<XC_API_TOKEN>"
```

**Using Provider Configuration:**

```hcl
provider "xcsh" {
  api_url   = var.xcsh_api_url
  api_token = var.xcsh_api_token
}
```

### Method 2: P12 Certificate Authentication (Recommended for Production)

P12 certificates provide mutual TLS (mTLS) authentication, where both client and server verify each other's identity.

**Using Environment Variables:**

```bash
export XCSH_API_URL="https://<XC_TENANT>.console.ves.volterra.io"
export XCSH_P12_FILE="/path/to/credentials.p12"
export XCSH_P12_PASSWORD="<XC_P12_PASSWORD>"  # pragma: allowlist secret
```

**Using Provider Configuration:**

```hcl
provider "xcsh" {
  api_url      = var.xcsh_api_url
  api_p12_file = var.xcsh_api_p12_file
  p12_password = var.xcsh_p12_password
}
```

### Method 3: PEM Certificate Authentication (Derived from P12)

PEM authentication uses separate certificate and private key files. Because F5 Distributed Cloud generates P12 certificates, you must extract PEM files using OpenSSL when required by external tooling.

**When to use this method:** Use PEM authentication only when your tooling specifically requires PEM format rather than P12.

**Step 1: Extract PEM files from P12:**

```bash
# Create a directory for certificates
mkdir -p certs

# Extract the certificate
openssl pkcs12 -in ~/<XC_TENANT>.console.ves.volterra.io.api-creds.p12 \
  -nodes -nokeys -out certs/xcsh.cert
# Enter Import Password: <XC_P12_PASSWORD>

# Extract the private key
openssl pkcs12 -in ~/<XC_TENANT>.console.ves.volterra.io.api-creds.p12 \
  -nodes -nocerts -out certs/xcsh.key
# Enter Import Password: <XC_P12_PASSWORD>
```

**Step 2: Configure the provider:**

**Using Environment Variables:**

```bash
export XCSH_API_URL="https://<XC_TENANT>.console.ves.volterra.io"
export XCSH_CERT="/path/to/certs/xcsh.cert"
export XCSH_KEY="/path/to/certs/xcsh.key"
```

**Using Provider Configuration:**

```hcl
provider "xcsh" {
  api_url  = var.xcsh_api_url
  api_cert = var.xcsh_api_cert
  api_key  = var.xcsh_api_key
}
```

~> **Note:** For server certificate verification, specify a CA certificate using the `XCSH_CACERT` environment variable or the `api_ca_cert` provider attribute.

## Environment Variable Reference

| Variable            | Description                                    | Required                       |
| ------------------- | ---------------------------------------------- | ------------------------------ |
| `XCSH_API_URL`      | F5 Distributed Cloud tenant API URL            | Yes                            |
| `XCSH_API_TOKEN`    | API token for bearer authentication            | One of: token, P12, or PEM     |
| `XCSH_P12_FILE`     | Path to P12 certificate file                   | With `XCSH_P12_PASSWORD`       |
| `XCSH_P12_PASSWORD` | Password for P12 file                          | With `XCSH_P12_FILE`           |
| `XCSH_CERT`         | Path to PEM certificate file                   | With `XCSH_KEY`                |
| `XCSH_KEY`          | Path to PEM private key file                   | With `XCSH_CERT`               |
| `XCSH_CACERT`       | Path to CA certificate for server verification | No                             |

**Adding to Shell Profile:**

```bash
# Add to ~/.bashrc or ~/.zshrc
export XCSH_API_URL="https://<XC_TENANT>.console.ves.volterra.io"
export XCSH_API_TOKEN="<XC_API_TOKEN>"
```

Then reload: `source ~/.zshrc` or `source ~/.bashrc`

### Authentication Priority

When multiple authentication methods are configured, the provider evaluates them in this order:

1. **P12 Certificate** — Evaluated when `api_p12_file` is set (requires `p12_password`)
2. **PEM Certificate** — Evaluated when both `api_cert` and `api_key` are set
3. **API Token** — Evaluated when `api_token` is set
4. **Error** — Returned when no valid credentials are provided

## CI/CD Integration

### GitHub Actions with API Token

```yaml
name: Terraform Deploy

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  terraform:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: hashicorp/setup-terraform@v3
        with:
          terraform_version: "1.8.0"

      - name: Terraform Init
        run: terraform init

      - name: Terraform Plan
        env:
          XCSH_API_URL: ${{ secrets.XCSH_API_URL }}
          XCSH_API_TOKEN: ${{ secrets.XCSH_API_TOKEN }}
        run: terraform plan -out=tfplan

      - name: Terraform Apply
        if: github.ref == 'refs/heads/main' && github.event_name == 'push'
        env:
          XCSH_API_URL: ${{ secrets.XCSH_API_URL }}
          XCSH_API_TOKEN: ${{ secrets.XCSH_API_TOKEN }}
        run: terraform apply -auto-approve tfplan
```

**GitHub Secrets to configure:**

| Secret Name        | Value                                         |
| ------------------ | --------------------------------------------- |
| `XCSH_API_URL`     | `https://<XC_TENANT>.console.ves.volterra.io` |
| `XCSH_API_TOKEN`   | `<XC_API_TOKEN>`                              |

### GitHub Actions with P12 Certificate

For production deployments requiring mTLS:

```yaml
name: Terraform Deploy with P12

on:
  push:
    branches: [main]

jobs:
  terraform:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: hashicorp/setup-terraform@v3
        with:
          terraform_version: "1.8.0"

      - name: Setup P12 Certificate
        run: |
          echo "${{ secrets.XCSH_P12_BASE64 }}" | base64 -d > /tmp/xcsh-credentials.p12
          chmod 600 /tmp/xcsh-credentials.p12

      - name: Terraform Init
        run: terraform init

      - name: Terraform Plan
        env:
          XCSH_API_URL: ${{ secrets.XCSH_API_URL }}
          XCSH_P12_FILE: /tmp/xcsh-credentials.p12
          XCSH_P12_PASSWORD: ${{ secrets.XCSH_P12_PASSWORD }}
        run: terraform plan -out=tfplan

      - name: Terraform Apply
        env:
          XCSH_API_URL: ${{ secrets.XCSH_API_URL }}
          XCSH_P12_FILE: /tmp/xcsh-credentials.p12
          XCSH_P12_PASSWORD: ${{ secrets.XCSH_P12_PASSWORD }}
        run: terraform apply -auto-approve tfplan

      - name: Cleanup
        if: always()
        run: rm -f /tmp/xcsh-credentials.p12
```

**Encoding P12 for GitHub Secrets:**

```bash
# On macOS
base64 -i credentials.p12 | pbcopy

# On Linux
base64 -w 0 credentials.p12
```

**GitHub Secrets to configure:**

| Secret Name         | Value                                         |
| ------------------- | --------------------------------------------- |
| `XCSH_API_URL`      | `https://<XC_TENANT>.console.ves.volterra.io` |
| `XCSH_P12_BASE64`   | Base64-encoded P12 file contents              |
| `XCSH_P12_PASSWORD` | Password for the P12 file                     |

## Security Best Practices

- **Never commit credentials** to version control. Add `*.tfvars`, `*.p12`, and `*.pem` to `.gitignore`.
- **Use environment variables** for sensitive values in local development.
- **Use secret managers or CI/CD secrets** for automated pipelines.
- **Limit credential scope** using Service Credentials with specific roles and namespaces.

### Choosing the Right Method

| Use Case              | Recommended Method  | Reason                    |
| --------------------- | ------------------- | ------------------------- |
| Local development     | API Token           | Simplest setup            |
| CI/CD pipelines       | P12 Certificate     | mTLS security             |
| Production automation | Service Credentials | Role-based access control |

## Troubleshooting

### Authentication Failed (401 Unauthorized)

1. Verify that the API URL does **not** include the `/api` suffix (for example, `https://<XC_TENANT>.console.ves.volterra.io`).
2. Verify that the token has not expired.
3. Verify that the token was copied without leading or trailing whitespace.
4. Verify that the environment variables are exported:

```bash
echo "$XCSH_API_URL"
echo "$XCSH_API_TOKEN"
```

### Certificate Verification Failed

1. Verify that the P12 password is correct.
2. Verify that the file path is accessible and correctly referenced.
3. Test the P12 file locally:

```bash
openssl pkcs12 -in credentials.p12 -nokeys -info
```

### Permission Denied (403 Forbidden)

1. Verify that the credential has required permissions for the target resources.
2. For Service Credentials, check assigned roles and namespace access.
3. Some operations require specific system roles (for example, tenant administration).

### Environment Variables Not Working

1. Verify that variables are exported:

```bash
export XCSH_API_TOKEN="<XC_API_TOKEN>"  # Correct
XCSH_API_TOKEN="<XC_API_TOKEN>"         # Is not visible to child processes
```

2. Verify that variable names match the exact required casing (`XCSH_API_URL`, `XCSH_API_TOKEN`).
3. Check for hidden or unexpected characters in variable values.

## Revoking Credentials

1. Navigate to **Administration** → **Personal Management** → **Credentials** (or **IAM** → **Service Credentials**).
2. Locate the credential to revoke.
3. Select **Actions** (three dots) → **Force Expiry**.

## Next Steps

- [HTTP Load Balancer Guide](http-loadbalancer.md) — Deploy your first load balancer
- [Blindfold Functions Guide](blindfold.md) — Secure secret management

## Support

- [Provider Documentation](../index.md)
- [F5 Distributed Cloud Documentation](https://docs.cloud.f5.com/)
- [F5 Credentials Guide](https://docs.cloud.f5.com/docs-v2/administration/how-tos/user-mgmt/Credentials)
- [GitHub Issues](https://github.com/f5-sales-demo/terraform-provider-xcsh/issues)
