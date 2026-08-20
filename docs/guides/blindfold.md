---
page_title: "Guide: Blindfold Secret Management Functions"
subcategory: "Guides"
description: |-
  Securely encrypt sensitive data using F5 Distributed Cloud Blindfold functions.
  Covers TLS certificates, cloud credentials, and container registry authentication.
---

# Blindfold Secret Management Functions

This guide explains how to use F5 Distributed Cloud Blindfold functions to encrypt sensitive data locally before transmitting it to the control plane. Topics covered include:

- **Encrypt TLS private keys** — Configure certificate key storage
- **Protect cloud credentials** — Secure AWS, Azure, and Google Cloud credentials
- **Secure container registries** — Authenticate private container image pulls
- **Understand the security model** — Apply client-side encryption without plaintext transmission

## Overview

F5 Distributed Cloud Secret Management (Blindfold) provides client-side encryption for sensitive data. The blindfold functions encrypt secrets locally using RSA-OAEP with SHA-256, ensuring that **plaintext secrets never leave your local environment unencrypted**.

### How It Works

```text
┌─────────────────────────────────────────────────────────────────────┐
│                     Your Local Machine                              │
│                                                                     │
│  ┌──────────────┐    ┌──────────────────┐    ┌─────────────────┐   │
│  │ Secret       │───►│ Blindfold        │───►│ Encrypted       │   │
│  │ (plaintext)  │    │ Function         │    │ Ciphertext      │   │
│  └──────────────┘    │ (RSA-OAEP)       │    └────────┬────────┘   │
│                      └────────┬─────────┘             │            │
│                               │                       │            │
│                      ┌────────▼─────────┐             │            │
│                      │ F5 Distributed   │             │            │
│                      │ Cloud Public Key │             │            │
│                      │ (fetched once)   │             │            │
│                      └──────────────────┘             │            │
└──────────────────────────────────────────────────────│────────────┘
                                                        │
                                                        ▼
                                          ┌─────────────────────────┐
                                          │   F5 Distributed Cloud  │
                                          │   (stores ciphertext    │
                                          │    only)                │
                                          └─────────────────────────┘
```

### Security Properties

- **Local encryption**: Secrets are encrypted on your local machine before transmission.
- **RSA-OAEP with SHA-256**: Industry-standard asymmetric encryption.
- **Policy-controlled decryption**: Only authorized F5 Distributed Cloud services can decrypt the payload.
- **No plaintext storage**: The control plane never receives or stores plaintext secrets.

## Prerequisites

Before you begin, ensure you have:

