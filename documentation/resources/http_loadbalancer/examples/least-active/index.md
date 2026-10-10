---
page_title: "Least active"
subcategory: "Load Balancing"
description: "Least active for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": ["least-active"], "body_bytes": 1134, "body_sha256": "sha256:b58850c5326d3ef2211bbb4b938dd222e36d7d82210d05be9488b4b12289a68a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:70324220d9c9fa5c1c5615baf4370d340fb061db4d54cc32d3c9ec11039e1ef2", "source_path": "examples/resources/xcsh_http_loadbalancer/least-active.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:example:least-active", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "documentation/resources/http_loadbalancer/examples/least-active/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3320011031003002-3333211212322330-0030003110001012-1023320020003123-1221211321102033-3132310113012033-1031032132322023-2031323311321203", "registry_path": "docs/guides/resources--http_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["least-active"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/least-active/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Least active for xcsh_http_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Least active

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
- Least active

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/least-active.tf`; digest `sha256:70324220d9c9fa5c1c5615baf4370d340fb061db4d54cc32d3c9ec11039e1ef2`.

```terraform
# LeastActive — Acceptance-test-derived Configuration
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

  least_active = {}

  advertise_on_public_default_vip = {}
}
```
