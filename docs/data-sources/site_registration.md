---
page_title: "xcsh_site_registration landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registration landing."
---

# xcsh_site_registration landing

<a id="canonical-29e210d357a5bdcac6bb8895520d01f565fe4c298fee3fb79b2fce630978c912"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-371f76364ce820097831e543ca3edda015d5e9c2a0ac545e6a35c43c0fa85e45"></a>

## xcsh_site_registration — xcsh_site_registration / 7faa2e1f1f9d / 2

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

<a id="canonical-8881388f09285cfcef7e652e0cde87a575493ec6807441f3649a47d3294860e5"></a>

## Prerequisites — xcsh_site_registration / 7faa2e1f1f9d / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2d831152a1e172c90087e87e819d953409cd61ce4cf5948a935ce58856abaae3"></a>

## Minimal configuration — xcsh_site_registration / 7faa2e1f1f9d / 4

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

<a id="canonical-85516e9fc12b44b69316b242e2a9ed26e7dc7afef6f72b8b7c36711b2704047f"></a>

## Root configuration — xcsh_site_registration / 7faa2e1f1f9d / 5

Required root properties: `site_name`. Full root flags and choices appear in the property reference.

<a id="canonical-661af57be09ac942638df6a8a6362a178a14eabe833c2439de113889b2d90327"></a>

## Next pages — xcsh_site_registration / 7faa2e1f1f9d / 6

- [Property reference](../guides/data-sources--site_registration--reference--group-001.md#canonical-b894edd5f71c8e0dd62fda7420a5fb0ac3186df32f033ed0b50b08a89478d898)
- [Examples](../guides/data-sources--site_registration--examples--group-001.md#canonical-2ffc7a37abcd0056a9f2d90bd2820c24ce6c8a60078d36dfa1a9cc17da2afecf)
