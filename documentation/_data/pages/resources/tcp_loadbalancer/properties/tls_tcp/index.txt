---
page_title: "tls_tcp"
subcategory: "Load Balancing"
description: "Choice for selecting TLS over TCP proxy with bring your own certificates."
xcsh_docs: {"aliases": ["tls tcp"], "body_bytes": 1957, "body_sha256": "sha256:ae8b7e0738d13037bd57aedcf2efac7d903e007eab979a0176b1f60d3563778a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params", "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:reference", "path": "documentation/resources/tcp_loadbalancer/properties/tls_tcp/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120", "registry_path": "docs/guides/resources--tcp_loadbalancer--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_tcp:ConflictingObjectAttributes:tls_cert_params,tls_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_tcp:ConflictingObjectAttributes:tls_cert_params,tls_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_tcp"], "schema_version": 1, "sections": [{"aliases": ["tls tcp tls cert params"], "anchor": "section", "description": "Select TLS Parameters and Certificates.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_tcp.tls_cert_params:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:no_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_tcp.tls_cert_params:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:use_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_tcp.tls_cert_params:RequiredObjectAttributes:certificates", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:certificates", "type": "requires"}], "schema_path": ["tls_tcp", "tls_cert_params"], "syntax": "block", "type": "object"}, {"aliases": ["tls tcp tls parameters"], "anchor": "section", "description": "Inline TLS parameters.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_tcp.tls_parameters:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:no_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_tcp.tls_parameters:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:use_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_tcp.tls_parameters:RequiredObjectAttributes:tls_certificates", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates", "type": "requires"}], "schema_path": ["tls_tcp", "tls_parameters"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/tls_tcp/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Choice for selecting TLS over TCP proxy with bring your own certificates.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/)
- tls_tcp

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting TLS over TCP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("tls_cert_params",
    "tls_parameters")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

Terraform syntax:

```terraform
tls_tcp {
  # Configure direct properties listed below.
}
```

## Direct properties

- [tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_cert_params/): complete subsection reference.

- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/): complete subsection reference.

## Next pages

- [tls_tcp.tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_cert_params/)
- [tls_tcp.tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/)
- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
