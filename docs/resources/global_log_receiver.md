---
page_title: "xcsh_global_log_receiver"
subcategory: ""
description: "xcsh_global_log_receiver for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1291, "body_sha256": "sha256:3d31df09bb3e4641c278bba21c717629004dcaea7816a4fe0e48aeae1ce42d37", "canonical_id": "xcsh-docs:resources:global_log_receiver:fundamentals", "child_ids": ["xcsh-docs:resources:global_log_receiver:reference", "xcsh-docs:resources:global_log_receiver:examples", "xcsh-docs:resources:global_log_receiver:import", "xcsh-docs:resources:global_log_receiver:timeouts"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:fundamentals", "parent_id": null, "path": "docs/resources/global_log_receiver.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_global_log_receiver for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_global_log_receiver

Breadcrumbs:

- xcsh_global_log_receiver

Manages new Global Log Receiver object in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# GlobalLogReceiver Resource Example
# Manages new Global Log Receiver object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic GlobalLogReceiver configuration
resource "xcsh_global_log_receiver" "example" {
  name      = "example-global-log-receiver"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--global_log_receiver--reference.md)
- [Examples](../guides/resources--global_log_receiver--examples.md)
- [Import](../guides/resources--global_log_receiver--import.md)
- [Timeouts](../guides/resources--global_log_receiver--timeouts.md)
