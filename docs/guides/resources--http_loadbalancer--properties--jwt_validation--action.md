---
page_title: "jwt_validation.action"
subcategory: "Load Balancing"
description: "jwt_validation.action for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1483, "body_sha256": "sha256:28c1d15425bcbc7d7bbcf3b2a223d42fb952b1567e29b4dac46198c99884a438", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:action", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:action:block", "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:action:report"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:action", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation", "path": "docs/guides/resources--http_loadbalancer--properties--jwt_validation--action.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation", "action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/jwt_validation/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.action for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# jwt_validation.action

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [jwt_validation](resources--http_loadbalancer--properties--jwt_validation.md)
- jwt_validation.action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Action

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "report")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"block\",\"report\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [block](resources--http_loadbalancer--properties--jwt_validation--action--block.md): complete subsection reference.

- [report](resources--http_loadbalancer--properties--jwt_validation--action--report.md): complete subsection reference.

## Next pages

- [jwt_validation.action.block](resources--http_loadbalancer--properties--jwt_validation--action--block.md)
- [jwt_validation.action.report](resources--http_loadbalancer--properties--jwt_validation--action--report.md)
- [jwt_validation](resources--http_loadbalancer--properties--jwt_validation.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
