---
page_title: "xcsh_subnet examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_subnet examples."
---

# xcsh_subnet examples

<a id="canonical-79d8e319da7c705c323908e69b3050b9dbfa8c2f97c84a0b8f524d43826ab865"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36dbadc8c5681442e6555a8422c7ea4cd76755d6a6095fe07f3edddd7bbafec6"></a>

## Examples — Examples / 3129f84e31c4 / 2

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)
- Examples

<a id="canonical-cdd4b46a684a6d28f3b30ee3299ac26a7b3e7c83f79e50f666b25f48402e4df0"></a>

## Complete configurations — Examples / 3129f84e31c4 / 3

- [Data source](data-sources--subnet--examples--group-001.md#canonical-989fb206ead55a15c640db062b7222260cdd55581c4955b8894ecf1ba746a81a): valid configuration.

<a id="canonical-d30bf62c63e065ddaa2b6913d99011d6d33c39e0ddc2197609c1e96ea49cab62"></a>

## Next pages — Examples / 3129f84e31c4 / 4

- [Data source](data-sources--subnet--examples--group-001.md#canonical-989fb206ead55a15c640db062b7222260cdd55581c4955b8894ecf1ba746a81a)
- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)

<a id="canonical-989fb206ead55a15c640db062b7222260cdd55581c4955b8894ecf1ba746a81a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fae685966ee87f4c875f033590aed874ea27056a0d8def0cef4fa5ae4ce1ac36"></a>

## Data source — Data source / 39266f9e0f4e / 2

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)
- [Examples](data-sources--subnet--examples--group-001.md#canonical-79d8e319da7c705c323908e69b3050b9dbfa8c2f97c84a0b8f524d43826ab865)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_subnet/data-source.tf`; digest `sha256:1dd541aa45894ccb54e6cdac49fe4ca6130605728a4cc127960dcf649a0838dd`.

```terraform
# Subnet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Subnet by name
data "xcsh_subnet" "example" {
  name      = "example-subnet"
  namespace = "staging"
}

output "subnet_id" {
  value = data.xcsh_subnet.example.id
}
```

<a id="canonical-54fbfbe15625e4f43971fb10e173ca695ad1dc6baa4bd09a3e124eaf2b6684c5"></a>

## Next pages — Data source / 39266f9e0f4e / 3

- [Examples](data-sources--subnet--examples--group-001.md#canonical-79d8e319da7c705c323908e69b3050b9dbfa8c2f97c84a0b8f524d43826ab865)
- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)
