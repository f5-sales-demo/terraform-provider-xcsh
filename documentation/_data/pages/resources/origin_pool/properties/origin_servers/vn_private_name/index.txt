---
page_title: "origin_servers.vn_private_name"
subcategory: "Load Balancing"
description: "Specify origin server with DNS name on Virtual Network."
xcsh_docs: {"aliases": ["origin servers vn private name"], "body_bytes": 3168, "body_sha256": "sha256:c250852194f5f07470aa0bfe0abf2700826badda7458293e7990ae4fc6ba679c", "capabilities": ["load-balancing", "load-balancing.backend-servers"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_name:private_network"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_name", "parent_id": "xcsh-docs:resources:origin_pool:properties:origin_servers", "path": "documentation/resources/origin_pool/properties/origin_servers/vn_private_name/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1013210030120200-0300013133323110-0211222301220003-2111231123223232-2130003221211000-0010213100230320-3322112211323121-2102120111321112", "registry_path": "docs/guides/resources--origin_pool--reference--group-002.md", "relationships": [{"anchor": "schema-origin_servers--vn_private_name--dns_name", "enforcement": "provider-schema", "group": "origin_servers.vn_private_name:RequiredObjectAttributes:dns_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_name", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "vn_private_name"], "schema_version": 1, "sections": [{"aliases": ["origin servers vn private name dns name"], "anchor": "schema-origin_servers--vn_private_name--dns_name", "description": "DNS Name", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_name", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "vn_private_name", "dns_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin servers vn private name private network"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_name:private_network", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-origin_servers--vn_private_name--private_network--name", "enforcement": "provider-schema", "group": "origin_servers.vn_private_name.private_network:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_name:private_network", "type": "requires"}], "schema_path": ["origin_servers", "vn_private_name", "private_network"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/origin_servers/vn_private_name/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "Specify origin server with DNS name on Virtual Network.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["origin_poolCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.vn_private_name

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/)
- origin_servers.vn_private_name

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with DNS name on Virtual Network.

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
vn_private_name {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-origin_servers--vn_private_name--dns_name"></a>

### dns_name property

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [private_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/vn_private_name/private_network/): complete subsection reference.

## Next pages

- [origin_servers.vn_private_name.private_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/vn_private_name/private_network/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
