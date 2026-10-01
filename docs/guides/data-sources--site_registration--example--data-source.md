---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_site_registration."
xcsh_docs: {"aliases": [], "body_bytes": 2395, "body_sha256": "sha256:7f4dbb4c09c872a8375ca9cb76d2df75f65e119e21a7fa9d971e7babd1abe4be", "canonical_id": "xcsh-docs:data-sources:site_registration:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:61863eead29a93d9046672296900d790c36ec0a2b1c8200672a0a876134858bb", "source_path": "examples/data-sources/xcsh_site_registration/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_registration:example:data-source", "parent_id": "xcsh-docs:data-sources:site_registration:examples", "path": "docs/guides/data-sources--site_registration--example--data-source.md", "provider_name": "site_registration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registration/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_site_registration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_site_registration](../data-sources/site_registration.md)
- [Examples](data-sources--site_registration--examples.md)
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

## Next pages

- [Examples](data-sources--site_registration--examples.md)
- [xcsh_site_registration](../data-sources/site_registration.md)
