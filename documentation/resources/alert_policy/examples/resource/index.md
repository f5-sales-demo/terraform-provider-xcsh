---
page_title: "Resource"
subcategory: "Monitoring"
description: "Resource for xcsh_alert_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1134, "body_sha256": "sha256:fb0728c223519dc4e17e3f5f7e99cb82e3a02e74b18efae83b05eb3616868345", "child_ids": [], "collection_id": "xcsh-docs:resources:alert_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:02e1db745e00ae3c15ffd89e3e1addc3f550309288f544ced75c1c7b1645161e", "source_path": "examples/resources/xcsh_alert_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:alert_policy:example:resource", "parent_id": "xcsh-docs:resources:alert_policy:examples", "path": "documentation/resources/alert_policy/examples/resource/index.md", "provider_name": "alert_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_alert_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_alert_policy/resource.tf`; digest `sha256:02e1db745e00ae3c15ffd89e3e1addc3f550309288f544ced75c1c7b1645161e`.

```terraform
# AlertPolicy Resource Example
# Manages new Alert Policy Object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertPolicy configuration
resource "xcsh_alert_policy" "example" {
  name      = "example-alert-policy"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/examples/)
- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
