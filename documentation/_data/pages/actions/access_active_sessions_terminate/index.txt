---
page_title: "xcsh_access_active_sessions_terminate"
subcategory: ""
description: "Terminates active access sessions in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["access active sessions terminate"], "body_bytes": 1409, "body_sha256": "sha256:f3f2ffdc77129acc6b3bdd42b34f91d72401bb9375ff4974fa0cd4ea43fcb1bb", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:actions:access_active_sessions_terminate:reference", "xcsh-docs:actions:access_active_sessions_terminate:examples", "xcsh-docs:actions:access_active_sessions_terminate:lifecycle"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:access_active_sessions_terminate:collection", "completeness": "complete", "id": "xcsh-docs:actions:access_active_sessions_terminate:fundamentals", "parent_id": "xcsh-docs:actions:xcsh:navigation", "path": "documentation/actions/access_active_sessions_terminate/index.md", "product": "distributed-cloud", "provider_name": "access_active_sessions_terminate", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "actions", "registry_anchor": "canonical-2232333201313313-0211033213132122-1112020213103320-3212121022102212-1220232111200021-2230200120313012-2130013111233211-1310221022123111", "registry_path": "docs/actions/access_active_sessions_terminate.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/access_active_sessions_terminate/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Terminates active access sessions in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_access_active_sessions_terminate

Breadcrumbs:

- xcsh_access_active_sessions_terminate

Terminates active access sessions in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AccessActiveSessionsTerminate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_access_active_sessions_terminate" "example" {
  config {
    namespace = "example-value"
  }
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/lifecycle/)
