---
page_title: "jwt_validation.target"
subcategory: "Load Balancing"
description: "jwt_validation.target for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2651, "body_sha256": "sha256:71a78e82c6eb677e55d2e2406927e30552240cc8e3ccd3391b4e934fa25e76a7", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:target:all_endpoint", "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:target:api_groups", "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:target:base_paths"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:target", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation", "path": "documentation/resources/cdn_loadbalancer/properties/jwt_validation/target/index.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["jwt_validation", "target"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/jwt_validation/target/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.target for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.target

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/)
- jwt_validation.target

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Define endpoints for which JWT token validation will be performed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_endpoint",
    "api_groups"),
  validators.ConflictingObjectAttributes("all_endpoint",
    "base_paths"),
  validators.ConflictingObjectAttributes("api_groups",
    "base_paths")}
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
  "x-ves-oneof-field-target": "[\"all_endpoint\",\"api_groups\",\"base_paths\"]"
}
```

Terraform syntax:

```terraform
target {
  # Configure direct properties listed below.
}
```

## Direct properties

- [all_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/target/all_endpoint/): complete subsection reference.

- [api_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/target/api_groups/): complete subsection reference.

- [base_paths](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/target/base_paths/): complete subsection reference.

## Next pages

- [jwt_validation.target.all_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/target/all_endpoint/)
- [jwt_validation.target.api_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/target/api_groups/)
- [jwt_validation.target.base_paths](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/target/base_paths/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
