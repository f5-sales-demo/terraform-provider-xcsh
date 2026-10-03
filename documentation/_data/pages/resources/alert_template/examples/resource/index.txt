---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_alert_template."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1376, "body_sha256": "sha256:979b02b69708e539876fa3b62288a46de2aa762fd2f8fe00d1f67e85a1830f0d", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_template:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:fca951c776be93b0c4b2e24ed4f4a814383784763f4abf35cbade1403b91638c", "source_path": "examples/resources/xcsh_alert_template/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:alert_template:example:resource", "parent_id": "xcsh-docs:resources:alert_template:examples", "path": "documentation/resources/alert_template/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "alert_template", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0121320133030232-0210111213113303-1023020123030013-3101102121221121-0210121030213133-0130000311123201-2131310233122023-0200133320302103", "registry_path": "docs/guides/resources--alert_template--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_template/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_alert_template.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["alert_templateCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_template/examples/)
- [xcsh_alert_template](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_template/)
