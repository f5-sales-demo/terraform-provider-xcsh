---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1089, "body_sha256": "sha256:55faee41ea3e3af6df560eb15934311246e4b395b7e5a5d96369f4196a015142", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:436b4e693b05709f4b55969967c5b9c433a210d85502a116c164d87cdf218449", "source_path": "examples/data-sources/xcsh_bigip_http_proxy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bigip_http_proxy:example:data-source", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:examples", "path": "documentation/data-sources/bigip_http_proxy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2233132123121133-0223233312131211-0310031101201023-0331131311311010-3003110020312301-1210331233113033-3133220231013312-3202333232300233", "registry_path": "docs/guides/data-sources--bigip_http_proxy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_bigip_http_proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bigip_http_proxy/data-source.tf`; digest `sha256:436b4e693b05709f4b55969967c5b9c433a210d85502a116c164d87cdf218449`.

```terraform
# BigIPHTTPProxy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BigIPHTTPProxy by name
data "xcsh_bigip_http_proxy" "example" {
  name      = "example-bigip-http-proxy"
  namespace = "staging"
}

output "bigip_http_proxy_id" {
  value = data.xcsh_bigip_http_proxy.example.id
}
```
