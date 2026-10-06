---
page_title: "xcsh_authorization_server"
subcategory: ""
description: "Manages authorization_server creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["authorization server"], "body_bytes": 1761, "body_sha256": "sha256:8d659aea0a7286743e44e3afe6700892d83c01fc673ab6f88718310a755df97d", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:authorization_server:reference", "xcsh-docs:resources:authorization_server:examples", "xcsh-docs:resources:authorization_server:import", "xcsh-docs:resources:authorization_server:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:authorization_server:collection", "completeness": "complete", "id": "xcsh-docs:resources:authorization_server:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/authorization_server/index.md", "product": "distributed-cloud", "provider_name": "authorization_server", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3123011103301332-2020213322103110-2220312130201323-0120021121102210-3010302120211132-3100321223002321-1020323131322210-3122032302302331", "registry_path": "docs/resources/authorization_server.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authorization_server/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Manages authorization_server creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.", "tasks": ["authentication", "configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["authorization_serverCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_authorization_server

Breadcrumbs:

- xcsh_authorization_server

Manages authorization\_server creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `jwks_uri`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authorization_server/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authorization_server/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authorization_server/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authorization_server/lifecycle/timeouts/)
