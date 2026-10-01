---
page_title: "xcsh_bgp_asn_set landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_asn_set landing."
---

# xcsh_bgp_asn_set landing

<a id="canonical-117a5d23afbac64f2130fcb3963e2ecf60d3582aab779009e560eb05521c7e3f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2eb00530e88047dc2a3a548af58fd3486b24335bf0fd7d542aa089160cda88f"></a>

## xcsh_bgp_asn_set — xcsh_bgp_asn_set / 3996d4d25970 / 2

Breadcrumbs:

- xcsh_bgp_asn_set

Manages bgp\_asn\_set creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-36e1cd15f3b547a189fe63f48dd70a13162735430e824ec5b26fa519ca29412d"></a>

## Prerequisites — xcsh_bgp_asn_set / 3996d4d25970 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-70b88b7abc62192ffc523054539a701bb0aa20a703fd2a4dd5d78b8d37199a1a"></a>

## Minimal configuration — xcsh_bgp_asn_set / 3996d4d25970 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BGPAsnSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGPAsnSet by name
data "xcsh_bgp_asn_set" "example" {
  name      = "example-bgp-asn-set"
  namespace = "staging"
}

output "bgp_asn_set_id" {
  value = data.xcsh_bgp_asn_set.example.id
}
```

<a id="canonical-789de1da094f200875fed9cd76c86908c1d66dc5318ac95bca374dfc70d78c48"></a>

## Root configuration — xcsh_bgp_asn_set / 3996d4d25970 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-18b3f4aac929209848528998bcf891201ea7067d799179ea8bc49487c8e9bedc"></a>

## Next pages — xcsh_bgp_asn_set / 3996d4d25970 / 6

- [Property reference](../guides/data-sources--bgp_asn_set--reference--group-001.md#canonical-f77d4e7e206bcb563c10db6a1d66e0f7227add89a6be618d75ba79917484a4eb)
- [Examples](../guides/data-sources--bgp_asn_set--examples--group-001.md#canonical-19443d73f29eec1400dbb793db9a20a8b2687458ef1ceba1d5cacf5041ca414d)
