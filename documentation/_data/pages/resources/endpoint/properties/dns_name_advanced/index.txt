---
page_title: "dns_name_advanced"
subcategory: "Networking"
description: "Specifies name and TTL used for DNS resolution."
xcsh_docs: {"aliases": ["dns name advanced"], "body_bytes": 3814, "body_sha256": "sha256:887e1decf792d0b443d4fa8b1fedc9d154e294595662ddf4f3458b7e97aa6632", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:resources:endpoint:properties:dns_name_advanced", "parent_id": "xcsh-docs:resources:endpoint:reference", "path": "documentation/resources/endpoint/properties/dns_name_advanced/index.md", "product": "distributed-cloud", "provider_name": "endpoint", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3031010133132331-3222110131233333-0211132022132131-2121101130201212-0013311333303003-1112332032030033-1130323022230323-2033313033231121", "registry_path": "docs/guides/resources--endpoint--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dns_name_advanced"], "schema_version": 1, "sections": [{"aliases": ["dns name advanced name"], "anchor": "schema-dns_name_advanced--name", "description": "Endpoint's IP address is discovered using DNS name resolution. The name given here is fully qualified domain name.", "document_id": "xcsh-docs:resources:endpoint:properties:dns_name_advanced", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dns_name_advanced", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["dns name advanced refresh interval"], "anchor": "schema-dns_name_advanced--refresh_interval", "description": "Exclusive with Interval for DNS refresh in seconds.", "document_id": "xcsh-docs:resources:endpoint:properties:dns_name_advanced", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dns_name_advanced", "refresh_interval"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/properties/dns_name_advanced/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Specifies name and TTL used for DNS resolution.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["endpointCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dns_name_advanced

Breadcrumbs:

- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/)
- dns_name_advanced

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
dns_name_advanced {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-dns_name_advanced--name"></a>

### name property

Type: `"string"`. Optional.

Endpoint's IP address is discovered using DNS name resolution. The name given here is fully
qualified domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Type: `"number"`. Optional.

Exclusive with \[\] Interval for DNS refresh in seconds.

Upstream description:

Exclusive with \[\] Interval for DNS refresh in seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(10, 604800),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/)
- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/)
