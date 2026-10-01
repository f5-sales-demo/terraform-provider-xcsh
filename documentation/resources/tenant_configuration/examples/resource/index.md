---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_tenant_configuration."
xcsh_docs: {"aliases": [], "body_bytes": 1368, "body_sha256": "sha256:791ef781da87c35be08e0815c772b461703230cdc1d8c6daaabbe6a35cb2c8b2", "child_ids": [], "collection_id": "xcsh-docs:resources:tenant_configuration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:649c10c3cc2ff997128da4475c6c82fbfad1826ae6d86d2d3e3332f24b98496b", "source_path": "examples/resources/xcsh_tenant_configuration/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:tenant_configuration:example:resource", "parent_id": "xcsh-docs:resources:tenant_configuration:examples", "path": "documentation/resources/tenant_configuration/examples/resource/index.md", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tenant_configuration/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_tenant_configuration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
