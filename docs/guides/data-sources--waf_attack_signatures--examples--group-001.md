---
page_title: "xcsh_waf_attack_signatures examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_attack_signatures examples."
---

# xcsh_waf_attack_signatures examples

<a id="canonical-406df1b7f29a187b4eac690cc0fb9071ea95aff78f5a777330f7d2afc40cd572"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fcfea70c78dc5494e74c16e70cf36c0ebdb4b5e3eaa80c22ef5104f9447eff59"></a>

## Examples — Examples / 172f14abb9d3 / 2

Breadcrumbs:

- [xcsh_waf_attack_signatures](../data-sources/waf_attack_signatures.md#canonical-4263664dad62728ed7ee4477dbb626c200a26b4fc673ffeb3c69ebfb8fec8479)
- Examples

<a id="canonical-072a8704d408e504c499b32926b3e17e41c1f3eed78a9c3eed9aa3eebd0757c1"></a>

## Complete configurations — Examples / 172f14abb9d3 / 3

- [Data source](data-sources--waf_attack_signatures--examples--group-001.md#canonical-b5e950da7c1a6bcbab06ee6dc99e9104cec71a1a14bc0b897af35ec1d5cb6ea7): valid configuration.

<a id="canonical-be517cc1530b5d8d8b68db5ee00a83924dc20b4e7eef3e468bc201223f5989bb"></a>

## Next pages — Examples / 172f14abb9d3 / 4

- [Data source](data-sources--waf_attack_signatures--examples--group-001.md#canonical-b5e950da7c1a6bcbab06ee6dc99e9104cec71a1a14bc0b897af35ec1d5cb6ea7)
- [xcsh_waf_attack_signatures](../data-sources/waf_attack_signatures.md#canonical-4263664dad62728ed7ee4477dbb626c200a26b4fc673ffeb3c69ebfb8fec8479)

<a id="canonical-b5e950da7c1a6bcbab06ee6dc99e9104cec71a1a14bc0b897af35ec1d5cb6ea7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-775e32ec70c1cc8f1858120e7232bf105246b3f0e656ee90b3ef209e3cc50eb6"></a>

## Data source — Data source / 100204803f70 / 2

Breadcrumbs:

- [xcsh_waf_attack_signatures](../data-sources/waf_attack_signatures.md#canonical-4263664dad62728ed7ee4477dbb626c200a26b4fc673ffeb3c69ebfb8fec8479)
- [Examples](data-sources--waf_attack_signatures--examples--group-001.md#canonical-406df1b7f29a187b4eac690cc0fb9071ea95aff78f5a777330f7d2afc40cd572)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_waf_attack_signatures/data-source.tf`; digest `sha256:5857e356bd3dea3a26e0c6c17fd28eb9757681e0e4bfa5915066871a8e38e467`.

```terraform
# WAFAttackSignatures DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_attack_signatures" "example" {
}

output "waf_attack_signatures_result" {
  value = data.xcsh_waf_attack_signatures.example
}
```

<a id="canonical-fdccf469728d037139ff5770af623dae706a6d8e28060a1fbec5f4f60b17a2a8"></a>

## Next pages — Data source / 100204803f70 / 3

- [Examples](data-sources--waf_attack_signatures--examples--group-001.md#canonical-406df1b7f29a187b4eac690cc0fb9071ea95aff78f5a777330f7d2afc40cd572)
- [xcsh_waf_attack_signatures](../data-sources/waf_attack_signatures.md#canonical-4263664dad62728ed7ee4477dbb626c200a26b4fc673ffeb3c69ebfb8fec8479)
