---
page_title: "xcsh_subnet examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_subnet examples."
---

# xcsh_subnet examples

<a id="canonical-1321312032030121-3122133013001130-0302032100203212-2123030011002321-3123332220300233-2113302010220023-2033110210311003-2002122223201211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-2030233332213333-3213211200222203-3230133323032230-1012203021120130-0010113332012110-2331021111131003-0132121132331301-3312323301120210)
- Examples

<a id="canonical-0312312322313020-3011122001101002-3212111111222010-0202301332221030-3113121311113112-2212002111333200-1333033231313131-1323232233323012"></a>

### Complete configurations for `xcsh_subnet`

- [Data source](data-sources--subnet--examples--group-001.md#canonical-2120213323020012-3222311111220111-3012100031230012-0223130202020212-0030313111111120-0130102111112320-2021103230330123-2213101222200122): valid configuration.

<a id="canonical-2120213323020012-3222311111220111-3012100031230012-0223130202020212-0030313111111120-0130102111112320-2021103230330123-2213101222200122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-2030233332213333-3213211200222203-3230133323032230-1012203021120130-0010113332012110-2331021111131003-0132121132331301-3312323301120210)
- [Examples](data-sources--subnet--examples--group-001.md#canonical-1321312032030121-3122133013001130-0302032100203212-2123030011002321-3123332220300233-2113302010220023-2033110210311003-2002122223201211)
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
