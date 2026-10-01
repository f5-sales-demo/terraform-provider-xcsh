---
page_title: "Multiple origins"
subcategory: "Load Balancing"
description: "Multiple origins for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1258, "body_sha256": "sha256:b393abaef2885da3005849724a886578b2dbcb6536b34737c1b5e47a27a84385", "canonical_id": "xcsh-docs:resources:origin_pool:example:multiple-origins", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:bd8a184b9c325d4b7310968c1937c2417c67cf0e0a8113c2bf868b7ccb8905ff", "source_path": "examples/resources/xcsh_origin_pool/multiple-origins.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:origin_pool:example:multiple-origins", "parent_id": "xcsh-docs:resources:origin_pool:examples", "path": "docs/guides/resources--origin_pool--example--multiple-origins.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["multiple-origins"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/examples/multiple-origins/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Multiple origins for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Multiple origins

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Examples](resources--origin_pool--examples.md)
- Multiple origins

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/multiple-origins.tf`; digest `sha256:bd8a184b9c325d4b7310968c1937c2417c67cf0e0a8113c2bf868b7ccb8905ff`.

```terraform
# MultipleOrigins — Acceptance-test-derived Configuration
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

  port = 443

  origin_servers {
    public_name {
      dns_name = "backend1.example.com"
    }
  }

  origin_servers {
    public_name {
      dns_name = "backend2.example.com"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}
```

## Next pages

- [Examples](resources--origin_pool--examples.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
