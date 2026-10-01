---
page_title: "xcsh_k8s_pod_security_admission examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_admission examples."
---

# xcsh_k8s_pod_security_admission examples

<a id="canonical-f399425534f21107dce7b00193dd22af67c7bc17963da3aa4ed37698c8e54efa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce61c5646c9aa7464dcc929e4ac16198dac2a491b34f3c0854c51c21eddb6485"></a>

## Examples — Examples / 312fd611c2c2 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)
- Examples

<a id="canonical-fe367c853175b657a143b6ce9b105417d53eb5e140e8828dbf3c0471646bf044"></a>

## Complete configurations — Examples / 312fd611c2c2 / 3

- [Resource](resources--k8s_pod_security_admission--examples--group-001.md#canonical-628db6f211a620db9f88befaedfa86244da6fdcb5a6ec658f73b1ec07bb3de22): valid configuration.

<a id="canonical-44f2e1105bf80926355395e9834759397890f4f37435c66d5095b82678e84042"></a>

## Next pages — Examples / 312fd611c2c2 / 4

- [Resource](resources--k8s_pod_security_admission--examples--group-001.md#canonical-628db6f211a620db9f88befaedfa86244da6fdcb5a6ec658f73b1ec07bb3de22)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)

<a id="canonical-628db6f211a620db9f88befaedfa86244da6fdcb5a6ec658f73b1ec07bb3de22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c3669bb4257fd35714a6211b7500e3352bdc76bb47ac06dbcec5fb476721f95"></a>

## Resource — Resource / 078335b85146 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)
- [Examples](resources--k8s_pod_security_admission--examples--group-001.md#canonical-f399425534f21107dce7b00193dd22af67c7bc17963da3aa4ed37698c8e54efa)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_k8s_pod_security_admission/resource.tf`; digest `sha256:55b7bf85163231f6d1cb9fd77f147ddf87ba9bdf58b1aa2d6af5bb2885858dc2`.

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

<a id="canonical-3ad4c0305285b41623c73585a83c9d59c7af32d3a7fcd7d43c0a5403b9c3618a"></a>

## Next pages — Resource / 078335b85146 / 3

- [Examples](resources--k8s_pod_security_admission--examples--group-001.md#canonical-f399425534f21107dce7b00193dd22af67c7bc17963da3aa4ed37698c8e54efa)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)
