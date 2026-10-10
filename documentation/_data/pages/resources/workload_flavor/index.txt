---
page_title: "xcsh_workload_flavor"
subcategory: ""
description: "Manages workload_flavor in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["workload flavor"], "body_bytes": 1511, "body_sha256": "sha256:c596e0403a4993086489401b2018ffe2e0cbd22373e902219eac79f939b634e6", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload_flavor:reference", "xcsh-docs:resources:workload_flavor:examples", "xcsh-docs:resources:workload_flavor:import", "xcsh-docs:resources:workload_flavor:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:workload_flavor:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload_flavor:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/workload_flavor/index.md", "product": "distributed-cloud", "provider_name": "workload_flavor", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1031333230130321-2120200203322003-3020210212312201-1302323133331112-0230023310233322-3012020133012201-2023100110323022-2002102111310001", "registry_path": "docs/resources/workload_flavor.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload_flavor/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Manages workload_flavor in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["workload_flavorCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload_flavor/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload_flavor/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload_flavor/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload_flavor/lifecycle/timeouts/)
