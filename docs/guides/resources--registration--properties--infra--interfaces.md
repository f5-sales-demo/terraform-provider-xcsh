---
page_title: "infra.interfaces"
subcategory: ""
description: "infra.interfaces for xcsh_registration."
xcsh_docs: {"aliases": [], "body_bytes": 1060, "body_sha256": "sha256:866528fa68a329110488106056364218a25f49749a3425002516925ab75c89ba", "canonical_id": "xcsh-docs:resources:registration:properties:infra:interfaces", "child_ids": [], "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:infra:interfaces", "parent_id": "xcsh-docs:resources:registration:properties:infra", "path": "docs/guides/resources--registration--properties--infra--interfaces.md", "provider_name": "registration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["infra", "interfaces"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/infra/interfaces/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "infra.interfaces for xcsh_registration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.interfaces

Breadcrumbs:

- [xcsh_registration](../resources/registration.md)
- [Property reference](resources--registration--reference.md)
- [infra](resources--registration--properties--infra.md)
- infra.interfaces

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
interfaces {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [infra](resources--registration--properties--infra.md)
- [xcsh_registration](../resources/registration.md)
