---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_access_active_sessions."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1058, "body_sha256": "sha256:520be803c8df52d5a68811edb79fee2ef3153e43af2fe5bea979221e92117a9b", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:access_active_sessions:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9f2701c9ce37e0c0584f554dedc7733344fe85678c9dadc11b3997db86cac357", "source_path": "examples/data-sources/xcsh_access_active_sessions/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:access_active_sessions:example:data-source", "parent_id": "xcsh-docs:data-sources:access_active_sessions:examples", "path": "documentation/data-sources/access_active_sessions/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "access_active_sessions", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-1202122232020100-1122112103210323-1133013331211222-3120012223010033-0112221002020012-2230332303101231-1121101202231232-2213301123311001", "registry_path": "docs/guides/data-sources--access_active_sessions--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/access_active_sessions/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_access_active_sessions.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_access_active_sessions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_sessions/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_sessions/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_access_active_sessions/data-source.tf`; digest `sha256:9f2701c9ce37e0c0584f554dedc7733344fe85678c9dadc11b3997db86cac357`.

```terraform
# AccessActiveSessions DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_access_active_sessions" "example" {
  namespace = "example-value"
}

output "access_active_sessions_result" {
  value = data.xcsh_access_active_sessions.example
}
```
