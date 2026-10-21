---
page_title: "remote_ip.ip.ipv4"
subcategory: ""
description: "IPv4 Address in dot-decimal notation."
xcsh_docs: {"aliases": ["remote ip ip ipv4"], "body_bytes": 1745, "body_sha256": "sha256:e0843b273f6574f6f7d2f0a0d7eac5ebd0bbecdfc69d1976149ce528359452a5", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip:ipv4", "parent_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip", "path": "documentation/data-sources/tunnel/properties/remote_ip/ip/ipv4/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3110233220123120-2201113132033203-3301120103133021-2000121133221013-1131100110221232-2300231321113120-1203120102313010-3013331101110223", "registry_path": "docs/guides/data-sources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["remote_ip", "ip", "ipv4"], "schema_version": 1, "sections": [{"aliases": ["remote ip ip ipv4 addr"], "anchor": "schema-remote_ip--ip--ipv4--addr", "description": "IPv4 Address in string form with dot-decimal notation.", "document_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip:ipv4", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["remote_ip", "ip", "ipv4", "addr"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/remote_ip/ip/ipv4/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "IPv4 Address in dot-decimal notation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["tunnelCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# remote_ip.ip.ipv4

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- [remote_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/)
- [remote_ip.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/ip/)
- remote_ip.ip.ipv4

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-remote_ip--ip--ipv4--addr"></a>

### addr property

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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
      "validatedAt": "2026-10-21T11:00:09+00:00"
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
