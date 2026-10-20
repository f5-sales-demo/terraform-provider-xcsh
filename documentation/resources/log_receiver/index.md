---
page_title: "xcsh_log_receiver"
subcategory: "Monitoring"
description: "Manages new Log Receiver object in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["log receiver"], "body_bytes": 1544, "body_sha256": "sha256:878b600aa0cb97fb32d6f612b38a47a425dff2d66ec510176cb4e08ad8e8fe00", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:log_receiver:reference", "xcsh-docs:resources:log_receiver:examples", "xcsh-docs:resources:log_receiver:import", "xcsh-docs:resources:log_receiver:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:log_receiver:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/log_receiver/index.md", "product": "distributed-cloud", "provider_name": "log_receiver", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003", "registry_path": "docs/resources/log_receiver.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/log_receiver/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Manages new Log Receiver object in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["log_receiverCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_log_receiver

Breadcrumbs:

- xcsh_log_receiver

Manages new Log Receiver object in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/lifecycle/timeouts/)
