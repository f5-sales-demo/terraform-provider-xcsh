---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1329, "body_sha256": "sha256:dbbcfea3fb0f2e0f3ce3943d4e9adfe4223f33206334e21b97135e5c48f1765e", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:436b4e693b05709f4b55969967c5b9c433a210d85502a116c164d87cdf218449", "source_path": "examples/data-sources/xcsh_bigip_http_proxy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bigip_http_proxy:example:data-source", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:examples", "path": "documentation/data-sources/bigip_http_proxy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2233132123121133-0223233312131211-0310031101201023-0331131311311010-3003110020312301-1210331233113033-3133220231013312-3202333232300233", "registry_path": "docs/guides/data-sources--bigip_http_proxy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_bigip_http_proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/examples/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
