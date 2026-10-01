---
page_title: "xcsh_k8s_pod_security_admission landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_admission landing."
---

# xcsh_k8s_pod_security_admission landing

<a id="canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b416fec1331a90bc82134aa9f1e28a9f9933d95d62e8f9eea14a1c6b9e60b38d"></a>

## xcsh_k8s_pod_security_admission — xcsh_k8s_pod_security_admission / a68502ec0bf6 / 2

Breadcrumbs:

- xcsh_k8s_pod_security_admission

Manages k8s\_pod\_security\_admission will create the object in the storage backend in F5
Distributed Cloud.

<a id="canonical-5e6e3edbc41ae8b8095f6084f744982caa461eb8e3dbd170b6cbfe6f481bfb56"></a>

## Prerequisites — xcsh_k8s_pod_security_admission / a68502ec0bf6 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-bcc1008729b51393a98bbf8b5ce3f57ae916b45b8e8f3ffafc60d681cb3e332c"></a>

## Minimal configuration — xcsh_k8s_pod_security_admission / a68502ec0bf6 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SPodSecurityAdmission Resource Example
# Manages k8s_pod_security_admission will create the object in the storage backend in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic K8SPodSecurityAdmission configuration
resource "xcsh_k8s_pod_security_admission" "example" {
  name      = "example-k8s-pod-security-admission"
  namespace = "system"
}
```

<a id="canonical-312e05c8be3fa6b275c2e136c695529a96e234e072cbb35313658e5f2a2bf878"></a>

## Root configuration — xcsh_k8s_pod_security_admission / a68502ec0bf6 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-1bb2e5626e68949d126c0c0477733e0248cba63d9c64f82047d653d131cc33da"></a>

## Next pages — xcsh_k8s_pod_security_admission / a68502ec0bf6 / 6

- [Property reference](../guides/resources--k8s_pod_security_admission--reference--group-001.md#canonical-aaff46c769c7779ed88f3eb28bd97bdfe97115cc0982f7d385c2509e2607313c)
- [Examples](../guides/resources--k8s_pod_security_admission--examples--group-001.md#canonical-f399425534f21107dce7b00193dd22af67c7bc17963da3aa4ed37698c8e54efa)
- [Import](../guides/resources--k8s_pod_security_admission--lifecycle--group-001.md#canonical-7eea5c62d6c0e1d4ff1bcae58c0507b74ca9fdffc63f281506649b85dfc25582)
- [Timeouts](../guides/resources--k8s_pod_security_admission--lifecycle--group-001.md#canonical-a6c18e2cd84cc7172dd4b2f66fb22400ca424e2b705f4c9726f7cd440271e120)
