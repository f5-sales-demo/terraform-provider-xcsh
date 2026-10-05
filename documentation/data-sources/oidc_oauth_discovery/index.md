---
page_title: "xcsh_oidc_oauth_discovery"
subcategory: ""
description: "Resource creation operation."
xcsh_docs: {"aliases": ["oidc oauth discovery"], "body_bytes": 1258, "body_sha256": "sha256:2c12f9c01e6830c54b73692d547f8eda177f2cf87b195b850d66d033cf1a6e4e", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:oidc_oauth_discovery:reference", "xcsh-docs:data-sources:oidc_oauth_discovery:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:oidc_oauth_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:oidc_oauth_discovery:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/oidc_oauth_discovery/index.md", "product": "distributed-cloud", "provider_name": "oidc_oauth_discovery", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1212212133022200-1202310302030030-1123103212311111-3102001300103310-0102100100032330-3310033310100102-1213021002223010-3132011230221122", "registry_path": "docs/data-sources/oidc_oauth_discovery.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/oidc_oauth_discovery/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource creation operation.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_oidc_oauth_discovery

Breadcrumbs:

- xcsh_oidc_oauth_discovery

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# OIDCOauthDiscovery DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_oidc_oauth_discovery" "example" {
  namespace = "example-value"
}

output "oidc_oauth_discovery_result" {
  value = data.xcsh_oidc_oauth_discovery.example
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/examples/)
