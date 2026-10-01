---
page_title: "rules"
subcategory: ""
description: "rules for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 4140, "body_sha256": "sha256:f64f5b33b8c3b0efcd4f0338d07544ce1a6265073f8e629460408efde547e2c2", "canonical_id": "xcsh-docs:data-sources:nat_policy:properties:rules", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules:action", "xcsh-docs:data-sources:nat_policy:properties:rules:cloud_connect", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria", "xcsh-docs:data-sources:nat_policy:properties:rules:disable_spec", "xcsh-docs:data-sources:nat_policy:properties:rules:enable", "xcsh-docs:data-sources:nat_policy:properties:rules:node_interface", "xcsh-docs:data-sources:nat_policy:properties:rules:segment", "xcsh-docs:data-sources:nat_policy:properties:rules:virtual_network"], "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules", "parent_id": "xcsh-docs:data-sources:nat_policy:reference", "path": "docs/guides/data-sources--nat_policy--properties--rules.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md)
- [Property reference](data-sources--nat_policy--reference.md)
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

- [action](data-sources--nat_policy--properties--rules--action.md): complete subsection reference.

- [cloud_connect](data-sources--nat_policy--properties--rules--cloud_connect.md): complete subsection reference.

- [criteria](data-sources--nat_policy--properties--rules--criteria.md): complete subsection reference.

- [disable_spec](data-sources--nat_policy--properties--rules--disable_spec.md): complete subsection reference.

- [enable](data-sources--nat_policy--properties--rules--enable.md): complete subsection reference.

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

- [node_interface](data-sources--nat_policy--properties--rules--node_interface.md): complete subsection reference.

- [segment](data-sources--nat_policy--properties--rules--segment.md): complete subsection reference.

- [virtual_network](data-sources--nat_policy--properties--rules--virtual_network.md): complete subsection reference.

## Next pages

- [rules.action](data-sources--nat_policy--properties--rules--action.md)
- [rules.cloud_connect](data-sources--nat_policy--properties--rules--cloud_connect.md)
- [rules.criteria](data-sources--nat_policy--properties--rules--criteria.md)
- [rules.disable_spec](data-sources--nat_policy--properties--rules--disable_spec.md)
- [rules.enable](data-sources--nat_policy--properties--rules--enable.md)
- [rules.node_interface](data-sources--nat_policy--properties--rules--node_interface.md)
- [rules.segment](data-sources--nat_policy--properties--rules--segment.md)
- [rules.virtual_network](data-sources--nat_policy--properties--rules--virtual_network.md)
- [Property reference](data-sources--nat_policy--reference.md)
- [xcsh_nat_policy](../data-sources/nat_policy.md)
