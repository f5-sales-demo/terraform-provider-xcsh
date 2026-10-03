---
page_title: "rule_list.rules.protocol_port_range"
subcategory: ""
description: "Protocol and Port ranges."
xcsh_docs: {"aliases": ["rule list rules protocol port range"], "body_bytes": 3885, "body_sha256": "sha256:63c7ecd7c199b465dbb64ff328429855494e4eb16df5c0c78195c810b540591b", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:protocol_port_range", "parent_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules", "path": "documentation/resources/enhanced_firewall_policy/properties/rule_list/rules/protocol_port_range/index.md", "product": "distributed-cloud", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2002103111113300-0332102033002033-2132302000221023-3221111201100303-0110202033011102-0333312301123312-1122301212333330-3331110232302232", "registry_path": "docs/guides/resources--enhanced_firewall_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "protocol_port_range"], "schema_version": 1, "sections": [{"aliases": ["rule list rules protocol port range port ranges"], "anchor": "schema-rule_list--rules--protocol_port_range--port_ranges", "description": "List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:protocol_port_range", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "protocol_port_range", "port_ranges"], "syntax": "attribute", "type": "list"}, {"aliases": ["rule list rules protocol port range protocol"], "anchor": "schema-rule_list--rules--protocol_port_range--protocol", "description": "Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:protocol_port_range", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "protocol_port_range", "protocol"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/enhanced_firewall_policy/properties/rule_list/rules/protocol_port_range/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Protocol and Port ranges.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.protocol_port_range

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/)
- rule_list.rules.protocol_port_range

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Protocol and Port. Protocol and Port ranges.

Upstream description:

Protocol and Port ranges.

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
protocol_port_range {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rule_list--rules--protocol_port_range--port_ranges"></a>

### port_ranges property

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="schema-rule_list--rules--protocol_port_range--protocol"></a>

### protocol property

Type: `"string"`. Optional.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ALL",
    "TCP",
    "UDP",
    "ICMP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

## Next pages

- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/)
- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/)
