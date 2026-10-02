---
page_title: "blocked_clients.http_header"
subcategory: "Load Balancing"
description: "Request header name and value pairs."
xcsh_docs: {"aliases": ["blocked clients http header"], "body_bytes": 1761, "body_sha256": "sha256:c695ebb571bac7177acf603bb89f1d2b7695abb9824168839a0b31d51b6aa27c", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:blocked_clients:http_header:headers"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:blocked_clients:http_header", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:blocked_clients", "path": "documentation/resources/cdn_loadbalancer/properties/blocked_clients/http_header/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0001133200202201-0112321323013102-2223132032103103-3300233111313011-3331213132131322-0202013020301310-3320013131330120-0022033221232033", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-007.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "blocked_clients.http_header:RequiredObjectAttributes:headers", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:blocked_clients:http_header:headers", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["blocked_clients", "http_header"], "schema_version": 1, "sections": [{"aliases": ["headers"], "anchor": "section", "description": "List of HTTP header name and value pairs.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:blocked_clients:http_header:headers", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-blocked_clients--http_header--headers--exact", "enforcement": "provider-schema", "group": "blocked_clients.http_header.headers:ConflictingListObjectAttributes:exact,presence", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:blocked_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-blocked_clients--http_header--headers--exact", "enforcement": "provider-schema", "group": "blocked_clients.http_header.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:blocked_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-blocked_clients--http_header--headers--presence", "enforcement": "provider-schema", "group": "blocked_clients.http_header.headers:ConflictingListObjectAttributes:exact,presence", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:blocked_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-blocked_clients--http_header--headers--presence", "enforcement": "provider-schema", "group": "blocked_clients.http_header.headers:ConflictingListObjectAttributes:presence,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:blocked_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-blocked_clients--http_header--headers--regex", "enforcement": "provider-schema", "group": "blocked_clients.http_header.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:blocked_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-blocked_clients--http_header--headers--regex", "enforcement": "provider-schema", "group": "blocked_clients.http_header.headers:ConflictingListObjectAttributes:presence,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:blocked_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-blocked_clients--http_header--headers--name", "enforcement": "provider-schema", "group": "blocked_clients.http_header.headers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:blocked_clients:http_header:headers", "type": "requires"}], "schema_path": ["blocked_clients", "http_header", "headers"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/blocked_clients/http_header/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Request header name and value pairs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blocked_clients.http_header

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [blocked_clients](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/blocked_clients/)
- blocked_clients.http_header

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

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/blocked_clients/http_header/headers/): complete subsection reference.

## Next pages

- [blocked_clients.http_header.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/blocked_clients/http_header/headers/)
- [blocked_clients](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/blocked_clients/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
