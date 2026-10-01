---
page_title: "xcsh_waf_latest_signatures_version examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_latest_signatures_version examples."
---

# xcsh_waf_latest_signatures_version examples

<a id="canonical-9dfaa27c323ff0ee73075c8652c72ac3d5ff126248800053866038ab4e0686d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b45a0f4af724bf182d990952ada2757131f6da532af1b3bb04a9326686c6d785"></a>

## Examples — Examples / 98a49412892e / 2

Breadcrumbs:

- [xcsh_waf_latest_signatures_version](../data-sources/waf_latest_signatures_version.md#canonical-e82588039d40a0bd7b98f558c6e993e8c040b7d365e0b5c93ac34a46905a3b8c)
- Examples

<a id="canonical-177e2a1fbb9266c3b9c6865080356e4a873d25e297c6f4eece8be71d2da23c2c"></a>

## Complete configurations — Examples / 98a49412892e / 3

- [Data source](data-sources--waf_latest_signatures_version--examples--group-001.md#canonical-facbb24541ce9194427b933d14a868a6f6bda491aa85ea664d413f1f3cc3d2d1): valid configuration.

<a id="canonical-68ab5726e306969d14ef28f4df3256b864cb8ccdd5a1c4d6f109530a13521f98"></a>

## Next pages — Examples / 98a49412892e / 4

- [Data source](data-sources--waf_latest_signatures_version--examples--group-001.md#canonical-facbb24541ce9194427b933d14a868a6f6bda491aa85ea664d413f1f3cc3d2d1)
- [xcsh_waf_latest_signatures_version](../data-sources/waf_latest_signatures_version.md#canonical-e82588039d40a0bd7b98f558c6e993e8c040b7d365e0b5c93ac34a46905a3b8c)

<a id="canonical-facbb24541ce9194427b933d14a868a6f6bda491aa85ea664d413f1f3cc3d2d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07bba6e4cbe8b99e11b6d3fa4befc38b1273469875bf4ac7ad215c5914ea4a51"></a>

## Data source — Data source / 8fa001c0fd3c / 2

Breadcrumbs:

- [xcsh_waf_latest_signatures_version](../data-sources/waf_latest_signatures_version.md#canonical-e82588039d40a0bd7b98f558c6e993e8c040b7d365e0b5c93ac34a46905a3b8c)
- [Examples](data-sources--waf_latest_signatures_version--examples--group-001.md#canonical-9dfaa27c323ff0ee73075c8652c72ac3d5ff126248800053866038ab4e0686d4)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_waf_latest_signatures_version/data-source.tf`; digest `sha256:6a87350aef0ac879fd92c4cae37ff6595f45e1a414503ebb0313bae760a73e29`.

```terraform
# WAFLatestSignaturesVersion DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_latest_signatures_version" "example" {
}

output "waf_latest_signatures_version_result" {
  value = data.xcsh_waf_latest_signatures_version.example
}
```

<a id="canonical-354404c380e7e9062e7c7640d7ce79e6ff11e59dd65380ee2a561a60d5fe180c"></a>

## Next pages — Data source / 8fa001c0fd3c / 3

- [Examples](data-sources--waf_latest_signatures_version--examples--group-001.md#canonical-9dfaa27c323ff0ee73075c8652c72ac3d5ff126248800053866038ab4e0686d4)
- [xcsh_waf_latest_signatures_version](../data-sources/waf_latest_signatures_version.md#canonical-e82588039d40a0bd7b98f558c6e993e8c040b7d365e0b5c93ac34a46905a3b8c)
