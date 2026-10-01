---
page_title: "xcsh_k8s_pod_security_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_policy landing."
---

# xcsh_k8s_pod_security_policy landing

<a id="canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4046dc4c6d18458deada5d991762cd9b23a614ce2d72dea8ccc981745194631c"></a>

## xcsh_k8s_pod_security_policy — xcsh_k8s_pod_security_policy / 3474840c213b / 2

Breadcrumbs:

- xcsh_k8s_pod_security_policy

Manages k8s\_pod\_security\_policy will create the object in the storage backend for namespace
metadata.namespace in F5 Distributed Cloud.

<a id="canonical-5954035eb55f8b6228e5f28de5b9c0fd496a46bd5f604bc4d04d6ce84b4f765d"></a>

## Prerequisites — xcsh_k8s_pod_security_policy / 3474840c213b / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-8d2daca0bec2319aaaaaf546e2319d07b16af950fc9124ab94f0458e9e1b25f8"></a>

## Minimal configuration — xcsh_k8s_pod_security_policy / 3474840c213b / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SPodSecurityPolicy Resource Example
# Manages k8s_pod_security_policy will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic K8SPodSecurityPolicy configuration
resource "xcsh_k8s_pod_security_policy" "example" {
  name      = "example-k8s-pod-security-policy"
  namespace = "staging"
}
```

<a id="canonical-e584d663f20e640599d0e26563c7ce6fb2a304d0e423841ad95fc33c87bcc492"></a>

## Root configuration — xcsh_k8s_pod_security_policy / 3474840c213b / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-719d48c038ec6099120bf94cdca71c591a6e23867fc36343214518b42bf6db6b"></a>

## Next pages — xcsh_k8s_pod_security_policy / 3474840c213b / 6

- [Property reference](../guides/resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [Examples](../guides/resources--k8s_pod_security_policy--examples--group-001.md#canonical-ffe90ec21ce5cd40a9ee9e2ea88847034ff59a568bd6d1da400ed3e4c87cde08)
- [Import](../guides/resources--k8s_pod_security_policy--lifecycle--group-001.md#canonical-73770d1cf2d9ebfb256664f3d351accf0de191b389d3cd05b36d2b72cc9c7ec8)
- [Timeouts](../guides/resources--k8s_pod_security_policy--lifecycle--group-001.md#canonical-bca9fcb9957fd2fd221b0100c0036fe5c4d3efe82e46b18ccfeeec0c2333daee)
