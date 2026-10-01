---
page_title: "xcsh_alert_gen_policy"
subcategory: ""
description: "xcsh_alert_gen_policy for xcsh_alert_gen_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1346, "body_sha256": "sha256:08629b09060224c28155a39f35c17a8b738c2a51cd4f5480b2c531e5208eb768", "canonical_id": "xcsh-docs:resources:alert_gen_policy:fundamentals", "child_ids": ["xcsh-docs:resources:alert_gen_policy:reference", "xcsh-docs:resources:alert_gen_policy:examples", "xcsh-docs:resources:alert_gen_policy:import", "xcsh-docs:resources:alert_gen_policy:timeouts"], "collection_id": "xcsh-docs:resources:alert_gen_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_gen_policy:fundamentals", "parent_id": null, "path": "docs/resources/alert_gen_policy.md", "provider_name": "alert_gen_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_gen_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_alert_gen_policy for xcsh_alert_gen_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_gen_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_alert_gen_policy

Breadcrumbs:

- xcsh_alert_gen_policy

Manages Alert Generation Policy in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AlertGenPolicy Resource Example
# Manages Alert Generation Policy in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertGenPolicy configuration
resource "xcsh_alert_gen_policy" "example" {
  name      = "example-alert-gen-policy"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--alert_gen_policy--reference.md)
- [Examples](../guides/resources--alert_gen_policy--examples.md)
- [Import](../guides/resources--alert_gen_policy--import.md)
- [Timeouts](../guides/resources--alert_gen_policy--timeouts.md)
