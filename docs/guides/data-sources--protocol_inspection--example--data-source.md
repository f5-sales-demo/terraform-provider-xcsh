---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_protocol_inspection."
xcsh_docs: {"aliases": [], "body_bytes": 1164, "body_sha256": "sha256:8cdaf3022acc6b7f8216b31b0d1c43f5ccc13b37541864a55f8373da8a408cfa", "canonical_id": "xcsh-docs:data-sources:protocol_inspection:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:protocol_inspection:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8e357f2e82e2e660e9b2335aec0a0bf6fb5efafdd6746da23a0a982fc6492237", "source_path": "examples/data-sources/xcsh_protocol_inspection/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:protocol_inspection:example:data-source", "parent_id": "xcsh-docs:data-sources:protocol_inspection:examples", "path": "docs/guides/data-sources--protocol_inspection--example--data-source.md", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_inspection/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_protocol_inspection.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md)
- [Examples](data-sources--protocol_inspection--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_protocol_inspection/data-source.tf`; digest `sha256:8e357f2e82e2e660e9b2335aec0a0bf6fb5efafdd6746da23a0a982fc6492237`.

```terraform
# ProtocolInspection Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtocolInspection by name
data "xcsh_protocol_inspection" "example" {
  name      = "example-protocol-inspection"
  namespace = "staging"
}

output "protocol_inspection_id" {
  value = data.xcsh_protocol_inspection.example.id
}
```

## Next pages

- [Examples](data-sources--protocol_inspection--examples.md)
- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md)
