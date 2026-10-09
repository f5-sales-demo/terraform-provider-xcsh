---
page_title: "xcsh_alert_receiver"
subcategory: ""
description: "Manages new Alert Receiver object in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["alert receiver"], "body_bytes": 1534, "body_sha256": "sha256:7bf0cfaf84622141a4ae0dddcd3ce8868bcd304344fafb47a1abaa78a7418df9", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:alert_receiver:reference", "xcsh-docs:resources:alert_receiver:examples", "xcsh-docs:resources:alert_receiver:import", "xcsh-docs:resources:alert_receiver:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/alert_receiver/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311", "registry_path": "docs/resources/alert_receiver.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Manages new Alert Receiver object in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_alert_receiver

Breadcrumbs:

- xcsh_alert_receiver

Manages new Alert Receiver object in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AlertReceiver Resource Example
# Manages new Alert Receiver object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertReceiver configuration
resource "xcsh_alert_receiver" "example" {
  name      = "example-alert-receiver"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/lifecycle/timeouts/)
