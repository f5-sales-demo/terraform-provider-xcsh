---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_site_signatures_update."
xcsh_docs: {"aliases": ["action"], "body_bytes": 1196, "body_sha256": "sha256:bd26bcf1be0726e4c30759e6f70caf70228a97804c08792a663cb12b27c1965d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:site_signatures_update:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:b8fc93938ea82d6d38f39cc7c48dce93f107714b65d36ac9a46938636cbed4f9", "source_path": "examples/actions/xcsh_site_signatures_update/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:site_signatures_update:example:action", "parent_id": "xcsh-docs:actions:site_signatures_update:examples", "path": "documentation/actions/site_signatures_update/examples/action/index.md", "product": "distributed-cloud", "provider_name": "site_signatures_update", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "actions", "registry_anchor": "canonical-3020300320201011-0002300300301300-2133123121323132-3231032221131331-1133221011312332-0201133122101132-2303110212031300-0323022321011103", "registry_path": "docs/guides/actions--site_signatures_update--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_signatures_update/examples/action/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Action for xcsh_site_signatures_update.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
