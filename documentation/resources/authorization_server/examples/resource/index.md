---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_authorization_server."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1415, "body_sha256": "sha256:b70065cc4f83a9ea31085e1bbcc86abfdc96cd1ae9a4ddd7ac6826fef0e43532", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:authorization_server:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:bbb05ccd40cd71dbeef4a9b6ed31c1cec2684fa158b0357f99c63ef6b31e6af5", "source_path": "examples/resources/xcsh_authorization_server/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:authorization_server:example:resource", "parent_id": "xcsh-docs:resources:authorization_server:examples", "path": "documentation/resources/authorization_server/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "authorization_server", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1110123012323002-2000131221012030-2202032322113010-0301211100213311-0132010130001223-3203220221023213-3131223222131113-1301133220321013", "registry_path": "docs/guides/resources--authorization_server--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authorization_server/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_authorization_server.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["authorization_serverCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_authorization_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authorization_server/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authorization_server/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_authorization_server/resource.tf`; digest `sha256:bbb05ccd40cd71dbeef4a9b6ed31c1cec2684fa158b0357f99c63ef6b31e6af5`.

```terraform
# AuthorizationServer Resource Example
# Manages authorization_server creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AuthorizationServer configuration
resource "xcsh_authorization_server" "example" {
  name      = "example-authorization-server"
  namespace = "staging"

  jwks_uri = "example-value"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authorization_server/examples/)
- [xcsh_authorization_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authorization_server/)
