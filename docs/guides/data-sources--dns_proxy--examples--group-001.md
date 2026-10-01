---
page_title: "xcsh_dns_proxy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_proxy examples."
---

# xcsh_dns_proxy examples

<a id="canonical-d78830af6a6ac5d2c3f1d06ae4583b8e36443c2efd4148677695d0e063893f86"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf209918f54f6413fe7b8b64a0124b688f04510f6351b300c8e9f6a3b3abb224"></a>

## Examples — Examples / 4485fa6fa50e / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- Examples

<a id="canonical-b9bd1b3e5b62c839aaa7c92edae6bcf17c40e26c756be0798c891f45110943e2"></a>

## Complete configurations — Examples / 4485fa6fa50e / 3

- [Data source](data-sources--dns_proxy--examples--group-001.md#canonical-ee5ffaa9b42993d06ea8e787c97a294a934f1a53a918a22c3ed39001fa78a4d8): valid configuration.

<a id="canonical-1e1751f03917659aecfd23af4fb685984836f90b1d6b5f9cd541858bb99ca294"></a>

## Next pages — Examples / 4485fa6fa50e / 4

- [Data source](data-sources--dns_proxy--examples--group-001.md#canonical-ee5ffaa9b42993d06ea8e787c97a294a934f1a53a918a22c3ed39001fa78a4d8)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-ee5ffaa9b42993d06ea8e787c97a294a934f1a53a918a22c3ed39001fa78a4d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41b9091ed98a4a10c6e02312fa91a258b9f0d5afc20e06ae95726ce6895be300"></a>

## Data source — Data source / 9eddb18f9098 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Examples](data-sources--dns_proxy--examples--group-001.md#canonical-d78830af6a6ac5d2c3f1d06ae4583b8e36443c2efd4148677695d0e063893f86)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_proxy/data-source.tf`; digest `sha256:8864854e827d63992024dc1c04e3c3a45871e2f9bca0bf3baae302eae31eb760`.

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

<a id="canonical-fc46c5bd2f783c26aeef2491ed50ef9d2891b7901cf18f31984bb0404a7a300f"></a>

## Next pages — Data source / 9eddb18f9098 / 3

- [Examples](data-sources--dns_proxy--examples--group-001.md#canonical-d78830af6a6ac5d2c3f1d06ae4583b8e36443c2efd4148677695d0e063893f86)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
