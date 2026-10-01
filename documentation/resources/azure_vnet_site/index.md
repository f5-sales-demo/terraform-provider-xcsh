---
page_title: "xcsh_azure_vnet_site"
subcategory: "Infrastructure"
description: "xcsh_azure_vnet_site for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1935, "body_sha256": "sha256:3901de4376bf173a9503c7c70e86e40493b4923adeaee15e271ffafa9481708d", "child_ids": ["xcsh-docs:resources:azure_vnet_site:reference", "xcsh-docs:resources:azure_vnet_site:examples", "xcsh-docs:resources:azure_vnet_site:import", "xcsh-docs:resources:azure_vnet_site:timeouts"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:fundamentals", "parent_id": null, "path": "documentation/resources/azure_vnet_site/index.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_azure_vnet_site for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/lifecycle/timeouts/)
