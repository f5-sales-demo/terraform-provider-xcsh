---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_infraprotect_synchronize_configuration."
xcsh_docs: {"aliases": ["action"], "body_bytes": 1340, "body_sha256": "sha256:d4593cb480193028e02e4586ef680dba485dcb583cbe1a0e41bed11a8b1d490a", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:infraprotect_synchronize_configuration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5708bb8f8ebc531162d8b17ecbd5421cebdfe14f2d0a5a47b3ca315eedfc8227", "source_path": "examples/actions/xcsh_infraprotect_synchronize_configuration/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:infraprotect_synchronize_configuration:example:action", "parent_id": "xcsh-docs:actions:infraprotect_synchronize_configuration:examples", "path": "documentation/actions/infraprotect_synchronize_configuration/examples/action/index.md", "product": "distributed-cloud", "provider_name": "infraprotect_synchronize_configuration", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "actions", "registry_anchor": "canonical-1013120113203013-0132301301023223-3203220333212301-2331233010311231-2121001222121010-0103302311222323-0221311222302123-0132211200320311", "registry_path": "docs/guides/actions--infraprotect_synchronize_configuration--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/infraprotect_synchronize_configuration/examples/action/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Action for xcsh_infraprotect_synchronize_configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/infraprotect_synchronize_configuration/examples/)
- [xcsh_infraprotect_synchronize_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/infraprotect_synchronize_configuration/)
