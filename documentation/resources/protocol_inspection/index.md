---
page_title: "xcsh_protocol_inspection"
subcategory: ""
description: "xcsh_protocol_inspection for xcsh_protocol_inspection."
xcsh_docs: {"aliases": [], "body_bytes": 1650, "body_sha256": "sha256:fbca740458930089a73a22dfa0bf80be3f2bf6d2b673257d160407eb07965d3f", "child_ids": ["xcsh-docs:resources:protocol_inspection:reference", "xcsh-docs:resources:protocol_inspection:examples", "xcsh-docs:resources:protocol_inspection:import", "xcsh-docs:resources:protocol_inspection:timeouts"], "collection_id": "xcsh-docs:resources:protocol_inspection:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_inspection:fundamentals", "parent_id": null, "path": "documentation/resources/protocol_inspection/index.md", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_inspection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_protocol_inspection for xcsh_protocol_inspection.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_protocol_inspection

Breadcrumbs:

- xcsh_protocol_inspection

Manages Protocol Inspection Specification in a given namespace. If one already exists it will give
an error in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtocolInspection Resource Example
# Manages Protocol Inspection Specification in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtocolInspection configuration
resource "xcsh_protocol_inspection" "example" {
  name      = "example-protocol-inspection"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/lifecycle/timeouts/)
