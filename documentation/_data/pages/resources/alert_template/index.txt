---
page_title: "xcsh_alert_template"
subcategory: ""
description: "Manages Domain to protect in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["alert template"], "body_bytes": 1688, "body_sha256": "sha256:a6412019ecd71ec6d2093764f71e02e764acc78826abd3153c3fc31a6090eab3", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:alert_template:reference", "xcsh-docs:resources:alert_template:examples", "xcsh-docs:resources:alert_template:import", "xcsh-docs:resources:alert_template:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_template:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_template:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/alert_template/index.md", "product": "distributed-cloud", "provider_name": "alert_template", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2113333112122122-3132222003103103-2101331030303103-1200332320133001-1230313330112233-3120111103131333-0332111212113231-2112301302213113", "registry_path": "docs/resources/alert_template.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_template/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages Domain to protect in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["alert_templateCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_alert_template

Breadcrumbs:

- xcsh_alert_template

Manages Domain to protect in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AlertTemplate Resource Example
# Manages Domain to protect in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertTemplate configuration
resource "xcsh_alert_template" "example" {
  name      = "example-alert-template"
  namespace = "staging"

  alert_message         = "example-value"
  alert_message_details = "example-value"
  alert_name            = "example-value"
}
```

## Root configuration

Required root properties: `alert_message`, `alert_message_details`, `alert_name`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_template/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_template/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_template/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_template/lifecycle/timeouts/)
