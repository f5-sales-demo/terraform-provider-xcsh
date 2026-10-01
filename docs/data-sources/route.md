---
page_title: "xcsh_route landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route landing."
---

# xcsh_route landing

<a id="canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2c4b71f115b18f8085b80486d432f3437c17a97e108524f007455225517cbcd"></a>

## xcsh_route — xcsh_route / 15d657738e87 / 2

Breadcrumbs:

- xcsh_route

Manages route object in a given namespace. Route object is list of route rules. Each rule has match
condition to match incoming requests and actions to take on matching requests in F5 Distributed
Cloud.

<a id="canonical-790bdaad8fd41950c02e18a353ae650353e1e467ff57e2b1f435612ac26a178c"></a>

## Prerequisites — xcsh_route / 15d657738e87 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-383d0650b4eca485f3f5f69d027d19a940632d57c31633019af6067662d22362"></a>

## Minimal configuration — xcsh_route / 15d657738e87 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Route Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Route by name
data "xcsh_route" "example" {
  name      = "example-route"
  namespace = "staging"
}

output "route_id" {
  value = data.xcsh_route.example.id
}
```

<a id="canonical-02a5da6357b480fb0b5feb9f86ca7a54bbf19d8103b8cf8582e821696e4300e2"></a>

## Root configuration — xcsh_route / 15d657738e87 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-ff67fb0a828e7fa7b6dc3b1480feaec5370b81854a4179a51fc9737bb65c9af4"></a>

## Next pages — xcsh_route / 15d657738e87 / 6

- [Property reference](../guides/data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [Examples](../guides/data-sources--route--examples--group-001.md#canonical-9391d20cfe41560e191bcd91f8e4b8422c2d0ad54f96e4235a550fc4f43c4447)
