---
page_title: "origin_pool.origin_servers.public_name"
subcategory: "Load Balancing"
description: "Specify origin server with public DNS name."
xcsh_docs: {"aliases": ["backend servers", "origin pool origin servers public name", "origin servers", "upstream servers"], "body_bytes": 4469, "body_sha256": "sha256:e3d5bed686db29391f64b020e43aa12adcb1feeced5371329e79b03bf9dcc873", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_name", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers", "path": "documentation/resources/cdn_loadbalancer/properties/origin_pool/origin_servers/public_name/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3000211111131131-0021133131203022-2211200013220222-3101213003033323-3331012011322300-3320121322302033-3213312313302023-0313132230123213", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-012.md", "relationships": [{"anchor": "schema-origin_pool--origin_servers--public_name--dns_name", "enforcement": "provider-schema", "group": "origin_pool.origin_servers.public_name:RequiredObjectAttributes:dns_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_name", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pool", "origin_servers", "public_name"], "schema_version": 1, "sections": [{"aliases": ["dns name"], "anchor": "schema-origin_pool--origin_servers--public_name--dns_name", "description": "DNS Name", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_name", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "origin_servers", "public_name", "dns_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["refresh interval"], "anchor": "schema-origin_pool--origin_servers--public_name--refresh_interval", "description": "Interval for DNS refresh in seconds. Max value is 7 days as per https://datatracker.ietf.org/doc/HTML/rfc8767.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:origin_servers:public_name", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "origin_servers", "public_name", "refresh_interval"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/origin_pool/origin_servers/public_name/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Specify origin server with public DNS name.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pool.origin_servers.public_name

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/)
- [origin_pool.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/origin_servers/)
- origin_pool.origin_servers.public_name

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public DNS name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name")}
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
public_name {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-origin_pool--origin_servers--public_name--dns_name"></a>

### dns_name property

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-origin_pool--origin_servers--public_name--refresh_interval"></a>

### refresh_interval property

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 604800},
  ),
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

## Next pages

- [origin_pool.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/origin_servers/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
