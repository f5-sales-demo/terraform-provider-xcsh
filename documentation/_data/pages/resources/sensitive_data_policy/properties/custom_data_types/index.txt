---
page_title: "custom_data_types"
subcategory: "Security"
description: "Select your custom data types to be monitored in the API discovery."
xcsh_docs: {"aliases": ["custom data types"], "body_bytes": 2070, "body_sha256": "sha256:798a211ec873fde51f41eb86c7eaeb57c5845317fee0d8017608154b7968e5e2", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:sensitive_data_policy:properties:custom_data_types:custom_data_type_ref"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:sensitive_data_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:sensitive_data_policy:properties:custom_data_types", "parent_id": "xcsh-docs:resources:sensitive_data_policy:reference", "path": "documentation/resources/sensitive_data_policy/properties/custom_data_types/index.md", "product": "distributed-cloud", "provider_name": "sensitive_data_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2100333133332323-2001302233101200-2010020300300200-0130311032100233-1233203010002032-3302003112301112-0311303331023330-1331100121002231", "registry_path": "docs/guides/resources--sensitive_data_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_data_types"], "schema_version": 1, "sections": [{"aliases": ["custom data type ref"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:sensitive_data_policy:properties:custom_data_types:custom_data_type_ref", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-custom_data_types--custom_data_type_ref--name", "enforcement": "provider-schema", "group": "custom_data_types.custom_data_type_ref:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:sensitive_data_policy:properties:custom_data_types:custom_data_type_ref", "type": "requires"}], "schema_path": ["custom_data_types", "custom_data_type_ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/sensitive_data_policy/properties/custom_data_types/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Select your custom data types to be monitored in the API discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["sensitive_data_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_data_types

Breadcrumbs:

- [xcsh_sensitive_data_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/sensitive_data_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/sensitive_data_policy/properties/)
- custom_data_types

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Select your custom data types to be monitored in the API discovery. Defaults to \`\[\]\`. Server
applies default when omitted.

Upstream description:

Select your custom data types to be monitored in the API discovery.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
custom_data_types {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_data_type_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/sensitive_data_policy/properties/custom_data_types/custom_data_type_ref/): complete subsection reference.

## Next pages

- [custom_data_types.custom_data_type_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/sensitive_data_policy/properties/custom_data_types/custom_data_type_ref/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/sensitive_data_policy/properties/)
- [xcsh_sensitive_data_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/sensitive_data_policy/)
