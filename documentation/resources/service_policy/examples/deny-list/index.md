---
page_title: "Deny list"
subcategory: "Security"
description: "Deny list for xcsh_service_policy."
xcsh_docs: {"aliases": ["deny-list"], "body_bytes": 1338, "body_sha256": "sha256:2a460ca002e53f6517bfdb83836ae574004ac92f4bc1432836052444fcbf0c1d", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:6000e015862d1bb25f8e6b5bc94664685fdab920f11f2608c4f47416c65ff393", "source_path": "examples/resources/xcsh_service_policy/deny-list.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:service_policy:example:deny-list", "parent_id": "xcsh-docs:resources:service_policy:examples", "path": "documentation/resources/service_policy/examples/deny-list/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1020102021132323-1231031122030310-2211013123032313-2032332201010211-2203101100303003-3003023331132200-0000320211100120-2102211112223003", "registry_path": "docs/guides/resources--service_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["deny-list"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/examples/deny-list/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Deny list for xcsh_service_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Deny list

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/examples/)
- Deny list

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy/deny-list.tf`; digest `sha256:6000e015862d1bb25f8e6b5bc94664685fdab920f11f2608c4f47416c65ff393`.

```terraform
# DenyList — Acceptance-test-derived Configuration
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

resource "xcsh_service_policy" "test" {
  name      = "example"
  namespace = "system"

  deny_list {
    prefix_list {
      prefixes = ["172.16.0.0/12"]
    }
    default_action_allow = {}
  }

  any_server = {}
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/examples/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
