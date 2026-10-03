---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_tenant_configuration."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1368, "body_sha256": "sha256:791ef781da87c35be08e0815c772b461703230cdc1d8c6daaabbe6a35cb2c8b2", "capabilities": ["administration"], "category": "administration", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tenant_configuration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:649c10c3cc2ff997128da4475c6c82fbfad1826ae6d86d2d3e3332f24b98496b", "source_path": "examples/resources/xcsh_tenant_configuration/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:tenant_configuration:example:resource", "parent_id": "xcsh-docs:resources:tenant_configuration:examples", "path": "documentation/resources/tenant_configuration/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0133311010123103-1113111120010223-2203223032033002-3120223101003000-2230220203303012-0210313323231230-3012203203332021-3022120230310103", "registry_path": "docs/guides/resources--tenant_configuration--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tenant_configuration/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_tenant_configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_tenant_configuration/resource.tf`; digest `sha256:649c10c3cc2ff997128da4475c6c82fbfad1826ae6d86d2d3e3332f24b98496b`.

```terraform
# TenantConfiguration Resource Example
# Manages a Tenant Configuration resource in F5 Distributed Cloud for tenant configuration specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic TenantConfiguration configuration
resource "xcsh_tenant_configuration" "example" {
  name      = "example-tenant-configuration"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/examples/)
- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/)
