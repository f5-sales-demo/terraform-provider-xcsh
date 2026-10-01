---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_access_active_sessions_terminate."
xcsh_docs: {"aliases": [], "body_bytes": 1079, "body_sha256": "sha256:bea97c4865ee72269861b0c4168aea7ab2536f098086b060faf7118f6cca84cc", "canonical_id": "xcsh-docs:actions:access_active_sessions_terminate:example:action", "child_ids": [], "collection_id": "xcsh-docs:actions:access_active_sessions_terminate:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e3e7c43f13f1ace8b437c1c8885aa667a9add78ae921c28c7a137377165d7c29", "source_path": "examples/actions/xcsh_access_active_sessions_terminate/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:access_active_sessions_terminate:example:action", "parent_id": "xcsh-docs:actions:access_active_sessions_terminate:examples", "path": "docs/guides/actions--access_active_sessions_terminate--example--action.md", "provider_name": "access_active_sessions_terminate", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "actions", "publishing_destination": "registry", "role": "example", "schema_path": ["action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/access_active_sessions_terminate/examples/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Action for xcsh_access_active_sessions_terminate.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Action

Breadcrumbs:

- [xcsh_access_active_sessions_terminate](../actions/access_active_sessions_terminate.md)
- [Examples](actions--access_active_sessions_terminate--examples.md)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_access_active_sessions_terminate/action.tf`; digest `sha256:e3e7c43f13f1ace8b437c1c8885aa667a9add78ae921c28c7a137377165d7c29`.

```terraform
# AccessActiveSessionsTerminate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_access_active_sessions_terminate" "example" {
  config {
    namespace = "example-value"
  }
}
```

## Next pages

- [Examples](actions--access_active_sessions_terminate--examples.md)
- [xcsh_access_active_sessions_terminate](../actions/access_active_sessions_terminate.md)
