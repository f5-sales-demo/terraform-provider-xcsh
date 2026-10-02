---
page_title: "xcsh_ike1"
subcategory: ""
description: "Manages a Ike1 resource in F5 Distributed Cloud for ike phase1 profile specification. configuration."
xcsh_docs: {"aliases": ["ike1"], "body_bytes": 1492, "body_sha256": "sha256:52edb3b0acf3474f913dfe9ee2d570ab2999854ab0eac6d277d1a2c57fc53c1c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:ike1:reference", "xcsh-docs:resources:ike1:examples", "xcsh-docs:resources:ike1:import", "xcsh-docs:resources:ike1:timeouts"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike1:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike1:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/ike1/index.md", "product": "distributed-cloud", "provider_name": "ike1", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130", "registry_path": "docs/resources/ike1.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike1/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Ike1 resource in F5 Distributed Cloud for ike phase1 profile specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ike1CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_ike1

Breadcrumbs:

- xcsh_ike1

Manages a Ike1 resource in F5 Distributed Cloud for ike phase1 profile specification. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Ike1 Resource Example
# Manages a Ike1 resource in F5 Distributed Cloud for ike phase1 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Ike1 configuration
resource "xcsh_ike1" "example" {
  name      = "example-ike1"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/lifecycle/timeouts/)
