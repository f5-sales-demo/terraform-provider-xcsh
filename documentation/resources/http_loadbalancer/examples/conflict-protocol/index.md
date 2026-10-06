---
page_title: "Conflict protocol"
subcategory: "Load Balancing"
description: "Conflict protocol for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": ["conflict-protocol"], "body_bytes": 1313, "body_sha256": "sha256:1cac8cfca0e20ddaae033128cb111629738b6cedda477add7f72b0f54d913b74", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "expected conflict", "sha256": "sha256:536a4332401ed618974908dabf261dd097e7687d8727db003cdf77a23e114526", "source_path": "examples/resources/xcsh_http_loadbalancer/conflict-protocol.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:negative-example:conflict-protocol", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "documentation/resources/http_loadbalancer/examples/conflict-protocol/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1223222112033332-2232002203232333-3101301121330131-2213023323032210-0223113110311202-2031310032310212-1003230110132011-0113230321111130", "registry_path": "docs/guides/resources--http_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "negative-example", "schema_path": ["conflict-protocol"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/conflict-protocol/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Conflict protocol for xcsh_http_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Conflict protocol

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
- Conflict protocol

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **expected conflict**.

Source: `examples/resources/xcsh_http_loadbalancer/conflict-protocol.tf`; digest `sha256:536a4332401ed618974908dabf261dd097e7687d8727db003cdf77a23e114526`.

```terraform
# ConflictProtocol — Negative Configuration Example
# Acceptance-test-derived conflict fixture; not a successful configuration.

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
