---
page_title: "xcsh_site_upgrade_os examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_upgrade_os examples."
---

# xcsh_site_upgrade_os examples

<a id="canonical-0001330122130010-1310122032120301-2021030020203122-2110211300130112-2221200012332023-2000332233021110-2101033001012101-1320301121012202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_site_upgrade_os](../actions/site_upgrade_os.md#canonical-0021103023333201-3221302120300023-0003110201010333-0330313301123130-0121011331210220-0002322000123231-1001203221200011-1100310302120313)
- Examples

<a id="canonical-0201012013122321-0232200002021112-3111132323001012-2010131331233011-3122310030320311-0330001333113210-0323000323200213-0231122212021213"></a>

### Complete configurations for `xcsh_site_upgrade_os`

- [Action](actions--site_upgrade_os--examples--group-001.md#canonical-2020233322102210-1200223201102011-1232100213021313-1223022003212330-3022013310331221-2123133322200210-3201330022133300-0130322311002100): valid configuration.

<a id="canonical-2020233322102210-1200223201102011-1232100213021313-1223022003212330-3022013310331221-2123133322200210-3201330022133300-0130322311002100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Action example

Breadcrumbs:

- [xcsh_site_upgrade_os](../actions/site_upgrade_os.md#canonical-0021103023333201-3221302120300023-0003110201010333-0330313301123130-0121011331210220-0002322000123231-1001203221200011-1100310302120313)
- [Examples](actions--site_upgrade_os--examples--group-001.md#canonical-0001330122130010-1310122032120301-2021030020203122-2110211300130112-2221200012332023-2000332233021110-2101033001012101-1320301121012202)
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
