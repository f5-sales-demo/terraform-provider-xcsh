---
page_title: "xcsh_endpoint"
subcategory: "Networking"
description: "Manages endpoint will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["endpoint"], "body_bytes": 1634, "body_sha256": "sha256:5e8e199c2ca5c4cacdca9b9410d1c677101a1babad3803c07d3c17847e627d30", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:endpoint:reference", "xcsh-docs:resources:endpoint:examples", "xcsh-docs:resources:endpoint:import", "xcsh-docs:resources:endpoint:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:resources:endpoint:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/endpoint/index.md", "product": "distributed-cloud", "provider_name": "endpoint", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012", "registry_path": "docs/resources/endpoint.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Manages endpoint will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["endpointCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/lifecycle/timeouts/)
