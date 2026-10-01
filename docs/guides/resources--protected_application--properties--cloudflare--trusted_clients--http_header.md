---
page_title: "cloudflare.trusted_clients.http_header"
subcategory: ""
description: "cloudflare.trusted_clients.http_header for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 1630, "body_sha256": "sha256:5b9889ff46ebfcc3f8d2e6ce7d5752fc766972f2f710bc93bd3657bf0bd6c137", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header:headers"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients", "path": "docs/guides/resources--protected_application--properties--cloudflare--trusted_clients--http_header.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudflare", "trusted_clients", "http_header"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/trusted_clients/http_header/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudflare.trusted_clients.http_header for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.trusted_clients.http_header

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- [cloudflare](resources--protected_application--properties--cloudflare.md)
- [cloudflare.trusted_clients](resources--protected_application--properties--cloudflare--trusted_clients.md)
- cloudflare.trusted_clients.http_header

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header.

Upstream description:

Request header name and value pairs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("headers")}
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
http_header {
  # Configure direct properties listed below.
}
```

## Direct properties

- [headers](resources--protected_application--properties--cloudflare--trusted_clients--http_header--headers.md): complete subsection reference.

## Next pages

- [cloudflare.trusted_clients.http_header.headers](resources--protected_application--properties--cloudflare--trusted_clients--http_header--headers.md)
- [cloudflare.trusted_clients](resources--protected_application--properties--cloudflare--trusted_clients.md)
- [xcsh_protected_application](../resources/protected_application.md)
