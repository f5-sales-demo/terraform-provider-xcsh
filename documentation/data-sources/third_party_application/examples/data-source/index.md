---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_third_party_application."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1159, "body_sha256": "sha256:9e13993c4d88d2b34a953a5e9120992c885b0465bc095110c304b3cc21eb31e5", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e3e758943d32635744f0c37b6cc645f6b182639afd8477f0071a1bfaa5279e53", "source_path": "examples/data-sources/xcsh_third_party_application/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:third_party_application:example:data-source", "parent_id": "xcsh-docs:data-sources:third_party_application:examples", "path": "documentation/data-sources/third_party_application/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "third_party_application", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0200101300201233-2010323323200232-1302213031311002-0003000231113013-2132310012222121-2002232120011103-0022033012313231-1122120330222102", "registry_path": "docs/guides/data-sources--third_party_application--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/examples/data-source/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Data source for xcsh_third_party_application.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_third_party_application/data-source.tf`; digest `sha256:e3e758943d32635744f0c37b6cc645f6b182639afd8477f0071a1bfaa5279e53`.

```terraform
# ThirdPartyApplication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ThirdPartyApplication by name
data "xcsh_third_party_application" "example" {
  name      = "example-third-party-application"
  namespace = "staging"
}

output "third_party_application_id" {
  value = data.xcsh_third_party_application.example.id
}
```
