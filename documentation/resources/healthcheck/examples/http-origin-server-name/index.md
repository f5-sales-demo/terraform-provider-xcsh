---
page_title: "Http origin server name"
subcategory: "Monitoring"
description: "Http origin server name for xcsh_healthcheck."
xcsh_docs: {"aliases": ["backend servers", "http-origin-server-name", "origin servers", "upstream servers"], "body_bytes": 1450, "body_sha256": "sha256:81d42512e85526afa524b88bbc84dddefb1d5662a84cc7b77dfdef367268a457", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:afde24e907d59c17ebd5e445c7c324b1ec4a0b83e3d3650974943970ea9a430a", "source_path": "examples/resources/xcsh_healthcheck/http-origin-server-name.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:healthcheck:example:http-origin-server-name", "parent_id": "xcsh-docs:resources:healthcheck:examples", "path": "documentation/resources/healthcheck/examples/http-origin-server-name/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3302000302103322-3133323211121002-2013031220203013-3203002312133122-2020010112003311-3130323123110230-3023002320103222-3122012102133021", "registry_path": "docs/guides/resources--healthcheck--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["http-origin-server-name"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/examples/http-origin-server-name/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Http origin server name for xcsh_healthcheck.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Http origin server name

Breadcrumbs:

- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/examples/)
- Http origin server name

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/http-origin-server-name.tf`; digest `sha256:afde24e907d59c17ebd5e445c7c324b1ec4a0b83e3d3650974943970ea9a430a`.

```terraform
# HttpOriginServerName — Acceptance-test-derived Configuration
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

resource "xcsh_healthcheck" "test" {
  name      = "example"
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  http_health_check {
    path                   = "example-value"
    use_origin_server_name = {}
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/examples/)
- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
