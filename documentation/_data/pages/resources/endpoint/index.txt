---
page_title: "xcsh_endpoint"
subcategory: "Networking"
description: "Manages endpoint will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["backend servers", "endpoint", "origin servers", "upstream servers"], "body_bytes": 1621, "body_sha256": "sha256:86eaf31993a19631fc19d7fb6a42cf5bf7468c5b29124851acea0862472d4b48", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:endpoint:reference", "xcsh-docs:resources:endpoint:examples", "xcsh-docs:resources:endpoint:import", "xcsh-docs:resources:endpoint:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:resources:endpoint:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/endpoint/index.md", "product": "distributed-cloud", "provider_name": "endpoint", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012", "registry_path": "docs/resources/endpoint.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages endpoint will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["endpointCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_endpoint

Breadcrumbs:

- xcsh_endpoint

Manages endpoint will create the object in the storage backend for namespace metadata.namespace in
F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Endpoint Resource Example
# Manages endpoint will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Endpoint configuration
resource "xcsh_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/lifecycle/timeouts/)
