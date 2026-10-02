---
page_title: "Data source"
subcategory: "Load Balancing"
description: "Data source for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1331, "body_sha256": "sha256:6b59bc539430afecbede8d52b3086773ed4ed79328cf2b1f6efefc8c878164ed", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:065d14809a577f5f36891bc47ab4a39d45652a5da6b5d83b3e19904703a77bc5", "source_path": "examples/data-sources/xcsh_cdn_loadbalancer/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cdn_loadbalancer:example:data-source", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:examples", "path": "documentation/data-sources/cdn_loadbalancer/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3123312220133013-2100223020113230-3103230121001312-0033233223213201-1200310122132110-1132320102221330-0213322300103033-0200200103313012", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_cdn_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cdn_loadbalancer/data-source.tf`; digest `sha256:065d14809a577f5f36891bc47ab4a39d45652a5da6b5d83b3e19904703a77bc5`.

```terraform
# CDNLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CDNLoadBalancer by name
data "xcsh_cdn_loadbalancer" "example" {
  name      = "example-cdn-loadbalancer"
  namespace = "staging"
}

output "cdn_loadbalancer_id" {
  value = data.xcsh_cdn_loadbalancer.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/examples/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
