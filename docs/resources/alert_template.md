---
page_title: "xcsh_alert_template"
subcategory: ""
description: "xcsh_alert_template for xcsh_alert_template."
xcsh_docs: {"aliases": [], "body_bytes": 1499, "body_sha256": "sha256:490f43cfa56949a0ace2bbab82598d4a72c8046a5d7dccb0bf90ad6bceef2615", "canonical_id": "xcsh-docs:resources:alert_template:fundamentals", "child_ids": ["xcsh-docs:resources:alert_template:reference", "xcsh-docs:resources:alert_template:examples", "xcsh-docs:resources:alert_template:import", "xcsh-docs:resources:alert_template:timeouts"], "collection_id": "xcsh-docs:resources:alert_template:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_template:fundamentals", "parent_id": null, "path": "docs/resources/alert_template.md", "provider_name": "alert_template", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_template/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_alert_template for xcsh_alert_template.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_templateCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [Property reference](../guides/resources--alert_template--reference.md)
- [Examples](../guides/resources--alert_template--examples.md)
- [Import](../guides/resources--alert_template--import.md)
- [Timeouts](../guides/resources--alert_template--timeouts.md)
