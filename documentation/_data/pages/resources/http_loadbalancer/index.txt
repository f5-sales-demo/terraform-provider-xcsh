---
page_title: "xcsh_http_loadbalancer"
subcategory: "Load Balancing"
description: "Manages an HTTP Load Balancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with routing and security controls."
xcsh_docs: {"aliases": ["http loadbalancer"], "body_bytes": 2323, "body_sha256": "sha256:c6d9e2f80ff47be12a5a531fd88b3729eec2260a5e5aba2bfcf18af78359be39", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:reference", "xcsh-docs:resources:http_loadbalancer:examples", "xcsh-docs:resources:http_loadbalancer:import", "xcsh-docs:resources:http_loadbalancer:timeouts", "xcsh-docs:resources:http_loadbalancer:lifecycle"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/http_loadbalancer/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203", "registry_path": "docs/resources/http_loadbalancer.md", "relationships": [{"anchor": "", "enforcement": "upstream-advisory", "source": "receipt-pinned-dependency:required", "target_id": "xcsh-docs:resources:origin_pool:fundamentals", "type": "advisory"}, {"anchor": "", "enforcement": "upstream-advisory", "source": "receipt-pinned-dependency:optional", "target_id": "xcsh-docs:resources:healthcheck:fundamentals", "type": "advisory"}, {"anchor": "", "enforcement": "upstream-advisory", "source": "receipt-pinned-dependency:optional", "target_id": "xcsh-docs:resources:app_firewall:fundamentals", "type": "advisory"}, {"anchor": "", "enforcement": "upstream-advisory", "source": "receipt-pinned-dependency:optional", "target_id": "xcsh-docs:resources:certificate:fundamentals", "type": "advisory"}, {"anchor": "", "enforcement": "upstream-advisory", "source": "receipt-pinned-dependency:optional", "target_id": "xcsh-docs:resources:rate_limiter:fundamentals", "type": "advisory"}, {"anchor": "", "enforcement": "upstream-advisory", "source": "receipt-pinned-dependency:optional", "target_id": "xcsh-docs:resources:service_policy:fundamentals", "type": "advisory"}], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Manages an HTTP Load Balancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with routing and security controls.", "tasks": ["configuration"], "unresolved_relationships": [{"name": "bot_defense_policy", "source": "receipt-pinned-dependency:optional"}], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_http_loadbalancer

Breadcrumbs:

- xcsh_http_loadbalancer

Manages an HTTP Load Balancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic
with routing and security controls.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `origin_pool`.

Optional integrations: `healthcheck`, `app_firewall`, `certificate`, `rate_limiter`, `service_policy`, `bot_defense_policy`.

- origin_pool: Backend servers for traffic distribution

- app_firewall: WAF protection (requires WAAP subscription)

- healthcheck: Monitor backend availability

- certificate: TLS termination for HTTPS

- rate_limiter: Protect against traffic spikes

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `domains`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/lifecycle/timeouts/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/lifecycle/)
