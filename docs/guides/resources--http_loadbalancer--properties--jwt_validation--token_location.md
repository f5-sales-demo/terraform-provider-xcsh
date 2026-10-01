---
page_title: "jwt_validation.token_location"
subcategory: "Load Balancing"
description: "jwt_validation.token_location for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1357, "body_sha256": "sha256:12bcd7ad1ee703437bbf2814fe8364280cbc47431859363079385a600085ba44", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:token_location", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:token_location:bearer_token"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:token_location", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation", "path": "docs/guides/resources--http_loadbalancer--properties--jwt_validation--token_location.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation", "token_location"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/jwt_validation/token_location/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.token_location for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.token_location

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [jwt_validation](resources--http_loadbalancer--properties--jwt_validation.md)
- jwt_validation.token_location

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for token location.

Upstream description:

Location of JWT in HTTP request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-token_location": "[\"bearer_token\"]"
}
```

Terraform syntax:

```terraform
token_location {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bearer_token](resources--http_loadbalancer--properties--jwt_validation--token_location--bearer_token.md): complete subsection reference.

## Next pages

- [jwt_validation.token_location.bearer_token](resources--http_loadbalancer--properties--jwt_validation--token_location--bearer_token.md)
- [jwt_validation](resources--http_loadbalancer--properties--jwt_validation.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
