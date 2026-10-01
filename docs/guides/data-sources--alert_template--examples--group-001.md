---
page_title: "xcsh_alert_template examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_template examples."
---

# xcsh_alert_template examples

<a id="canonical-e7c73e3ee522555b1d0dfc331cfd8a97f27c510ca20b448538b9f423afd2f12d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24d6dee1567c9a944b9a0c4f857ecfd21e3bdddddfec4a7793ef66e645e4934c"></a>

## Examples — Examples / e3e9c88a4b88 / 2

Breadcrumbs:

- [xcsh_alert_template](../data-sources/alert_template.md#canonical-43748ee092eb77be7c417e05e7d416bc1aead9dc6699db467e2027d12e27cf88)
- Examples

<a id="canonical-7cfd080713746feb92f9ce210432bd1df336233212bab0c4019cc9b83a4c6cec"></a>

## Complete configurations — Examples / e3e9c88a4b88 / 3

- [Data source](data-sources--alert_template--examples--group-001.md#canonical-9bf5075faa2a5ea4d5421b9ab80cfe8a5c060ce79e29c8d06a539ad5d0b2fa81): valid configuration.

<a id="canonical-d8dfb460b00341b5c2b759432770afdb7f2545814a6a75dbe657a2f371a12065"></a>

## Next pages — Examples / e3e9c88a4b88 / 4

- [Data source](data-sources--alert_template--examples--group-001.md#canonical-9bf5075faa2a5ea4d5421b9ab80cfe8a5c060ce79e29c8d06a539ad5d0b2fa81)
- [xcsh_alert_template](../data-sources/alert_template.md#canonical-43748ee092eb77be7c417e05e7d416bc1aead9dc6699db467e2027d12e27cf88)

<a id="canonical-9bf5075faa2a5ea4d5421b9ab80cfe8a5c060ce79e29c8d06a539ad5d0b2fa81"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c932f34ffabff5cc5b5021ddc0e2b8b34fea92e83faa10e79ed878fb9ef39580"></a>

## Data source — Data source / f55e7c6fb88c / 2

Breadcrumbs:

- [xcsh_alert_template](../data-sources/alert_template.md#canonical-43748ee092eb77be7c417e05e7d416bc1aead9dc6699db467e2027d12e27cf88)
- [Examples](data-sources--alert_template--examples--group-001.md#canonical-e7c73e3ee522555b1d0dfc331cfd8a97f27c510ca20b448538b9f423afd2f12d)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_alert_template/data-source.tf`; digest `sha256:38bcb6351cb774c27b9a7119ac259c1eaff5879d005e5da296e0eb47b5f478f7`.

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

<a id="canonical-bfaca92bf9b61cf03441ddaa65924e249ef918e7cec0ca3377b643a13d902a87"></a>

## Next pages — Data source / f55e7c6fb88c / 3

- [Examples](data-sources--alert_template--examples--group-001.md#canonical-e7c73e3ee522555b1d0dfc331cfd8a97f27c510ca20b448538b9f423afd2f12d)
- [xcsh_alert_template](../data-sources/alert_template.md#canonical-43748ee092eb77be7c417e05e7d416bc1aead9dc6699db467e2027d12e27cf88)
