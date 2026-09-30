---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1441, "body_sha256": "sha256:739ce22740c88276d47e965be5ec565c7bbcb13d25d0c8d58a191e7e7f777ab4", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_global_log_receiver:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e6481bf54911d3f66c740203a948d54d5c759ec25e22c585e36cf6e6833c987c", "source_path": "examples/data-sources/xcsh_network_global_log_receiver/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_global_log_receiver:example:data-source", "parent_id": "xcsh-docs:data-sources:network_global_log_receiver:examples", "path": "documentation/data-sources/network_global_log_receiver/examples/data-source/index.md", "provider_name": "network_global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_global_log_receiver/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_network_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_network_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_log_receiver/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_log_receiver/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_global_log_receiver/data-source.tf`; digest `sha256:e6481bf54911d3f66c740203a948d54d5c759ec25e22c585e36cf6e6833c987c`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_global_log_receiver" "receivers" {}

# This example chooses TLS syslog on TCP 6514. The manifest supplies only
# destinations; choose the port required by the configured log receiver.
output "tls_syslog_egress" {
  value = {
    direction    = "egress"
    protocol     = "tcp"
    port         = 6514
    destinations = data.xcsh_network_global_log_receiver.receivers.cidr_blocks
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_log_receiver/examples/)
- [xcsh_network_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_log_receiver/)
