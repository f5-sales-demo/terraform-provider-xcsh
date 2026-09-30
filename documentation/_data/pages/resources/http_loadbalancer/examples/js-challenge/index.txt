---
page_title: "Js challenge"
subcategory: "Load Balancing"
description: "Js challenge for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1327, "body_sha256": "sha256:c4da2fcc3917a0b9e2369b64b13b380866647623fe4fbe5a051717efab585dc7", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:d2d6735f1a5073a1a55a04da2ce4e35a1688038eb16678dc8de5248c6cb894e4", "source_path": "examples/resources/xcsh_http_loadbalancer/js-challenge.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:example:js-challenge", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "documentation/resources/http_loadbalancer/examples/js-challenge/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["js-challenge"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/js-challenge/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Js challenge for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Js challenge

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
- Js challenge

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/js-challenge.tf`; digest `sha256:d2d6735f1a5073a1a55a04da2ce4e35a1688038eb16678dc8de5248c6cb894e4`.

```terraform
# JsChallenge — Acceptance-test-derived Configuration
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

  js_challenge {
    js_script_delay = 5000
    cookie_expiry   = 3600
  }

  advertise_on_public_default_vip = {}
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/examples/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
