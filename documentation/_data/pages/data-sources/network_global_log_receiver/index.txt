---
page_title: "xcsh_network_global_log_receiver"
subcategory: ""
description: "Global Log Receiver destinations. Published source entries mix CIDRs and individual IPv4 addresses. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest."
xcsh_docs: {"aliases": ["network global log receiver"], "body_bytes": 1706, "body_sha256": "sha256:1e099b15e6fc98ce348b0ad160c8f77c9d4e1af03e1f961c99b103822b8a96cb", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:network_global_log_receiver:reference", "xcsh-docs:data-sources:network_global_log_receiver:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_global_log_receiver:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/network_global_log_receiver/index.md", "product": "distributed-cloud", "provider_name": "network_global_log_receiver", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0211012321131230-2013331210022130-2003210011323120-3212212113100323-3012123311310320-2030302302201323-1222230131301120-2020212031212211", "registry_path": "docs/data-sources/network_global_log_receiver.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_global_log_receiver/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Global Log Receiver destinations. Published source entries mix CIDRs and individual IPv4 addresses. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_log_receiver/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_log_receiver/examples/)
