---
page_title: "With labels"
subcategory: ""
description: "With labels for xcsh_user_identification."
xcsh_docs: {"aliases": [], "body_bytes": 1246, "body_sha256": "sha256:2fb00df452d9ed07a8cdfda245c2fe2f7382965131172907330a15e7e8554dba", "child_ids": [], "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:ba484df9375d0da693ea56c4aaf4288ad5030a88885849721fac1dec608ba5a3", "source_path": "examples/resources/xcsh_user_identification/with-labels.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:user_identification:example:with-labels", "parent_id": "xcsh-docs:resources:user_identification:examples", "path": "documentation/resources/user_identification/examples/with-labels/index.md", "provider_name": "user_identification", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["with-labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/examples/with-labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With labels for xcsh_user_identification.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["user_identificationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# With labels

Breadcrumbs:

- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/examples/)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/with-labels.tf`; digest `sha256:ba484df9375d0da693ea56c4aaf4288ad5030a88885849721fac1dec608ba5a3`.

```terraform
# WithLabels — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name      = "example"
  namespace = "system"

  labels = {
    example-key = "example-value"
  }

  rules {
    client_ip = {}
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/examples/)
- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/)
