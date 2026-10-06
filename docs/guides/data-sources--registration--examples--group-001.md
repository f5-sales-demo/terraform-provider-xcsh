---
page_title: "xcsh_registration examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_registration examples."
---

# xcsh_registration examples

<a id="canonical-0300203131021300-1230011003310011-0211010322121322-1122310000121231-2202100310321000-2020031132033031-3220322213011231-2031002333213313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- Examples

<a id="canonical-1010023213121120-3223011332223200-2030033321101321-3301331011200123-3011211101110013-2012011023100201-0313133031333011-0301303120303322"></a>

### Complete configurations for `xcsh_registration`

- [Data source](data-sources--registration--examples--group-001.md#canonical-3002131203233232-1020331232213323-2011131113233312-3303311323331302-3133132013002211-2233013223212112-0310212021011031-1101133030233022): valid configuration.

<a id="canonical-3002131203233232-1020331232213323-2011131113233312-3303311323331302-3133132013002211-2233013223212112-0310212021011031-1101133030233022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Examples](data-sources--registration--examples--group-001.md#canonical-0300203131021300-1230011003310011-0211010322121322-1122310000121231-2202100310321000-2020031132033031-3220322213011231-2031002333213313)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_registration/data-source.tf`; digest `sha256:887504fe79e23b9f0b1ae62f2ad6a04c459d805004c72b40be81620b6a16261f`.

```terraform
# Registration Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Registration by name
data "xcsh_registration" "example" {
  name      = "example-registration"
  namespace = "staging"
}

output "registration_id" {
  value = data.xcsh_registration.example.id
}
```
