---
page_title: "create_cloud_hosted"
subcategory: ""
description: "create_cloud_hosted for xcsh_bot_infrastructure."
xcsh_docs: {"aliases": [], "body_bytes": 1807, "body_sha256": "sha256:34505ee6a37588c5cc33abf978f39674a5ce54fb6ace4c88bd3f46ea58f35031", "canonical_id": "xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted", "child_ids": ["xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted:production", "xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted:testing"], "collection_id": "xcsh-docs:data-sources:bot_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted", "parent_id": "xcsh-docs:data-sources:bot_infrastructure:reference", "path": "docs/guides/data-sources--bot_infrastructure--properties--create_cloud_hosted.md", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["create_cloud_hosted"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_infrastructure/properties/create_cloud_hosted/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "create_cloud_hosted for xcsh_bot_infrastructure.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# create_cloud_hosted

Breadcrumbs:

- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md)
- [Property reference](data-sources--bot_infrastructure--reference.md)
- create_cloud_hosted

<a id="section"></a>

Type: `"single"`. Computed.

F5 Cloud Hosted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type_choice": "[\"production\",\"testing\"]"
}
```

## Direct properties

<a id="schema-create_cloud_hosted--ip_addresses"></a>

### ip_addresses property

Type: `["list", "string"]`. Computed.

Only traffic from these IP addresses is allowed to access this Bot Defense infrastructure.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

- [production](data-sources--bot_infrastructure--properties--create_cloud_hosted--production.md): complete subsection reference.

- [testing](data-sources--bot_infrastructure--properties--create_cloud_hosted--testing.md): complete subsection reference.

## Next pages

- [create_cloud_hosted.production](data-sources--bot_infrastructure--properties--create_cloud_hosted--production.md)
- [create_cloud_hosted.testing](data-sources--bot_infrastructure--properties--create_cloud_hosted--testing.md)
- [Property reference](data-sources--bot_infrastructure--reference.md)
- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md)
