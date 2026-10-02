---
page_title: "api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata"
subcategory: "Load Balancing"
description: "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs."
xcsh_docs: {"aliases": ["api specification validation custom list fall through mode fall through mode custom open api validation rules metadata"], "body_bytes": 5636, "body_sha256": "sha256:39eabef5b7d4011d2ff8d54f300290cc709854b362a947b76a518f549b0658a3", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:metadata", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules", "path": "documentation/resources/http_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/metadata/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0012120101300020-3311220132210321-0221033232332120-3303132320332132-0321333303120231-1110202220212213-2020131213102012-0211012112303213", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-009.md", "relationships": [{"anchor": "schema-api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--metadata--name", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:metadata", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode", "fall_through_mode_custom", "open_api_validation_rules", "metadata"], "schema_version": 1, "sections": [{"aliases": ["description spec"], "anchor": "schema-api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--metadata--description_spec", "description": "Description. Human readable description.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:metadata", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode", "fall_through_mode_custom", "open_api_validation_rules", "metadata", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--metadata--name", "description": "This is the name of the message. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:metadata", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode", "fall_through_mode_custom", "open_api_validation_rules", "metadata", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/metadata/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/)
- [api_specification.validation_custom_list.fall_through_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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
metadata {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--metadata--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="schema-api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--metadata--name"></a>

### name property

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

## Next pages

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
