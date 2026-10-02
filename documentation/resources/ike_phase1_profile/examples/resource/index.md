---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_ike_phase1_profile."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1513, "body_sha256": "sha256:5471004c5cc60b81960e135f0747e31eabf06fb54a0e43e8a0a6eefc53d53833", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike_phase1_profile:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8b0906176c52164cc70b7c3ef3103bf969192548192507a8a9df9fa2b687acf6", "source_path": "examples/resources/xcsh_ike_phase1_profile/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:ike_phase1_profile:example:resource", "parent_id": "xcsh-docs:resources:ike_phase1_profile:examples", "path": "documentation/resources/ike_phase1_profile/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2321102110313013-1133331010000002-2232020212000012-1003221021302333-3001330231230132-0112111103213021-0113020023311031-1003110212100232", "registry_path": "docs/guides/resources--ike_phase1_profile--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike_phase1_profile/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_ike_phase1_profile.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_ike_phase1_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_ike_phase1_profile/resource.tf`; digest `sha256:8b0906176c52164cc70b7c3ef3103bf969192548192507a8a9df9fa2b687acf6`.

```terraform
# IKEPhase1Profile Resource Example
# Manages a IKE Phase1 Profile resource in F5 Distributed Cloud for ike phase1 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic IKEPhase1Profile configuration
resource "xcsh_ike_phase1_profile" "example" {
  name      = "example-ike-phase1-profile"
  namespace = "staging"

  authentication_algos = ["example-value"]
  dh_group             = ["example-value"]
  encryption_algos     = ["example-value"]
  prf                  = ["example-value"]
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/examples/)
- [xcsh_ike_phase1_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/)
