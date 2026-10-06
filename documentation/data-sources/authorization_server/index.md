---
page_title: "xcsh_authorization_server"
subcategory: ""
description: "Reads Authorization Server information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["authorization server"], "body_bytes": 1405, "body_sha256": "sha256:7c365346a49e6d1a48c120e0db19f490e39940425ad080fbffa67f1812c6d218", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:authorization_server:reference", "xcsh-docs:data-sources:authorization_server:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:authorization_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:authorization_server:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/authorization_server/index.md", "product": "distributed-cloud", "provider_name": "authorization_server", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3021220012103201-0130032102230113-0221231210301022-2323112200300110-3212010013210200-3300200232211023-2113102021313101-1310130133021322", "registry_path": "docs/data-sources/authorization_server.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/authorization_server/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Reads Authorization Server information from F5 Distributed Cloud.", "tasks": ["authentication", "configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["authorization_serverCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_authorization_server

Breadcrumbs:

- xcsh_authorization_server

Reads Authorization Server information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authorization_server/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authorization_server/examples/)
