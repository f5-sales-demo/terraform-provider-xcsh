---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_advertise_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1178, "body_sha256": "sha256:0e1a479c6cc26220e92aa3ff3ed8fd1007d6ae15142abbe5d0ed62ac2a180eba", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9e9301fb78d45973929d2d48eacb74419a673af009dabc983da6f509200329a9", "source_path": "examples/resources/xcsh_advertise_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:advertise_policy:example:resource", "parent_id": "xcsh-docs:resources:advertise_policy:examples", "path": "documentation/resources/advertise_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1130302312201231-2213332000223123-2201301302211121-1220330133313300-3102122020131231-2010302111013101-2113320200031320-0222303330222321", "registry_path": "docs/guides/resources--advertise_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Resource for xcsh_advertise_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_advertise_policy/resource.tf`; digest `sha256:9e9301fb78d45973929d2d48eacb74419a673af009dabc983da6f509200329a9`.

```terraform
# AdvertisePolicy Resource Example
# Manages a Advertise Policy resource in F5 Distributed Cloud for advertise_policy object controls how and where a service represented by a given virtual_host object is advertised to consumers.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AdvertisePolicy configuration
resource "xcsh_advertise_policy" "example" {
  name      = "example-advertise-policy"
  namespace = "staging"
}
```
