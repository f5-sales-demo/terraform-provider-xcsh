---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_advertise_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1412, "body_sha256": "sha256:d0891a5913c641993eb4af61ed7251400b73a6f014d5e321073c631a3163da6e", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9e9301fb78d45973929d2d48eacb74419a673af009dabc983da6f509200329a9", "source_path": "examples/resources/xcsh_advertise_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:advertise_policy:example:resource", "parent_id": "xcsh-docs:resources:advertise_policy:examples", "path": "documentation/resources/advertise_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1130302312201231-2213332000223123-2201301302211121-1220330133313300-3102122020131231-2010302111013101-2113320200031320-0222303330222321", "registry_path": "docs/guides/resources--advertise_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_advertise_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/examples/)
- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/)
