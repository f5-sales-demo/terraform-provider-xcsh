---
page_title: "rules"
subcategory: ""
description: "List of rules to apply under the NAT Policy. Rule that matches first would be applied."
xcsh_docs: {"aliases": ["rules"], "body_bytes": 6151, "body_sha256": "sha256:02d90670a81f20cc5adc9ebce0b619b85c84c120d378a419ed3b8b5e50aa602d", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:action", "xcsh-docs:resources:nat_policy:properties:rules:cloud_connect", "xcsh-docs:resources:nat_policy:properties:rules:criteria", "xcsh-docs:resources:nat_policy:properties:rules:disable_spec", "xcsh-docs:resources:nat_policy:properties:rules:enable", "xcsh-docs:resources:nat_policy:properties:rules:node_interface", "xcsh-docs:resources:nat_policy:properties:rules:segment", "xcsh-docs:resources:nat_policy:properties:rules:virtual_network"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules", "parent_id": "xcsh-docs:resources:nat_policy:reference", "path": "documentation/resources/nat_policy/properties/rules/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132", "registry_path": "docs/guides/resources--nat_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cloud_connect,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:cloud_connect", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cloud_connect,segment", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:cloud_connect", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cloud_connect,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:cloud_connect", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:enable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cloud_connect,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:node_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:node_interface,segment", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:node_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:node_interface,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:node_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cloud_connect,segment", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:node_interface,segment", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:segment,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cloud_connect,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:virtual_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:node_interface,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:virtual_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:segment,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:virtual_network", "type": "conflicts"}, {"anchor": "schema-rules--name", "enforcement": "provider-schema", "group": "rules:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules"], "schema_version": 1, "sections": [{"aliases": ["action"], "anchor": "section", "description": "Action to apply on the packet if the NAT rule is applied.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:action", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rules--action--virtual_cidr", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:dynamic,virtual_cidr", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:dynamic,virtual_cidr", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic", "type": "conflicts"}], "schema_path": ["rules", "action"], "syntax": "block", "type": "object"}, {"aliases": ["cloud connect"], "anchor": "section", "description": "Reference to Cloud connect Object.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:cloud_connect", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.cloud_connect:RequiredObjectAttributes:refs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:cloud_connect:refs", "type": "requires"}], "schema_path": ["rules", "cloud_connect"], "syntax": "block", "type": "object"}, {"aliases": ["criteria"], "anchor": "section", "description": "Match criteria of the packet to apply the NAT Rule.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:any,icmp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:any,tcp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:any,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:any,icmp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:icmp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:icmp,tcp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:icmp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:icmp,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:icmp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:site_local_inside_network,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:site_local_inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:site_local_inside_network,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:site_local_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:any,tcp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:icmp,tcp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:tcp,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:any,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:icmp,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:tcp,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp", "type": "conflicts"}], "schema_path": ["rules", "criteria"], "syntax": "block", "type": "object"}, {"aliases": ["disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:disable_spec", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:enable", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "enable"], "syntax": "attribute", "type": "object"}, {"aliases": ["name"], "anchor": "schema-rules--name", "description": "Name of the Rule.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["node interface"], "anchor": "section", "description": "On multinode site, this type holds the information about per node interfaces.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:node_interface", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "node_interface"], "syntax": "block", "type": "object"}, {"aliases": ["segment"], "anchor": "section", "description": "Reference to Segment Object.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:segment", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.segment:RequiredObjectAttributes:refs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:segment:refs", "type": "requires"}], "schema_path": ["rules", "segment"], "syntax": "block", "type": "object"}, {"aliases": ["virtual network"], "anchor": "section", "description": "Carries the reference to virtual network.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:virtual_network", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "virtual_network"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of rules to apply under the NAT Policy. Rule that matches first would be applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nat_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
