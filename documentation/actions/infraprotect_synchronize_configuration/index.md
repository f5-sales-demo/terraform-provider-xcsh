---
page_title: "xcsh_infraprotect_synchronize_configuration"
subcategory: ""
description: "Synchronizes infrastructure protection configuration."
xcsh_docs: {"aliases": ["infraprotect synchronize configuration"], "body_bytes": 1447, "body_sha256": "sha256:3ec6b5f30030555f7792c944908cf031508f8c036f8ff82735154d8781fc4041", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:actions:infraprotect_synchronize_configuration:reference", "xcsh-docs:actions:infraprotect_synchronize_configuration:examples", "xcsh-docs:actions:infraprotect_synchronize_configuration:lifecycle"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:infraprotect_synchronize_configuration:collection", "completeness": "complete", "id": "xcsh-docs:actions:infraprotect_synchronize_configuration:fundamentals", "parent_id": "xcsh-docs:actions:xcsh:navigation", "path": "documentation/actions/infraprotect_synchronize_configuration/index.md", "product": "distributed-cloud", "provider_name": "infraprotect_synchronize_configuration", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "actions", "registry_anchor": "canonical-3003320123032222-1321103312000302-3220231112323333-3132221121231130-0320003131013001-3130000003113133-2111213321103223-2223330322121111", "registry_path": "docs/actions/infraprotect_synchronize_configuration.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/infraprotect_synchronize_configuration/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Synchronizes infrastructure protection configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_infraprotect_synchronize_configuration

Breadcrumbs:

- xcsh_infraprotect_synchronize_configuration

Synchronizes infrastructure protection configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# InfraprotectSynchronizeConfiguration Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_infraprotect_synchronize_configuration" "example" {
  config {
    namespace = "example-value"
  }
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/infraprotect_synchronize_configuration/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/infraprotect_synchronize_configuration/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/infraprotect_synchronize_configuration/lifecycle/)
