---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_authorization_server."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1131, "body_sha256": "sha256:0429731fe6a47da518fc460c546ebb6489e8c5e44c4370cdabd16250b8815369", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:authorization_server:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:268fe1444519cbb845d8fcb9ea236d26293e2aeafb324fad41e08e7bc571ffcb", "source_path": "examples/data-sources/xcsh_authorization_server/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:authorization_server:example:data-source", "parent_id": "xcsh-docs:data-sources:authorization_server:examples", "path": "documentation/data-sources/authorization_server/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "authorization_server", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3033220020101301-1200103222123230-3302210133123003-1010001223210330-2100220321112211-3302123133112100-3031103103010010-1313301033330030", "registry_path": "docs/guides/data-sources--authorization_server--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/authorization_server/examples/data-source/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Data source for xcsh_authorization_server.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["authorization_serverCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_authorization_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authorization_server/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authorization_server/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_authorization_server/data-source.tf`; digest `sha256:268fe1444519cbb845d8fcb9ea236d26293e2aeafb324fad41e08e7bc571ffcb`.

```terraform
# AuthorizationServer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AuthorizationServer by name
data "xcsh_authorization_server" "example" {
  name      = "example-authorization-server"
  namespace = "staging"
}

output "authorization_server_id" {
  value = data.xcsh_authorization_server.example.id
}
```
