---
page_title: "proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip"
subcategory: ""
description: "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name."
xcsh_docs: {"aliases": ["proxy advertisement advertise custom advertise where advertise v6 on public public ip"], "body_bytes": 6395, "body_sha256": "sha256:eb0e33d6aca2aaab5c5b2ab3f115edc0e642069a6c7ca998195144d2abb164a2", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_v6_on_public:public_ip", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_v6_on_public", "path": "documentation/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_v6_on_public/public_ip/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2130113210010113-2031300331003133-2332130320133102-1001003131231212-2303001312120330-2033310013100033-1122202011312311-3133132113323230", "registry_path": "docs/guides/data-sources--dns_proxy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "advertise_v6_on_public", "public_ip"], "schema_version": 1, "sections": [{"aliases": ["proxy advertisement advertise custom advertise where advertise v6 on public public ip name"], "anchor": "schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_v6_on_public:public_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "advertise_v6_on_public", "public_ip", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["proxy advertisement advertise custom advertise where advertise v6 on public public ip namespace"], "anchor": "schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_v6_on_public:public_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "advertise_v6_on_public", "public_ip", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["proxy advertisement advertise custom advertise where advertise v6 on public public ip tenant"], "anchor": "schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_v6_on_public:public_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "advertise_v6_on_public", "public_ip", "tenant"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_v6_on_public/public_ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- [proxy_advertisement](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/)
- [proxy_advertisement.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/)
- [proxy_advertisement.advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_v6_on_public/)
- proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="section"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

## Next pages

- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_v6_on_public/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
