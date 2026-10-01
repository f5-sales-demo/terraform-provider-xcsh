---
page_title: "log_receiver_with_net"
subcategory: ""
description: "log_receiver_with_net for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 3024, "body_sha256": "sha256:b7a2e7d9df70ffb149ccb186bbe14e88a58680436e35799685fe5899a64f48c6", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:log_receiver", "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:use_management_network", "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:use_slo_sli"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/log_receiver_with_net/index.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["log_receiver_with_net"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/log_receiver_with_net/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "log_receiver_with_net for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
