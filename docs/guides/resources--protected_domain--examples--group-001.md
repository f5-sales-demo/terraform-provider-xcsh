---
page_title: "xcsh_protected_domain examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_domain examples."
---

# xcsh_protected_domain examples

<a id="canonical-1031221233100201-2313210011123002-1020031220131123-3322003310200131-1011111321033102-1321001001033030-1200120331033032-2002012333003231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_protected_domain](../resources/protected_domain.md#canonical-3133012330313310-3233013213001311-1212311200333023-2121112332020022-2232211202113213-1220123021122333-3111220011333221-3312301231323110)
- Examples

<a id="canonical-0312300001100102-0013002203233032-0002030221131333-2011221100323020-3121100330330210-3122320330031202-1030313213322223-2102133302312001"></a>

### Complete configurations for `xcsh_protected_domain`

- [Resource](resources--protected_domain--examples--group-001.md#canonical-2233220302022313-3333220310330113-3201111301233032-0312132232211230-2012203102002121-2222212020020301-3220020221200310-0133203022130133): valid configuration.

<a id="canonical-2233220302022313-3333220310330113-3201111301233032-0312132232211230-2012203102002121-2222212020020301-3220020221200310-0133203022130133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_protected_domain](../resources/protected_domain.md#canonical-3133012330313310-3233013213001311-1212311200333023-2121112332020022-2232211202113213-1220123021122333-3111220011333221-3312301231323110)
- [Examples](resources--protected_domain--examples--group-001.md#canonical-1031221233100201-2313210011123002-1020031220131123-3322003310200131-1011111321033102-1321001001033030-1200120331033032-2002012333003231)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_protected_domain/resource.tf`; digest `sha256:a41646b64f77c1c10c9bfb255c8aeb3c51bbd0f1f5060b667473b425512ade92`.

```terraform
# ProtectedDomain Resource Example
# Manages Domain to protect in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtectedDomain configuration
resource "xcsh_protected_domain" "example" {
  name      = "example-protected-domain"
  namespace = "staging"

  protected_domain = "example.com"
}
```
