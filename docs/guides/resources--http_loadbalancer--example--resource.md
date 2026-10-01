---
page_title: "Resource"
subcategory: "Load Balancing"
description: "Resource for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1192, "body_sha256": "sha256:3b89de38ca8de186c1ee0c5b811f0fe2a57ca8c307b4fc0a61d7f500c887362d", "canonical_id": "xcsh-docs:resources:http_loadbalancer:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:2cc83254a22eded7b491fec9fc0411929e49400732480241f5c8b07d373a95c1", "source_path": "examples/resources/xcsh_http_loadbalancer/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:example:resource", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "docs/guides/resources--http_loadbalancer--example--resource.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Examples](resources--http_loadbalancer--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/resource.tf`; digest `sha256:2cc83254a22eded7b491fec9fc0411929e49400732480241f5c8b07d373a95c1`.

```terraform
# HTTPLoadBalancer Resource Example
# Manages a HTTP Load Balancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with routing and security controls.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic HTTPLoadBalancer configuration
resource "xcsh_http_loadbalancer" "example" {
  name      = "example-http-loadbalancer"
  namespace = "staging"

  domains = ["example-value"]
}
```

## Next pages

- [Examples](resources--http_loadbalancer--examples.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
