---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_authorization_server."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1383, "body_sha256": "sha256:b0543e4fff1b93102fb07c13bbb67b97ad8c89a7296b7ab04a6cbc38c2a6df79", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:authorization_server:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:268fe1444519cbb845d8fcb9ea236d26293e2aeafb324fad41e08e7bc571ffcb", "source_path": "examples/data-sources/xcsh_authorization_server/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:authorization_server:example:data-source", "parent_id": "xcsh-docs:data-sources:authorization_server:examples", "path": "documentation/data-sources/authorization_server/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "authorization_server", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3033220020101301-1200103222123230-3302210133123003-1010001223210330-2100220321112211-3302123133112100-3031103103010010-1313301033330030", "registry_path": "docs/guides/data-sources--authorization_server--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/authorization_server/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_authorization_server.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["authorization_serverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authorization_server/examples/)
- [xcsh_authorization_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authorization_server/)
