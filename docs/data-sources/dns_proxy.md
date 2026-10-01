---
page_title: "xcsh_dns_proxy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_proxy landing."
---

# xcsh_dns_proxy landing

<a id="canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9f1d2d2431591adacaaf1ab7ab10116848e55a601813add65eb8cd323bff8d9"></a>

## xcsh_dns_proxy — xcsh_dns_proxy / be8b851e479e / 2

Breadcrumbs:

- xcsh_dns_proxy

Manages DNS Proxy in a given namespace. If one already exists it will give an error in F5
Distributed Cloud.

<a id="canonical-4d6a74336fa3ffd12a9783c7610906eb49fae7bf69afe15cbf4a78a473999b0f"></a>

## Prerequisites — xcsh_dns_proxy / be8b851e479e / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-10f97d29bb82504ed45092aaf4b0a9b2ee5489439e474099264e69b94512e4b8"></a>

## Minimal configuration — xcsh_dns_proxy / be8b851e479e / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSProxy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSProxy by name
data "xcsh_dns_proxy" "example" {
  name      = "example-dns-proxy"
  namespace = "system"
}

output "dns_proxy_id" {
  value = data.xcsh_dns_proxy.example.id
}
```

<a id="canonical-be56d19d4301360ad772ae7cc5afc0cd7129558869ea3e1971bc72162c5a22f3"></a>

## Root configuration — xcsh_dns_proxy / be8b851e479e / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-51a62addf2240134bc3c1518ff298dd158e529d579a3377628f7d09dc7ad699c"></a>

## Next pages — xcsh_dns_proxy / be8b851e479e / 6

- [Property reference](../guides/data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [Examples](../guides/data-sources--dns_proxy--examples--group-001.md#canonical-d78830af6a6ac5d2c3f1d06ae4583b8e36443c2efd4148677695d0e063893f86)
