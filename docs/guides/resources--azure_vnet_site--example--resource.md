---
page_title: "Resource"
subcategory: "Infrastructure"
description: "Resource for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1231, "body_sha256": "sha256:f2b44ff53cf1bd154587047dca36c96ff0151330f6c831d7b795ad3318f11207", "canonical_id": "xcsh-docs:resources:azure_vnet_site:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ba22e659e1847f8f68a606a677c06fc95a71eb38fbebadbdb91aec35b0057b71", "source_path": "examples/resources/xcsh_azure_vnet_site/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:azure_vnet_site:example:resource", "parent_id": "xcsh-docs:resources:azure_vnet_site:examples", "path": "docs/guides/resources--azure_vnet_site--example--resource.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Examples](resources--azure_vnet_site--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_azure_vnet_site/resource.tf`; digest `sha256:ba22e659e1847f8f68a606a677c06fc95a71eb38fbebadbdb91aec35b0057b71`.

```terraform
# AzureVNETSite Resource Example
# Manages a Azure VNET Site resource in F5 Distributed Cloud for deploying F5 sites within Azure Virtual Network environments.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AzureVNETSite configuration
resource "xcsh_azure_vnet_site" "example" {
  name      = "example-azure-vnet-site"
  namespace = "system"

  machine_type   = "example-value"
  resource_group = "example-value"
  ssh_key        = "example-value"
}
```

## Next pages

- [Examples](resources--azure_vnet_site--examples.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
