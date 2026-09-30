---
page_title: "virtual_server.connection_rate_limit_mode.per_virtual_server_source_address"
subcategory: ""
description: "virtual_server.connection_rate_limit_mode.per_virtual_server_source_address for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 2023, "body_sha256": "sha256:296f0a6abb8fc68e9c348f02240dbca78f47d586b5a0be2446f0b73f0f1beb09", "canonical_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_source_address", "child_ids": [], "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_source_address", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode", "path": "docs/guides/data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_address.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_virtual_server_source_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_source_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.connection_rate_limit_mode.per_virtual_server_source_address for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# virtual_server.connection_rate_limit_mode.per_virtual_server_source_address

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md)
- [Property reference](data-sources--application_profiles--reference.md)
- [virtual_server](data-sources--application_profiles--properties--virtual_server.md)
- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode.md)
- virtual_server.connection_rate_limit_mode.per_virtual_server_source_address

<a id="section"></a>

Type: `"single"`. Computed.

Source Address Mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

<a id="schema-virtual_server--connection_rate_limit_mode--per_virtual_server_source_address--source_mask"></a>

### source_mask property

Type: `"number"`. Computed.

Configuration parameter for source mask.

Upstream description:

Configuration parameter for source mask

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

## Next pages

- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode.md)
- [xcsh_application_profiles](../data-sources/application_profiles.md)
