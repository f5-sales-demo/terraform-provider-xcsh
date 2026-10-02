---
page_title: "rules"
subcategory: ""
description: "List of rules to apply under the NAT Policy. Rule that matches first would be applied."
xcsh_docs: {"aliases": ["rules"], "body_bytes": 5148, "body_sha256": "sha256:82d483dc8a40a10938c584d0fd3c326a16efb2990c5b1a365cd7acfe607e9a95", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules:action", "xcsh-docs:data-sources:nat_policy:properties:rules:cloud_connect", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria", "xcsh-docs:data-sources:nat_policy:properties:rules:disable_spec", "xcsh-docs:data-sources:nat_policy:properties:rules:enable", "xcsh-docs:data-sources:nat_policy:properties:rules:node_interface", "xcsh-docs:data-sources:nat_policy:properties:rules:segment", "xcsh-docs:data-sources:nat_policy:properties:rules:virtual_network"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules", "parent_id": "xcsh-docs:data-sources:nat_policy:reference", "path": "documentation/data-sources/nat_policy/properties/rules/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230", "registry_path": "docs/guides/data-sources--nat_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules"], "schema_version": 1, "sections": [{"aliases": ["action"], "anchor": "section", "description": "Action to apply on the packet if the NAT rule is applied.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:action", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "action"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloud connect"], "anchor": "section", "description": "Reference to Cloud connect Object.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:cloud_connect", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "cloud_connect"], "syntax": "attribute", "type": "object"}, {"aliases": ["criteria"], "anchor": "section", "description": "Match criteria of the packet to apply the NAT Rule.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "criteria"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:disable_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:enable", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "enable"], "syntax": "attribute", "type": "object"}, {"aliases": ["name"], "anchor": "schema-rules--name", "description": "Name of the Rule.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["node interface"], "anchor": "section", "description": "On multinode site, this type holds the information about per node interfaces.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:node_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "node_interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["segment"], "anchor": "section", "description": "Reference to Segment Object.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:segment", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "segment"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual network"], "anchor": "section", "description": "Carries the reference to virtual network.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:virtual_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "virtual_network"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of rules to apply under the NAT Policy. Rule that matches first would be applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nat_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/)
- rules

<a id="section"></a>

Type: `"list"`. Computed.

List of rules to apply under the NAT Policy. Rule that matches first would be applied.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64"
  }
}
```

## Direct properties

- [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/): complete subsection reference.

- [cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/cloud_connect/): complete subsection reference.

- [criteria](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/): complete subsection reference.

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/disable_spec/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/enable/): complete subsection reference.

<a id="schema-rules--name"></a>

### name property

Type: `"string"`. Computed.

Name. Name of the Rule.

Upstream description:

Name of the Rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/node_interface/): complete subsection reference.

- [segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/segment/): complete subsection reference.

- [virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/virtual_network/): complete subsection reference.

## Next pages

- [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/)
- [rules.cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/cloud_connect/)
- [rules.criteria](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/)
- [rules.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/disable_spec/)
- [rules.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/enable/)
- [rules.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/node_interface/)
- [rules.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/segment/)
- [rules.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/virtual_network/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
