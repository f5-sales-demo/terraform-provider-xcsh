---
page_title: "xcsh_oidc_oauth_discovery"
subcategory: ""
description: "xcsh_oidc_oauth_discovery for xcsh_oidc_oauth_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1173, "body_sha256": "sha256:c95a9163f544571037ae57cc7307ddffc0feccd58b6113bdb40c9656ed1ce5a7", "canonical_id": "xcsh-docs:data-sources:oidc_oauth_discovery:fundamentals", "child_ids": ["xcsh-docs:data-sources:oidc_oauth_discovery:reference", "xcsh-docs:data-sources:oidc_oauth_discovery:examples"], "collection_id": "xcsh-docs:data-sources:oidc_oauth_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:oidc_oauth_discovery:fundamentals", "parent_id": null, "path": "docs/data-sources/oidc_oauth_discovery.md", "provider_name": "oidc_oauth_discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/oidc_oauth_discovery/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_oidc_oauth_discovery for xcsh_oidc_oauth_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [Property reference](../guides/data-sources--oidc_oauth_discovery--reference.md)
- [Examples](../guides/data-sources--oidc_oauth_discovery--examples.md)
