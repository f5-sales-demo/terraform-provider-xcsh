---
page_title: "xcsh_forward_proxy_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_forward_proxy_policy examples."
---

# xcsh_forward_proxy_policy examples

<a id="canonical-f108bba99b58a1f0aa3abba2accf4a02569830c4a7fc769e98e7fe3de1c6625e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0c3e3a7b038b9888b176aea57ab84971904f4bcf4889d397a0d2e6c3ce89f52"></a>

## Examples — Examples / 0a7663e6c108 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- Examples

<a id="canonical-f0ff6ad2c908030c4f3f4627f90d9769848b4749a36d5a9bc76b9570d49d1436"></a>

## Complete configurations — Examples / 0a7663e6c108 / 3

- [Data source](data-sources--forward_proxy_policy--examples--group-001.md#canonical-9a658d15971c7df0c9ef1ba090d639f1ffc9fa6843a7d826dba6d2fa9ef4b541): valid configuration.

<a id="canonical-a5b2deb45089cfcc63f8ba7496602412a577c3b987a47dae53631d14791710fa"></a>

## Next pages — Examples / 0a7663e6c108 / 4

- [Data source](data-sources--forward_proxy_policy--examples--group-001.md#canonical-9a658d15971c7df0c9ef1ba090d639f1ffc9fa6843a7d826dba6d2fa9ef4b541)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-9a658d15971c7df0c9ef1ba090d639f1ffc9fa6843a7d826dba6d2fa9ef4b541"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72709dafdbef7a85b54f0847bc49125623533a262255f7c9c08ecb078fda643d"></a>

## Data source — Data source / 110b1ac7310a / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Examples](data-sources--forward_proxy_policy--examples--group-001.md#canonical-f108bba99b58a1f0aa3abba2accf4a02569830c4a7fc769e98e7fe3de1c6625e)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_forward_proxy_policy/data-source.tf`; digest `sha256:3fff12bb997c990a591800c80ace62a7790c01337e27ba68116856af1dfaf950`.

```terraform
# ForwardProxyPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ForwardProxyPolicy by name
data "xcsh_forward_proxy_policy" "example" {
  name      = "example-forward-proxy-policy"
  namespace = "staging"
}

output "forward_proxy_policy_id" {
  value = data.xcsh_forward_proxy_policy.example.id
}
```

<a id="canonical-c9264eb8e5e2ab6a24f8bd621887f538ac78002b777535dd04ed2ea38b485053"></a>

## Next pages — Data source / 110b1ac7310a / 3

- [Examples](data-sources--forward_proxy_policy--examples--group-001.md#canonical-f108bba99b58a1f0aa3abba2accf4a02569830c4a7fc769e98e7fe3de1c6625e)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
