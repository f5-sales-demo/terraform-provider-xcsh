---
page_title: "With origin pool"
subcategory: "Load Balancing"
description: "With origin pool for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "upstream servers", "with-origin-pool"], "body_bytes": 1759, "body_sha256": "sha256:3fdb524a318a00dfcc5ed0bbd5743cd2ad549ddd867bcd239a5580f6e0291dbc", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:108dbeb11bdb7cc69645334a91b9dce00f52f21d75be901d99e0c95036f9f959", "source_path": "examples/resources/xcsh_http_loadbalancer/with-origin-pool.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:example:with-origin-pool", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "documentation/resources/http_loadbalancer/examples/with-origin-pool/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1021031001203102-1232322101103101-3000301221322230-3313031312223200-3223100131230111-0333323031311201-0101001312210011-1013201123131330", "registry_path": "docs/guides/resources--http_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["with-origin-pool"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/with-origin-pool/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "With origin pool for xcsh_http_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
