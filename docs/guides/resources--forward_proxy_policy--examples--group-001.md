---
page_title: "xcsh_forward_proxy_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_forward_proxy_policy examples."
---

# xcsh_forward_proxy_policy examples

<a id="canonical-2030102031112111-2133212023230213-0100322013131112-3131222303110302-1133313322002110-2320102233123123-2003020020003131-1100313123230120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- Examples

<a id="canonical-3103132112203300-2123311210300012-3320022013231222-1210312230322202-3001213233002110-1231312202223300-0120003233331000-3333312021013331"></a>

### Complete configurations for `xcsh_forward_proxy_policy`

- [Resource](resources--forward_proxy_policy--examples--group-001.md#canonical-1333002001312121-0210033011100310-2331003122231131-2003002003022312-3113033121200010-1120120033200012-0021311013010101-2303222112331101): valid configuration.

<a id="canonical-1333002001312121-0210033011100310-2331003122231131-2003002003022312-3113033121200010-1120120033200012-0021311013010101-2303222112331101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Examples](resources--forward_proxy_policy--examples--group-001.md#canonical-2030102031112111-2133212023230213-0100322013131112-3131222303110302-1133313322002110-2320102233123123-2003020020003131-1100313123230120)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_forward_proxy_policy/resource.tf`; digest `sha256:bcc594a880ce1beadc720b2bd1c76066016eb6d5489e15cf1b1295e1e8d40584`.

```terraform
# ForwardProxyPolicy Resource Example
# Manages a Forward Proxy Policy resource in F5 Distributed Cloud for forward proxy policy specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ForwardProxyPolicy configuration
resource "xcsh_forward_proxy_policy" "example" {
  name      = "example-forward-proxy-policy"
  namespace = "staging"
}
```
