---
page_title: "proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site"
subcategory: ""
description: "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name."
xcsh_docs: {"aliases": ["proxy advertisement advertise custom advertise where vk8s service virtual site"], "body_bytes": 6990, "body_sha256": "sha256:02c7d3be2934639842c21982c5f659fb494bbc360456130975a7570b6a6b7932", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service:virtual_site", "parent_id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service", "path": "documentation/resources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/virtual_site/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3103101011303203-2022103310012021-3001313311211213-3322321020321310-0220133030010133-1112333000112311-3301302213102023-0231221312201301", "registry_path": "docs/guides/resources--dns_proxy--reference--group-002.md", "relationships": [{"anchor": "schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--name", "enforcement": "provider-schema", "group": "proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service:virtual_site", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "vk8s_service", "virtual_site"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service:virtual_site", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "vk8s_service", "virtual_site", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service:virtual_site", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "vk8s_service", "virtual_site", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["tenant"], "anchor": "schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service:virtual_site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "vk8s_service", "virtual_site", "tenant"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/virtual_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [proxy_advertisement](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/)
- [proxy_advertisement.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_custom/)
- [proxy_advertisement.advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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
virtual_site {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--name"></a>

### name property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
