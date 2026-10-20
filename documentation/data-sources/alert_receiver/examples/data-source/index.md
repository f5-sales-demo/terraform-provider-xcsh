---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_alert_receiver."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1071, "body_sha256": "sha256:ba519bbc5d6bf26c574026c1182d794cef33d53e62ddced671c6de0feebb4b0d", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:96027faa6721d178ff8fb480c66120ec2d45036c783323915216275133e6a1d6", "source_path": "examples/data-sources/xcsh_alert_receiver/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:alert_receiver:example:data-source", "parent_id": "xcsh-docs:data-sources:alert_receiver:examples", "path": "documentation/data-sources/alert_receiver/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3221110320020111-3312110013000223-1332222210300322-1033301303313312-1122112222311311-1303332201220113-0031310322112310-2103120203003001", "registry_path": "docs/guides/data-sources--alert_receiver--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/examples/data-source/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Data source for xcsh_alert_receiver.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_alert_receiver/data-source.tf`; digest `sha256:96027faa6721d178ff8fb480c66120ec2d45036c783323915216275133e6a1d6`.

```terraform
# AlertReceiver Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertReceiver by name
data "xcsh_alert_receiver" "example" {
  name      = "example-alert-receiver"
  namespace = "staging"
}

output "alert_receiver_id" {
  value = data.xcsh_alert_receiver.example.id
}
```
