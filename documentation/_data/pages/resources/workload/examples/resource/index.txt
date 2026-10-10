---
page_title: "Resource"
subcategory: "Container"
description: "Resource for xcsh_workload."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 990, "body_sha256": "sha256:6f362bbe4358abac637a35fd3575225f7488667f826d40db6a0d9dc903b7d1d8", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:95bc586cf100a6ef22147aba592992dc22645f4b290f9226bf45165f3414e58c", "source_path": "examples/resources/xcsh_workload/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:workload:example:resource", "parent_id": "xcsh-docs:resources:workload:examples", "path": "documentation/resources/workload/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3202301020033123-1000311020333023-2200330011313010-3211330331312122-3231102022332233-0023313320023323-2321332001221132-0001132011323320", "registry_path": "docs/guides/resources--workload--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/examples/resource/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Resource for xcsh_workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["workloadCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_workload/resource.tf`; digest `sha256:95bc586cf100a6ef22147aba592992dc22645f4b290f9226bf45165f3414e58c`.

```terraform
# Workload Resource Example
# Manages a Workload resource in F5 Distributed Cloud for workload.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Workload configuration
resource "xcsh_workload" "example" {
  name      = "example-workload"
  namespace = "staging"
}
```
