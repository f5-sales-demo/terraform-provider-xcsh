---
page_title: "xcsh_site_registration"
subcategory: ""
description: "Resolves the runtime registration of a site's Customer Edge (CE) node in F5 Distributed Cloud. A registration is named `r-<uuid>`, **not** after the site it belongs to, so it cannot be read by site name. This data source lists the registrations belonging to a site and returns the one that matches, giving you the name"
xcsh_docs: {"aliases": ["site registration"], "body_bytes": 3657, "body_sha256": "sha256:70cd741990ff947ec32973892eeb18e265687fc267ce8bb149a3014a4c16d213", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registration:reference", "xcsh-docs:data-sources:site_registration:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registration:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/site_registration/index.md", "product": "distributed-cloud", "provider_name": "site_registration", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-0221320201003103-1113221123313022-3012232320202111-1102003100013311-1211333210300221-2033323203332313-2123023330321203-0021132030210102", "registry_path": "docs/data-sources/site_registration.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registration/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Resolves the runtime registration of a site's Customer Edge (CE) node in F5 Distributed Cloud. A registration is named `r-<uuid>`, **not** after the site it belongs to, so it cannot be read by site name. This data source lists the registrations belonging to a site and returns the one that matches, giving you the name", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_site_registration

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

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

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

## Root configuration

Required root properties: `site_name`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registration/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registration/examples/)
