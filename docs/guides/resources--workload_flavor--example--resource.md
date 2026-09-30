---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_workload_flavor."
xcsh_docs: {"aliases": [], "body_bytes": 952, "body_sha256": "sha256:dfe9f5c132c5ad1b52143ab31204fc4043e6582484e65a6132879ec3a4429ff3", "canonical_id": "xcsh-docs:resources:workload_flavor:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:workload_flavor:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3a521cc20a27b0aff22bd7904a8747719375d23df70f4cf67625e0e522097422", "source_path": "examples/resources/xcsh_workload_flavor/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:workload_flavor:example:resource", "parent_id": "xcsh-docs:resources:workload_flavor:examples", "path": "docs/guides/resources--workload_flavor--example--resource.md", "provider_name": "workload_flavor", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload_flavor/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_workload_flavor.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workload_flavorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_workload_flavor](../resources/workload_flavor.md)
- [Examples](resources--workload_flavor--examples.md)
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

## Next pages

- [Examples](resources--workload_flavor--examples.md)
- [xcsh_workload_flavor](../resources/workload_flavor.md)
