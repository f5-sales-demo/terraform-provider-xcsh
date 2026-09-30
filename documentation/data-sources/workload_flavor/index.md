---
page_title: "xcsh_workload_flavor"
subcategory: ""
description: "xcsh_workload_flavor for xcsh_workload_flavor."
xcsh_docs: {"aliases": [], "body_bytes": 1213, "body_sha256": "sha256:ff5169b9105df432f414ded13b29dbaade1312f77db8e5873a51fb6dbab46403", "child_ids": ["xcsh-docs:data-sources:workload_flavor:reference", "xcsh-docs:data-sources:workload_flavor:examples"], "collection_id": "xcsh-docs:data-sources:workload_flavor:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload_flavor:fundamentals", "parent_id": null, "path": "documentation/data-sources/workload_flavor/index.md", "provider_name": "workload_flavor", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload_flavor/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_workload_flavor for xcsh_workload_flavor.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workload_flavorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
