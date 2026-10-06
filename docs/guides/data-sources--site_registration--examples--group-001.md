---
page_title: "xcsh_site_registration examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registration examples."
---

# xcsh_site_registration examples

<a id="canonical-0233333013220313-2223303100001112-2221330231210023-3102200200300210-3032123020221200-0013203103123133-2201222130300113-3122022233323033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_site_registration](../data-sources/site_registration.md#canonical-0221320201003103-1113221123313022-3012232320202111-1102003100013311-1211333210300221-2033323203332313-2123023330321203-0021132030210102)
- Examples

<a id="canonical-0310030012121012-3321302030132121-0113022220321332-0021300112020000-2002232233313230-3312012213310033-2111023012102121-1322030321012103"></a>

### Complete configurations for `xcsh_site_registration`

- [Data source](data-sources--site_registration--examples--group-001.md#canonical-3211121201123100-2310210110202213-3000102233112120-1213031221230023-1322202313211333-0003030101333102-1221023223213103-3220310310022310): valid configuration.

<a id="canonical-3211121201123100-2310210110202213-3000102233112120-1213031221230023-1322202313211333-0003030101333102-1221023223213103-3220310310022310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_site_registration](../data-sources/site_registration.md#canonical-0221320201003103-1113221123313022-3012232320202111-1102003100013311-1211333210300221-2033323203332313-2123023330321203-0021132030210102)
- [Examples](data-sources--site_registration--examples--group-001.md#canonical-0233333013220313-2223303100001112-2221330231210023-3102200200300210-3032123020221200-0013203103123133-2201222130300113-3122022233323033)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_registration/data-source.tf`; digest `sha256:61863eead29a93d9046672296900d790c36ec0a2b1c8200672a0a876134858bb`.

```terraform
# Example: Resolve a Customer Edge registration so it can be approved
#
# A registration is named "r-<uuid>", NOT after the site, so it cannot be read
# by site name. This data source finds the registration that belongs to a site.
#
# It returns found = false — with no error — until the CE has booted and
# registered, so an approval can safely be gated on it.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registration" "ce" {
  site_name = "my-ce-site"
  namespace = "system"
}

# A multi-node site returns one registration per node; pick one by hostname.
data "xcsh_site_registration" "ha_ce_node_0" {
  site_name = "my-ha-ce-site"
  hostname  = "master-0"
}

output "registration_name" {
  description = "Registration name (r-<uuid>) to approve, null until the CE registers"
  value       = data.xcsh_site_registration.ce.name
}

output "registration_state" {
  description = "Current registration state (PENDING, ONLINE, ...)"
  value       = data.xcsh_site_registration.ce.state
}

output "ha_node_0_registration_name" {
  description = "Registration name of the master-0 node of the three-node site"
  value       = data.xcsh_site_registration.ha_ce_node_0.name
}

# Approve the registration only once it exists — on the first apply the CE has
# not registered yet, so nothing is planned; re-apply after the node boots.
resource "xcsh_registration_approval" "ce" {
  count = data.xcsh_site_registration.ce.found ? 1 : 0

  name         = data.xcsh_site_registration.ce.name
  namespace    = data.xcsh_site_registration.ce.namespace
  cluster_size = 1
}
```
