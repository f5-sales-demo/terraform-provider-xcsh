---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_tenant_configuration."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1131, "body_sha256": "sha256:fbbf73cbbd1a7b387cbd77076167af9357adc6a1ab6fb6eb82aed2ad831930ae", "capabilities": ["administration"], "category": "administration", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tenant_configuration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:82e6c50873cbaa19717f3b339ef9fae150aab727845b88eb85553bf86a8cc317", "source_path": "examples/data-sources/xcsh_tenant_configuration/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:tenant_configuration:example:data-source", "parent_id": "xcsh-docs:data-sources:tenant_configuration:examples", "path": "documentation/data-sources/tenant_configuration/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2231303120332320-0132300223331320-1323001201212331-3123023223222100-3323110013200100-2202223200203323-2010113102333232-1023130203323132", "registry_path": "docs/guides/data-sources--tenant_configuration--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tenant_configuration/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_tenant_configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
