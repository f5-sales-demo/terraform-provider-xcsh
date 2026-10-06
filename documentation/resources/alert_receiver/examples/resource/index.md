---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_alert_receiver."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1029, "body_sha256": "sha256:7bc71bd24b03d72d1b8fb5f50581b1fdd1c504a3cb26587c0e7c5fd5127cf211", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7520b9716f7e91da22e687315c0f06dd552d7dfa1c32467bdbbdf9c32f5c2efb", "source_path": "examples/resources/xcsh_alert_receiver/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:alert_receiver:example:resource", "parent_id": "xcsh-docs:resources:alert_receiver:examples", "path": "documentation/resources/alert_receiver/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0302211312132003-2302010323020101-0131303320113013-1001122332121003-1012202002120331-1200011020020323-3122133110000031-2221032233231212", "registry_path": "docs/guides/resources--alert_receiver--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/examples/resource/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Resource for xcsh_alert_receiver.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_alert_receiver/resource.tf`; digest `sha256:7520b9716f7e91da22e687315c0f06dd552d7dfa1c32467bdbbdf9c32f5c2efb`.

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
