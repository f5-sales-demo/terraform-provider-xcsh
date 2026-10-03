---
page_title: "Resource"
subcategory: "Infrastructure"
description: "Resource for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1437, "body_sha256": "sha256:5b536b7b6c1c1498bdaada6ace6c48f8630ab44af7e69043089f637a0e3b2e62", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ba22e659e1847f8f68a606a677c06fc95a71eb38fbebadbdb91aec35b0057b71", "source_path": "examples/resources/xcsh_azure_vnet_site/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:azure_vnet_site:example:resource", "parent_id": "xcsh-docs:resources:azure_vnet_site:examples", "path": "documentation/resources/azure_vnet_site/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0320202033320013-1230211022102023-1200103332102301-2033023222001330-0003210321001133-1022123332223023-0111322022010222-0112300302112333", "registry_path": "docs/guides/resources--azure_vnet_site--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/examples/resource/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Resource for xcsh_azure_vnet_site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/examples/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
