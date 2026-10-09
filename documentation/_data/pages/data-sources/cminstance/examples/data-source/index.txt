---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cminstance."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1033, "body_sha256": "sha256:344ee279ebf4e094594564830786da6276b97a0a05b360954d70a61e89537485", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cminstance:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:833c4b202ec741b02e9ef290bb6512a37d3156dd48242949d7ad995d973193c7", "source_path": "examples/data-sources/xcsh_cminstance/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cminstance:example:data-source", "parent_id": "xcsh-docs:data-sources:cminstance:examples", "path": "documentation/data-sources/cminstance/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "cminstance", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1303001130110232-3311203011112302-0331320312021002-2222332203313102-1300223231331331-2122011210121101-2121012030322331-2030110303322011", "registry_path": "docs/guides/data-sources--cminstance--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cminstance/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_cminstance.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["cminstanceCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_cminstance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cminstance/data-source.tf`; digest `sha256:833c4b202ec741b02e9ef290bb6512a37d3156dd48242949d7ad995d973193c7`.

```terraform
# Cminstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Cminstance by name
data "xcsh_cminstance" "example" {
  name      = "example-cminstance"
  namespace = "staging"
}

output "cminstance_id" {
  value = data.xcsh_cminstance.example.id
}
```
