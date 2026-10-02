---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_cloud_link."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1228, "body_sha256": "sha256:60e13b97f210d620d77658c8bcbac3e368d0f43450801e5b2ccfc9aa0ed12164", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1be93a99c9a3175561f8ebf72f83a5762c7f54a0a05e770026c4c08e2199ba00", "source_path": "examples/resources/xcsh_cloud_link/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cloud_link:example:resource", "parent_id": "xcsh-docs:resources:cloud_link:examples", "path": "documentation/resources/cloud_link/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1010233102112003-1120030021313310-2013011023303020-3201222103232003-2211131202122210-2010333023323331-2231022231232122-3023120032022030", "registry_path": "docs/guides/resources--cloud_link--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_cloud_link.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cloud_link/resource.tf`; digest `sha256:1be93a99c9a3175561f8ebf72f83a5762c7f54a0a05e770026c4c08e2199ba00`.

```terraform
# CloudLink Resource Example
# Manages new CloudLink with configured parameters in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudLink configuration
resource "xcsh_cloud_link" "example" {
  name      = "example-cloud-link"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/examples/)
- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
