---
page_title: "xcsh_mitigated_domain"
subcategory: ""
description: "xcsh_mitigated_domain for xcsh_mitigated_domain."
xcsh_docs: {"aliases": [], "body_bytes": 1482, "body_sha256": "sha256:e8ba0a1a7f788d5c9bf67a8cdd78d4e7fa879b853ae5e35b7c0dbaf1b487d722", "child_ids": ["xcsh-docs:resources:mitigated_domain:reference", "xcsh-docs:resources:mitigated_domain:examples", "xcsh-docs:resources:mitigated_domain:import", "xcsh-docs:resources:mitigated_domain:timeouts"], "collection_id": "xcsh-docs:resources:mitigated_domain:collection", "completeness": "complete", "id": "xcsh-docs:resources:mitigated_domain:fundamentals", "parent_id": null, "path": "documentation/resources/mitigated_domain/index.md", "provider_name": "mitigated_domain", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/mitigated_domain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_mitigated_domain for xcsh_mitigated_domain.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["mitigated_domainCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_mitigated_domain

Breadcrumbs:

- xcsh_mitigated_domain

Manages Mitigated Domain in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# MitigatedDomain Resource Example
# Manages Mitigated Domain in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic MitigatedDomain configuration
resource "xcsh_mitigated_domain" "example" {
  name      = "example-mitigated-domain"
  namespace = "staging"

  mitigated_domain = "example-value"
}
```

## Root configuration

Required root properties: `mitigated_domain`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/mitigated_domain/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/mitigated_domain/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/mitigated_domain/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/mitigated_domain/lifecycle/timeouts/)
