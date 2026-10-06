---
page_title: "With description"
subcategory: "Load Balancing"
description: "With description for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": ["with-description"], "body_bytes": 1667, "body_sha256": "sha256:81b65b7f47f57db14d7602cb348ad95d54e301752109beaf61712d7f5045c8c5", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:5684729ca1aa2364ebb7a66f862baa44ce77ded7c47b4875b858bfb0298f3b7a", "source_path": "examples/resources/xcsh_tcp_loadbalancer/with-description.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:tcp_loadbalancer:example:with-description", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:examples", "path": "documentation/resources/tcp_loadbalancer/examples/with-description/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2332333133312100-2313121003203331-2133101333332332-1022213331101312-0002032003132003-3030312102032312-0030333122031122-0021031112031122", "registry_path": "docs/guides/resources--tcp_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["with-description"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/examples/with-description/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "With description for xcsh_tcp_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With description

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/examples/)
- With description

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_tcp_loadbalancer/with-description.tf`; digest `sha256:5684729ca1aa2364ebb7a66f862baa44ce77ded7c47b4875b858bfb0298f3b7a`.

```terraform
# WithDescription — Acceptance-test-derived Configuration
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
  name        = "example"
  namespace   = "system"
  description = "example-value"

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

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
