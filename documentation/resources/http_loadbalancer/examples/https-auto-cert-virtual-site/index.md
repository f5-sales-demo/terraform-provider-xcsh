---
page_title: "Https auto cert virtual site"
subcategory: "Load Balancing"
description: "Https auto cert virtual site for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": ["https-auto-cert-virtual-site"], "body_bytes": 1650, "body_sha256": "sha256:5405527b9993c9229db9486dc3d2f017733f271393341bbd2c5e4ea4a225a239", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:88746cf0029b0d4fda1c11f433ad73b5d62d553eb9971ee08c2988ec5e8c543d", "source_path": "examples/resources/xcsh_http_loadbalancer/https-auto-cert-virtual-site.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:example:https-auto-cert-virtual-site", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "documentation/resources/http_loadbalancer/examples/https-auto-cert-virtual-site/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2102202300103120-3301312001320200-0322123133022133-3201323000313020-2112130222312120-3120103101103003-3000003011113212-2310111232112030", "registry_path": "docs/guides/resources--http_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["https-auto-cert-virtual-site"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/https-auto-cert-virtual-site/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Https auto cert virtual site for xcsh_http_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Https auto cert virtual site

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
