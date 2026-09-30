---
page_title: "xcsh_workload_flavor"
subcategory: ""
description: "xcsh_workload_flavor for xcsh_workload_flavor."
xcsh_docs: {"aliases": [], "body_bytes": 1210, "body_sha256": "sha256:6618abedf67c998fd3d2bc0d02c2e8b121d5e231029badbcd005895bfadf9a04", "canonical_id": "xcsh-docs:resources:workload_flavor:fundamentals", "child_ids": ["xcsh-docs:resources:workload_flavor:reference", "xcsh-docs:resources:workload_flavor:examples", "xcsh-docs:resources:workload_flavor:import", "xcsh-docs:resources:workload_flavor:timeouts"], "collection_id": "xcsh-docs:resources:workload_flavor:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload_flavor:fundamentals", "parent_id": null, "path": "docs/resources/workload_flavor.md", "provider_name": "workload_flavor", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload_flavor/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_workload_flavor for xcsh_workload_flavor.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workload_flavorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

- [Property reference](../guides/resources--workload_flavor--reference.md)
- [Examples](../guides/resources--workload_flavor--examples.md)
- [Import](../guides/resources--workload_flavor--import.md)
- [Timeouts](../guides/resources--workload_flavor--timeouts.md)
