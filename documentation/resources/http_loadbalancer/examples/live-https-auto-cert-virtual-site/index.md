---
page_title: "Live https auto cert virtual site"
subcategory: "Load Balancing"
description: "Live https auto cert virtual site for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": ["live-https-auto-cert-virtual-site"], "body_bytes": 1924, "body_sha256": "sha256:2bc4dcf66090a7bc1c2b79896b5cf3e65e3bba6ea386d4c6a3b5f89624d6a2b5", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:d4a0031e7002f14871ec0e167d4c2e4fbb19bca356c59205380d515763ed3321", "source_path": "examples/resources/xcsh_http_loadbalancer/live-https-auto-cert-virtual-site.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:example:live-https-auto-cert-virtual-site", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "documentation/resources/http_loadbalancer/examples/live-https-auto-cert-virtual-site/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2323223122222222-3210000131013200-0233201102303202-2101213100320203-0110303302122200-0313333001332300-3011033100330122-3110032010230012", "registry_path": "docs/guides/resources--http_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["live-https-auto-cert-virtual-site"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/live-https-auto-cert-virtual-site/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Live https auto cert virtual site for xcsh_http_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
