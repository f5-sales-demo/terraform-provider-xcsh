---
page_title: "xcsh_azure_vnet_site"
subcategory: "Infrastructure"
description: "xcsh_azure_vnet_site for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1746, "body_sha256": "sha256:f0f0e66ea2487d4123cb12c19d125e1bc7ddb8edcebb8fb77e2a24a6dec7bc45", "canonical_id": "xcsh-docs:resources:azure_vnet_site:fundamentals", "child_ids": ["xcsh-docs:resources:azure_vnet_site:reference", "xcsh-docs:resources:azure_vnet_site:examples", "xcsh-docs:resources:azure_vnet_site:import", "xcsh-docs:resources:azure_vnet_site:timeouts"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:fundamentals", "parent_id": null, "path": "docs/resources/azure_vnet_site.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_azure_vnet_site for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_azure_vnet_site

Breadcrumbs:

- xcsh_azure_vnet_site

Manages a Azure VNET Site resource in F5 Distributed Cloud for deploying F5 sites within Azure
Virtual Network environments.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cloud_credentials`.

- cloud_credentials: Azure authentication for deployment

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `machine_type`, `name`, `resource_group`, `ssh_key`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--azure_vnet_site--reference.md)
- [Examples](../guides/resources--azure_vnet_site--examples.md)
- [Import](../guides/resources--azure_vnet_site--import.md)
- [Timeouts](../guides/resources--azure_vnet_site--timeouts.md)
