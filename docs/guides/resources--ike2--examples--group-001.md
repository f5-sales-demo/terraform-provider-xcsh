---
page_title: "xcsh_ike2 examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike2 examples."
---

# xcsh_ike2 examples

<a id="canonical-0111201230203323-2121321302031333-0133222302222032-3320123331001032-3023312131303101-0131233020020330-3233122333320210-0213113213030103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md#canonical-3103220130122111-0023320210302011-3122102311032022-0223033112130223-1201210032310333-1323122321230203-0200202200211032-1120020123001232)
- Examples

<a id="canonical-1200210123121121-1321333330132221-3000010320312020-3221200312300002-2030211210300332-2200012110321222-2211030200020133-3013313313103331"></a>

### Complete configurations for `xcsh_ike2`

- [Resource](resources--ike2--examples--group-001.md#canonical-3211023310301031-0330202110310232-1122212202130031-3311003323310133-1320332122302023-0213113311301321-1312233300101022-1122221312211022): valid configuration.

<a id="canonical-3211023310301031-0330202110310232-1122212202130031-3311003323310133-1320332122302023-0213113311301321-1312233300101022-1122221312211022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md#canonical-3103220130122111-0023320210302011-3122102311032022-0223033112130223-1201210032310333-1323122321230203-0200202200211032-1120020123001232)
- [Examples](resources--ike2--examples--group-001.md#canonical-0111201230203323-2121321302031333-0133222302222032-3320123331001032-3023312131303101-0131233020020330-3233122333320210-0213113213030103)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_ike2/resource.tf`; digest `sha256:9af0540e9ef3cddbd2cf33f981bb394c77fd860b251d272dd666ef02cd5bc017`.

```terraform
# Ike2 Resource Example
# Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Ike2 configuration
resource "xcsh_ike2" "example" {
  name      = "example-ike2"
  namespace = "staging"
}
```
