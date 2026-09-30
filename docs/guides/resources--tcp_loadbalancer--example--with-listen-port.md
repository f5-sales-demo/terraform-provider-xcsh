---
page_title: "With listen port"
subcategory: "Load Balancing"
description: "With listen port for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1544, "body_sha256": "sha256:2d676d7d67829d7362cc99b681039cfde68d141991d653e863449284581abf40", "canonical_id": "xcsh-docs:resources:tcp_loadbalancer:example:with-listen-port", "child_ids": [], "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:418207da4ff4973dbac54d60c3cd0f60db919fc512e95339fa402a4a71300827", "source_path": "examples/resources/xcsh_tcp_loadbalancer/with-listen-port.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:tcp_loadbalancer:example:with-listen-port", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:examples", "path": "docs/guides/resources--tcp_loadbalancer--example--with-listen-port.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["with-listen-port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/examples/with-listen-port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With listen port for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# With listen port

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
- [Examples](resources--tcp_loadbalancer--examples.md)
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

## Next pages

- [Examples](resources--tcp_loadbalancer--examples.md)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
