---
page_title: "xcsh_tenant_configuration"
subcategory: ""
description: "Reads Tenant Configuration information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["tenant configuration"], "body_bytes": 1405, "body_sha256": "sha256:571b9e63e921115da2bc199829d3af874f79e799a0241fa151e3912ba23533e1", "capabilities": ["administration"], "category": "administration", "child_ids": ["xcsh-docs:data-sources:tenant_configuration:reference", "xcsh-docs:data-sources:tenant_configuration:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tenant_configuration:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/tenant_configuration/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1230121131133133-0211313013212213-1101132331100023-2203020001230102-0301100132002332-3002232111300122-2302311100323013-2003322300110201", "registry_path": "docs/data-sources/tenant_configuration.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tenant_configuration/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Reads Tenant Configuration information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_tenant_configuration

Breadcrumbs:

- xcsh_tenant_configuration

Reads Tenant Configuration information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/examples/)
