---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1162, "body_sha256": "sha256:63ad76209f465ffebfa7034dcdfe851f36aaf401ec6adc64334cf01dc71f40f9", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9f73258c5164e08d5a0687330200706260aacebe3599fa365152a1976e19940b", "source_path": "examples/data-sources/xcsh_global_log_receiver/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:global_log_receiver:example:data-source", "parent_id": "xcsh-docs:data-sources:global_log_receiver:examples", "path": "docs/guides/data-sources--global_log_receiver--example--data-source.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Examples](data-sources--global_log_receiver--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_global_log_receiver/data-source.tf`; digest `sha256:9f73258c5164e08d5a0687330200706260aacebe3599fa365152a1976e19940b`.

```terraform
# GlobalLogReceiver Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing GlobalLogReceiver by name
data "xcsh_global_log_receiver" "example" {
  name      = "example-global-log-receiver"
  namespace = "staging"
}

output "global_log_receiver_id" {
  value = data.xcsh_global_log_receiver.example.id
}
```

## Next pages

- [Examples](data-sources--global_log_receiver--examples.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
