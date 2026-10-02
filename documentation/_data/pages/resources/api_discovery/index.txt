---
page_title: "xcsh_api_discovery"
subcategory: ""
description: "Manages API discovery creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["api discovery"], "body_bytes": 1621, "body_sha256": "sha256:fd6e0048f94073192e79922dda324e8faf59212d730bc4d78da7aca0caea09b0", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:api_discovery:reference", "xcsh-docs:resources:api_discovery:examples", "xcsh-docs:resources:api_discovery:import", "xcsh-docs:resources:api_discovery:timeouts"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_discovery:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/api_discovery/index.md", "product": "distributed-cloud", "provider_name": "api_discovery", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310", "registry_path": "docs/resources/api_discovery.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_discovery/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages API discovery creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_api_discovery

Breadcrumbs:

- xcsh_api_discovery

Manages API discovery creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APIDiscovery Resource Example
# Manages API discovery creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APIDiscovery configuration
resource "xcsh_api_discovery" "example" {
  name      = "example-api-discovery"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/lifecycle/timeouts/)
