---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_access_active_sessions_terminate."
xcsh_docs: {"aliases": ["action"], "body_bytes": 1285, "body_sha256": "sha256:dd2ab5eec0a6938bb4a41e80fb187fd3e145ff3bc6b0db276736c75dd842ed14", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:access_active_sessions_terminate:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e3e7c43f13f1ace8b437c1c8885aa667a9add78ae921c28c7a137377165d7c29", "source_path": "examples/actions/xcsh_access_active_sessions_terminate/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:access_active_sessions_terminate:example:action", "parent_id": "xcsh-docs:actions:access_active_sessions_terminate:examples", "path": "documentation/actions/access_active_sessions_terminate/examples/action/index.md", "product": "distributed-cloud", "provider_name": "access_active_sessions_terminate", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "actions", "registry_anchor": "canonical-0323022220300023-0020313113010222-2300010131030002-0211131012302230-2213331102112032-2023220223020302-0230331230022310-0020130211311311", "registry_path": "docs/guides/actions--access_active_sessions_terminate--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/access_active_sessions_terminate/examples/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Action for xcsh_access_active_sessions_terminate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Action

Breadcrumbs:

- [xcsh_access_active_sessions_terminate](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/examples/)
- [xcsh_access_active_sessions_terminate](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/)
