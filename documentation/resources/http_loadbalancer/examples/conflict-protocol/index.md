---
page_title: "Conflict protocol"
subcategory: "Load Balancing"
description: "Conflict protocol for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": ["conflict-protocol"], "body_bytes": 1550, "body_sha256": "sha256:584c28ec4b156074ff17ab70a5c3b435806f2e2845133e427488ceae61341817", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "expected conflict", "sha256": "sha256:536a4332401ed618974908dabf261dd097e7687d8727db003cdf77a23e114526", "source_path": "examples/resources/xcsh_http_loadbalancer/conflict-protocol.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:negative-example:conflict-protocol", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "documentation/resources/http_loadbalancer/examples/conflict-protocol/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1223222112033332-2232002203232333-3101301121330131-2213023323032210-0223113110311202-2031310032310212-1003230110132011-0113230321111130", "registry_path": "docs/guides/resources--http_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "negative-example", "schema_path": ["conflict-protocol"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/conflict-protocol/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Conflict protocol for xcsh_http_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
