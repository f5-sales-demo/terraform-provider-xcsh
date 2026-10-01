---
page_title: "jwt_validation.target"
subcategory: "Load Balancing"
description: "jwt_validation.target for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2113, "body_sha256": "sha256:2d3f6c45718480c55c1fd3fd8a5f29d9a76707f676db53732c472dbcf8625442", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:target", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:target:all_endpoint", "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:target:api_groups", "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:target:base_paths"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:target", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation", "path": "docs/guides/resources--http_loadbalancer--properties--jwt_validation--target.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation", "target"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/jwt_validation/target/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.target for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.target

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [jwt_validation](resources--http_loadbalancer--properties--jwt_validation.md)
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

- [all_endpoint](resources--http_loadbalancer--properties--jwt_validation--target--all_endpoint.md): complete subsection reference.

- [api_groups](resources--http_loadbalancer--properties--jwt_validation--target--api_groups.md): complete subsection reference.

- [base_paths](resources--http_loadbalancer--properties--jwt_validation--target--base_paths.md): complete subsection reference.

## Next pages

- [jwt_validation.target.all_endpoint](resources--http_loadbalancer--properties--jwt_validation--target--all_endpoint.md)
- [jwt_validation.target.api_groups](resources--http_loadbalancer--properties--jwt_validation--target--api_groups.md)
- [jwt_validation.target.base_paths](resources--http_loadbalancer--properties--jwt_validation--target--base_paths.md)
- [jwt_validation](resources--http_loadbalancer--properties--jwt_validation.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
