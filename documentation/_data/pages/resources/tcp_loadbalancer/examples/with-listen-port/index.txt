---
page_title: "With listen port"
subcategory: "Load Balancing"
description: "With listen port for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": ["with-listen-port"], "body_bytes": 1615, "body_sha256": "sha256:d7e8467814ab15710b332b44c6dda65c111e5348eb98553f0e38a6e7e5100780", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:418207da4ff4973dbac54d60c3cd0f60db919fc512e95339fa402a4a71300827", "source_path": "examples/resources/xcsh_tcp_loadbalancer/with-listen-port.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:tcp_loadbalancer:example:with-listen-port", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:examples", "path": "documentation/resources/tcp_loadbalancer/examples/with-listen-port/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1301223213021230-0103300221332112-2210310131020331-0222120223000330-1210132111201310-0001333321020232-2231032313320311-1033021130000002", "registry_path": "docs/guides/resources--tcp_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["with-listen-port"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/examples/with-listen-port/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "With listen port for xcsh_tcp_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With listen port

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/examples/)
- With listen port

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_tcp_loadbalancer/with-listen-port.tf`; digest `sha256:418207da4ff4973dbac54d60c3cd0f60db919fc512e95339fa402a4a71300827`.

```terraform
# WithListenPort — Acceptance-test-derived Configuration
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

resource "xcsh_origin_pool" "test" {
  name      = "example-pool"
  namespace = "system"
  port      = 443

  origin_servers {
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}

resource "xcsh_tcp_loadbalancer" "test" {
  name      = "example"
  namespace = "system"

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  domains     = ["example.example.com"]
  listen_port = 443

  tcp = {}
  sni = {}

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
