---
page_title: "ipsec.ipsec_tunnel_parameters.segment"
subcategory: ""
description: "Reference to Segment Object."
xcsh_docs: {"aliases": ["ipsec ipsec tunnel parameters segment"], "body_bytes": 1984, "body_sha256": "sha256:6420b3f72f84fc7972ee6fdc1efe8248d0073b857ac6c66a1e9072d931a4250e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment:refs"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment", "parent_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "path": "documentation/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1322200131300021-2200232001312301-2123232000201310-3131302201230021-0000003021103031-1222031021130120-1133230031113323-3321303122021132", "registry_path": "docs/guides/resources--external_connector--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters.segment:RequiredObjectAttributes:refs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment:refs", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ipsec", "ipsec_tunnel_parameters", "segment"], "schema_version": 1, "sections": [{"aliases": ["refs"], "anchor": "section", "description": "Reference to Segment Object.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment:refs", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "segment", "refs"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Reference to Segment Object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["external_connectorCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ipsec_tunnel_parameters.segment

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/)
- [ipsec.ipsec_tunnel_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/)
- ipsec.ipsec_tunnel_parameters.segment

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Segment Reference Type. Reference to Segment Object.

Upstream description:

Reference to Segment Object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("refs")}
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
segment {
  # Configure direct properties listed below.
}
```

## Direct properties

- [refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/refs/): complete subsection reference.

## Next pages

- [ipsec.ipsec_tunnel_parameters.segment.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/refs/)
- [ipsec.ipsec_tunnel_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/)
- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
