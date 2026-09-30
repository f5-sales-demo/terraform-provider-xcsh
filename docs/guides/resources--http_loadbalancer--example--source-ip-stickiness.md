---
page_title: "Source ip stickiness"
subcategory: "Load Balancing"
description: "Source ip stickiness for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1105, "body_sha256": "sha256:328517e6687087ae3d29276a27812dc28ca5a67ca37135146fe551e0fbc3ddb3", "canonical_id": "xcsh-docs:resources:http_loadbalancer:example:source-ip-stickiness", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:739e9d735e5602a3b26f0711cc74f66e68a750409c6a597ce93f240dc10a68a7", "source_path": "examples/resources/xcsh_http_loadbalancer/source-ip-stickiness.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:example:source-ip-stickiness", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "docs/guides/resources--http_loadbalancer--example--source-ip-stickiness.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["source-ip-stickiness"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/source-ip-stickiness/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Source ip stickiness for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Source ip stickiness

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Examples](resources--http_loadbalancer--examples.md)
- Source ip stickiness

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/source-ip-stickiness.tf`; digest `sha256:739e9d735e5602a3b26f0711cc74f66e68a750409c6a597ce93f240dc10a68a7`.

```terraform
# SourceIpStickiness — Acceptance-test-derived Configuration
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

  source_ip_stickiness = {}

  advertise_on_public_default_vip = {}
}
```

## Next pages

- [Examples](resources--http_loadbalancer--examples.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
