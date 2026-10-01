---
page_title: "xcsh_service_policy_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_set examples."
---

# xcsh_service_policy_set examples

<a id="canonical-64a0cc77566504934e629646dfc340565e8343844ecc4a17ecf571158b2ff726"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1ae93a13e931232345274b32c54900a2eb64c0f3083d996fc38b749800852ae"></a>

## Examples — Examples / ecd193b75234 / 2

Breadcrumbs:

- [xcsh_service_policy_set](../data-sources/service_policy_set.md#canonical-52eaa9a13b0184a72441569b070040914137b030f2dff5fcdbb25301ba8d99b6)
- Examples

<a id="canonical-4d01de651c16007e3bc376b5eacb5c826226dace7c421e88f99a3968a1af7328"></a>

## Complete configurations — Examples / ecd193b75234 / 3

- [Data source](data-sources--service_policy_set--examples--group-001.md#canonical-3c4b4b1a6b8c968d232e0c8d376c1a53ad031d05b32d72c8890e5d3c7f39adb2): valid configuration.

<a id="canonical-fef60f4ed21feb87c922b2e5a83f4404e14afafd62dc6e19a7e7a015fbdffe7c"></a>

## Next pages — Examples / ecd193b75234 / 4

- [Data source](data-sources--service_policy_set--examples--group-001.md#canonical-3c4b4b1a6b8c968d232e0c8d376c1a53ad031d05b32d72c8890e5d3c7f39adb2)
- [xcsh_service_policy_set](../data-sources/service_policy_set.md#canonical-52eaa9a13b0184a72441569b070040914137b030f2dff5fcdbb25301ba8d99b6)

<a id="canonical-3c4b4b1a6b8c968d232e0c8d376c1a53ad031d05b32d72c8890e5d3c7f39adb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57e7064f783d168b941c094ccb4b64fe996dd0f25643c186bef9ecf2a23fdccf"></a>

## Data source — Data source / 3406c9e03159 / 2

Breadcrumbs:

- [xcsh_service_policy_set](../data-sources/service_policy_set.md#canonical-52eaa9a13b0184a72441569b070040914137b030f2dff5fcdbb25301ba8d99b6)
- [Examples](data-sources--service_policy_set--examples--group-001.md#canonical-64a0cc77566504934e629646dfc340565e8343844ecc4a17ecf571158b2ff726)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_service_policy_set/data-source.tf`; digest `sha256:3bd6ef68376941c1abe8ba51cff222b7a95e9a7e72218b34827623539153d92a`.

```terraform
# ServicePolicySet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicySet by name
data "xcsh_service_policy_set" "example" {
  name      = "example-service-policy-set"
  namespace = "staging"
}

output "service_policy_set_id" {
  value = data.xcsh_service_policy_set.example.id
}
```

<a id="canonical-827e222fc76413d6ddd64272ad157446c611b93bf85b6a0d53acd1c5803c941c"></a>

## Next pages — Data source / 3406c9e03159 / 3

- [Examples](data-sources--service_policy_set--examples--group-001.md#canonical-64a0cc77566504934e629646dfc340565e8343844ecc4a17ecf571158b2ff726)
- [xcsh_service_policy_set](../data-sources/service_policy_set.md#canonical-52eaa9a13b0184a72441569b070040914137b030f2dff5fcdbb25301ba8d99b6)
