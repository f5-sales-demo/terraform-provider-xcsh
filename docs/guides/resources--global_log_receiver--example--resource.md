---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1109, "body_sha256": "sha256:588f33d92936c15f6a045de17120466cab35ac363da32b307454daa8d13b2532", "canonical_id": "xcsh-docs:resources:global_log_receiver:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1d1066d34a2e865fbb60f6e1a0d99d1dfeb43aa5b8be24b88b8a109599e81802", "source_path": "examples/resources/xcsh_global_log_receiver/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:global_log_receiver:example:resource", "parent_id": "xcsh-docs:resources:global_log_receiver:examples", "path": "docs/guides/resources--global_log_receiver--example--resource.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Examples](resources--global_log_receiver--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_global_log_receiver/resource.tf`; digest `sha256:1d1066d34a2e865fbb60f6e1a0d99d1dfeb43aa5b8be24b88b8a109599e81802`.

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

## Next pages

- [Examples](resources--global_log_receiver--examples.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
