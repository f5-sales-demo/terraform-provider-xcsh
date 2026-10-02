---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_access_active_sessions."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1316, "body_sha256": "sha256:aedb4f9f153d826d8ddb842dd94011f545391c1dfe4f13fd420ba697a8938839", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:access_active_sessions:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9f2701c9ce37e0c0584f554dedc7733344fe85678c9dadc11b3997db86cac357", "source_path": "examples/data-sources/xcsh_access_active_sessions/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:access_active_sessions:example:data-source", "parent_id": "xcsh-docs:data-sources:access_active_sessions:examples", "path": "documentation/data-sources/access_active_sessions/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "access_active_sessions", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1202122232020100-1122112103210323-1133013331211222-3120012223010033-0112221002020012-2230332303101231-1121101202231232-2213301123311001", "registry_path": "docs/guides/data-sources--access_active_sessions--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/access_active_sessions/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_access_active_sessions.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_sessions/examples/)
- [xcsh_access_active_sessions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_sessions/)
