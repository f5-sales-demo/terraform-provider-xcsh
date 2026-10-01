---
page_title: "xcsh_app_firewall"
subcategory: "Security"
description: "xcsh_app_firewall for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1620, "body_sha256": "sha256:2dc0bbcd9d6791e681ee41f54acf3b9350b39cd8967dfbf3ee84f96caa208914", "child_ids": ["xcsh-docs:resources:app_firewall:reference", "xcsh-docs:resources:app_firewall:examples", "xcsh-docs:resources:app_firewall:import", "xcsh-docs:resources:app_firewall:timeouts"], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:fundamentals", "parent_id": null, "path": "documentation/resources/app_firewall/index.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_app_firewall for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_app_firewall

Breadcrumbs:

- xcsh_app_firewall

Manages Application Firewall in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Optional integrations: `service_policy`.

- service_policy: Fine-grained access control rules

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/lifecycle/timeouts/)
