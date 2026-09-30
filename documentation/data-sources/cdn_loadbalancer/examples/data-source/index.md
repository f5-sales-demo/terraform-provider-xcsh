---
page_title: "Data source"
subcategory: "Load Balancing"
description: "Data source for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1232, "body_sha256": "sha256:8305df0f528c32e6a1d70b93c7a2710964296f85df60964e42d8f26cf750a8cc", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:065d14809a577f5f36891bc47ab4a39d45652a5da6b5d83b3e19904703a77bc5", "source_path": "examples/data-sources/xcsh_cdn_loadbalancer/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cdn_loadbalancer:example:data-source", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:examples", "path": "documentation/data-sources/cdn_loadbalancer/examples/data-source/index.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
