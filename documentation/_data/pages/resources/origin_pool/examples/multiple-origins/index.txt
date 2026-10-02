---
page_title: "Multiple origins"
subcategory: "Load Balancing"
description: "Multiple origins for xcsh_origin_pool."
xcsh_docs: {"aliases": ["multiple-origins"], "body_bytes": 1464, "body_sha256": "sha256:81734acc3ac79aaaa091b6e5cc21e48cc638adfe998a22898c6d3b01a65425e2", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:bd8a184b9c325d4b7310968c1937c2417c67cf0e0a8113c2bf868b7ccb8905ff", "source_path": "examples/resources/xcsh_origin_pool/multiple-origins.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:origin_pool:example:multiple-origins", "parent_id": "xcsh-docs:resources:origin_pool:examples", "path": "documentation/resources/origin_pool/examples/multiple-origins/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3233022010022103-0131333231223300-2112303032010131-2303302301330233-2232110202113013-3300113312232223-3100330113013101-1232322222220303", "registry_path": "docs/guides/resources--origin_pool--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["multiple-origins"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/examples/multiple-origins/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Multiple origins for xcsh_origin_pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["origin_poolCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/examples/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
