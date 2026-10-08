---
page_title: "Source ip stickiness"
subcategory: "Load Balancing"
description: "Source ip stickiness for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": ["source-ip-stickiness"], "body_bytes": 1173, "body_sha256": "sha256:c014b59c1d8d4bbdbd2994816b75bc890fb985ef9e46d539404dcd8351632ab2", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:739e9d735e5602a3b26f0711cc74f66e68a750409c6a597ce93f240dc10a68a7", "source_path": "examples/resources/xcsh_http_loadbalancer/source-ip-stickiness.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:example:source-ip-stickiness", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "documentation/resources/http_loadbalancer/examples/source-ip-stickiness/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1123330203031122-0213112100022302-2131233022233300-2211010100313212-3202302233332233-0021210322133220-1232023231200121-3211022211030331", "registry_path": "docs/guides/resources--http_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["source-ip-stickiness"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/source-ip-stickiness/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Source ip stickiness for xcsh_http_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Source ip stickiness

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
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
