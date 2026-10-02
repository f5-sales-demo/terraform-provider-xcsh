---
page_title: "api_specification.validation_custom_list.open_api_validation_rules"
subcategory: "Load Balancing"
description: "Rule or policy definition"
xcsh_docs: {"aliases": ["api specification validation custom list open api validation rules"], "body_bytes": 8197, "body_sha256": "sha256:d62b9727f89168afd0c438220637cf0e243b5927a84e4796126d2341dec71ecd", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:any_domain", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:api_endpoint", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:metadata", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list", "path": "documentation/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-006.md", "relationships": [{"anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--api_group", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules:ConflictingListObjectAttributes:api_endpoint,api_group", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "type": "conflicts"}, {"anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--api_group", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules:ConflictingListObjectAttributes:api_group,base_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "type": "conflicts"}, {"anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--base_path", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules:ConflictingListObjectAttributes:api_endpoint,base_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "type": "conflicts"}, {"anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--base_path", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules:ConflictingListObjectAttributes:api_group,base_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "type": "conflicts"}, {"anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--specific_domain", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules:ConflictingListObjectAttributes:any_domain,specific_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules:ConflictingListObjectAttributes:any_domain,specific_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules:ConflictingListObjectAttributes:api_endpoint,api_group", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:api_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules:ConflictingListObjectAttributes:api_endpoint,base_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:api_endpoint", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules"], "schema_version": 1, "sections": [{"aliases": ["any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:any_domain", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["api endpoint"], "anchor": "section", "description": "This defines API endpoint.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:api_endpoint", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--api_endpoint--path", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.api_endpoint:RequiredObjectAttributes:path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:api_endpoint", "type": "requires"}], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "api_endpoint"], "syntax": "block", "type": "object"}, {"aliases": ["api group"], "anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--api_group", "description": "Exclusive with The API group which this validation applies to.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "api_group"], "syntax": "attribute", "type": "string"}, {"aliases": ["base path"], "anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--base_path", "description": "Exclusive with The base path which this validation applies to.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "base_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--metadata--name", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:metadata", "type": "requires"}], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["specific domain"], "anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--specific_domain", "description": "Exclusive with The rule will apply for a specific domain.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "specific_domain"], "syntax": "attribute", "type": "string"}, {"aliases": ["validation mode"], "anchor": "section", "description": "Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode:ConflictingObjectAttributes:response_validation_mode_active,skip_response_validation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode:ConflictingObjectAttributes:response_validation_mode_active,skip_response_validation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_response_validation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode:ConflictingObjectAttributes:skip_validation,validation_mode_active", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_validation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode:ConflictingObjectAttributes:skip_validation,validation_mode_active", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active", "type": "conflicts"}], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Rule or policy definition", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.open_api_validation_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/)
- api_specification.validation_custom_list.open_api_validation_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Validation List. Rule or policy definition

Upstream description:

Rule or policy definition

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "api_group"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "base_path"),
  validators.ConflictingListObjectAttributes("api_group",
    "base_path")}
```

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

Terraform syntax:

```terraform
open_api_validation_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/any_domain/): complete subsection reference.

- [api_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/api_endpoint/): complete subsection reference.

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--api_group"></a>

### api_group property

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/metadata/): complete subsection reference.

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--specific_domain"></a>

### specific_domain property

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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

- [validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.open_api_validation_rules.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/any_domain/)
- [api_specification.validation_custom_list.open_api_validation_rules.api_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/api_endpoint/)
- [api_specification.validation_custom_list.open_api_validation_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/metadata/)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
