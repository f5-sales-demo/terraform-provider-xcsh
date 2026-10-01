---
page_title: "xcsh_alert_template landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_template landing."
---

# xcsh_alert_template landing

<a id="canonical-43748ee092eb77be7c417e05e7d416bc1aead9dc6699db467e2027d12e27cf88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-39682ebad3e1683cf7304d22f9efa64d295c9b8b931c6119b801ee776c6977a1"></a>

## xcsh_alert_template — xcsh_alert_template / f3d006fb9aad / 2

Breadcrumbs:

- xcsh_alert_template

Manages Domain to protect in F5 Distributed Cloud.

<a id="canonical-66b735a55a397eb507f6434f96a9865913dc46e71e2e45149026d4e75f7ae651"></a>

## Prerequisites — xcsh_alert_template / f3d006fb9aad / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-d78d0d34826c9d859b9d09c8381cffc9b96235d5070e6ceb16e599d19fd5667c"></a>

## Minimal configuration — xcsh_alert_template / f3d006fb9aad / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AlertTemplate Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertTemplate by name
data "xcsh_alert_template" "example" {
  name      = "example-alert-template"
  namespace = "staging"
}

output "alert_template_id" {
  value = data.xcsh_alert_template.example.id
}
```

<a id="canonical-d566aca7ae185bb8b7c27a8c972474c1226239542589073a66a9f59446a2351b"></a>

## Root configuration — xcsh_alert_template / f3d006fb9aad / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-bd277cd842c23d37ae077ecb78b2080e1a0327c2fa8cd6d47702c6bdbd767cee"></a>

## Next pages — xcsh_alert_template / f3d006fb9aad / 6

- [Property reference](../guides/data-sources--alert_template--reference--group-001.md#canonical-e6acd3bb08b31083263c203b17dd9b602b533a79339a12f5c11f80d073103560)
- [Examples](../guides/data-sources--alert_template--examples--group-001.md#canonical-e7c73e3ee522555b1d0dfc331cfd8a97f27c510ca20b448538b9f423afd2f12d)
