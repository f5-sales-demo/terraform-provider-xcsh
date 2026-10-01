---
page_title: "xcsh_ike1 examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike1 examples."
---

# xcsh_ike1 examples

<a id="canonical-defb7512b14d447bc7b57fa79136b54efef5377b69c4d47d608382ae823c9fe4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b0fec500b2ef2e42f3503d40ca42d2b696d0316c1c2c4c4b6e1a4a5460a824f"></a>

## Examples — Examples / 204b7e01d797 / 2

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md#canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7)
- Examples

<a id="canonical-db71cd89a730ce509a555fe914f1fb984669ea60a5a39d32e86e471501718125"></a>

## Complete configurations — Examples / 204b7e01d797 / 3

- [Data source](data-sources--ike1--examples--group-001.md#canonical-ca07bfe181aedfc5e4d4dd64cfba298209ecb8ac8225572cd77c009327d35470): valid configuration.

<a id="canonical-165bd1e32fb2d8a237b551911b217b34f58c851a49904fe72061199d8e597eb0"></a>

## Next pages — Examples / 204b7e01d797 / 4

- [Data source](data-sources--ike1--examples--group-001.md#canonical-ca07bfe181aedfc5e4d4dd64cfba298209ecb8ac8225572cd77c009327d35470)
- [xcsh_ike1](../data-sources/ike1.md#canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7)

<a id="canonical-ca07bfe181aedfc5e4d4dd64cfba298209ecb8ac8225572cd77c009327d35470"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e7601b614f1ed38e94708ea1846ecab1848ec4d65cb245e1c6d9795e5ac1355"></a>

## Data source — Data source / a9bb6673d8eb / 2

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md#canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7)
- [Examples](data-sources--ike1--examples--group-001.md#canonical-defb7512b14d447bc7b57fa79136b54efef5377b69c4d47d608382ae823c9fe4)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_ike1/data-source.tf`; digest `sha256:1f5aa2a482784f4287f2518ad47fd88a6918d0ecdd01b6d56a7fd34d407e402f`.

```terraform
# Ike1 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Ike1 by name
data "xcsh_ike1" "example" {
  name      = "example-ike1"
  namespace = "staging"
}

output "ike1_id" {
  value = data.xcsh_ike1.example.id
}
```

<a id="canonical-653e0952638b4c3a63d8caca24e7ae307e225901fb46b7fc84c261ab443bb262"></a>

## Next pages — Data source / a9bb6673d8eb / 3

- [Examples](data-sources--ike1--examples--group-001.md#canonical-defb7512b14d447bc7b57fa79136b54efef5377b69c4d47d608382ae823c9fe4)
- [xcsh_ike1](../data-sources/ike1.md#canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7)
