---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_protocol_policer."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1090, "body_sha256": "sha256:c68787429aea39872062d208e1d818d466bb43d8a91c5b09af8e19c10b36d3af", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protocol_policer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:fc068c58763daba5264bf1b0ce09146319c11b6d3fcae995dc7232b3a411a2d4", "source_path": "examples/data-sources/xcsh_protocol_policer/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:protocol_policer:example:data-source", "parent_id": "xcsh-docs:data-sources:protocol_policer:examples", "path": "documentation/data-sources/protocol_policer/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3022121233330123-2320113200101313-3122301123313301-2302322123030003-0122010101021121-0032023223210113-2201130230232121-2132031103112132", "registry_path": "docs/guides/data-sources--protocol_policer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_policer/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Data source for xcsh_protocol_policer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/examples/)
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
