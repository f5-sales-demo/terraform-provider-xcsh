---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_access_active_session_terminate."
xcsh_docs: {"aliases": ["action"], "body_bytes": 1308, "body_sha256": "sha256:24eba07232aa25b6f07f586bf13069d2559bbf807ffc528e47abaa55ade97797", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:access_active_session_terminate:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1bbb807c7b4d5333250c2b20cf18f88631715519e818829cb7766f98ee751aa7", "source_path": "examples/actions/xcsh_access_active_session_terminate/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:access_active_session_terminate:example:action", "parent_id": "xcsh-docs:actions:access_active_session_terminate:examples", "path": "documentation/actions/access_active_session_terminate/examples/action/index.md", "product": "distributed-cloud", "provider_name": "access_active_session_terminate", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "actions", "registry_anchor": "canonical-2110223031311333-3310300201232301-3301310220122231-2200000222230101-2013031200123120-1223030202323330-1122220301101011-2132332032023110", "registry_path": "docs/guides/actions--access_active_session_terminate--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/access_active_session_terminate/examples/action/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Action for xcsh_access_active_session_terminate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Action

Breadcrumbs:

- [xcsh_access_active_session_terminate](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_session_terminate/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_session_terminate/examples/)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_access_active_session_terminate/action.tf`; digest `sha256:1bbb807c7b4d5333250c2b20cf18f88631715519e818829cb7766f98ee751aa7`.

```terraform
# AccessActiveSessionTerminate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_access_active_session_terminate" "example" {
  config {
    id        = "example-value"
    namespace = "example-value"
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_session_terminate/examples/)
- [xcsh_access_active_session_terminate](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_session_terminate/)
