---
page_title: "log_receiver_with_net"
subcategory: ""
description: "Select log receiver for logs streaming with network option."
xcsh_docs: {"aliases": ["log receiver with net"], "body_bytes": 3024, "body_sha256": "sha256:b7a2e7d9df70ffb149ccb186bbe14e88a58680436e35799685fe5899a64f48c6", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:log_receiver", "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:use_management_network", "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:use_slo_sli"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/log_receiver_with_net/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2121032223220303-3321210233231023-1001332020100333-3000331113000212-2133111021333102-2331032012311032-2322012211021133-3030101010311030", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-012.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "log_receiver_with_net:ConflictingObjectAttributes:use_management_network,use_slo_sli", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:use_management_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "log_receiver_with_net:ConflictingObjectAttributes:use_management_network,use_slo_sli", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:use_slo_sli", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["log_receiver_with_net"], "schema_version": 1, "sections": [{"aliases": ["log receiver"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:log_receiver", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-log_receiver_with_net--log_receiver--name", "enforcement": "provider-schema", "group": "log_receiver_with_net.log_receiver:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:log_receiver", "type": "requires"}], "schema_path": ["log_receiver_with_net", "log_receiver"], "syntax": "block", "type": "object"}, {"aliases": ["use management network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:use_management_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["log_receiver_with_net", "use_management_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["use slo sli"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:use_slo_sli", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["log_receiver_with_net", "use_slo_sli"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/log_receiver_with_net/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Select log receiver for logs streaming with network option.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# log_receiver_with_net

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- log_receiver_with_net

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: log\_receiver\_with\_net, logs\_streaming\_disabled\] Select log receiver for logs
streaming with network option.

Upstream description:

Select log receiver for logs streaming with network option.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("use_management_network",
    "use_slo_sli")}
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
  "x-ves-oneof-field-network_choice": "[\"use_management_network\",\"use_slo_sli\"]"
}
```

OneOf alternatives in this subsection:

- [log_receiver_with_net](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/log_receiver_with_net/#section)
- [logs_streaming_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/logs_streaming_disabled/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
log_receiver_with_net {
  # Configure direct properties listed below.
}
```

## Direct properties

- [log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/log_receiver_with_net/log_receiver/): complete subsection reference.

- [use_management_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/log_receiver_with_net/use_management_network/): complete subsection reference.

- [use_slo_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/log_receiver_with_net/use_slo_sli/): complete subsection reference.

## Next pages

- [log_receiver_with_net.log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/log_receiver_with_net/log_receiver/)
- [log_receiver_with_net.use_management_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/log_receiver_with_net/use_management_network/)
- [log_receiver_with_net.use_slo_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/log_receiver_with_net/use_slo_sli/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
