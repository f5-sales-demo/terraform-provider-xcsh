---
page_title: "local_ip.ip_address.ip_address.ipv4"
subcategory: ""
description: "IPv4 Address in dot-decimal notation."
xcsh_docs: {"aliases": ["local ip ip address ip address ipv4"], "body_bytes": 2264, "body_sha256": "sha256:aeb58ded0bd5421ee625e393e6c36bda5f0b38db12c6e0bbdd35864807b116a2", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:ipv4", "parent_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address", "path": "documentation/resources/tunnel/properties/local_ip/ip_address/ip_address/ipv4/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3002322102010000-2000122100201021-1013220022011213-1103122211213231-3221313033132330-3201012030112122-0112000032131332-2030102123001222", "registry_path": "docs/guides/resources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_ip", "ip_address", "ip_address", "ipv4"], "schema_version": 1, "sections": [{"aliases": ["local ip ip address ip address ipv4 addr"], "anchor": "schema-local_ip--ip_address--ip_address--ipv4--addr", "description": "IPv4 Address in string form with dot-decimal notation.", "document_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:ipv4", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_ip", "ip_address", "ip_address", "ipv4", "addr"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/local_ip/ip_address/ip_address/ipv4/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "IPv4 Address in dot-decimal notation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["tunnelCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip.ip_address.ip_address.ipv4

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- [local_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/)
- [local_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/)
- [local_ip.ip_address.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/)
- local_ip.ip_address.ip_address.ipv4

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Additional upstream details:

IPv4 Address in dot-decimal notation.

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
ipv4 {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-local_ip--ip_address--ip_address--ipv4--addr"></a>

### addr property

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```
