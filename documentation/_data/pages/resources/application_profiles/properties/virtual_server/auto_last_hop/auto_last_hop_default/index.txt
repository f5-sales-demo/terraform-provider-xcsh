---
page_title: "virtual_server.auto_last_hop.auto_last_hop_default"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["virtual server auto last hop auto last hop default"], "body_bytes": 1267, "body_sha256": "sha256:855e9f0aa5a9e7c14eeb4829aa816289ccdc2c475a646580bd35103328aa3605", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_default", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop", "path": "documentation/resources/application_profiles/properties/virtual_server/auto_last_hop/auto_last_hop_default/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3300231123130002-0223100030003331-3103032001323132-0123301233320131-2132201322230132-1322110221110023-3231131230133312-0213310310001103", "registry_path": "docs/guides/resources--application_profiles--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "auto_last_hop", "auto_last_hop_default"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/auto_last_hop/auto_last_hop_default/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["application_profilesCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.auto_last_hop.auto_last_hop_default

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [virtual_server.auto_last_hop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/auto_last_hop/)
- virtual_server.auto_last_hop.auto_last_hop_default

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for auto last hop default.

Additional upstream details:

This can be used for messages where no values are needed.

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
auto_last_hop_default = {}
```

This is an empty object or choice marker. It has no direct properties.
