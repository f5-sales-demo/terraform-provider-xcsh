---
page_title: "Data source"
subcategory: "Load Balancing"
description: "Data source for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1344, "body_sha256": "sha256:a8caba3781f863886f941e13a246a14ee977b646350347706d27234a93a39087", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f7d3fa4bb60b803988578a72b5dac4a76f304bac2925ae4fdfe075100bed37ad", "source_path": "examples/data-sources/xcsh_http_loadbalancer/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:http_loadbalancer:example:data-source", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:examples", "path": "documentation/data-sources/http_loadbalancer/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3130300133300312-1122213020302223-1201233323112022-2232022310002133-0210131131322231-2010030002222320-2320131131002120-3230202332313130", "registry_path": "docs/guides/data-sources--http_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_http_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_http_loadbalancer/data-source.tf`; digest `sha256:f7d3fa4bb60b803988578a72b5dac4a76f304bac2925ae4fdfe075100bed37ad`.

```terraform
# HTTPLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing HTTPLoadBalancer by name
data "xcsh_http_loadbalancer" "example" {
  name      = "example-http-loadbalancer"
  namespace = "staging"
}

output "http_loadbalancer_id" {
  value = data.xcsh_http_loadbalancer.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/examples/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
