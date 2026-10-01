---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1024, "body_sha256": "sha256:0c6b0cedc607dba179752b5ef98d627ba425c77d2c10339cb7821a3dd9f25a16", "canonical_id": "xcsh-docs:resources:app_firewall:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:03e5e1d1ea7b4ce0a0005da7cf6bf27eeca5913c46ae5552f44fffd430a9cb75", "source_path": "examples/resources/xcsh_app_firewall/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:resource", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "docs/guides/resources--app_firewall--example--resource.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md)
- [Examples](resources--app_firewall--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/resource.tf`; digest `sha256:03e5e1d1ea7b4ce0a0005da7cf6bf27eeca5913c46ae5552f44fffd430a9cb75`.

```terraform
# AppFirewall Resource Example
# Manages Application Firewall in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppFirewall configuration
resource "xcsh_app_firewall" "example" {
  name      = "example-app-firewall"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--app_firewall--examples.md)
- [xcsh_app_firewall](../resources/app_firewall.md)
