---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_alert_template."
xcsh_docs: {"aliases": [], "body_bytes": 1170, "body_sha256": "sha256:c80ef80c38c019fba71bbb17b618f6d43d4b14eff87514c9610747609f84a41f", "canonical_id": "xcsh-docs:resources:alert_template:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:alert_template:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:fca951c776be93b0c4b2e24ed4f4a814383784763f4abf35cbade1403b91638c", "source_path": "examples/resources/xcsh_alert_template/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:alert_template:example:resource", "parent_id": "xcsh-docs:resources:alert_template:examples", "path": "docs/guides/resources--alert_template--example--resource.md", "provider_name": "alert_template", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_template/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_alert_template.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_templateCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_alert_template](../resources/alert_template.md)
- [Examples](resources--alert_template--examples.md)
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

- [Examples](resources--alert_template--examples.md)
- [xcsh_alert_template](../resources/alert_template.md)
