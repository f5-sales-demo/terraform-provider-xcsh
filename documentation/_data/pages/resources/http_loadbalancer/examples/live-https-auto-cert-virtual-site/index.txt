---
page_title: "Live https auto cert virtual site"
subcategory: "Load Balancing"
description: "Live https auto cert virtual site for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": ["live-https-auto-cert-virtual-site"], "body_bytes": 1687, "body_sha256": "sha256:3932b2deabfd1b04641baeacd1f77e37ffb505c150fb1b1f5b7ce68744dc6f6c", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:d4a0031e7002f14871ec0e167d4c2e4fbb19bca356c59205380d515763ed3321", "source_path": "examples/resources/xcsh_http_loadbalancer/live-https-auto-cert-virtual-site.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:example:live-https-auto-cert-virtual-site", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "documentation/resources/http_loadbalancer/examples/live-https-auto-cert-virtual-site/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2323223122222222-3210000131013200-0233201102303202-2101213100320203-0110303302122200-0313333001332300-3011033100330122-3110032010230012", "registry_path": "docs/guides/resources--http_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["live-https-auto-cert-virtual-site"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/live-https-auto-cert-virtual-site/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Live https auto cert virtual site for xcsh_http_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Live https auto cert virtual site

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
- Live https auto cert virtual site

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/live-https-auto-cert-virtual-site.tf`; digest `sha256:d4a0031e7002f14871ec0e167d4c2e4fbb19bca356c59205380d515763ed3321`.

```terraform
# LiveHTTPSAutoCertVirtualSite — Acceptance-test-derived Configuration
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

resource "xcsh_virtual_site" "test" {
  name      = "example-description"
  namespace = "example-value"
  site_type = "CUSTOMER_EDGE"

  site_selector {
    expressions = ["site_type=customer_edge"]
  }
}

resource "xcsh_http_loadbalancer" "test" {
  depends_on = [xcsh_virtual_site.test]
  name       = "example"
  namespace  = "example-value"
  domains    = ["test.example.com"]

  https_auto_cert {}

  advertise_custom {
    advertise_where {
      virtual_site {
        network = "SITE_NETWORK_INSIDE_AND_OUTSIDE"
        virtual_site {
          name      = xcsh_virtual_site.test.name
          namespace = "example-value"
        }
      }
      use_default_port = {}
    }
  }
}
```
