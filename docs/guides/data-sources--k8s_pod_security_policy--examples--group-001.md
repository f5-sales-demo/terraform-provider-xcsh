---
page_title: "xcsh_k8s_pod_security_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_policy examples."
---

# xcsh_k8s_pod_security_policy examples

<a id="canonical-94881518bbc37debae300d89fbdc06722ac20767c40925cad906ca38d7b59aa1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7afa0925cd63c703ce5486a827ad40b50b8bdec171654c5f2934c2782c4c3249"></a>

## Examples — Examples / b52aad1aee66 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- Examples

<a id="canonical-2f0509a00f86794a62b8e3b483fb76f1106ffe1eebc2a8558617d71fd7d094bf"></a>

## Complete configurations — Examples / b52aad1aee66 / 3

- [Data source](data-sources--k8s_pod_security_policy--examples--group-001.md#canonical-eebf55ff7166f8d775f10a2af9b20f8abcba75282d31d27ff0c560e1173a681a): valid configuration.

<a id="canonical-a159808c3ed1a1181720d042cb82cdf74537b57de4a240e6acb724b5033ea983"></a>

## Next pages — Examples / b52aad1aee66 / 4

- [Data source](data-sources--k8s_pod_security_policy--examples--group-001.md#canonical-eebf55ff7166f8d775f10a2af9b20f8abcba75282d31d27ff0c560e1173a681a)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-eebf55ff7166f8d775f10a2af9b20f8abcba75282d31d27ff0c560e1173a681a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71f30bd15a439ba9ef17d612ab94de58e34af9e80c963556e76081d410e788aa"></a>

## Data source — Data source / 7b5ec18ff5cd / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Examples](data-sources--k8s_pod_security_policy--examples--group-001.md#canonical-94881518bbc37debae300d89fbdc06722ac20767c40925cad906ca38d7b59aa1)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_k8s_pod_security_policy/data-source.tf`; digest `sha256:466f0cfc6eb4d12538a2c33e13b91f9270feaa389452e6596f117ab99a547c9b`.

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

<a id="canonical-746cb562f26343571df5f4489302395dcc5f21485486d5eae115b85c4f37d967"></a>

## Next pages — Data source / 7b5ec18ff5cd / 3

- [Examples](data-sources--k8s_pod_security_policy--examples--group-001.md#canonical-94881518bbc37debae300d89fbdc06722ac20767c40925cad906ca38d7b59aa1)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
