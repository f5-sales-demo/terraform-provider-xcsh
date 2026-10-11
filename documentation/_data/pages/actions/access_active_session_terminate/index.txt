---
page_title: "xcsh_access_active_session_terminate"
subcategory: ""
description: "Terminates one active access session in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["access active session terminate"], "body_bytes": 1443, "body_sha256": "sha256:8980efaf82a3c406ea541773ebf2c9fb1bb1b2abc2955b71287d33e3d225d097", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:actions:access_active_session_terminate:reference", "xcsh-docs:actions:access_active_session_terminate:examples", "xcsh-docs:actions:access_active_session_terminate:lifecycle"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:access_active_session_terminate:collection", "completeness": "complete", "id": "xcsh-docs:actions:access_active_session_terminate:fundamentals", "parent_id": "xcsh-docs:actions:xcsh:navigation", "path": "documentation/actions/access_active_session_terminate/index.md", "product": "distributed-cloud", "provider_name": "access_active_session_terminate", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "actions", "registry_anchor": "canonical-0233012110120301-2112002303202012-1202012200302211-1131130003102111-0002212030201223-3120132220230001-3123332221233323-3312211113313020", "registry_path": "docs/actions/access_active_session_terminate.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/access_active_session_terminate/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Terminates one active access session in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_access_active_session_terminate

Breadcrumbs:

- xcsh_access_active_session_terminate

Terminates one active access session in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AccessActiveSessionTerminate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_access_active_session_terminate" "example" {
  config {
    id        = "example-value"
    namespace = "example-value"
  }
}
```

## Root configuration

Required root properties: `id`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_session_terminate/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_session_terminate/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_session_terminate/lifecycle/)
