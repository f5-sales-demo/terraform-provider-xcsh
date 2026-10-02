---
page_title: "Labels update"
subcategory: "Load Balancing"
description: "Labels update for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": ["labels-update"], "body_bytes": 1435, "body_sha256": "sha256:332dcd57a458cf1cec4a7963b88464f94a59f1391c3b061a7b28d9b998fafd8e", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:047c3e407fe8f84c02b646d4800c1eace6de8efae5137aae3f4ce8865c70e147", "source_path": "examples/resources/xcsh_http_loadbalancer/labels-update.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:example:labels-update", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "documentation/resources/http_loadbalancer/examples/labels-update/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1310030100232322-2010311111022213-1201131221102123-0031113301000323-0000022033010223-0233103123001000-1122231311231323-1301311130320302", "registry_path": "docs/guides/resources--http_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["labels-update"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/labels-update/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Labels update for xcsh_http_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
