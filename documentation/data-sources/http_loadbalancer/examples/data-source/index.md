---
page_title: "Data source"
subcategory: "Load Balancing"
description: "Data source for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1344, "body_sha256": "sha256:a8caba3781f863886f941e13a246a14ee977b646350347706d27234a93a39087", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f7d3fa4bb60b803988578a72b5dac4a76f304bac2925ae4fdfe075100bed37ad", "source_path": "examples/data-sources/xcsh_http_loadbalancer/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:http_loadbalancer:example:data-source", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:examples", "path": "documentation/data-sources/http_loadbalancer/examples/data-source/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
