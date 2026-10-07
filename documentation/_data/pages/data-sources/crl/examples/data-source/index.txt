---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_crl."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 963, "body_sha256": "sha256:7dc85fe758d0c6e5be09f6f040e3ca0544e7dbb0d15d3f6ab8141a05a930418d", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:crl:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7250b8110197f6a15add9c7d7bb8ba0193edd8faf34e61501e055b382c218d8a", "source_path": "examples/data-sources/xcsh_crl/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:crl:example:data-source", "parent_id": "xcsh-docs:data-sources:crl:examples", "path": "documentation/data-sources/crl/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "crl", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3020121233030311-3122033322301333-2312211200310011-2031300000230123-2113303100023233-1310201220030010-2102030210123201-0312131312201322", "registry_path": "docs/guides/data-sources--crl--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/crl/examples/data-source/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Data source for xcsh_crl.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["crlCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/crl/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/crl/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_crl/data-source.tf`; digest `sha256:7250b8110197f6a15add9c7d7bb8ba0193edd8faf34e61501e055b382c218d8a`.

```terraform
# CRL Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CRL by name
data "xcsh_crl" "example" {
  name      = "example-crl"
  namespace = "staging"
}

output "crl_id" {
  value = data.xcsh_crl.example.id
}
```
