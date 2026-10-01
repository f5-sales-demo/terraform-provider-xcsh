---
page_title: "xcsh_virtual_site"
subcategory: "Infrastructure"
description: "xcsh_virtual_site for xcsh_virtual_site."
xcsh_docs: {"aliases": [], "body_bytes": 1561, "body_sha256": "sha256:1c0ed1bdc8b6ffb7ce3ea755c556af14ac0a6cca8b07cea59cfa10ab6082191f", "child_ids": ["xcsh-docs:resources:virtual_site:reference", "xcsh-docs:resources:virtual_site:examples", "xcsh-docs:resources:virtual_site:import", "xcsh-docs:resources:virtual_site:timeouts"], "collection_id": "xcsh-docs:resources:virtual_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_site:fundamentals", "parent_id": null, "path": "documentation/resources/virtual_site/index.md", "provider_name": "virtual_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_virtual_site for xcsh_virtual_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_virtual_site

Breadcrumbs:

- xcsh_virtual_site

Manages virtual site object in given namespace in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualSite Resource Example
# Manages virtual site object in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualSite configuration
resource "xcsh_virtual_site" "example" {
  name      = "example-virtual-site"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_site/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_site/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_site/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_site/lifecycle/timeouts/)
