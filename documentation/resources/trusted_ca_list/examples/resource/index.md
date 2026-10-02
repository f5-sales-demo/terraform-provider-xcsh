---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_trusted_ca_list."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1317, "body_sha256": "sha256:d9e9cab79e96ac2459302e46a8bd23ae91ad8429e1fe64340a1627d47455a65d", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:trusted_ca_list:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:98ef39907d2777d4837a31daa9bde5346e1139b6bb15f38e7c1bd69f44aac328", "source_path": "examples/resources/xcsh_trusted_ca_list/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:trusted_ca_list:example:resource", "parent_id": "xcsh-docs:resources:trusted_ca_list:examples", "path": "documentation/resources/trusted_ca_list/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "trusted_ca_list", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2103021213110323-1232123021100101-0100111023203031-2120302130122313-1013031001112213-0112213033332111-0122031310202201-0001213323302301", "registry_path": "docs/guides/resources--trusted_ca_list--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/trusted_ca_list/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_trusted_ca_list.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["trusted_ca_listCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/trusted_ca_list/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/trusted_ca_list/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_trusted_ca_list/resource.tf`; digest `sha256:98ef39907d2777d4837a31daa9bde5346e1139b6bb15f38e7c1bd69f44aac328`.

```terraform
# TrustedCAList Resource Example
# Manages a Trusted CA List resource in F5 Distributed Cloud for trusted certificate authority list management.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic TrustedCAList configuration
resource "xcsh_trusted_ca_list" "example" {
  name      = "example-trusted-ca-list"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/trusted_ca_list/examples/)
- [xcsh_trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/trusted_ca_list/)
