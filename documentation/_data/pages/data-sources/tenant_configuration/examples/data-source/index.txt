---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_tenant_configuration."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1383, "body_sha256": "sha256:c44ce0e0480827879b3d469359128eb3b990660e29e2cabf8f298bcea235ec32", "capabilities": ["administration"], "category": "administration", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tenant_configuration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:82e6c50873cbaa19717f3b339ef9fae150aab727845b88eb85553bf86a8cc317", "source_path": "examples/data-sources/xcsh_tenant_configuration/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:tenant_configuration:example:data-source", "parent_id": "xcsh-docs:data-sources:tenant_configuration:examples", "path": "documentation/data-sources/tenant_configuration/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2231303120332320-0132300223331320-1323001201212331-3123023223222100-3323110013200100-2202223200203323-2010113102333232-1023130203323132", "registry_path": "docs/guides/data-sources--tenant_configuration--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tenant_configuration/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_tenant_configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_tenant_configuration/data-source.tf`; digest `sha256:82e6c50873cbaa19717f3b339ef9fae150aab727845b88eb85553bf86a8cc317`.

```terraform
# TenantConfiguration Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing TenantConfiguration by name
data "xcsh_tenant_configuration" "example" {
  name      = "example-tenant-configuration"
  namespace = "staging"
}

output "tenant_configuration_id" {
  value = data.xcsh_tenant_configuration.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/examples/)
- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/)
