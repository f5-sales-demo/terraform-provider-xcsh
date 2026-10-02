---
page_title: "log_receiver_with_net"
subcategory: ""
description: "Select log receiver for logs streaming with network option."
xcsh_docs: {"aliases": ["log receiver with net"], "body_bytes": 2747, "body_sha256": "sha256:b5118aaef7d1b560e322c5298b605c1146427415198317182cb388d2be79e62e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:log_receiver_with_net:log_receiver", "xcsh-docs:data-sources:securemesh_site_v2:properties:log_receiver_with_net:use_management_network", "xcsh-docs:data-sources:securemesh_site_v2:properties:log_receiver_with_net:use_slo_sli"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:log_receiver_with_net", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "documentation/data-sources/securemesh_site_v2/properties/log_receiver_with_net/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3023110201200001-2313230223011122-3303112211100000-3132120202123300-1210303000011110-0232103302130300-0331113010221003-0103313111112332", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["log_receiver_with_net"], "schema_version": 1, "sections": [{"aliases": ["log receiver"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:log_receiver_with_net:log_receiver", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["log_receiver_with_net", "log_receiver"], "syntax": "attribute", "type": "object"}, {"aliases": ["use management network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:log_receiver_with_net:use_management_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["log_receiver_with_net", "use_management_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["use slo sli"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:log_receiver_with_net:use_slo_sli", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["log_receiver_with_net", "use_slo_sli"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/log_receiver_with_net/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Select log receiver for logs streaming with network option.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# log_receiver_with_net

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- log_receiver_with_net

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: log\_receiver\_with\_net, logs\_streaming\_disabled\] Select log receiver for logs
streaming with network option.

Upstream description:

Select log receiver for logs streaming with network option.

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

- [log_receiver_with_net](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/log_receiver_with_net/#section)
- [logs_streaming_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/logs_streaming_disabled/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/log_receiver_with_net/log_receiver/): complete subsection reference.

- [use_management_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/log_receiver_with_net/use_management_network/): complete subsection reference.

- [use_slo_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/log_receiver_with_net/use_slo_sli/): complete subsection reference.

## Next pages

- [log_receiver_with_net.log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/log_receiver_with_net/log_receiver/)
- [log_receiver_with_net.use_management_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/log_receiver_with_net/use_management_network/)
- [log_receiver_with_net.use_slo_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/log_receiver_with_net/use_slo_sli/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
