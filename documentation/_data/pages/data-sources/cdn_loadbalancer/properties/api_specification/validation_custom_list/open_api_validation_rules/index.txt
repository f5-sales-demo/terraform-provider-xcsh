---
page_title: "api_specification.validation_custom_list.open_api_validation_rules"
subcategory: "Load Balancing"
description: "Rule or policy definition"
xcsh_docs: {"aliases": ["api specification validation custom list open api validation rules"], "body_bytes": 7280, "body_sha256": "sha256:5af4e72ffa7551b9c25164e41b3ba88bcb793e845741883945c3f368dd5c42ea", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:any_domain", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:api_endpoint", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:metadata", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list", "path": "documentation/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3031231222230321-1030300023112022-1013022221003000-1110012111020320-2023221202120210-1211111233202113-1322002311121130-1132231003302213", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules"], "schema_version": 1, "sections": [{"aliases": ["any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:any_domain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["api endpoint"], "anchor": "section", "description": "This defines API endpoint.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:api_endpoint", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "api_endpoint"], "syntax": "attribute", "type": "object"}, {"aliases": ["api group"], "anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--api_group", "description": "Exclusive with The API group which this validation applies to.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "api_group"], "syntax": "attribute", "type": "string"}, {"aliases": ["base path"], "anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--base_path", "description": "Exclusive with The base path which this validation applies to.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "base_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["specific domain"], "anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--specific_domain", "description": "Exclusive with The rule will apply for a specific domain.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "specific_domain"], "syntax": "attribute", "type": "string"}, {"aliases": ["validation mode"], "anchor": "section", "description": "Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Rule or policy definition", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.open_api_validation_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/)
- api_specification.validation_custom_list.open_api_validation_rules

<a id="section"></a>

Type: `"list"`. Computed.

Validation List. Rule or policy definition

Upstream description:

Rule or policy definition

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/any_domain/): complete subsection reference.

- [api_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/api_endpoint/): complete subsection reference.

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--api_group"></a>

### api_group property

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--base_path"></a>

### base_path property

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/metadata/): complete subsection reference.

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--specific_domain"></a>

### specific_domain property

Type: `"string"`. Computed.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.open_api_validation_rules.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/any_domain/)
- [api_specification.validation_custom_list.open_api_validation_rules.api_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/api_endpoint/)
- [api_specification.validation_custom_list.open_api_validation_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/metadata/)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
