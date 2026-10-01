---
page_title: "xcsh_network_global_log_receiver"
subcategory: ""
description: "xcsh_network_global_log_receiver for xcsh_network_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1621, "body_sha256": "sha256:1304d97e7e87b0c55af40061e5dbc9b0a86393505e6ccd2563b34a2cda1e8234", "canonical_id": "xcsh-docs:data-sources:network_global_log_receiver:fundamentals", "child_ids": ["xcsh-docs:data-sources:network_global_log_receiver:reference", "xcsh-docs:data-sources:network_global_log_receiver:examples"], "collection_id": "xcsh-docs:data-sources:network_global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_global_log_receiver:fundamentals", "parent_id": null, "path": "docs/data-sources/network_global_log_receiver.md", "provider_name": "network_global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_global_log_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_network_global_log_receiver for xcsh_network_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_global_log_receiver

Breadcrumbs:

- xcsh_network_global_log_receiver

Global Log Receiver destinations. Published source entries mix CIDRs and individual IPv4 addresses.
Values are bundled from the pinned OpenAPI release; this data source performs no network request.
Ports and traffic direction are not encoded in the manifest.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--network_global_log_receiver--reference.md)
- [Examples](../guides/data-sources--network_global_log_receiver--examples.md)
