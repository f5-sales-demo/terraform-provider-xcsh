---
page_title: "dns_name_advanced"
subcategory: "Networking"
description: "Specifies name and TTL used for DNS resolution."
xcsh_docs: {"aliases": ["dns name advanced"], "body_bytes": 3420, "body_sha256": "sha256:061655d8124fe577758d9218cfcf0d00d97e99fac51d030d0edcf9be70553f44", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:endpoint:properties:dns_name_advanced", "parent_id": "xcsh-docs:data-sources:endpoint:reference", "path": "documentation/data-sources/endpoint/properties/dns_name_advanced/index.md", "product": "distributed-cloud", "provider_name": "endpoint", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2012123113222233-3233120330213010-2000031000100133-0113201221232212-3303012201020021-3213302333103112-1001012322323113-3100023101323323", "registry_path": "docs/guides/data-sources--endpoint--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dns_name_advanced"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-dns_name_advanced--name", "description": "Endpoint's IP address is discovered using DNS name resolution. The name given here is fully qualified domain name.", "document_id": "xcsh-docs:data-sources:endpoint:properties:dns_name_advanced", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dns_name_advanced", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["refresh interval"], "anchor": "schema-dns_name_advanced--refresh_interval", "description": "Exclusive with Interval for DNS refresh in seconds.", "document_id": "xcsh-docs:data-sources:endpoint:properties:dns_name_advanced", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dns_name_advanced", "refresh_interval"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/endpoint/properties/dns_name_advanced/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specifies name and TTL used for DNS resolution.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["endpointCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dns_name_advanced

Breadcrumbs:

- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/)
- dns_name_advanced

<a id="section"></a>

Type: `"single"`. Computed.

Specifies name and TTL used for DNS resolution.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ttl_choice": "[\"refresh_interval\"]"
}
```

## Direct properties

<a id="schema-dns_name_advanced--name"></a>

### name property

Type: `"string"`. Computed.

Endpoint's IP address is discovered using DNS name resolution. The name given here is fully
qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "format": "hostname",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 256,
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-dns_name_advanced--refresh_interval"></a>

### refresh_interval property

Type: `"number"`. Computed.

Exclusive with \[\] Interval for DNS refresh in seconds.

Upstream description:

Exclusive with \[\] Interval for DNS refresh in seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 10
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "604800"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/)
- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/)
