---
page_title: "infra.interfaces"
subcategory: ""
description: "infra.interfaces for xcsh_registration."
xcsh_docs: {"aliases": [], "body_bytes": 1003, "body_sha256": "sha256:1d0d99d1b07841bf2ba91995f5c1399aedb614accabcdcaf3e1acf28ba884d39", "canonical_id": "xcsh-docs:data-sources:registration:properties:infra:interfaces", "child_ids": [], "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:infra:interfaces", "parent_id": "xcsh-docs:data-sources:registration:properties:infra", "path": "docs/guides/data-sources--registration--properties--infra--interfaces.md", "provider_name": "registration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["infra", "interfaces"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/infra/interfaces/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "infra.interfaces for xcsh_registration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.interfaces

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md)
- [Property reference](data-sources--registration--reference.md)
- [infra](data-sources--registration--properties--infra.md)
- infra.interfaces

<a id="section"></a>

Type: `"single"`. Computed.

Machine interfaces present during registration time.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [infra](data-sources--registration--properties--infra.md)
- [xcsh_registration](../data-sources/registration.md)
