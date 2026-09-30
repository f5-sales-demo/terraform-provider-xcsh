---
page_title: "api_testing.domains.credentials.api_key.value"
subcategory: "Load Balancing"
description: "api_testing.domains.credentials.api_key.value for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2042, "body_sha256": "sha256:d7d7cae8b4aa1ff40b062ac2ba001702f3d9607d5b099b7960ab091997822030", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:api_key:value", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:api_key:value:blindfold_secret_info", "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:api_key:value:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:api_key:value", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:api_key", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_testing--domains--credentials--api_key--value.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_testing", "domains", "credentials", "api_key", "value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_testing/domains/credentials/api_key/value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_testing.domains.credentials.api_key.value for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_testing.domains.credentials.api_key.value

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_testing](data-sources--http_loadbalancer--properties--api_testing.md)
- [api_testing.domains](data-sources--http_loadbalancer--properties--api_testing--domains.md)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--properties--api_testing--domains--credentials.md)
- [api_testing.domains.credentials.api_key](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--api_key.md)
- api_testing.domains.credentials.api_key.value

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

## Direct properties

- [blindfold_secret_info](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--api_key--value--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--api_key--value--clear_secret_info.md): complete subsection reference.

## Next pages

- [api_testing.domains.credentials.api_key.value.blindfold_secret_info](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--api_key--value--blindfold_secret_info.md)
- [api_testing.domains.credentials.api_key.value.clear_secret_info](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--api_key--value--clear_secret_info.md)
- [api_testing.domains.credentials.api_key](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--api_key.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
