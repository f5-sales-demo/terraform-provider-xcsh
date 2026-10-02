---
page_title: "xcsh_bigip_http_proxy"
subcategory: ""
description: "Manages BIG-IP HTTP Proxy in a given namespace. If one already exists, it will give an error in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["bigip http proxy"], "body_bytes": 1587, "body_sha256": "sha256:5cb0c769c7a1461a90d3d6933ccb679d48ffbd1a9cd6f98ec677837560a65918", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:reference", "xcsh-docs:resources:bigip_http_proxy:examples", "xcsh-docs:resources:bigip_http_proxy:import", "xcsh-docs:resources:bigip_http_proxy:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/bigip_http_proxy/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111", "registry_path": "docs/resources/bigip_http_proxy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages BIG-IP HTTP Proxy in a given namespace. If one already exists, it will give an error in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bigip_http_proxy

Breadcrumbs:

- xcsh_bigip_http_proxy

Manages BIG-IP HTTP Proxy in a given namespace. If one already exists, it will give an error in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BigIPHTTPProxy Resource Example
# Manages BIG-IP HTTP Proxy in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BigIPHTTPProxy configuration
resource "xcsh_bigip_http_proxy" "example" {
  name      = "example-bigip-http-proxy"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/lifecycle/timeouts/)
