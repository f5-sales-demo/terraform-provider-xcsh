---
page_title: "Public ip"
subcategory: "Load Balancing"
description: "Public ip for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1124, "body_sha256": "sha256:bdd5c47f4b989aa7bc668379d8bf50713cf89885c7e3a28c0dc1255d6acdbee6", "canonical_id": "xcsh-docs:resources:origin_pool:example:public-ip", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:1dade891c9c99d46213e707ba46387f7c7b5eed90053ec8616b5c928911c7d54", "source_path": "examples/resources/xcsh_origin_pool/public-ip.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:origin_pool:example:public-ip", "parent_id": "xcsh-docs:resources:origin_pool:examples", "path": "docs/guides/resources--origin_pool--example--public-ip.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["public-ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/examples/public-ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Public ip for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Public ip

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Examples](resources--origin_pool--examples.md)
- Public ip

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/public-ip.tf`; digest `sha256:1dade891c9c99d46213e707ba46387f7c7b5eed90053ec8616b5c928911c7d54`.

```terraform
# PublicIp — Acceptance-test-derived Configuration
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

  port = 8080

  origin_servers {
    public_ip {
      ip = "192.0.2.1"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}
```

## Next pages

- [Examples](resources--origin_pool--examples.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
