---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_access_active_session_terminate."
xcsh_docs: {"aliases": [], "body_bytes": 1308, "body_sha256": "sha256:24eba07232aa25b6f07f586bf13069d2559bbf807ffc528e47abaa55ade97797", "child_ids": [], "collection_id": "xcsh-docs:actions:access_active_session_terminate:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1bbb807c7b4d5333250c2b20cf18f88631715519e818829cb7766f98ee751aa7", "source_path": "examples/actions/xcsh_access_active_session_terminate/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:access_active_session_terminate:example:action", "parent_id": "xcsh-docs:actions:access_active_session_terminate:examples", "path": "documentation/actions/access_active_session_terminate/examples/action/index.md", "provider_name": "access_active_session_terminate", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "actions", "role": "example", "schema_path": ["action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/access_active_session_terminate/examples/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Action for xcsh_access_active_session_terminate.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
