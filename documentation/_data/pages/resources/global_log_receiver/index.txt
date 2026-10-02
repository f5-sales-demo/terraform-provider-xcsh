---
page_title: "xcsh_global_log_receiver"
subcategory: ""
description: "Manages new Global Log Receiver object in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["global log receiver"], "body_bytes": 1579, "body_sha256": "sha256:c6aecf14693c6171d1fb8181c92127a91e94b119474d4880c3a9db332a4015d4", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:reference", "xcsh-docs:resources:global_log_receiver:examples", "xcsh-docs:resources:global_log_receiver:import", "xcsh-docs:resources:global_log_receiver:timeouts"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/global_log_receiver/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211", "registry_path": "docs/resources/global_log_receiver.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages new Global Log Receiver object in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/lifecycle/timeouts/)
