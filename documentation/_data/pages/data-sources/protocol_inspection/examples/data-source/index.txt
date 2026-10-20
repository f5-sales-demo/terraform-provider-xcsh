---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_protocol_inspection."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1121, "body_sha256": "sha256:4fface66b4abadcbef2f85ab997952118e97ee86739a809dde17f5ff3c8b6d55", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protocol_inspection:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8e357f2e82e2e660e9b2335aec0a0bf6fb5efafdd6746da23a0a982fc6492237", "source_path": "examples/data-sources/xcsh_protocol_inspection/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:protocol_inspection:example:data-source", "parent_id": "xcsh-docs:data-sources:protocol_inspection:examples", "path": "documentation/data-sources/protocol_inspection/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-2321033120121020-3101302133330101-2131210100112013-1223322021203030-3211213013202023-3112213231210012-2132113032320231-1203012130102321", "registry_path": "docs/guides/data-sources--protocol_inspection--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_inspection/examples/data-source/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Data source for xcsh_protocol_inspection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/examples/)
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
