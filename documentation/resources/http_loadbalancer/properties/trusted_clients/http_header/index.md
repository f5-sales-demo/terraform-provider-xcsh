---
page_title: "trusted_clients.http_header"
subcategory: "Load Balancing"
description: "Request header name and value pairs."
xcsh_docs: {"aliases": ["trusted clients http header"], "body_bytes": 1770, "body_sha256": "sha256:e1bc91d77da3c1291248afe037e6136411ce17a0a6545896687423ef373ddba2", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:trusted_clients:http_header:headers"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:trusted_clients:http_header", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:trusted_clients", "path": "documentation/resources/http_loadbalancer/properties/trusted_clients/http_header/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0130212200202101-2221112101301312-2221330231323300-2033221313303030-0332100303130330-3000322033023322-2211001320111312-0312213323202112", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-027.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "trusted_clients.http_header:RequiredObjectAttributes:headers", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:trusted_clients:http_header:headers", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["trusted_clients", "http_header"], "schema_version": 1, "sections": [{"aliases": ["trusted clients http header headers"], "anchor": "section", "description": "List of HTTP header name and value pairs.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:trusted_clients:http_header:headers", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-trusted_clients--http_header--headers--exact", "enforcement": "provider-schema", "group": "trusted_clients.http_header.headers:ConflictingListObjectAttributes:exact,presence", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:trusted_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-trusted_clients--http_header--headers--exact", "enforcement": "provider-schema", "group": "trusted_clients.http_header.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:trusted_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-trusted_clients--http_header--headers--presence", "enforcement": "provider-schema", "group": "trusted_clients.http_header.headers:ConflictingListObjectAttributes:exact,presence", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:trusted_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-trusted_clients--http_header--headers--presence", "enforcement": "provider-schema", "group": "trusted_clients.http_header.headers:ConflictingListObjectAttributes:presence,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:trusted_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-trusted_clients--http_header--headers--regex", "enforcement": "provider-schema", "group": "trusted_clients.http_header.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:trusted_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-trusted_clients--http_header--headers--regex", "enforcement": "provider-schema", "group": "trusted_clients.http_header.headers:ConflictingListObjectAttributes:presence,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:trusted_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-trusted_clients--http_header--headers--name", "enforcement": "provider-schema", "group": "trusted_clients.http_header.headers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:trusted_clients:http_header:headers", "type": "requires"}], "schema_path": ["trusted_clients", "http_header", "headers"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/trusted_clients/http_header/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Request header name and value pairs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# trusted_clients.http_header

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [trusted_clients](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/trusted_clients/)
- trusted_clients.http_header

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

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/trusted_clients/http_header/headers/): complete subsection reference.

## Next pages

- [trusted_clients.http_header.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/trusted_clients/http_header/headers/)
- [trusted_clients](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/trusted_clients/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
