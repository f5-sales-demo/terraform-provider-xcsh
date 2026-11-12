---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_filter_set."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1018, "body_sha256": "sha256:dc453a8e3445fbe23ea677f778b41257e9712e839066a5904c11fc15bf2dccd8", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:filter_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:034749c02446cbe5e781e600fc1e24249a070ab78e205de7268da87fdb24c610", "source_path": "examples/resources/xcsh_filter_set/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:filter_set:example:resource", "parent_id": "xcsh-docs:resources:filter_set:examples", "path": "documentation/resources/filter_set/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "filter_set", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-2313012200213310-2312010001331203-2320031131200101-3010310321330001-2010300010123111-0202311102003120-3021311011232101-0012112122113133", "registry_path": "docs/guides/resources--filter_set--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/filter_set/examples/resource/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Resource for xcsh_filter_set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["filter_setCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_filter_set/resource.tf`; digest `sha256:034749c02446cbe5e781e600fc1e24249a070ab78e205de7268da87fdb24c610`.

```terraform
# FilterSet Resource Example
# Manages specification in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic FilterSet configuration
resource "xcsh_filter_set" "example" {
  name      = "example-filter-set"
  namespace = "staging"

  context_key = "example-value"
}
```
