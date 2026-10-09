---
page_title: "Multiple origins"
subcategory: "Load Balancing"
description: "Multiple origins for xcsh_origin_pool."
xcsh_docs: {"aliases": ["multiple-origins"], "body_bytes": 1245, "body_sha256": "sha256:0627b4f95b33dec77e33c911a6ed43ac1396baf1a1faf4aa03b7b7cf009d6334", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:bd8a184b9c325d4b7310968c1937c2417c67cf0e0a8113c2bf868b7ccb8905ff", "source_path": "examples/resources/xcsh_origin_pool/multiple-origins.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:origin_pool:example:multiple-origins", "parent_id": "xcsh-docs:resources:origin_pool:examples", "path": "documentation/resources/origin_pool/examples/multiple-origins/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3233022010022103-0131333231223300-2112303032010131-2303302301330233-2232110202113013-3300113312232223-3100330113013101-1232322222220303", "registry_path": "docs/guides/resources--origin_pool--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["multiple-origins"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/examples/multiple-origins/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Multiple origins for xcsh_origin_pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["origin_poolCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Multiple origins

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/examples/)
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
