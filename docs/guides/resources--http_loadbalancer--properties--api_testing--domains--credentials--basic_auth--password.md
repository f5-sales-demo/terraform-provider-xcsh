---
page_title: "api_testing.domains.credentials.basic_auth.password"
subcategory: "Load Balancing"
description: "api_testing.domains.credentials.basic_auth.password for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2371, "body_sha256": "sha256:356c9fad3aeea0be7e5a53cddf808bab19f39401755e24379de9ea71ef7e8a9b", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:basic_auth:password", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:basic_auth:password:blindfold_secret_info", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:basic_auth:password:clear_secret_info"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:basic_auth:password", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:basic_auth", "path": "docs/guides/resources--http_loadbalancer--properties--api_testing--domains--credentials--basic_auth--password.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_testing", "domains", "credentials", "basic_auth", "password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_testing/domains/credentials/basic_auth/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_testing.domains.credentials.basic_auth.password for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_testing.domains.credentials.basic_auth.password

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_testing](resources--http_loadbalancer--properties--api_testing.md)
- [api_testing.domains](resources--http_loadbalancer--properties--api_testing--domains.md)
- [api_testing.domains.credentials](resources--http_loadbalancer--properties--api_testing--domains--credentials.md)
- [api_testing.domains.credentials.basic_auth](resources--http_loadbalancer--properties--api_testing--domains--credentials--basic_auth.md)
- api_testing.domains.credentials.basic_auth.password

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--http_loadbalancer--properties--api_testing--domains--credentials--basic_auth--password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--properties--api_testing--domains--credentials--basic_auth--password--clear_secret_info.md): complete subsection reference.

## Next pages

- [api_testing.domains.credentials.basic_auth.password.blindfold_secret_info](resources--http_loadbalancer--properties--api_testing--domains--credentials--basic_auth--password--blindfold_secret_info.md)
- [api_testing.domains.credentials.basic_auth.password.clear_secret_info](resources--http_loadbalancer--properties--api_testing--domains--credentials--basic_auth--password--clear_secret_info.md)
- [api_testing.domains.credentials.basic_auth](resources--http_loadbalancer--properties--api_testing--domains--credentials--basic_auth.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
