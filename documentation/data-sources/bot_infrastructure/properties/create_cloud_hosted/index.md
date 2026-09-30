---
page_title: "create_cloud_hosted"
subcategory: ""
description: "create_cloud_hosted for xcsh_bot_infrastructure."
xcsh_docs: {"aliases": [], "body_bytes": 2116, "body_sha256": "sha256:ac5eb2877daa98ca2997ec2f929102a542e534f4ac9cc1caddd2bda57c902434", "child_ids": ["xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted:production", "xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted:testing"], "collection_id": "xcsh-docs:data-sources:bot_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted", "parent_id": "xcsh-docs:data-sources:bot_infrastructure:reference", "path": "documentation/data-sources/bot_infrastructure/properties/create_cloud_hosted/index.md", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["create_cloud_hosted"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_infrastructure/properties/create_cloud_hosted/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "create_cloud_hosted for xcsh_bot_infrastructure.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# create_cloud_hosted

Breadcrumbs:

- [xcsh_bot_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/properties/)
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

- [production](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/properties/create_cloud_hosted/production/): complete subsection reference.

- [testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/properties/create_cloud_hosted/testing/): complete subsection reference.

## Next pages

- [create_cloud_hosted.production](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/properties/create_cloud_hosted/production/)
- [create_cloud_hosted.testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/properties/create_cloud_hosted/testing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/properties/)
- [xcsh_bot_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/)
