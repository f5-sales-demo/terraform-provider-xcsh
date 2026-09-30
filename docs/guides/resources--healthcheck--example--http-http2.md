---
page_title: "Http http2"
subcategory: "Monitoring"
description: "Http http2 for xcsh_healthcheck."
xcsh_docs: {"aliases": [], "body_bytes": 1107, "body_sha256": "sha256:a50474f5381e6ddc873d9f32cd30c146c76a20633348024fae85993ee168dbe1", "canonical_id": "xcsh-docs:resources:healthcheck:example:http-http2", "child_ids": [], "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:7586adcf4b2cb6423fa64e91dd0205c03d4130e45c575432d8fff8d1cf0d1e8f", "source_path": "examples/resources/xcsh_healthcheck/http-http2.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:healthcheck:example:http-http2", "parent_id": "xcsh-docs:resources:healthcheck:examples", "path": "docs/guides/resources--healthcheck--example--http-http2.md", "provider_name": "healthcheck", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["http-http2"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/examples/http-http2/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Http http2 for xcsh_healthcheck.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Http http2

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md)
- [Examples](resources--healthcheck--examples.md)
- Http http2

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/http-http2.tf`; digest `sha256:7586adcf4b2cb6423fa64e91dd0205c03d4130e45c575432d8fff8d1cf0d1e8f`.

```terraform
# HttpHttp2 — Acceptance-test-derived Configuration
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
    path        = "example-value"
    host_header = "example.com"
    use_http2   = true
  }
}
```

## Next pages

- [Examples](resources--healthcheck--examples.md)
- [xcsh_healthcheck](../resources/healthcheck.md)
