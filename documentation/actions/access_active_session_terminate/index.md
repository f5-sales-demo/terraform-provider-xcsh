---
page_title: "xcsh_access_active_session_terminate"
subcategory: ""
description: "Resource deletion operation."
xcsh_docs: {"aliases": ["access active session terminate"], "body_bytes": 1397, "body_sha256": "sha256:241e52a228834e9a85f07538c480113306f5ce68900d5c83460eb745fb8e78de", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:actions:access_active_session_terminate:reference", "xcsh-docs:actions:access_active_session_terminate:examples", "xcsh-docs:actions:access_active_session_terminate:lifecycle"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:access_active_session_terminate:collection", "completeness": "complete", "id": "xcsh-docs:actions:access_active_session_terminate:fundamentals", "parent_id": "xcsh-docs:actions:xcsh:navigation", "path": "documentation/actions/access_active_session_terminate/index.md", "product": "distributed-cloud", "provider_name": "access_active_session_terminate", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "actions", "registry_anchor": "canonical-0233012110120301-2112002303202012-1202012200302211-1131130003102111-0002212030201223-3120132220230001-3123332221233323-3312211113313020", "registry_path": "docs/actions/access_active_session_terminate.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/access_active_session_terminate/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource deletion operation.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_access_active_session_terminate

Breadcrumbs:

- xcsh_access_active_session_terminate

Resource deletion operation.

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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_session_terminate/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_session_terminate/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_session_terminate/lifecycle/)
