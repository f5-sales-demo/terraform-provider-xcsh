---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_protocol_policer."
xcsh_docs: {"aliases": [], "body_bytes": 1124, "body_sha256": "sha256:8db412e0f392f5be4f568b62104c4aea55a67eabff87c33fb6cda69c0d329172", "canonical_id": "xcsh-docs:data-sources:protocol_policer:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:protocol_policer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:fc068c58763daba5264bf1b0ce09146319c11b6d3fcae995dc7232b3a411a2d4", "source_path": "examples/data-sources/xcsh_protocol_policer/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:protocol_policer:example:data-source", "parent_id": "xcsh-docs:data-sources:protocol_policer:examples", "path": "docs/guides/data-sources--protocol_policer--example--data-source.md", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_policer/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_protocol_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md)
- [Examples](data-sources--protocol_policer--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_protocol_policer/data-source.tf`; digest `sha256:fc068c58763daba5264bf1b0ce09146319c11b6d3fcae995dc7232b3a411a2d4`.

```terraform
# ProtocolPolicer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtocolPolicer by name
data "xcsh_protocol_policer" "example" {
  name      = "example-protocol-policer"
  namespace = "system"
}

output "protocol_policer_id" {
  value = data.xcsh_protocol_policer.example.id
}
```

## Next pages

- [Examples](data-sources--protocol_policer--examples.md)
- [xcsh_protocol_policer](../data-sources/protocol_policer.md)
