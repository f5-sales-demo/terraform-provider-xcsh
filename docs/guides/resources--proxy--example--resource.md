---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 888, "body_sha256": "sha256:05c08f78e1311f5c25adcad768580ae9225cb1afbc0582002e225d075bb23370", "canonical_id": "xcsh-docs:resources:proxy:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8d3ef9470714fab754f0f7631e40f5f93a8f5acfb3c6def25dff74cba83c2ef7", "source_path": "examples/resources/xcsh_proxy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:proxy:example:resource", "parent_id": "xcsh-docs:resources:proxy:examples", "path": "docs/guides/resources--proxy--example--resource.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Examples](resources--proxy--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_proxy/resource.tf`; digest `sha256:8d3ef9470714fab754f0f7631e40f5f93a8f5acfb3c6def25dff74cba83c2ef7`.

```terraform
# Proxy Resource Example
# Manages a Proxy resource in F5 Distributed Cloud for tcp loadbalancer create specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Proxy configuration
resource "xcsh_proxy" "example" {
  name      = "example-proxy"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--proxy--examples.md)
- [xcsh_proxy](../resources/proxy.md)
