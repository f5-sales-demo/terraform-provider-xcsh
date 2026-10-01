---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1073, "body_sha256": "sha256:c3b47b81664544afa38f020d7ddb88459a2700fc6c5cc5ee963e5a3687595239", "canonical_id": "xcsh-docs:data-sources:app_firewall:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3f7d455a2134b0926217d55dc017e3580269822e7d265dbe83b586db534e3931", "source_path": "examples/data-sources/xcsh_app_firewall/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:app_firewall:example:data-source", "parent_id": "xcsh-docs:data-sources:app_firewall:examples", "path": "docs/guides/data-sources--app_firewall--example--data-source.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md)
- [Examples](data-sources--app_firewall--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_firewall/data-source.tf`; digest `sha256:3f7d455a2134b0926217d55dc017e3580269822e7d265dbe83b586db534e3931`.

```terraform
# AppFirewall Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppFirewall by name
data "xcsh_app_firewall" "example" {
  name      = "example-app-firewall"
  namespace = "staging"
}

output "app_firewall_id" {
  value = data.xcsh_app_firewall.example.id
}
```

## Next pages

- [Examples](data-sources--app_firewall--examples.md)
- [xcsh_app_firewall](../data-sources/app_firewall.md)
