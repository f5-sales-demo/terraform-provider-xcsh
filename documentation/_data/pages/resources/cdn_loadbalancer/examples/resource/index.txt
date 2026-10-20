---
page_title: "Resource"
subcategory: "Load Balancing"
description: "Resource for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1137, "body_sha256": "sha256:d7c66a50db2c3db4f4fb1ad8c35a2aae0c25e59b0bc04efe1f861c063fd1296f", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9dbe7d10d66d8f67145599515660dfb082a890fef657898fc4e2002f5c72f6dd", "source_path": "examples/resources/xcsh_cdn_loadbalancer/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cdn_loadbalancer:example:resource", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:examples", "path": "documentation/resources/cdn_loadbalancer/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2233222001213003-1100302333323121-2230023101203200-1030021131132032-2210331310200012-1001001122001323-3312010220131211-0113301100112130", "registry_path": "docs/guides/resources--cdn_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/examples/resource/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Resource for xcsh_cdn_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cdn_loadbalancer/resource.tf`; digest `sha256:9dbe7d10d66d8f67145599515660dfb082a890fef657898fc4e2002f5c72f6dd`.

```terraform
# CDNLoadBalancer Resource Example
# Manages a CDN Load Balancer resource in F5 Distributed Cloud for content delivery and edge caching with load balancing.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CDNLoadBalancer configuration
resource "xcsh_cdn_loadbalancer" "example" {
  name      = "example-cdn-loadbalancer"
  namespace = "staging"

  domains = ["example-value"]
}
```
