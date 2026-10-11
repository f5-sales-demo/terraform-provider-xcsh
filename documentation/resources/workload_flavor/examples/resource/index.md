---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_workload_flavor."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1026, "body_sha256": "sha256:2052a070f0c203728284ff35e33ddd7c1af76ba813d3028c8413682a49b00b9b", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:workload_flavor:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3a521cc20a27b0aff22bd7904a8747719375d23df70f4cf67625e0e522097422", "source_path": "examples/resources/xcsh_workload_flavor/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:workload_flavor:example:resource", "parent_id": "xcsh-docs:resources:workload_flavor:examples", "path": "documentation/resources/workload_flavor/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "workload_flavor", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-3122300202300023-2031012101111123-0330001332122130-3331033120333023-3202103021033303-3031120222000320-1212110121311311-0223303331032102", "registry_path": "docs/guides/resources--workload_flavor--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload_flavor/examples/resource/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Resource for xcsh_workload_flavor.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["workload_flavorCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_workload_flavor](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload_flavor/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload_flavor/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_workload_flavor/resource.tf`; digest `sha256:3a521cc20a27b0aff22bd7904a8747719375d23df70f4cf67625e0e522097422`.

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
