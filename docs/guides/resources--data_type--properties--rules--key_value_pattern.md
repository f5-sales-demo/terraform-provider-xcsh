---
page_title: "rules.key_value_pattern"
subcategory: ""
description: "rules.key_value_pattern for xcsh_data_type."
xcsh_docs: {"aliases": [], "body_bytes": 1378, "body_sha256": "sha256:7afcfb8ccc32c8e5531da0a6b03775a26efb965e5346968bdf7a234e5d60d5fb", "canonical_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern", "child_ids": ["xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern", "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:value_pattern"], "collection_id": "xcsh-docs:resources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern", "parent_id": "xcsh-docs:resources:data_type:properties:rules", "path": "docs/guides/resources--data_type--properties--rules--key_value_pattern.md", "provider_name": "data_type", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "key_value_pattern"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_type/properties/rules/key_value_pattern/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.key_value_pattern for xcsh_data_type.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_typeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.key_value_pattern

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md)
- [Property reference](resources--data_type--reference.md)
- [rules](resources--data_type--properties--rules.md)
- rules.key_value_pattern

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Search for specific key &amp; value patterns in the specified sections.

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

Terraform syntax:

```terraform
key_value_pattern {
  # Configure direct properties listed below.
}
```

## Direct properties

- [key_pattern](resources--data_type--properties--rules--key_value_pattern--key_pattern.md): complete subsection reference.

- [value_pattern](resources--data_type--properties--rules--key_value_pattern--value_pattern.md): complete subsection reference.

## Next pages

- [rules.key_value_pattern.key_pattern](resources--data_type--properties--rules--key_value_pattern--key_pattern.md)
- [rules.key_value_pattern.value_pattern](resources--data_type--properties--rules--key_value_pattern--value_pattern.md)
- [rules](resources--data_type--properties--rules.md)
- [xcsh_data_type](../resources/data_type.md)
