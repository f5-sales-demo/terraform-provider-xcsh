---
page_title: "Resource"
subcategory: "Container"
description: "Resource for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1200, "body_sha256": "sha256:490127396ec861ccdae425a09d4f8dd1e7ad67eb50fc3f65fc4b4b1540f13684", "child_ids": [], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:95bc586cf100a6ef22147aba592992dc22645f4b290f9226bf45165f3414e58c", "source_path": "examples/resources/xcsh_workload/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:workload:example:resource", "parent_id": "xcsh-docs:resources:workload:examples", "path": "documentation/resources/workload/examples/resource/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/examples/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
