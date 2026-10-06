---
page_title: "xcsh_workload_flavor"
subcategory: ""
description: "Reads Workload Flavor information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["workload flavor"], "body_bytes": 1336, "body_sha256": "sha256:de3a41946e05cedde0402494578f518791312b2dfa7c0e72249da03973abb036", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload_flavor:reference", "xcsh-docs:data-sources:workload_flavor:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:workload_flavor:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload_flavor:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/workload_flavor/index.md", "product": "distributed-cloud", "provider_name": "workload_flavor", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2030203000323030-0232123232021133-3020001332010013-1013132311011132-1021330332030103-1003300002310323-0333232112231113-1231333120223030", "registry_path": "docs/data-sources/workload_flavor.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload_flavor/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Reads Workload Flavor information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workload_flavorCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_workload_flavor

Breadcrumbs:

- xcsh_workload_flavor

Reads Workload Flavor information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WorkloadFlavor Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing WorkloadFlavor by name
data "xcsh_workload_flavor" "example" {
  name      = "example-workload-flavor"
  namespace = "shared"
}

output "workload_flavor_id" {
  value = data.xcsh_workload_flavor.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload_flavor/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload_flavor/examples/)
