---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cdn_purge_command."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1099, "body_sha256": "sha256:ae4dbbde8c0f24bd004f41a1df8532a447d13c815337946c495bc707b25e0630", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cdn_purge_command:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:25cb8095a143a10866ff868fcae4396aaa9eca15b734c77fdd0d7054bddadd8a", "source_path": "examples/data-sources/xcsh_cdn_purge_command/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cdn_purge_command:example:data-source", "parent_id": "xcsh-docs:data-sources:cdn_purge_command:examples", "path": "documentation/data-sources/cdn_purge_command/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "cdn_purge_command", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-2211013302231123-0023210003112102-1111023330301123-0320113021003123-2323130002101121-2233213301221312-1231320333301230-0031221111310032", "registry_path": "docs/guides/data-sources--cdn_purge_command--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_purge_command/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Data source for xcsh_cdn_purge_command.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["cdn_purge_commandCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_cdn_purge_command](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_purge_command/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_purge_command/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cdn_purge_command/data-source.tf`; digest `sha256:25cb8095a143a10866ff868fcae4396aaa9eca15b734c77fdd0d7054bddadd8a`.

```terraform
# CDNPurgeCommand Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CDNPurgeCommand by name
data "xcsh_cdn_purge_command" "example" {
  name      = "example-cdn-purge-command"
  namespace = "staging"
}

output "cdn_purge_command_id" {
  value = data.xcsh_cdn_purge_command.example.id
}
```
