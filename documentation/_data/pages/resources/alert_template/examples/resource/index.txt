---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_alert_template."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1148, "body_sha256": "sha256:d4ca22d16b26470362f15b53e0f51608872341de05641c31026d5896b46c44c8", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_template:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:fca951c776be93b0c4b2e24ed4f4a814383784763f4abf35cbade1403b91638c", "source_path": "examples/resources/xcsh_alert_template/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:alert_template:example:resource", "parent_id": "xcsh-docs:resources:alert_template:examples", "path": "documentation/resources/alert_template/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "alert_template", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-0121320133030232-0210111213113303-1023020123030013-3101102121221121-0210121030213133-0130000311123201-2131310233122023-0200133320302103", "registry_path": "docs/guides/resources--alert_template--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_template/examples/resource/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Resource for xcsh_alert_template.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["alert_templateCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_alert_template](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_template/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_template/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_alert_template/resource.tf`; digest `sha256:fca951c776be93b0c4b2e24ed4f4a814383784763f4abf35cbade1403b91638c`.

```terraform
# AlertTemplate Resource Example
# Manages Domain to protect in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertTemplate configuration
resource "xcsh_alert_template" "example" {
  name      = "example-alert-template"
  namespace = "staging"

  alert_message         = "example-value"
  alert_message_details = "example-value"
  alert_name            = "example-value"
}
```
