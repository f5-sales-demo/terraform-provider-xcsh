---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_site_registration."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 2358, "body_sha256": "sha256:d1b4b03fbadc217be1c819792473fe014d1bcfccd0fd714057280319b0f6b9e0", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:61863eead29a93d9046672296900d790c36ec0a2b1c8200672a0a876134858bb", "source_path": "examples/data-sources/xcsh_site_registration/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_registration:example:data-source", "parent_id": "xcsh-docs:data-sources:site_registration:examples", "path": "documentation/data-sources/site_registration/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "site_registration", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3211121201123100-2310210110202213-3000102233112120-1213031221230023-1322202313211333-0003030101333102-1221023223213103-3220310310022310", "registry_path": "docs/guides/data-sources--site_registration--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registration/examples/data-source/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Data source for xcsh_site_registration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_site_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registration/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registration/examples/)
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
