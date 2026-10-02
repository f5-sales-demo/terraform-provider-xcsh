---
page_title: "xcsh_advertise_policy"
subcategory: ""
description: "Manages a Advertise Policy resource in F5 Distributed Cloud for advertise_policy object controls how and where a service represented by a given virtual_host object is advertised to consumers. configuration."
xcsh_docs: {"aliases": ["advertise policy"], "body_bytes": 1824, "body_sha256": "sha256:e9215049908a59cc2e2e83ba2a9c7d8613cec7e6dfcec1ff54b7c319a23a3515", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:advertise_policy:reference", "xcsh-docs:resources:advertise_policy:examples", "xcsh-docs:resources:advertise_policy:import", "xcsh-docs:resources:advertise_policy:timeouts"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:advertise_policy:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/advertise_policy/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000", "registry_path": "docs/resources/advertise_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Advertise Policy resource in F5 Distributed Cloud for advertise_policy object controls how and where a service represented by a given virtual_host object is advertised to consumers. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_advertise_policy

Breadcrumbs:

- xcsh_advertise_policy

Manages a Advertise Policy resource in F5 Distributed Cloud for advertise\_policy object controls
how and where a service represented by a given virtual\_host object is advertised to consumers.
configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AdvertisePolicy Resource Example
# Manages a Advertise Policy resource in F5 Distributed Cloud for advertise_policy object controls how and where a service represented by a given virtual_host object is advertised to consumers.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AdvertisePolicy configuration
resource "xcsh_advertise_policy" "example" {
  name      = "example-advertise-policy"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/lifecycle/timeouts/)
