---
page_title: "cloudflare.trusted_clients"
subcategory: ""
description: "cloudflare.trusted_clients for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 2786, "body_sha256": "sha256:45b8346b9ecf7208d62016547ca88288443e7c1908bd9dff0984952479ac645f", "canonical_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:trusted_clients", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudflare:trusted_clients:http_header", "xcsh-docs:data-sources:protected_application:properties:cloudflare:trusted_clients:metadata"], "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:trusted_clients", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare", "path": "docs/guides/data-sources--protected_application--properties--cloudflare--trusted_clients.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudflare", "trusted_clients"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudflare/trusted_clients/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudflare.trusted_clients for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# cloudflare.trusted_clients

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md)
- [Property reference](data-sources--protected_application--reference.md)
- [cloudflare](data-sources--protected_application--properties--cloudflare.md)
- cloudflare.trusted_clients

<a id="section"></a>

Type: `"list"`. Computed.

Define your allowlists to skip Bot Defense inference processing.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [http_header](data-sources--protected_application--properties--cloudflare--trusted_clients--http_header.md): complete subsection reference.

<a id="schema-cloudflare--trusted_clients--ip_prefix"></a>

### ip_prefix property

Type: `"string"`. Computed.

Exclusive with \[http\_header\] IP prefix string.

Upstream description:

Exclusive with \[http\_header\] IP prefix string.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [metadata](data-sources--protected_application--properties--cloudflare--trusted_clients--metadata.md): complete subsection reference.

## Next pages

- [cloudflare.trusted_clients.http_header](data-sources--protected_application--properties--cloudflare--trusted_clients--http_header.md)
- [cloudflare.trusted_clients.metadata](data-sources--protected_application--properties--cloudflare--trusted_clients--metadata.md)
- [cloudflare](data-sources--protected_application--properties--cloudflare.md)
- [xcsh_protected_application](../data-sources/protected_application.md)
