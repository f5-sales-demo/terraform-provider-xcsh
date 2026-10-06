---
page_title: "xcsh_proxy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy examples."
---

# xcsh_proxy examples

<a id="canonical-3101102230333321-3121202112111310-2102101121320321-1110131122110221-3333232113210233-3320003332133013-3212300101003213-0012020323223133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- Examples

<a id="canonical-2101111300022030-2032123022213231-1111021110201031-2002000122331321-2313000111222001-2211301330103120-1301210101131201-1320012213123212"></a>

### Complete configurations for `xcsh_proxy`

- [Data source](data-sources--proxy--examples--group-001.md#canonical-3011233233011100-1322002211231232-3311121212313310-1330022232310010-1231033133323023-2233122330313330-3203030201132033-1021101202221100): valid configuration.

<a id="canonical-3011233233011100-1322002211231232-3311121212313310-1330022232310010-1231033133323023-2233122330313330-3203030201132033-1021101202221100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Examples](data-sources--proxy--examples--group-001.md#canonical-3101102230333321-3121202112111310-2102101121320321-1110131122110221-3333232113210233-3320003332133013-3212300101003213-0012020323223133)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_proxy/data-source.tf`; digest `sha256:d2bc93690268dd4557ac75b114b1e4c362c8f8d843db7cc8d8dd1623ef5dbfe6`.

```terraform
# Proxy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Proxy by name
data "xcsh_proxy" "example" {
  name      = "example-proxy"
  namespace = "staging"
}

output "proxy_id" {
  value = data.xcsh_proxy.example.id
}
```
