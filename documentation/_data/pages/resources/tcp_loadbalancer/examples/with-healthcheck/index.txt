---
page_title: "With healthcheck"
subcategory: "Load Balancing"
description: "With healthcheck for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": ["with-healthcheck"], "body_bytes": 2084, "body_sha256": "sha256:a6b1f8480475771e98fd67d6af96f3556fe092bc5e02d090f462d83597847d02", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:fdae141ffa6040bc54466938c997fa7b252cc16638c171911504af08dee09ab5", "source_path": "examples/resources/xcsh_tcp_loadbalancer/with-healthcheck.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:tcp_loadbalancer:example:with-healthcheck", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:examples", "path": "documentation/resources/tcp_loadbalancer/examples/with-healthcheck/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1213131201102133-2000302012101102-2030113221011230-0312000232013011-0120022300223021-0301113103321321-3100211002203223-1000133203101100", "registry_path": "docs/guides/resources--tcp_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["with-healthcheck"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/examples/with-healthcheck/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "With healthcheck for xcsh_tcp_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/examples/)
- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
