---
page_title: "xcsh_workload_flavor"
subcategory: ""
description: "Manages workload_flavor in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["workload flavor"], "body_bytes": 1498, "body_sha256": "sha256:91d0d0b6a3d471c9236dd7a885738bf27d115e802802a38632dd44082fb7ff1f", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload_flavor:reference", "xcsh-docs:resources:workload_flavor:examples", "xcsh-docs:resources:workload_flavor:import", "xcsh-docs:resources:workload_flavor:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:workload_flavor:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload_flavor:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/workload_flavor/index.md", "product": "distributed-cloud", "provider_name": "workload_flavor", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1031333230130321-2120200203322003-3020210212312201-1302323133331112-0230023310233322-3012020133012201-2023100110323022-2002102111310001", "registry_path": "docs/resources/workload_flavor.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload_flavor/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages workload_flavor in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workload_flavorCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
# WorkloadFlavor Resource Example
# Manages workload_flavor in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic WorkloadFlavor configuration
resource "xcsh_workload_flavor" "example" {
  name      = "example-workload-flavor"
  namespace = "shared"
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload_flavor/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload_flavor/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload_flavor/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload_flavor/lifecycle/timeouts/)
