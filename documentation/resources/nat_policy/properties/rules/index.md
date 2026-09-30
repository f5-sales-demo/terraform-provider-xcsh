---
page_title: "rules"
subcategory: ""
description: "rules for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 6052, "body_sha256": "sha256:970dd0ed77792ade4b4b9636189b12861884e02c46b5870f298489098131cf21", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:action", "xcsh-docs:resources:nat_policy:properties:rules:cloud_connect", "xcsh-docs:resources:nat_policy:properties:rules:criteria", "xcsh-docs:resources:nat_policy:properties:rules:disable_spec", "xcsh-docs:resources:nat_policy:properties:rules:enable", "xcsh-docs:resources:nat_policy:properties:rules:node_interface", "xcsh-docs:resources:nat_policy:properties:rules:segment", "xcsh-docs:resources:nat_policy:properties:rules:virtual_network"], "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules", "parent_id": "xcsh-docs:resources:nat_policy:reference", "path": "documentation/resources/nat_policy/properties/rules/index.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/)
- rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of rules to apply under the NAT Policy. Rule that matches first would be applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("cloud_connect",
    "node_interface"),
  validators.ConflictingListObjectAttributes("cloud_connect",
    "segment"),
  validators.ConflictingListObjectAttributes("cloud_connect",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("disable_spec",
    "enable"),
  validators.ConflictingListObjectAttributes("node_interface",
    "segment"),
  validators.ConflictingListObjectAttributes("node_interface",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("segment",
    "virtual_network")}
```

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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/): complete subsection reference.

- [cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/cloud_connect/): complete subsection reference.

- [criteria](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/): complete subsection reference.

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/disable_spec/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/enable/): complete subsection reference.

<a id="schema-rules--name"></a>

### name property

Type: `"string"`. Optional.

Name. Name of the Rule.

Upstream description:

Name of the Rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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

- [node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/): complete subsection reference.

- [segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/segment/): complete subsection reference.

- [virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/virtual_network/): complete subsection reference.

## Next pages

- [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/)
- [rules.cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/cloud_connect/)
- [rules.criteria](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/)
- [rules.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/disable_spec/)
- [rules.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/enable/)
- [rules.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/)
- [rules.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/segment/)
- [rules.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/virtual_network/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
