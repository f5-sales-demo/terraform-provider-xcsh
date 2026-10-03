---
page_title: "xcsh_ike2"
subcategory: ""
description: "Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification. configuration."
xcsh_docs: {"aliases": ["ike2"], "body_bytes": 1492, "body_sha256": "sha256:15e96f5df7fbf0e80df4c06dc751e5c1e760a484486193641daa6592abbf0513", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:ike2:reference", "xcsh-docs:resources:ike2:examples", "xcsh-docs:resources:ike2:import", "xcsh-docs:resources:ike2:timeouts"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike2:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike2:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/ike2/index.md", "product": "distributed-cloud", "provider_name": "ike2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3103220130122111-0023320210302011-3122102311032022-0223033112130223-1201210032310333-1323122321230203-0200202200211032-1120020123001232", "registry_path": "docs/resources/ike2.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike2/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ike2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_ike2

Breadcrumbs:

- xcsh_ike2

Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Ike2 Resource Example
# Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Ike2 configuration
resource "xcsh_ike2" "example" {
  name      = "example-ike2"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/lifecycle/timeouts/)
