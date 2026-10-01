---
page_title: "xcsh_site_registration"
subcategory: ""
description: "xcsh_site_registration for xcsh_site_registration."
xcsh_docs: {"aliases": [], "body_bytes": 3559, "body_sha256": "sha256:1f906b1ae133709f0d882baf301905136fbe894a2e19668e697af6c221d952e5", "canonical_id": "xcsh-docs:data-sources:site_registration:fundamentals", "child_ids": ["xcsh-docs:data-sources:site_registration:reference", "xcsh-docs:data-sources:site_registration:examples"], "collection_id": "xcsh-docs:data-sources:site_registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registration:fundamentals", "parent_id": null, "path": "docs/data-sources/site_registration.md", "provider_name": "site_registration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registration/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_site_registration for xcsh_site_registration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Property reference](../guides/data-sources--site_registration--reference.md)
- [Examples](../guides/data-sources--site_registration--examples.md)
