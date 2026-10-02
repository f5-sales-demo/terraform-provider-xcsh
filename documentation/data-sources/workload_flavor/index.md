---
page_title: "xcsh_workload_flavor"
subcategory: ""
description: "Manages workload_flavor in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["workload flavor"], "body_bytes": 1312, "body_sha256": "sha256:fffa8352613084c40f0c823f3fdeb2ce9b7ecb7c497c7ab27f2b272869e42b42", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload_flavor:reference", "xcsh-docs:data-sources:workload_flavor:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:workload_flavor:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload_flavor:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/workload_flavor/index.md", "product": "distributed-cloud", "provider_name": "workload_flavor", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2030203000323030-0232123232021133-3020001332010013-1013132311011132-1021330332030103-1003300002310323-0333232112231113-1231333120223030", "registry_path": "docs/data-sources/workload_flavor.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload_flavor/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages workload_flavor in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workload_flavorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_workload_flavor

Breadcrumbs:

- xcsh_workload_flavor

Manages workload\_flavor in F5 Distributed Cloud.

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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload_flavor/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload_flavor/examples/)
