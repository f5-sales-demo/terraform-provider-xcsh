---
page_title: "xcsh_site_registrations_by_state examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registrations_by_state examples."
---

# xcsh_site_registrations_by_state examples

<a id="canonical-f3844d2c83ae807351e728918ca3aadf7ff8666e92b43c036423d51b385ded73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-774315652f797d2fd5c93671eb040d7d35dfe1bcd46ee91982e95151674c472f"></a>

## Examples — Examples / 4a8fd0238295 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- Examples

<a id="canonical-c52cad8a750a0c5019ba1a6d43e6c8759619ad8bf26987e0c4e352c19ca10833"></a>

## Complete configurations — Examples / 4a8fd0238295 / 3

- [Data source](data-sources--site_registrations_by_state--examples--group-001.md#canonical-9a2effbc9771c0fe71fe71fae81de3973461bc3a2548051bdbbecbde59c412e6): valid configuration.

<a id="canonical-462f1cff5d24aacb6c7765517f76213ba3c504e7bbdec23df0fefc70498c3dfd"></a>

## Next pages — Examples / 4a8fd0238295 / 4

- [Data source](data-sources--site_registrations_by_state--examples--group-001.md#canonical-9a2effbc9771c0fe71fe71fae81de3973461bc3a2548051bdbbecbde59c412e6)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-9a2effbc9771c0fe71fe71fae81de3973461bc3a2548051bdbbecbde59c412e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d181d0a45167923e86dcfcde25a5e698f27a32966f020a8a906b287989d140a0"></a>

## Data source — Data source / 83a7a829972c / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Examples](data-sources--site_registrations_by_state--examples--group-001.md#canonical-f3844d2c83ae807351e728918ca3aadf7ff8666e92b43c036423d51b385ded73)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_registrations_by_state/data-source.tf`; digest `sha256:87878119dc98d70aee1fa7c3e546432844de313704e34499df4b6dfdc776ce5a`.

```terraform
# SiteRegistrationsByState DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations_by_state" "example" {
  state = "NOTSET"
}

output "site_registrations_by_state_result" {
  value = data.xcsh_site_registrations_by_state.example
}
```

<a id="canonical-4553203e64b1a4905cb1c36455907be2fed0360d18142b85d0b09495d34fccf9"></a>

## Next pages — Data source / 83a7a829972c / 3

- [Examples](data-sources--site_registrations_by_state--examples--group-001.md#canonical-f3844d2c83ae807351e728918ca3aadf7ff8666e92b43c036423d51b385ded73)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
