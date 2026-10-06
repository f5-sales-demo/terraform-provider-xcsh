---
page_title: "xcsh_oidc_oauth_discovery"
subcategory: ""
description: "Reads OIDC/OAuth discovery information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["oidc oauth discovery"], "body_bytes": 1308, "body_sha256": "sha256:07882a7f01b1ebf46b26f3debedc2250d00190c52f39049030784ad40618ecbb", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:oidc_oauth_discovery:reference", "xcsh-docs:data-sources:oidc_oauth_discovery:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:oidc_oauth_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:oidc_oauth_discovery:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/oidc_oauth_discovery/index.md", "product": "distributed-cloud", "provider_name": "oidc_oauth_discovery", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1212212133022200-1202310302030030-1123103212311111-3102001300103310-0102100100032330-3310033310100102-1213021002223010-3132011230221122", "registry_path": "docs/data-sources/oidc_oauth_discovery.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/oidc_oauth_discovery/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Reads OIDC/OAuth discovery information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_oidc_oauth_discovery

Breadcrumbs:

- xcsh_oidc_oauth_discovery

Reads OIDC/OAuth discovery information from F5 Distributed Cloud.

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/examples/)
