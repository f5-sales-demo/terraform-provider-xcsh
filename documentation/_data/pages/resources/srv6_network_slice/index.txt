---
page_title: "xcsh_srv6_network_slice"
subcategory: ""
description: "Manages srv6_network_slice creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["srv6 network slice"], "body_bytes": 1719, "body_sha256": "sha256:fea96b6fd85b6863405cbcfd7edad71fb4775125a789a261b3ca7ae83d9dd64f", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:srv6_network_slice:reference", "xcsh-docs:resources:srv6_network_slice:examples", "xcsh-docs:resources:srv6_network_slice:import", "xcsh-docs:resources:srv6_network_slice:timeouts"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:srv6_network_slice:collection", "completeness": "complete", "id": "xcsh-docs:resources:srv6_network_slice:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/srv6_network_slice/index.md", "product": "distributed-cloud", "provider_name": "srv6_network_slice", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1003011003113030-2133000032112000-2110300131312112-0323202222213032-3010021101322220-3130233303030103-2300101230220121-0201100002310332", "registry_path": "docs/resources/srv6_network_slice.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/srv6_network_slice/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages srv6_network_slice creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["srv6_network_sliceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_srv6_network_slice

Breadcrumbs:

- xcsh_srv6_network_slice

Manages srv6\_network\_slice creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Srv6NetworkSlice Resource Example
# Manages srv6_network_slice creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Srv6NetworkSlice configuration
resource "xcsh_srv6_network_slice" "example" {
  name      = "example-srv6-network-slice"
  namespace = "system"

  sid_prefixes = ["example-value"]
}
```

## Root configuration

Required root properties: `name`, `sid_prefixes`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/srv6_network_slice/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/srv6_network_slice/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/srv6_network_slice/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/srv6_network_slice/lifecycle/timeouts/)
