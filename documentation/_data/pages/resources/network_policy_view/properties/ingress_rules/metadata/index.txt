---
page_title: "ingress_rules.metadata"
subcategory: ""
description: "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs."
xcsh_docs: {"aliases": ["ingress rules metadata"], "body_bytes": 4066, "body_sha256": "sha256:e88b8cb6256da34673dd518957bcdf42ebc4a351ca13aabf4c95472f67444936", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_view:properties:ingress_rules:metadata", "parent_id": "xcsh-docs:resources:network_policy_view:properties:ingress_rules", "path": "documentation/resources/network_policy_view/properties/ingress_rules/metadata/index.md", "product": "distributed-cloud", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3332022032030020-0212311133021021-1200223300231102-1111312312203303-2131211232300013-0010133103122333-0003312313302213-1111322322211300", "registry_path": "docs/guides/resources--network_policy_view--reference--group-001.md", "relationships": [{"anchor": "schema-ingress_rules--metadata--name", "enforcement": "provider-schema", "group": "ingress_rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:ingress_rules:metadata", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_rules", "metadata"], "schema_version": 1, "sections": [{"aliases": ["description spec"], "anchor": "schema-ingress_rules--metadata--description_spec", "description": "Description. Human readable description.", "document_id": "xcsh-docs:resources:network_policy_view:properties:ingress_rules:metadata", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_rules", "metadata", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-ingress_rules--metadata--name", "description": "This is the name of the message. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:network_policy_view:properties:ingress_rules:metadata", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_rules", "metadata", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_view/properties/ingress_rules/metadata/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_rules.metadata

Breadcrumbs:

- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/)
- [ingress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/ingress_rules/)
- ingress_rules.metadata

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

<a id="schema-ingress_rules--metadata--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="schema-ingress_rules--metadata--name"></a>

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

- [ingress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/ingress_rules/)
- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/)
