---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_infraprotect_synchronize_configuration."
xcsh_docs: {"aliases": ["action"], "body_bytes": 1044, "body_sha256": "sha256:4da2c8b3d1b05c564c7fa009a4ae1bd61f304f051c7a2e1208c2ab5751183b67", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:infraprotect_synchronize_configuration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5708bb8f8ebc531162d8b17ecbd5421cebdfe14f2d0a5a47b3ca315eedfc8227", "source_path": "examples/actions/xcsh_infraprotect_synchronize_configuration/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:infraprotect_synchronize_configuration:example:action", "parent_id": "xcsh-docs:actions:infraprotect_synchronize_configuration:examples", "path": "documentation/actions/infraprotect_synchronize_configuration/examples/action/index.md", "product": "distributed-cloud", "provider_name": "infraprotect_synchronize_configuration", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "actions", "registry_anchor": "canonical-1013120113203013-0132301301023223-3203220333212301-2331233010311231-2121001222121010-0103302311222323-0221311222302123-0132211200320311", "registry_path": "docs/guides/actions--infraprotect_synchronize_configuration--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/infraprotect_synchronize_configuration/examples/action/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Action for xcsh_infraprotect_synchronize_configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Action

Breadcrumbs:

- [xcsh_infraprotect_synchronize_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/infraprotect_synchronize_configuration/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/infraprotect_synchronize_configuration/examples/)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_infraprotect_synchronize_configuration/action.tf`; digest `sha256:5708bb8f8ebc531162d8b17ecbd5421cebdfe14f2d0a5a47b3ca315eedfc8227`.

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
