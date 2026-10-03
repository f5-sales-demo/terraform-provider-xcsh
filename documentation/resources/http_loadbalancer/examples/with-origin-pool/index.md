---
page_title: "With origin pool"
subcategory: "Load Balancing"
description: "With origin pool for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "upstream servers", "with-origin-pool"], "body_bytes": 1759, "body_sha256": "sha256:3fdb524a318a00dfcc5ed0bbd5743cd2ad549ddd867bcd239a5580f6e0291dbc", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:108dbeb11bdb7cc69645334a91b9dce00f52f21d75be901d99e0c95036f9f959", "source_path": "examples/resources/xcsh_http_loadbalancer/with-origin-pool.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:example:with-origin-pool", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "documentation/resources/http_loadbalancer/examples/with-origin-pool/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1021031001203102-1232322101103101-3000301221322230-3313031312223200-3223100131230111-0333323031311201-0101001312210011-1013201123131330", "registry_path": "docs/guides/resources--http_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["with-origin-pool"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/with-origin-pool/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "With origin pool for xcsh_http_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With origin pool

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
- With origin pool

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/with-origin-pool.tf`; digest `sha256:108dbeb11bdb7cc69645334a91b9dce00f52f21d75be901d99e0c95036f9f959`.

```terraform
# WithOriginPool — Acceptance-test-derived Configuration
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
  name      = "example"
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

resource "xcsh_http_loadbalancer" "test" {
  name      = "example"
  namespace = "system"

  domains = ["test.example.com"]

  http {
    port = 80
  }

  default_route_pools {
    pool {
      name      = xcsh_origin_pool.test.name
      namespace = "system"
    }
    weight   = 1
    priority = 1
  }

  advertise_on_public_default_vip = {}
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
