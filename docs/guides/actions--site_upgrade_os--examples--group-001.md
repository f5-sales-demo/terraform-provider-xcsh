---
page_title: "xcsh_site_upgrade_os examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_upgrade_os examples."
---

# xcsh_site_upgrade_os examples

<a id="canonical-01f1a7047468e631893088da94970716a9806f8b80faf254913c119178c591a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-211876b92e802256d57bb0468477dbc5dad0ce353c07f5e43b03b8272d6a6267"></a>

## Examples — Examples / 046c3ad95041 / 2

Breadcrumbs:

- [xcsh_site_upgrade_os](../actions/site_upgrade_os.md#canonical-094cbfe1e9c98c0b0352113f3cdf16dc1917d92802e806ed418e980550d32637)
- Examples

<a id="canonical-06938436c8d384c0e43fe12380ef55863decc8b36c1dc2e014d4fbd88d8c379e"></a>

## Complete configurations — Examples / 046c3ad95041 / 3

- [Action](actions--site_upgrade_os--examples--group-001.md#canonical-88bfa4a460ae14856e4272776b2839bcca1f4f699b7fa824e1f0a7f01ceb5090): valid configuration.

<a id="canonical-96f87076d32a48a22acea2b671872d1523edc34767818e4c3c8d54bbc4df0671"></a>

## Next pages — Examples / 046c3ad95041 / 4

- [Action](actions--site_upgrade_os--examples--group-001.md#canonical-88bfa4a460ae14856e4272776b2839bcca1f4f699b7fa824e1f0a7f01ceb5090)
- [xcsh_site_upgrade_os](../actions/site_upgrade_os.md#canonical-094cbfe1e9c98c0b0352113f3cdf16dc1917d92802e806ed418e980550d32637)

<a id="canonical-88bfa4a460ae14856e4272776b2839bcca1f4f699b7fa824e1f0a7f01ceb5090"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a7d431c58b91bd71b1be6936666302c8b52e929f9fe67873a558eeee5f9b00f"></a>

## Action — Action / 4ce31cdab752 / 2

Breadcrumbs:

- [xcsh_site_upgrade_os](../actions/site_upgrade_os.md#canonical-094cbfe1e9c98c0b0352113f3cdf16dc1917d92802e806ed418e980550d32637)
- [Examples](actions--site_upgrade_os--examples--group-001.md#canonical-01f1a7047468e631893088da94970716a9806f8b80faf254913c119178c591a2)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_site_upgrade_os/action.tf`; digest `sha256:39e348141a455b8f8743fa7481580e9f9219671ca870734d07a968a5bcbd2acc`.

```terraform
# SiteUpgradeOS Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_site_upgrade_os" "example" {
  config {
    site       = "example-value"
    os_version = "example-value"
  }
}
```

<a id="canonical-59c0507817538c7588797e9df3650bb80ad5b171261cdd2962e34300dd73736c"></a>

## Next pages — Action / 4ce31cdab752 / 3

- [Examples](actions--site_upgrade_os--examples--group-001.md#canonical-01f1a7047468e631893088da94970716a9806f8b80faf254913c119178c591a2)
- [xcsh_site_upgrade_os](../actions/site_upgrade_os.md#canonical-094cbfe1e9c98c0b0352113f3cdf16dc1917d92802e806ed418e980550d32637)
