---
page_title: "Https auto cert"
subcategory: "Load Balancing"
description: "Https auto cert for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1431, "body_sha256": "sha256:2ff03ad2ae5e7722a96de44f076c4068b4e21f6a7ce74a16cd0e9c228ef15c7d", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:0d662f3b7cc004a1f0f8a4fb17d79cd9f87b6670b4fb87c8b2fac5e03df5a58a", "source_path": "examples/resources/xcsh_http_loadbalancer/https-auto-cert.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:example:https-auto-cert", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "documentation/resources/http_loadbalancer/examples/https-auto-cert/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["https-auto-cert"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/https-auto-cert/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Https auto cert for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Https auto cert

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
- Https auto cert

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/https-auto-cert.tf`; digest `sha256:0d662f3b7cc004a1f0f8a4fb17d79cd9f87b6670b4fb87c8b2fac5e03df5a58a`.

```terraform
# HttpsAutoCert — Acceptance-test-derived Configuration
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

  domains = ["test.example.com"]

  https_auto_cert {
    add_hsts                 = false
    no_mtls                  = {}
    default_header           = {}
    enable_path_normalize    = {}
    non_default_loadbalancer = {}
  }

  advertise_on_public_default_vip = {}
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