- **Terraform >= 1.8** — Provider-defined functions require Terraform 1.8 or later (download from <https://www.terraform.io/downloads>)
- **F5 Distributed Cloud Account** — Sign up at <https://www.f5.com/cloud/products/distributed-cloud-console>
- **API Credentials** — Configure API token or P12 certificate authentication (see the [Authentication Guide](authentication.md))

### Authentication Setup

Configure one of these authentication methods using environment variables:

#### Option 1: API Token (Recommended for development)

```bash
export XCSH_API_URL="https://<XC_TENANT>.console.ves.volterra.io"
export XCSH_API_TOKEN="<XC_API_TOKEN>"
```

#### Option 2: P12 Certificate (Recommended for production)

```bash
export XCSH_API_URL="https://<XC_TENANT>.console.ves.volterra.io"
export XCSH_P12_FILE="/path/to/credentials.p12"
export XCSH_P12_PASSWORD="<XC_P12_PASSWORD>"  # pragma: allowlist secret
```

-> **Tip:** Add these variables to your shell profile (`~/.bashrc` or `~/.zshrc`) for persistence across terminal sessions.

## Quick Start

### Step 1: Clone the Repository

```bash
git clone https://github.com/f5-sales-demo/terraform-provider-xcsh.git
cd terraform-provider-xcsh/examples/guides/blindfold
```

### Step 2: Configure Your Deployment

```bash
cp terraform.tfvars.example terraform.tfvars
```

Edit `terraform.tfvars` to enable the examples you want to test:

```hcl
# Enable the TLS certificate example
enable_certificate_example = true

# Namespace configuration
namespace_name   = "example-blindfold-test"
create_namespace = true
```

### Step 3: Generate Test Certificates (for certificate example)

```bash
openssl req -x509 -newkey rsa:2048 \
  -keyout certs/server.key \
  -out certs/server.crt \
  -days 365 -nodes \
  -subj "/CN=example.com"
```

### Step 4: Deploy

```bash
terraform init
terraform plan
terraform apply
```

Review the plan output, then type `yes` to confirm deployment.

### Step 5: Verify

1. Check the Terraform outputs for created resource names
2. Navigate to the F5XC Console
3. Verify the certificate shows encrypted private key (not plaintext)

## Understanding Blindfold Functions

The provider includes two blindfold functions:

### blindfold()

Encrypts base64-encoded plaintext:

```hcl
provider::xcsh::blindfold(plaintext, policy_name, namespace)
```

| Parameter     | Type   | Description                      |
| ------------- | ------ | -------------------------------- |
| `plaintext`   | string | Base64-encoded secret to encrypt |
| `policy_name` | string | Name of the SecretPolicy         |
| `namespace`   | string | Namespace containing the policy  |

**Example:**

```hcl
location = provider::xcsh::blindfold(
  base64encode(var.my_secret),
  "ves-io-allow-volterra",
  "shared"
)
```

### blindfold_file()

Reads a file and encrypts its contents:

```hcl
provider::xcsh::blindfold_file(path, policy_name, namespace)
```

| Parameter     | Type   | Description                     |
| ------------- | ------ | ------------------------------- |
| `path`        | string | Path to the file to encrypt     |
| `policy_name` | string | Name of the SecretPolicy        |
| `namespace`   | string | Namespace containing the policy |

**Example:**

```hcl
location = provider::xcsh::blindfold_file(
  "${path.module}/certs/server.key",
  "ves-io-allow-volterra",
  "shared"
)
```

### Built-in SecretPolicy

Every F5 Distributed Cloud tenant includes a default policy: `ves-io-allow-volterra` in the `shared` namespace. This policy allows F5 Distributed Cloud platform services to decrypt secrets.

## Use Cases

### TLS Certificate Private Key

Store TLS certificates for load balancers. Note that RSA TLS private keys are typically too large (>1000 bytes) for direct asymmetric blindfold encryption (~190 byte limit). Use `clear_secret_info` for TLS keys:

```hcl
resource "xcsh_certificate" "example" {
  name      = "example-certificate"
  namespace = "shared"

  certificate_url = "string:///${base64encode(file("${path.module}/certs/server.crt"))}"

  # TLS private keys exceed the RSA-OAEP direct encryption size limit
  # Use clear_secret_info — the key is transmitted securely over HTTPS (TLS)
  private_key {
    clear_secret_info {
      url = "string:///${base64encode(file("${path.module}/certs/server.key"))}"
    }
  }

  use_system_defaults {}
}
```

~> **Size Limitation:** Direct Blindfold encryption uses RSA-OAEP which limits plaintext to ~190 bytes for 2048-bit keys. TLS private keys exceed this limit. Use `clear_secret_info` for certificates — the payload is encrypted in transit over HTTPS. For secrets under 190 bytes (such as API keys, passwords, and client tokens), use the `blindfold()` function as shown below.

### AWS Cloud Credentials

Protect AWS secret access keys for VPC site deployments:

```hcl
resource "xcsh_cloud_credentials" "aws" {
  name      = "aws-credentials"
  namespace = "system"

  aws_secret_key {
    access_key = var.aws_access_key_id

    secret_key {
      blindfold_secret_info {
        location = provider::xcsh::blindfold(
          base64encode(var.aws_secret_access_key),
          "ves-io-allow-volterra",
          "shared"
        )
      }
    }
  }
}
```

### Azure Cloud Credentials

Secure Azure service principal client secrets:

```hcl
resource "xcsh_cloud_credentials" "azure" {
  name      = "azure-credentials"
  namespace = "system"

  azure_client_secret {
    subscription_id = var.azure_subscription_id
    tenant_id       = var.azure_tenant_id
    client_id       = var.azure_client_id

    client_secret {
      blindfold_secret_info {
        location = provider::xcsh::blindfold(
          base64encode(var.azure_client_secret),
          "ves-io-allow-volterra",
          "shared"
        )
      }
    }
  }
}
```

### GCP Cloud Credentials

Encrypt GCP service account JSON key files:

```hcl
resource "xcsh_cloud_credentials" "gcp" {
  name      = "gcp-credentials"
  namespace = "system"

  gcp_cred_file {
    credential_file {
      blindfold_secret_info {
        location = provider::xcsh::blindfold_file(
          var.gcp_credentials_file,
          "ves-io-allow-volterra",
          "shared"
        )
      }
    }
  }
}
```

### Container Registry Authentication

Protect container registry passwords for private image pulls:

```hcl
resource "xcsh_container_registry" "example" {
  name      = "docker-registry"
  namespace = "shared"

  registry  = "docker.io"
  user_name = var.registry_username

  password {
    blindfold_secret_info {
      location = provider::xcsh::blindfold(
        base64encode(var.registry_password),
        "ves-io-allow-volterra",
        "shared"
      )
    }
  }
}
```

## Configuration Options

### Using Custom SecretPolicies

While the built-in `ves-io-allow-volterra` policy works for most cases, you can create custom policies for fine-grained access control:

```hcl
locals {
  policy_name = "example-custom-policy"
  policy_ns   = "example-namespace"
}

# Reference your custom policy
location = provider::xcsh::blindfold(
  base64encode(var.secret),
  local.policy_name,
  local.policy_ns
)
```

### Encrypting Multiple Secrets with for_each

```hcl
variable "secrets" {
  type = map(string)
  default = {
    "api-key"    = "secret1"
    "auth-token" = "secret2"
  }
}

locals {
  encrypted_secrets = {
    for name, value in var.secrets :
    name => provider::xcsh::blindfold(
      base64encode(value),
      "ves-io-allow-volterra",
      "shared"
    )
  }
}
```

## Technical Details

### Size Limitations

RSA-OAEP encryption has a maximum plaintext size based on the key size:

| Key Size | Maximum Plaintext |
| -------- | ----------------- |
| 2048-bit | ~190 bytes        |
| 4096-bit | ~446 bytes        |

~> **Note:** If your secret exceeds the size limit, consider splitting it or using an alternative approach. The function returns an error message if the plaintext exceeds the maximum allowed size.

### Output Format

The blindfold functions return a sealed secret string with the `string:///` prefix followed by a base64-encoded JSON structure:

```text
string:///eyJrZXlfdmVyc2lvbiI6InYxLjIuMyIsInBvbGljeV9pZCI6InNoYXJlZC92ZXMtaW8tYWxsb3ctdm9sdGVycmEiLCJ0ZW5hbnQiOiJ5b3VyLXRlbmFudCIsImRhdGEiOiJBQkNERUYxMjM0NTY3ODkwLi4uIn0=
```

When base64-decoded, the sealed JSON contains these fields:

```json
{
  "key_version": "v1.2.3",
  "policy_id": "shared/ves-io-allow-volterra",
  "tenant": "<XC_TENANT>",
  "data": "ABCDEF1234567890..."
}
```

Field descriptions:

- `key_version`: Public key version used for encryption
- `policy_id`: Reference to the SecretPolicy (`namespace/name` format)
- `tenant`: Tenant or organization identifier
- `data`: Base64-encoded RSA-OAEP ciphertext

### Function Behavior

- **Non-deterministic output**: The same plaintext generates different ciphertext on each run due to randomized OAEP padding.
- **Network access required**: Functions fetch the active public key from the F5 Distributed Cloud API.
- **In-memory caching**: Public keys are cached for the duration of the Terraform execution.

## Troubleshooting

### Authentication Configuration Error

**Symptom:** Error message indicates missing or invalid authentication credentials.

**Solution:**

```bash
# Verify environment variables are set
echo "$XCSH_API_URL"
echo "$XCSH_API_TOKEN"  # or XCSH_P12_FILE

# Set them if missing
export XCSH_API_URL="https://<XC_TENANT>.console.ves.volterra.io"
export XCSH_API_TOKEN="<XC_API_TOKEN>"
```

### Policy Not Found

**Symptom:** Error indicating that the specified SecretPolicy was not found.

**Solutions:**

1. Use the built-in default policy:

   ```hcl
   policy_name = "ves-io-allow-volterra"
   namespace   = "shared"
   ```

2. Verify that any custom SecretPolicy exists in the specified namespace in the F5 Distributed Cloud Console.

### Plaintext Too Large

**Symptom:** Error indicates that the plaintext exceeds the maximum allowed size for RSA-OAEP encryption.

**Solutions:**

1. Check your secret payload size:

   ```bash
   wc -c < your-secret-file
   ```

2. For larger files (such as GCP credentials or TLS private keys):
   - Extract only the private key or secret token portion
   - Use `clear_secret_info` with TLS in transit

### File Not Found

**Symptom:** Error indicating file not found when calling `blindfold_file()`.

**Solutions:**

1. Use `${path.module}` for relative file paths:

   ```hcl
   location = provider::xcsh::blindfold_file(
     "${path.module}/certs/server.key",
     "ves-io-allow-volterra",
     "shared"
   )
   ```

2. Verify that the file exists and has appropriate read permissions.

### Invalid Base64 Encoding

**Symptom:** Error indicating invalid base64 encoding.

**Solution:** Ensure you base64-encode plaintext strings before passing them to `blindfold()`:

```hcl
# Correct
location = provider::xcsh::blindfold(
  base64encode(var.secret),
  "ves-io-allow-volterra",
  "shared"
)

# Incorrect
location = provider::xcsh::blindfold(
  var.secret,  # Raw plaintext string fails
  "ves-io-allow-volterra",
  "shared"
)
```

## Clean Up

To remove all resources created by this guide:

```bash
terraform destroy
```

Type `yes` when prompted to confirm resource destruction.

~> **Warning:** This command immediately destroys all managed resources in the plan. Ensure you have backups of any certificates or credentials before proceeding.

## Next Steps

- [Certificate Resource](../resources/certificate.md) — Full certificate management
- [Cloud Credentials Resource](../resources/cloud_credentials.md) — Cloud provider authentication
- [HTTP Load Balancer Guide](./http-loadbalancer.md) — Use certificates in load balancers
- [blindfold Function Reference](../functions/blindfold.md) — Function API details
- [blindfold_file Function Reference](../functions/blindfold_file.md) — Function API details

## Support

- [Provider Documentation](../index.md)
- [F5 Distributed Cloud Documentation](https://docs.cloud.f5.com/)
- [F5 Secret Management](https://docs.cloud.f5.com/docs/how-to/secrets-management)
- [GitHub Issues](https://github.com/f5-sales-demo/terraform-provider-xcsh/issues)
