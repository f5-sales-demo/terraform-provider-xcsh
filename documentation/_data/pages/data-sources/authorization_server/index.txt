---
page_title: "xcsh_authorization_server"
subcategory: ""
description: "Reads Authorization Server information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["authorization server"], "body_bytes": 1405, "body_sha256": "sha256:7c365346a49e6d1a48c120e0db19f490e39940425ad080fbffa67f1812c6d218", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:authorization_server:reference", "xcsh-docs:data-sources:authorization_server:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:authorization_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:authorization_server:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/authorization_server/index.md", "product": "distributed-cloud", "provider_name": "authorization_server", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3021220012103201-0130032102230113-0221231210301022-2323112200300110-3212010013210200-3300200232211023-2113102021313101-1310130133021322", "registry_path": "docs/data-sources/authorization_server.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/authorization_server/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Reads Authorization Server information from F5 Distributed Cloud.", "tasks": ["authentication", "configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["authorization_serverCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
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
