---
page_title: "Labels update"
subcategory: "Load Balancing"
description: "Labels update for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": ["labels-update"], "body_bytes": 1198, "body_sha256": "sha256:4c53ca0f772682118b5d1aa950e14c9759488a7c7e11c24650024157f6857442", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:047c3e407fe8f84c02b646d4800c1eace6de8efae5137aae3f4ce8865c70e147", "source_path": "examples/resources/xcsh_http_loadbalancer/labels-update.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:example:labels-update", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "documentation/resources/http_loadbalancer/examples/labels-update/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1310030100232322-2010311111022213-1201131221102123-0031113301000323-0000022033010223-0233103123001000-1122231311231323-1301311130320302", "registry_path": "docs/guides/resources--http_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["labels-update"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/labels-update/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Labels update for xcsh_http_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Labels update

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
- Labels update

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/labels-update.tf`; digest `sha256:047c3e407fe8f84c02b646d4800c1eace6de8efae5137aae3f4ce8865c70e147`.

```terraform
# LabelsUpdate — Acceptance-test-derived Configuration
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

  labels = {
    environment = "example-value"
    managed_by  = "terraform"
  }

  domains = ["test.example.com"]

  http {
    port = 80
  }

  advertise_on_public_default_vip = {}
}
```
