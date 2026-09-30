---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_policy_view."
xcsh_docs: {"aliases": [], "body_bytes": 1062, "body_sha256": "sha256:ef742037a5de4aa63ced796fef5d2319150f54f59343b7724b59a99c5e895a93", "canonical_id": "xcsh-docs:data-sources:network_policy_view:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_policy_view:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:961fb42feeb0126081bc0ae8691b145c93619816626daf020f96266ec095e50a", "source_path": "examples/data-sources/xcsh_network_policy_view/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_policy_view:example:data-source", "parent_id": "xcsh-docs:data-sources:network_policy_view:examples", "path": "docs/guides/data-sources--network_policy_view--example--data-source.md", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_view/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_network_policy_view.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md)
- [Examples](data-sources--network_policy_view--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_policy_view/data-source.tf`; digest `sha256:961fb42feeb0126081bc0ae8691b145c93619816626daf020f96266ec095e50a`.

```terraform
# NetworkPolicyView Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicyView by name
data "xcsh_network_policy_view" "example" {
  name      = "example-network-policy-view"
  namespace = "system"
}

output "network_policy_view_id" {
  value = data.xcsh_network_policy_view.example.id
}
```

## Next pages

- [Examples](data-sources--network_policy_view--examples.md)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md)
