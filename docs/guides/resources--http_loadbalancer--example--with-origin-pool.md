---
page_title: "With origin pool"
subcategory: "Load Balancing"
description: "With origin pool for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1553, "body_sha256": "sha256:fd25aecaba8c84c40e4d5396726eec9de8ea9bd7d819c5e12f2f2c676a914f74", "canonical_id": "xcsh-docs:resources:http_loadbalancer:example:with-origin-pool", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:108dbeb11bdb7cc69645334a91b9dce00f52f21d75be901d99e0c95036f9f959", "source_path": "examples/resources/xcsh_http_loadbalancer/with-origin-pool.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:example:with-origin-pool", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "docs/guides/resources--http_loadbalancer--example--with-origin-pool.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["with-origin-pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/with-origin-pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With origin pool for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With origin pool

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Examples](resources--http_loadbalancer--examples.md)
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

- [Examples](resources--http_loadbalancer--examples.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
