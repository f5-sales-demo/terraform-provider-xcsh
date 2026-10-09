---
page_title: "With healthcheck"
subcategory: "Load Balancing"
description: "With healthcheck for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": ["with-healthcheck"], "body_bytes": 1850, "body_sha256": "sha256:a82c618f3279ea7c34421d67e0fed5694e0111e526c38f3d9a0272903f5bc0f4", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:fdae141ffa6040bc54466938c997fa7b252cc16638c171911504af08dee09ab5", "source_path": "examples/resources/xcsh_tcp_loadbalancer/with-healthcheck.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:tcp_loadbalancer:example:with-healthcheck", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:examples", "path": "documentation/resources/tcp_loadbalancer/examples/with-healthcheck/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1213131201102133-2000302012101102-2030113221011230-0312000232013011-0120022300223021-0301113103321321-3100211002203223-1000133203101100", "registry_path": "docs/guides/resources--tcp_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["with-healthcheck"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/examples/with-healthcheck/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "With healthcheck for xcsh_tcp_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With healthcheck

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/examples/)
- With healthcheck

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_tcp_loadbalancer/with-healthcheck.tf`; digest `sha256:fdae141ffa6040bc54466938c997fa7b252cc16638c171911504af08dee09ab5`.

```terraform
# WithHealthcheck — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_healthcheck" "test" {
  name      = "example-hc"
  namespace = "system"

  healthy_threshold   = 3
  unhealthy_threshold = 1
  timeout             = 3
  interval            = 15

  tcp_health_check {}
}

resource "xcsh_origin_pool" "test" {
  name      = "example-pool"
  namespace = "system"
  port      = 443

  origin_servers {
    public_name {
      dns_name = "example.com"
    }
  }

  healthcheck {
    name      = xcsh_healthcheck.test.name
    namespace = "system"
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}

resource "xcsh_tcp_loadbalancer" "test" {
  name      = "example"
  namespace = "system"

  domains     = ["example.example.com"]
  listen_port = 443
  tcp         = {}
  sni         = {}

  origin_pools_weights {
    pool {
      name      = xcsh_origin_pool.test.name
      namespace = "system"
    }
    weight = 1
  }

  advertise_on_public_default_vip = {}
}
```
