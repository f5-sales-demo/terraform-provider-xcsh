---
page_title: "Https auto cert virtual site"
subcategory: "Load Balancing"
description: "Https auto cert virtual site for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1345, "body_sha256": "sha256:1ac42634ce4f0b5edfc3aac7c0285fa194360ba03fa4c0f2a71bf1404d2dbb4a", "canonical_id": "xcsh-docs:resources:http_loadbalancer:example:https-auto-cert-virtual-site", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:88746cf0029b0d4fda1c11f433ad73b5d62d553eb9971ee08c2988ec5e8c543d", "source_path": "examples/resources/xcsh_http_loadbalancer/https-auto-cert-virtual-site.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:example:https-auto-cert-virtual-site", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "docs/guides/resources--http_loadbalancer--example--https-auto-cert-virtual-site.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["https-auto-cert-virtual-site"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/https-auto-cert-virtual-site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Https auto cert virtual site for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Https auto cert virtual site

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Examples](resources--http_loadbalancer--examples.md)
- Https auto cert virtual site

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/https-auto-cert-virtual-site.tf`; digest `sha256:88746cf0029b0d4fda1c11f433ad73b5d62d553eb9971ee08c2988ec5e8c543d`.

```terraform
# HttpsAutoCertVirtualSite — Acceptance-test-derived Configuration
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
  namespace = "example-value"
  domains   = ["test.example.com"]

  https_auto_cert {}

  advertise_custom {
    advertise_where {
      virtual_site {
        network = "SITE_NETWORK_INSIDE_AND_OUTSIDE"
        virtual_site {
          name      = "example-description"
          namespace = "example-value"
        }
      }
      use_default_port = {}
    }
  }
}
```

## Next pages

- [Examples](resources--http_loadbalancer--examples.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
