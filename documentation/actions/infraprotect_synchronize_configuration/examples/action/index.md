---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_infraprotect_synchronize_configuration."
xcsh_docs: {"aliases": ["action"], "body_bytes": 1340, "body_sha256": "sha256:d4593cb480193028e02e4586ef680dba485dcb583cbe1a0e41bed11a8b1d490a", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:infraprotect_synchronize_configuration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5708bb8f8ebc531162d8b17ecbd5421cebdfe14f2d0a5a47b3ca315eedfc8227", "source_path": "examples/actions/xcsh_infraprotect_synchronize_configuration/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:infraprotect_synchronize_configuration:example:action", "parent_id": "xcsh-docs:actions:infraprotect_synchronize_configuration:examples", "path": "documentation/actions/infraprotect_synchronize_configuration/examples/action/index.md", "product": "distributed-cloud", "provider_name": "infraprotect_synchronize_configuration", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "actions", "registry_anchor": "canonical-1013120113203013-0132301301023223-3203220333212301-2331233010311231-2121001222121010-0103302311222323-0221311222302123-0132211200320311", "registry_path": "docs/guides/actions--infraprotect_synchronize_configuration--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/infraprotect_synchronize_configuration/examples/action/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Action for xcsh_infraprotect_synchronize_configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
