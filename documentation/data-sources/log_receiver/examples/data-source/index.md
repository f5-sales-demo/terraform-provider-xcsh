---
page_title: "Data source"
subcategory: "Monitoring"
description: "Data source for xcsh_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1180, "body_sha256": "sha256:aae7a69eb5ea358d5138bcfda37b40a22268b740adf6190ae9ad01a7cc58580b", "child_ids": [], "collection_id": "xcsh-docs:data-sources:log_receiver:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e0ffa47f4ff4996df5615eeba45c356df21a5deecfea79894da1f56537fef3b2", "source_path": "examples/data-sources/xcsh_log_receiver/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:log_receiver:example:data-source", "parent_id": "xcsh-docs:data-sources:log_receiver:examples", "path": "documentation/data-sources/log_receiver/examples/data-source/index.md", "provider_name": "log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/log_receiver/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_log_receiver/data-source.tf`; digest `sha256:e0ffa47f4ff4996df5615eeba45c356df21a5deecfea79894da1f56537fef3b2`.

```terraform
# LogReceiver Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing LogReceiver by name
data "xcsh_log_receiver" "example" {
  name      = "example-log-receiver"
  namespace = "staging"
}

output "log_receiver_id" {
  value = data.xcsh_log_receiver.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/examples/)
- [xcsh_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/)
