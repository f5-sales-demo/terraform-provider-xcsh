---
page_title: "xcsh_alert_receiver"
subcategory: ""
description: "xcsh_alert_receiver for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1332, "body_sha256": "sha256:c2b1be8323a8cf2483d5a84c80fcce59e4c997fd4143d155051bf66fece46080", "canonical_id": "xcsh-docs:resources:alert_receiver:fundamentals", "child_ids": ["xcsh-docs:resources:alert_receiver:reference", "xcsh-docs:resources:alert_receiver:examples", "xcsh-docs:resources:alert_receiver:import", "xcsh-docs:resources:alert_receiver:timeouts"], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:fundamentals", "parent_id": null, "path": "docs/resources/alert_receiver.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_alert_receiver for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_alert_receiver

Breadcrumbs:

- xcsh_alert_receiver

Manages new Alert Receiver object in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AlertReceiver Resource Example
# Manages new Alert Receiver object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertReceiver configuration
resource "xcsh_alert_receiver" "example" {
  name      = "example-alert-receiver"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--alert_receiver--reference.md)
- [Examples](../guides/resources--alert_receiver--examples.md)
- [Import](../guides/resources--alert_receiver--import.md)
- [Timeouts](../guides/resources--alert_receiver--timeouts.md)
