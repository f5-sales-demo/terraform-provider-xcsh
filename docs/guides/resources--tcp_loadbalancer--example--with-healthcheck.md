---
page_title: "With healthcheck"
subcategory: "Load Balancing"
description: "With healthcheck for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1878, "body_sha256": "sha256:ac64df6c63cb42708c38239355d53c1b3bdbea52c19e2faffcd2a880644a6676", "canonical_id": "xcsh-docs:resources:tcp_loadbalancer:example:with-healthcheck", "child_ids": [], "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:fdae141ffa6040bc54466938c997fa7b252cc16638c171911504af08dee09ab5", "source_path": "examples/resources/xcsh_tcp_loadbalancer/with-healthcheck.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:tcp_loadbalancer:example:with-healthcheck", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:examples", "path": "docs/guides/resources--tcp_loadbalancer--example--with-healthcheck.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["with-healthcheck"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/examples/with-healthcheck/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With healthcheck for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With healthcheck

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
- [Examples](resources--tcp_loadbalancer--examples.md)
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

- [Examples](resources--tcp_loadbalancer--examples.md)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
