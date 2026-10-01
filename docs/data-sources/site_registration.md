---
page_title: "xcsh_site_registration landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registration landing."
---

# xcsh_site_registration landing

<a id="canonical-0221320201003103-1113221123313022-3012232320202111-1102003100013311-1211333210300221-2033323203332313-2123023330321203-0021132030210102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313013313120312-1030322002000021-1320030132111003-3022033231312200-0111311132213002-2200223011101132-1222031130100330-0033222011321011"></a>

## xcsh_site_registration — xcsh_site_registration / 013301332131 / 2

Breadcrumbs:

- xcsh_site_registration

Resolves the runtime registration of a site's Customer Edge (CE) node in F5 Distributed Cloud.

A registration is named \`r-&lt;uuid&gt;\`, \*\*not\*\* after the site it belongs to, so it cannot
be read by site name. This data source lists the registrations belonging to a site and returns the
one that matches, giving you the name that \`xcsh\_registration\_approval\` requires.

The registration only exists once the CE has booted and registered with its token. Until then this
data source reports \`found = false\` \*\*without raising an error\*\*, so an approval can safely be
gated on it:

\`\`\`terraform data "xcsh\_site\_registration" "ce" \{ site\_name =
xcsh\_securemesh\_site\_v2.ce.name \}

resource "xcsh\_registration\_approval" "ce" \{ count = data.xcsh\_site\_registration.ce.found ? 1 :
0 name = data.xcsh\_site\_registration.ce.name namespace = "system" cluster\_size = 1 \} \`\`\`

\*\*Possible \`state\` values:\*\* \`NOTSET\`, \`NEW\`, \`APPROVED\`, \`ADMITTED\`, \`RETIRED\`,
\`FAILED\`, \`DONE\`, \`PENDING\`, \`ONLINE\`, \`UPGRADING\`, \`MAINTENANCE\`, \`FAILED\_INACTIVE\`.

<a id="canonical-2020200103202033-0021022011303330-3233133212110232-0030313220132211-1311102103323012-2000131010013303-1210212210133103-0221102012003211"></a>

## Prerequisites — xcsh_site_registration / 013301332131 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0231200301011102-2201320113023021-0000201332201332-2001213121110310-0021303112013032-1030331121102022-2103113032112020-1112222322223203"></a>

## Minimal configuration — xcsh_site_registration / 013301332131 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-2011110112322133-3001022310102312-2103011223021002-3202222132310212-3213313013223332-3312331302232023-1330031213010123-0213001000101333"></a>

## Root configuration — xcsh_site_registration / 013301332131 / 5

Required root properties: `site_name`. Full root flags and choices appear in the property reference.

<a id="canonical-1212012233111323-3200212230211002-1203203133122220-2212031202220113-2022011032222332-2003033002100321-3132010103202021-2302312100030213"></a>

## Next pages — xcsh_site_registration / 013301332131 / 6

- [Property reference](../guides/data-sources--site_registration--reference--group-001.md#canonical-2320211032313111-3313013020320031-3112023331221310-0200221133230022-3003012012313303-0233000303323100-2311002300202220-2110132031202120)
- [Examples](../guides/data-sources--site_registration--examples--group-001.md#canonical-0233333013220313-2223303100001112-2221330231210023-3102200200300210-3032123020221200-0013203103123133-2201222130300113-3122022233323033)
