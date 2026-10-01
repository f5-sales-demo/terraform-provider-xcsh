---
page_title: "xcsh_k8s_pod_security_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_policy landing."
---

# xcsh_k8s_pod_security_policy landing

<a id="canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e6b6271e6900d593e018012bbe5ee1bfc7359c7a3a1d2d235e081804dbbacd5"></a>

## xcsh_k8s_pod_security_policy — xcsh_k8s_pod_security_policy / 9ddb9b458d2f / 2

Breadcrumbs:

- xcsh_k8s_pod_security_policy

Manages k8s\_pod\_security\_policy will create the object in the storage backend for namespace
metadata.namespace in F5 Distributed Cloud.

<a id="canonical-08ff58a7d7c3fe7eb3f252150f92ccf55d9d27e8220ef8cf90ab1c991647efda"></a>

## Prerequisites — xcsh_k8s_pod_security_policy / 9ddb9b458d2f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3a3bace979c8739e1427671eb170db99ca64c16a2de5d9a617f6fe7b10c49f07"></a>

## Minimal configuration — xcsh_k8s_pod_security_policy / 9ddb9b458d2f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SPodSecurityPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing K8SPodSecurityPolicy by name
data "xcsh_k8s_pod_security_policy" "example" {
  name      = "example-k8s-pod-security-policy"
  namespace = "staging"
}

output "k8s_pod_security_policy_id" {
  value = data.xcsh_k8s_pod_security_policy.example.id
}
```

<a id="canonical-19b7d0cd97482d6ff38bea8df25ddcebd773991de5f1f08b0254ac11c94968d0"></a>

## Root configuration — xcsh_k8s_pod_security_policy / 9ddb9b458d2f / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-bdab74ddf73393b9c9b4b78c736ba7d281305c13dfce641d544df9276b41eaf9"></a>

## Next pages — xcsh_k8s_pod_security_policy / 9ddb9b458d2f / 6

- [Property reference](../guides/data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [Examples](../guides/data-sources--k8s_pod_security_policy--examples--group-001.md#canonical-94881518bbc37debae300d89fbdc06722ac20767c40925cad906ca38d7b59aa1)
