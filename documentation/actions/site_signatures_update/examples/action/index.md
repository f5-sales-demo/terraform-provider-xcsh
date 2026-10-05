---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_site_signatures_update."
xcsh_docs: {"aliases": ["action"], "body_bytes": 1196, "body_sha256": "sha256:bd26bcf1be0726e4c30759e6f70caf70228a97804c08792a663cb12b27c1965d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:site_signatures_update:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:b8fc93938ea82d6d38f39cc7c48dce93f107714b65d36ac9a46938636cbed4f9", "source_path": "examples/actions/xcsh_site_signatures_update/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:site_signatures_update:example:action", "parent_id": "xcsh-docs:actions:site_signatures_update:examples", "path": "documentation/actions/site_signatures_update/examples/action/index.md", "product": "distributed-cloud", "provider_name": "site_signatures_update", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "actions", "registry_anchor": "canonical-3020300320201011-0002300300301300-2133123121323132-3231032221131331-1133221011312332-0201133122101132-2303110212031300-0323022321011103", "registry_path": "docs/guides/actions--site_signatures_update--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_signatures_update/examples/action/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Action for xcsh_site_signatures_update.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Action

Breadcrumbs:

- [xcsh_site_signatures_update](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_signatures_update/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_signatures_update/examples/)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_site_signatures_update/action.tf`; digest `sha256:b8fc93938ea82d6d38f39cc7c48dce93f107714b65d36ac9a46938636cbed4f9`.

```terraform
# SiteSignaturesUpdate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_site_signatures_update" "example" {
  config {
    namespace = "example-value"
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_signatures_update/examples/)
- [xcsh_site_signatures_update](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_signatures_update/)
