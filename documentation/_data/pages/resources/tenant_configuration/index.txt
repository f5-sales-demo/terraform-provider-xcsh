---
page_title: "xcsh_tenant_configuration"
subcategory: ""
description: "Manages a Tenant Configuration resource in F5 Distributed Cloud for tenant configuration specification. configuration."
xcsh_docs: {"aliases": ["tenant configuration"], "body_bytes": 1699, "body_sha256": "sha256:5a1dd86aec6b9f7c5fa6cef142e3149c4ed9c9aff5d3c6f704037243a97d4981", "capabilities": ["administration"], "category": "administration", "child_ids": ["xcsh-docs:resources:tenant_configuration:reference", "xcsh-docs:resources:tenant_configuration:examples", "xcsh-docs:resources:tenant_configuration:import", "xcsh-docs:resources:tenant_configuration:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:resources:tenant_configuration:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/tenant_configuration/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310", "registry_path": "docs/resources/tenant_configuration.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tenant_configuration/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Manages a Tenant Configuration resource in F5 Distributed Cloud for tenant configuration specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_tenant_configuration

Breadcrumbs:

- xcsh_tenant_configuration

Manages a Tenant Configuration resource in F5 Distributed Cloud for tenant configuration
specification. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/lifecycle/timeouts/)
