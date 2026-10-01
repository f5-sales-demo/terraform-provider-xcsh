---
page_title: "xcsh_ike_phase1_profile examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase1_profile examples."
---

# xcsh_ike_phase1_profile examples

<a id="canonical-dd4f9b52237c7b7acd3c241cb898af511b1254c0fa3b3b89791832c09351fdea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2816f3e773fa4b6a55ab666addcd3cd9e2f3ea040c2270075fd31d6b72fc1e26"></a>

## Examples — Examples / 1bdfe0da8b8d / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b)
- Examples

<a id="canonical-6d2ffa538a24711c4229db3601535b261546d105dd323a49349c16b2ee10c9c2"></a>

## Complete configurations — Examples / 1bdfe0da8b8d / 3

- [Data source](data-sources--ike_phase1_profile--examples--group-001.md#canonical-fb918b6f589f19377ac593778da8484506fa1bcc59bd7695d5a2ca53a52a1d1c): valid configuration.

<a id="canonical-41480ae2eb613fd63a76eb6e82a2af24442ba7a87590bff184efa681d584b1c7"></a>

## Next pages — Examples / 1bdfe0da8b8d / 4

- [Data source](data-sources--ike_phase1_profile--examples--group-001.md#canonical-fb918b6f589f19377ac593778da8484506fa1bcc59bd7695d5a2ca53a52a1d1c)
- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b)

<a id="canonical-fb918b6f589f19377ac593778da8484506fa1bcc59bd7695d5a2ca53a52a1d1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72647552cdeffd34a2d320c272309c5fda26db9b16f62c59a447a35ea205fd77"></a>

## Data source — Data source / 5dfa06843207 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b)
- [Examples](data-sources--ike_phase1_profile--examples--group-001.md#canonical-dd4f9b52237c7b7acd3c241cb898af511b1254c0fa3b3b89791832c09351fdea)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_ike_phase1_profile/data-source.tf`; digest `sha256:a22af56dca6a3b92413f9588bd02d5a8a0c4c469599c9023c9feb46580028ab5`.

```terraform
# IKEPhase1Profile Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing IKEPhase1Profile by name
data "xcsh_ike_phase1_profile" "example" {
  name      = "example-ike-phase1-profile"
  namespace = "staging"
}

output "ike_phase1_profile_id" {
  value = data.xcsh_ike_phase1_profile.example.id
}
```

<a id="canonical-7b0beaf934da802eafeb19fdd0e047fb024724059bb4b464570ae3f7b92bbb6a"></a>

## Next pages — Data source / 5dfa06843207 / 3

- [Examples](data-sources--ike_phase1_profile--examples--group-001.md#canonical-dd4f9b52237c7b7acd3c241cb898af511b1254c0fa3b3b89791832c09351fdea)
- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b)
