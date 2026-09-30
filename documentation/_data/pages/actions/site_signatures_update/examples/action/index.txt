---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_site_signatures_update."
xcsh_docs: {"aliases": [], "body_bytes": 1097, "body_sha256": "sha256:dcc46dac83d5c6af7502e8b37b8b26caa26c80e5678782a1eaf04dfd75e926c7", "child_ids": [], "collection_id": "xcsh-docs:actions:site_signatures_update:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:b8fc93938ea82d6d38f39cc7c48dce93f107714b65d36ac9a46938636cbed4f9", "source_path": "examples/actions/xcsh_site_signatures_update/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:site_signatures_update:example:action", "parent_id": "xcsh-docs:actions:site_signatures_update:examples", "path": "documentation/actions/site_signatures_update/examples/action/index.md", "provider_name": "site_signatures_update", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "actions", "role": "example", "schema_path": ["action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_signatures_update/examples/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Action for xcsh_site_signatures_update.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
