---
page_title: "Resource"
subcategory: "Monitoring"
description: "Resource for xcsh_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1027, "body_sha256": "sha256:953b60bdd2a4b23e248aa326350c9c0edb9d316b73d116395c9b886cf88f3253", "canonical_id": "xcsh-docs:resources:log_receiver:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:log_receiver:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5e9b1e9ee7893ce8f260198c000519dd49a791826b61c5ded6feeb808a6b615d", "source_path": "examples/resources/xcsh_log_receiver/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:log_receiver:example:resource", "parent_id": "xcsh-docs:resources:log_receiver:examples", "path": "docs/guides/resources--log_receiver--example--resource.md", "provider_name": "log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/log_receiver/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md)
- [Examples](resources--log_receiver--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_log_receiver/resource.tf`; digest `sha256:5e9b1e9ee7893ce8f260198c000519dd49a791826b61c5ded6feeb808a6b615d`.

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

## Next pages

- [Examples](resources--log_receiver--examples.md)
- [xcsh_log_receiver](../resources/log_receiver.md)
