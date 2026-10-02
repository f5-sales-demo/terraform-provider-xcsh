---
page_title: "xcsh_proxy"
subcategory: ""
description: "Manages a Proxy resource in F5 Distributed Cloud for tcp loadbalancer create specification. configuration."
xcsh_docs: {"aliases": ["proxy"], "body_bytes": 1285, "body_sha256": "sha256:00813e2c5c3e1a8441944f0e23b7b7effef444dd47c60b8f4e1a4b47626ce9f7", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:proxy:reference", "xcsh-docs:data-sources:proxy:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/proxy/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033", "registry_path": "docs/data-sources/proxy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a Proxy resource in F5 Distributed Cloud for tcp loadbalancer create specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_proxy

Breadcrumbs:

- xcsh_proxy

Manages a Proxy resource in F5 Distributed Cloud for tcp loadbalancer create specification.
configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Proxy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Proxy by name
data "xcsh_proxy" "example" {
  name      = "example-proxy"
  namespace = "staging"
}

output "proxy_id" {
  value = data.xcsh_proxy.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/examples/)
