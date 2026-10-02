---
page_title: "custom_data_types"
subcategory: "Security"
description: "Select your custom data types to be monitored in the API discovery."
xcsh_docs: {"aliases": ["custom data types"], "body_bytes": 1962, "body_sha256": "sha256:0e2750bcda40b83af0a99c35f360d73181b67beda9deb3062116f11e3d08efd3", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:sensitive_data_policy:properties:custom_data_types:custom_data_type_ref"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:sensitive_data_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:sensitive_data_policy:properties:custom_data_types", "parent_id": "xcsh-docs:data-sources:sensitive_data_policy:reference", "path": "documentation/data-sources/sensitive_data_policy/properties/custom_data_types/index.md", "product": "distributed-cloud", "provider_name": "sensitive_data_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2213331322133102-0313231331030120-0210321330120000-3203032221020333-2021113001131100-2132023212032102-1123203103322320-1230012010303112", "registry_path": "docs/guides/data-sources--sensitive_data_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_data_types"], "schema_version": 1, "sections": [{"aliases": ["custom data type ref"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:sensitive_data_policy:properties:custom_data_types:custom_data_type_ref", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_data_types", "custom_data_type_ref"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/sensitive_data_policy/properties/custom_data_types/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Select your custom data types to be monitored in the API discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["sensitive_data_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_data_types

Breadcrumbs:

- [xcsh_sensitive_data_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/sensitive_data_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/sensitive_data_policy/properties/)
- custom_data_types

<a id="section"></a>

Type: `"list"`. Computed.

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

## Direct properties

- [custom_data_type_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/sensitive_data_policy/properties/custom_data_types/custom_data_type_ref/): complete subsection reference.

## Next pages

- [custom_data_types.custom_data_type_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/sensitive_data_policy/properties/custom_data_types/custom_data_type_ref/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/sensitive_data_policy/properties/)
- [xcsh_sensitive_data_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/sensitive_data_policy/)
