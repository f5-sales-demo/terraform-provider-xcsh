---
page_title: "xcsh_infraprotect_synchronize_configuration"
subcategory: ""
description: "Resource creation operation."
xcsh_docs: {"aliases": ["infraprotect synchronize configuration"], "body_bytes": 1409, "body_sha256": "sha256:7843a630e4460b4cc28448b6877b6c4aa0150fe5762c9f706bf0b1d2ebf9fee0", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:actions:infraprotect_synchronize_configuration:reference", "xcsh-docs:actions:infraprotect_synchronize_configuration:examples", "xcsh-docs:actions:infraprotect_synchronize_configuration:lifecycle"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:infraprotect_synchronize_configuration:collection", "completeness": "complete", "id": "xcsh-docs:actions:infraprotect_synchronize_configuration:fundamentals", "parent_id": "xcsh-docs:actions:xcsh:navigation", "path": "documentation/actions/infraprotect_synchronize_configuration/index.md", "product": "distributed-cloud", "provider_name": "infraprotect_synchronize_configuration", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "actions", "registry_anchor": "canonical-3003320123032222-1321103312000302-3220231112323333-3132221121231130-0320003131013001-3130000003113133-2111213321103223-2223330322121111", "registry_path": "docs/actions/infraprotect_synchronize_configuration.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/infraprotect_synchronize_configuration/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource creation operation.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_infraprotect_synchronize_configuration

Breadcrumbs:

- xcsh_infraprotect_synchronize_configuration

Resource creation operation.

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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/infraprotect_synchronize_configuration/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/infraprotect_synchronize_configuration/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/infraprotect_synchronize_configuration/lifecycle/)
