---
page_title: "xcsh_k8s_pod_security_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_policy examples."
---

# xcsh_k8s_pod_security_policy examples

<a id="canonical-ffe90ec21ce5cd40a9ee9e2ea88847034ff59a568bd6d1da400ed3e4c87cde08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-578f3dc5ff1eef1c00de735c8dc982dbbbbec7b7baa54bb99afc31b70d0411b1"></a>

## Examples — Examples / dc19b48542a1 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- Examples

<a id="canonical-0b1cab4fae1dc1ecde485e96c0ed7066fd1b60c92c4c11c78e5afca4798ec66c"></a>

## Complete configurations — Examples / dc19b48542a1 / 3

- [Resource](resources--k8s_pod_security_policy--examples--group-001.md#canonical-0166d380ff4021453e880714e8fb09bfc0f16278caa39ce6b987b03dd710e2ef): valid configuration.

<a id="canonical-b9ffc1bffd9a8219d0c85f4a812a262702d3a796eb066f55257a0c800c47ef64"></a>

## Next pages — Examples / dc19b48542a1 / 4

- [Resource](resources--k8s_pod_security_policy--examples--group-001.md#canonical-0166d380ff4021453e880714e8fb09bfc0f16278caa39ce6b987b03dd710e2ef)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-0166d380ff4021453e880714e8fb09bfc0f16278caa39ce6b987b03dd710e2ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d211a85682a51be40c78571280e145667f0dffc676fe5a9b577781a397b13e1a"></a>

## Resource — Resource / bbfeb75e66c3 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Examples](resources--k8s_pod_security_policy--examples--group-001.md#canonical-ffe90ec21ce5cd40a9ee9e2ea88847034ff59a568bd6d1da400ed3e4c87cde08)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_k8s_pod_security_policy/resource.tf`; digest `sha256:e574c58014b7e9d4bf5b442007d145c5a6006853ecb4a1aebacdb4b005ac5fa9`.

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

<a id="canonical-d755d7f1d3c7bed74df76a36827b1e833cc5f50ffb504b2d152a68c1344a5157"></a>

## Next pages — Resource / bbfeb75e66c3 / 3

- [Examples](resources--k8s_pod_security_policy--examples--group-001.md#canonical-ffe90ec21ce5cd40a9ee9e2ea88847034ff59a568bd6d1da400ed3e4c87cde08)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
