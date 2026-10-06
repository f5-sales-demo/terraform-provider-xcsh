---
page_title: "log_receiver_with_net"
subcategory: ""
description: "Select log receiver for logs streaming with network option."
xcsh_docs: {"aliases": ["log receiver with net"], "body_bytes": 2154, "body_sha256": "sha256:98cbd9fa383d52170da58b0e58e3bd4ff7f1956e1ec2aea3df7a9d8ff5b9de57", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:log_receiver", "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:use_management_network", "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:use_slo_sli"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/log_receiver_with_net/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2121032223220303-3321210233231023-1001332020100333-3000331113000212-2133111021333102-2331032012311032-2322012211021133-3030101010311030", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-011.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "log_receiver_with_net:ConflictingObjectAttributes:use_management_network,use_slo_sli", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:use_management_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "log_receiver_with_net:ConflictingObjectAttributes:use_management_network,use_slo_sli", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:use_slo_sli", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["log_receiver_with_net"], "schema_version": 1, "sections": [{"aliases": ["log receiver with net log receiver"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:log_receiver", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-log_receiver_with_net--log_receiver--name", "enforcement": "provider-schema", "group": "log_receiver_with_net.log_receiver:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:log_receiver", "type": "requires"}], "schema_path": ["log_receiver_with_net", "log_receiver"], "syntax": "block", "type": "object"}, {"aliases": ["log receiver with net use management network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:use_management_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["log_receiver_with_net", "use_management_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["log receiver with net use slo sli"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:use_slo_sli", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["log_receiver_with_net", "use_slo_sli"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/log_receiver_with_net/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Select log receiver for logs streaming with network option.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
