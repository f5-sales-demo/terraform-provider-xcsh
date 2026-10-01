---
page_title: "Port"
subcategory: "Load Balancing"
description: "Port for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1114, "body_sha256": "sha256:33df0c6a65368c9200fbda739acbdebb39d97ff9bfde4fba4a2c057af59366bc", "canonical_id": "xcsh-docs:resources:origin_pool:example:port", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:3852ca8c70ab22d421232a89d1c20cf56857ad2339faf797aa538bea24c6a139", "source_path": "examples/resources/xcsh_origin_pool/port.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:origin_pool:example:port", "parent_id": "xcsh-docs:resources:origin_pool:examples", "path": "docs/guides/resources--origin_pool--example--port.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/examples/port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Port for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Port

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Examples](resources--origin_pool--examples.md)
- Port

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/port.tf`; digest `sha256:3852ca8c70ab22d421232a89d1c20cf56857ad2339faf797aa538bea24c6a139`.

```terraform
# Port — Acceptance-test-derived Configuration
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
      dns_name = "example.com"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}
```

## Next pages

- [Examples](resources--origin_pool--examples.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
