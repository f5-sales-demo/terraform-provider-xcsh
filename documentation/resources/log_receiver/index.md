---
page_title: "xcsh_log_receiver"
subcategory: "Monitoring"
description: "xcsh_log_receiver for xcsh_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1432, "body_sha256": "sha256:a8d0cd85da565085bcc753eb07a75de92df263c09dfbd5c939672b7a9e276af5", "child_ids": ["xcsh-docs:resources:log_receiver:reference", "xcsh-docs:resources:log_receiver:examples", "xcsh-docs:resources:log_receiver:import", "xcsh-docs:resources:log_receiver:timeouts"], "collection_id": "xcsh-docs:resources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:log_receiver:fundamentals", "parent_id": null, "path": "documentation/resources/log_receiver/index.md", "provider_name": "log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/log_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_log_receiver for xcsh_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_log_receiver

Breadcrumbs:

- xcsh_log_receiver

Manages new Log Receiver object in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# LogReceiver Resource Example
# Manages new Log Receiver object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic LogReceiver configuration
resource "xcsh_log_receiver" "example" {
  name      = "example-log-receiver"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/lifecycle/timeouts/)
