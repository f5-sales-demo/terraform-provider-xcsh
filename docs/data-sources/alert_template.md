---
page_title: "xcsh_alert_template"
subcategory: ""
description: "xcsh_alert_template for xcsh_alert_template."
xcsh_docs: {"aliases": [], "body_bytes": 1133, "body_sha256": "sha256:d0799dfe74ec11f8447b99942e47ea0f8b965f0731b7a6348ac56c561c5409e6", "canonical_id": "xcsh-docs:data-sources:alert_template:fundamentals", "child_ids": ["xcsh-docs:data-sources:alert_template:reference", "xcsh-docs:data-sources:alert_template:examples"], "collection_id": "xcsh-docs:data-sources:alert_template:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_template:fundamentals", "parent_id": null, "path": "docs/data-sources/alert_template.md", "provider_name": "alert_template", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_template/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_alert_template for xcsh_alert_template.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_templateCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_alert_template

Breadcrumbs:

- xcsh_alert_template

Manages Domain to protect in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AlertTemplate Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertTemplate by name
data "xcsh_alert_template" "example" {
  name      = "example-alert-template"
  namespace = "staging"
}

output "alert_template_id" {
  value = data.xcsh_alert_template.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--alert_template--reference.md)
- [Examples](../guides/data-sources--alert_template--examples.md)
