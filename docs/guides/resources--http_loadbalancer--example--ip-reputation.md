---
page_title: "Ip reputation"
subcategory: "Load Balancing"
description: "Ip reputation for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1222, "body_sha256": "sha256:7c2e335a2fc640b27d64560489c45045bd58637537e4d6219c662e7cc9b5343a", "canonical_id": "xcsh-docs:resources:http_loadbalancer:example:ip-reputation", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:dcf6720963e3d25c8204a3d4b94665f5900cbfd21d3a740b92a547a3fa256c53", "source_path": "examples/resources/xcsh_http_loadbalancer/ip-reputation.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:example:ip-reputation", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "docs/guides/resources--http_loadbalancer--example--ip-reputation.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["ip-reputation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/ip-reputation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Ip reputation for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Ip reputation

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Examples](resources--http_loadbalancer--examples.md)
- Ip reputation

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/ip-reputation.tf`; digest `sha256:dcf6720963e3d25c8204a3d4b94665f5900cbfd21d3a740b92a547a3fa256c53`.

```terraform
# IpReputation — Acceptance-test-derived Configuration
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

resource "xcsh_http_loadbalancer" "test" {
  name      = "example"
  namespace = "system"
  domains   = ["test.example.com"]

  http {
    port = 80
  }

  enable_ip_reputation {
    ip_threat_categories = ["SPAM_SOURCES"]
  }

  advertise_on_public_default_vip = {}
}
```

## Next pages

- [Examples](resources--http_loadbalancer--examples.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
