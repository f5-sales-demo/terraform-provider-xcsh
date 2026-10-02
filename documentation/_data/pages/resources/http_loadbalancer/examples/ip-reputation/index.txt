---
page_title: "Ip reputation"
subcategory: "Load Balancing"
description: "Ip reputation for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": ["ip-reputation"], "body_bytes": 1428, "body_sha256": "sha256:ffe492b180822379ad2b22aefeae57bcd9147731ad14966ca8755af2837cae38", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:dcf6720963e3d25c8204a3d4b94665f5900cbfd21d3a740b92a547a3fa256c53", "source_path": "examples/resources/xcsh_http_loadbalancer/ip-reputation.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:example:ip-reputation", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "documentation/resources/http_loadbalancer/examples/ip-reputation/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3110122013200000-2033131031130223-0220320122010133-1210113221211021-2022112222220110-3312002203321312-2033013031121321-3030122033013033", "registry_path": "docs/guides/resources--http_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["ip-reputation"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/ip-reputation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Ip reputation for xcsh_http_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Ip reputation

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
- Ip reputation

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/ip-reputation.tf`; digest `sha256:dcf6720963e3d25c8204a3d4b94665f5900cbfd21d3a740b92a547a3fa256c53`.

```terraform
# IpReputation — Acceptance-test-derived Configuration
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

  enable_ip_reputation {
    ip_threat_categories = ["SPAM_SOURCES"]
  }

  advertise_on_public_default_vip = {}
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
